package oldvault

import (
	"github.com/hashicorp/vault/meta"
	"github.com/jaloren/testgomod/newvault"
)

// Commands provides the functions that back locksmith functionality
func Commands() *meta.Meta {
	newvault.NewClient()
	return &meta.Meta{}
}
