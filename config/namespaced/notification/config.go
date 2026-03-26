package notification

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "notification"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_notification_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NotificationPolicy"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_notification_policy_webhooks", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NotificationPolicyWebhooks"
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
