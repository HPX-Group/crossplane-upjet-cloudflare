package dns

import (
	"context"

	"github.com/crossplane/upjet/v2/pkg/config"
)

const shortGroup = "dns"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_dns_record", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DNSRecord"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		// The Cloudflare v5 Terraform provider stores the DNS record ID as a
		// composite "zone_id/dns_record_id" in the Terraform state and requires
		// that format for refresh and import operations.
		//
		// Without this fix, a fresh resource has external-name="" which causes
		// EnsureTFState to write id="" to the state file. Terraform then calls
		// the Cloudflare provider's Read function with id="" which fails with
		// "missing required dns_record_id parameter" instead of returning
		// not-found, so the resource never proceeds to Create.
		//
		// The fix: when external-name is empty (resource not yet created), return
		// a placeholder composite ID using the zone_id. The Cloudflare API returns
		// 404 for a non-existent record, which Terraform treats as "resource
		// removed", the refresh succeeds with empty state, and Observe correctly
		// returns ResourceDoesNotExist so Create can proceed.
		r.ExternalName = config.IdentifierFromProvider
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, parameters map[string]any, _ map[string]any) (string, error) {
			if externalName != "" {
				return externalName, nil
			}
			zoneID, _ := parameters["zone_id"].(string)
			if zoneID == "" {
				return "", nil
			}
			// Placeholder composite ID: triggers a 404 from the Cloudflare API
			// rather than the "missing parameter" error from an empty id.
			return zoneID + "/00000000-0000-0000-0000-000000000000", nil
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
