package workers

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "workers"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_worker", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Worker"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_worker_version", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WorkerVersion"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if v, ok := attr["jwt"]; ok {
				conn["jwt"] = []byte(v.(string))
			}
			if v, ok := attr["key_base64"]; ok {
				conn["key_base64"] = []byte(v.(string))
			}
			if v, ok := attr["key_jwk"]; ok {
				conn["key_jwk"] = []byte(v.(string))
			}
			if v, ok := attr["text"]; ok {
				conn["text"] = []byte(v.(string))
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("cloudflare_workers_script", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WorkersScript"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if v, ok := attr["jwt"]; ok {
				conn["jwt"] = []byte(v.(string))
			}
			if v, ok := attr["key_base64"]; ok {
				conn["key_base64"] = []byte(v.(string))
			}
			if v, ok := attr["key_jwk"]; ok {
				conn["key_jwk"] = []byte(v.(string))
			}
			if v, ok := attr["text"]; ok {
				conn["text"] = []byte(v.(string))
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("cloudflare_workers_cron_trigger", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WorkersCronTrigger"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_workers_custom_domain", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WorkersCustomDomain"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_workers_deployment", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WorkersDeployment"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_workers_for_platforms_dispatch_namespace", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WorkersForPlatformsDispatchNamespace"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_workers_kv", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WorkersKV"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.References["namespace_id"] = config.Reference{
			TerraformName: "cloudflare_workers_kv_namespace",
		}
	})

	p.AddResourceConfigurator("cloudflare_workers_kv_namespace", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WorkersKVNamespace"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_workers_route", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WorkersRoute"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_workers_script_subdomain", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WorkersScriptSubdomain"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_workflow", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Workflow"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	// Pages
	p.AddResourceConfigurator("cloudflare_pages_project", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PagesProject"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if v, ok := attr["value"]; ok {
				conn["value"] = []byte(v.(string))
			}
			if v, ok := attr["web_analytics_token"]; ok {
				conn["web_analytics_token"] = []byte(v.(string))
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("cloudflare_pages_domain", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PagesDomain"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.References["project_name"] = config.Reference{
			TerraformName: "cloudflare_pages_project",
		}
	})
}
