package alicloud

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth"
)

type Provider interface {
	Retrieve() (auth.Credential, error)
}

func NewChainProvider(providers []Provider) Provider {
	return &ChainProvider{
		Providers: providers,
	}
}

type ChainProvider struct {
	Providers []Provider
}

func (p *ChainProvider) Retrieve() (auth.Credential, error) {
	var lastErr error
	for _, p := range p.Providers {
		creds, err := p.Retrieve()
		if err == nil {
			return creds, nil
		}
		lastErr = err
	}
	return nil, lastErr
}
