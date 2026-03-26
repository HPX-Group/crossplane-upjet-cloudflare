package firewall

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "firewall"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_filter", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_firewall_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "FirewallRule"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["filter_id"] = config.Reference{
			TerraformName: "cloudflare_filter",
		}
	})

	p.AddResourceConfigurator("cloudflare_rate_limit", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RateLimit"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_ruleset", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Ruleset"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_managed_transforms", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ManagedTransforms"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_url_normalization_settings", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "URLNormalizationSettings"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_user_agent_blocking_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "UserAgentBlockingRule"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_page_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PageRule"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_bot_management", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "BotManagement"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_leaked_credential_check", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LeakedCredentialCheck"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_leaked_credential_check_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LeakedCredentialCheckRule"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_page_shield_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PageShieldPolicy"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_custom_pages", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "CustomPages"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})
}
