// Package shim re-exports the Cloudflare Terraform provider's factory function
// from its internal package. This is necessary because Go's internal/ package
// restriction prevents external modules from importing it directly.
//
// This package must live within the same Go module as the Cloudflare provider
// (github.com/cloudflare/terraform-provider-cloudflare) to access internal/.
package shim

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/hashicorp/terraform-plugin-framework/provider"
)

// NewProvider returns the Cloudflare Terraform provider factory function.
// The returned function creates a new provider.Provider instance each time
// it is called. The version parameter is embedded in the provider for
// user-agent and diagnostic purposes.
func NewProvider(version string) func() provider.Provider {
	return internal.NewProvider(version)
}
