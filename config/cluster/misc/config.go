package misc

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "misc"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_web3_hostname", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Web3Hostname"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_address_map", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AddressMap"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_byo_ip_prefix", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "BYOIPPrefix"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_regional_hostname", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RegionalHostname"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_snippet", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Snippet"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_snippet_rules", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SnippetRules"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_snippets", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Snippets"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_turnstile_widget", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "TurnstileWidget"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if v, ok := attr["secret"]; ok {
				conn["secret"] = []byte(v.(string))
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("cloudflare_web_analytics_site", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WebAnalyticsSite"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_web_analytics_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WebAnalyticsRule"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_queue", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Queue"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_queue_consumer", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "QueueConsumer"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.References["queue_id"] = config.Reference{
			TerraformName: "cloudflare_queue",
		}
	})

	p.AddResourceConfigurator("cloudflare_hyperdrive_config", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "HyperdriveConfig"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if v, ok := attr["access_client_secret"]; ok {
				conn["access_client_secret"] = []byte(v.(string))
			}
			if v, ok := attr["password"]; ok {
				conn["password"] = []byte(v.(string))
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("cloudflare_d1_database", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "D1Database"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_observatory_scheduled_test", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ObservatoryScheduledTest"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_cloudforce_one_request", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "CloudforceOneRequest"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_cloudforce_one_request_asset", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "CloudforceOneRequestAsset"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_cloudforce_one_request_message", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "CloudforceOneRequestMessage"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_cloudforce_one_request_priority", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "CloudforceOneRequestPriority"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_image", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Image"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_image_variant", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ImageVariant"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_calls_sfu_app", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "CallsSFUApp"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if v, ok := attr["secret"]; ok {
				conn["secret"] = []byte(v.(string))
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("cloudflare_calls_turn_app", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "CallsTURNApp"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if v, ok := attr["key"]; ok {
				conn["key"] = []byte(v.(string))
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("cloudflare_content_scanning", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ContentScanning"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_content_scanning_expression", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ContentScanningExpression"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_cloud_connector_rules", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "CloudConnectorRules"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_connectivity_directory_service", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ConnectivityDirectoryService"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_sso_connector", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SSOConnector"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_token_validation_config", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "TokenValidationConfig"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_token_validation_rules", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "TokenValidationRules"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "User"
	})

	p.AddResourceConfigurator("cloudflare_organization", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Organization"
	})

	p.AddResourceConfigurator("cloudflare_organization_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "OrganizationProfile"
	})
}
