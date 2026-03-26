package logpush

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "logpush"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_logpush_job", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LogpushJob"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if v, ok := attr["ownership_challenge"]; ok {
				conn["ownership_challenge"] = []byte(v.(string))
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("cloudflare_logpull_retention", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LogpullRetention"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_logpush_ownership_challenge", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LogpushOwnershipChallenge"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})
}
