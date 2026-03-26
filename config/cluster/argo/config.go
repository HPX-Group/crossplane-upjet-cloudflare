package argo

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "argo"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_argo_smart_routing", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ArgoSmartRouting"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_argo_tiered_caching", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ArgoTieredCaching"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_tiered_cache", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "TieredCache"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_regional_tiered_cache", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RegionalTieredCache"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	// Magic WAN
	p.AddResourceConfigurator("cloudflare_magic_wan_gre_tunnel", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "MagicWANGRETunnel"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_magic_wan_ipsec_tunnel", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "MagicWANIPSecTunnel"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if v, ok := attr["psk"]; ok {
				conn["psk"] = []byte(v.(string))
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("cloudflare_magic_wan_static_route", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "MagicWANStaticRoute"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	// Magic Network Monitoring
	p.AddResourceConfigurator("cloudflare_magic_network_monitoring_configuration", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "MagicNetworkMonitoringConfiguration"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_magic_network_monitoring_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "MagicNetworkMonitoringRule"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	// Magic Transit
	p.AddResourceConfigurator("cloudflare_magic_transit_connector", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "MagicTransitConnector"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if v, ok := attr["license_key"]; ok {
				conn["license_key"] = []byte(v.(string))
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("cloudflare_magic_transit_site", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "MagicTransitSite"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_magic_transit_site_acl", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "MagicTransitSiteACL"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_magic_transit_site_lan", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "MagicTransitSiteLAN"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_magic_transit_site_wan", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "MagicTransitSiteWAN"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})
}
