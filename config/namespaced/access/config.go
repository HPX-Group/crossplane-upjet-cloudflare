package access

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "access"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// cloudflare_access_rule is a standalone resource (not renamed to zero_trust_access_*)
	p.AddResourceConfigurator("cloudflare_access_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AccessRule"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_access_application", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustAccessApplication"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if v, ok := attr["client_secret"]; ok {
				conn["client_secret"] = []byte(v.(string))
			}
			if v, ok := attr["password"]; ok {
				conn["password"] = []byte(v.(string))
			}
			if v, ok := attr["token"]; ok {
				conn["token"] = []byte(v.(string))
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_access_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustAccessPolicy"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_access_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustAccessGroup"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_access_identity_provider", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustAccessIdentityProvider"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if v, ok := attr["client_secret"]; ok {
				conn["client_secret"] = []byte(v.(string))
			}
			if v, ok := attr["secret"]; ok {
				conn["secret"] = []byte(v.(string))
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_access_mtls_certificate", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustAccessMTLSCertificate"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_access_mtls_hostname_settings", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustAccessMTLSHostnameSettings"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_access_service_token", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustAccessServiceToken"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if v, ok := attr["client_secret"]; ok {
				conn["client_secret"] = []byte(v.(string))
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_access_key_configuration", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustAccessKeyConfiguration"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_access_custom_page", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustAccessCustomPage"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_access_tag", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustAccessTag"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_access_short_lived_certificate", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustAccessShortLivedCertificate"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_access_infrastructure_target", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustAccessInfrastructureTarget"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_access_ai_controls_mcp_portal", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustAccessAIControlsMCPPortal"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_access_ai_controls_mcp_server", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustAccessAIControlsMCPServer"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_organization", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustOrganization"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})
}
