package email

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "email"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_email_routing_address", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "EmailRoutingAddress"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_email_routing_catch_all", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "EmailRoutingCatchAll"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_email_routing_dns", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "EmailRoutingDNS"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_email_routing_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "EmailRoutingRule"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_email_routing_settings", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "EmailRoutingSettings"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	// Email Security
	p.AddResourceConfigurator("cloudflare_email_security_block_sender", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "EmailSecurityBlockSender"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_email_security_impersonation_registry", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "EmailSecurityImpersonationRegistry"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_email_security_trusted_domains", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "EmailSecurityTrustedDomains"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})
}
