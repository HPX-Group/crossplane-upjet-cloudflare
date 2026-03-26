package account

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "account"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_account", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})

	p.AddResourceConfigurator("cloudflare_account_member", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_account_token", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AccountToken"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if v, ok := attr["value"]; ok {
				conn["value"] = []byte(v.(string))
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("cloudflare_account_subscription", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AccountSubscription"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_api_token", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "APIToken"
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if v, ok := attr["value"]; ok {
				conn["value"] = []byte(v.(string))
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("cloudflare_registrar_domain", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RegistrarDomain"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})
}
