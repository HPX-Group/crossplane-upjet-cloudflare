package dns

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "dns"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_dns_record", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DNSRecord"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_dns_firewall", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DNSFirewall"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_account_dns_settings", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AccountDNSSettings"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_account_dns_settings_internal_view", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AccountDNSSettingsInternalView"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	// DNS Zone Transfers
	p.AddResourceConfigurator("cloudflare_dns_zone_transfers_acl", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DNSZoneTransfersACL"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_dns_zone_transfers_incoming", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DNSZoneTransfersIncoming"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_dns_zone_transfers_outgoing", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DNSZoneTransfersOutgoing"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_dns_zone_transfers_peer", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DNSZoneTransfersPeer"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_dns_zone_transfers_tsig", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DNSZoneTransfersTSIG"
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
}
