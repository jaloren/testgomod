package newvault

import (
	"github.com/hashicorp/vault/api"
)

// NewClient constructs an api.Client
func NewClient() (*api.Client, error) {
	config := api.DefaultConfig()
	vaultClient, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	return vaultClient, nil
}
