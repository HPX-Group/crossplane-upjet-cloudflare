package zerotrust

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "zerotrust"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_zero_trust_list", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustList"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_gateway_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustGatewayPolicy"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_gateway_settings", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustGatewaySettings"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_gateway_certificate", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustGatewayCertificate"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_gateway_logging", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustGatewayLogging"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_gateway_proxy_endpoint", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustGatewayProxyEndpoint"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_dns_location", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDNSLocation"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_tunnel_cloudflared", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustTunnelCloudflared"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if v, ok := attr["tunnel_secret"]; ok {
				conn["tunnel_secret"] = []byte(v.(string))
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_tunnel_cloudflared_config", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustTunnelCloudflaredConfig"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.References["tunnel_id"] = config.Reference{
			TerraformName: "cloudflare_zero_trust_tunnel_cloudflared",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_tunnel_cloudflared_route", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustTunnelCloudflaredRoute"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.References["tunnel_id"] = config.Reference{
			TerraformName: "cloudflare_zero_trust_tunnel_cloudflared",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_tunnel_cloudflared_virtual_network", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustTunnelCloudflaredVirtualNetwork"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_tunnel_warp_connector", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustTunnelWarpConnector"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if v, ok := attr["tunnel_secret"]; ok {
				conn["tunnel_secret"] = []byte(v.(string))
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_network_hostname_route", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustNetworkHostnameRoute"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	// Device resources
	p.AddResourceConfigurator("cloudflare_zero_trust_device_posture_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDevicePostureRule"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_device_posture_integration", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDevicePostureIntegration"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if v, ok := attr["access_client_secret"]; ok {
				conn["access_client_secret"] = []byte(v.(string))
			}
			if v, ok := attr["client_secret"]; ok {
				conn["client_secret"] = []byte(v.(string))
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_device_managed_networks", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDeviceManagedNetworks"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_device_settings", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDeviceSettings"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_device_custom_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDeviceCustomProfile"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_device_custom_profile_local_domain_fallback", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDeviceCustomProfileLocalDomainFallback"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_device_default_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDeviceDefaultProfile"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_device_default_profile_certificates", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDeviceDefaultProfileCertificates"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_device_default_profile_local_domain_fallback", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDeviceDefaultProfileLocalDomainFallback"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_dex_test", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDEXTest"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	// DLP resources
	p.AddResourceConfigurator("cloudflare_zero_trust_dlp_custom_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDLPCustomProfile"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_dlp_entry", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDLPEntry"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_dlp_custom_entry", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDLPCustomEntry"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_dlp_dataset", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDLPDataset"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_dlp_integration_entry", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDLPIntegrationEntry"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_dlp_predefined_entry", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDLPPredefinedEntry"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_dlp_predefined_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDLPPredefinedProfile"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	// Risk resources
	p.AddResourceConfigurator("cloudflare_zero_trust_risk_behavior", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustRiskBehavior"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_risk_scoring_integration", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustRiskScoringIntegration"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})
}
