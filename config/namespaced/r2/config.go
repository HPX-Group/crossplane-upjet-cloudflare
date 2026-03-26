package r2

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "r2"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_r2_bucket", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "R2Bucket"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_r2_custom_domain", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "R2CustomDomain"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.References["bucket_name"] = config.Reference{
			TerraformName: "cloudflare_r2_bucket",
		}
	})

	p.AddResourceConfigurator("cloudflare_r2_managed_domain", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "R2ManagedDomain"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.References["bucket_name"] = config.Reference{
			TerraformName: "cloudflare_r2_bucket",
		}
	})

	p.AddResourceConfigurator("cloudflare_r2_bucket_lock", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "R2BucketLock"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.References["bucket_name"] = config.Reference{
			TerraformName: "cloudflare_r2_bucket",
		}
	})

	p.AddResourceConfigurator("cloudflare_r2_bucket_lifecycle", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "R2BucketLifecycle"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.References["bucket_name"] = config.Reference{
			TerraformName: "cloudflare_r2_bucket",
		}
	})

	p.AddResourceConfigurator("cloudflare_r2_bucket_cors", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "R2BucketCORS"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.References["bucket_name"] = config.Reference{
			TerraformName: "cloudflare_r2_bucket",
		}
	})

	p.AddResourceConfigurator("cloudflare_r2_bucket_sippy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "R2BucketSippy"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.References["bucket_name"] = config.Reference{
			TerraformName: "cloudflare_r2_bucket",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if v, ok := attr["secret_access_key"]; ok {
				conn["secret_access_key"] = []byte(v.(string))
			}
			if v, ok := attr["private_key"]; ok {
				conn["private_key"] = []byte(v.(string))
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("cloudflare_r2_bucket_event_notification", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "R2BucketEventNotification"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.References["bucket_name"] = config.Reference{
			TerraformName: "cloudflare_r2_bucket",
		}
	})
}
