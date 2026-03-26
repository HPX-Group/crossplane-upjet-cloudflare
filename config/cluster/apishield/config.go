package apishield

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "apishield"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_api_shield", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "APIShield"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_api_shield_operation", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "APIShieldOperation"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_api_shield_schema", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "APIShieldSchema"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_api_shield_operation_schema_validation_settings", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "APIShieldOperationSchemaValidationSettings"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_api_shield_schema_validation_settings", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "APIShieldSchemaValidationSettings"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_api_shield_discovery_operation", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "APIShieldDiscoveryOperation"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_schema_validation_operation_settings", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SchemaValidationOperationSettings"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_schema_validation_schemas", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SchemaValidationSchemas"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_schema_validation_settings", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SchemaValidationSettings"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})
}
