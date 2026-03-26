package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Account & Organization Resources
	"cloudflare_account":              config.IdentifierFromProvider,
	"cloudflare_account_member":       config.IdentifierFromProvider,
	"cloudflare_account_token":        config.IdentifierFromProvider,
	"cloudflare_account_subscription": config.IdentifierFromProvider,
	"cloudflare_api_token":            config.IdentifierFromProvider,
	"cloudflare_organization":         config.IdentifierFromProvider,
	"cloudflare_organization_profile": config.IdentifierFromProvider,
	"cloudflare_registrar_domain":     config.IdentifierFromProvider,
	"cloudflare_user":                 config.IdentifierFromProvider,
	"cloudflare_sso_connector":        config.IdentifierFromProvider,

	// Zone & DNS Resources
	"cloudflare_zone":                          config.IdentifierFromProvider,
	"cloudflare_zone_setting":                  config.IdentifierFromProvider,
	"cloudflare_zone_dnssec":                   config.IdentifierFromProvider,
	"cloudflare_zone_lockdown":                 config.IdentifierFromProvider,
	"cloudflare_zone_hold":                     config.IdentifierFromProvider,
	"cloudflare_zone_cache_reserve":            config.IdentifierFromProvider,
	"cloudflare_zone_cache_variants":           config.IdentifierFromProvider,
	"cloudflare_zone_dns_settings":             config.IdentifierFromProvider,
	"cloudflare_zone_subscription":             config.IdentifierFromProvider,
	"cloudflare_dns_record":                    config.IdentifierFromProvider,
	"cloudflare_dns_firewall":                  config.IdentifierFromProvider,
	"cloudflare_account_dns_settings":          config.IdentifierFromProvider,
	"cloudflare_account_dns_settings_internal_view": config.IdentifierFromProvider,
	"cloudflare_dns_zone_transfers_acl":        config.IdentifierFromProvider,
	"cloudflare_dns_zone_transfers_incoming":   config.IdentifierFromProvider,
	"cloudflare_dns_zone_transfers_outgoing":   config.IdentifierFromProvider,
	"cloudflare_dns_zone_transfers_peer":       config.IdentifierFromProvider,
	"cloudflare_dns_zone_transfers_tsig":       config.IdentifierFromProvider,

	// Custom Hostnames & SSL/TLS
	"cloudflare_custom_hostname":                        config.IdentifierFromProvider,
	"cloudflare_custom_hostname_fallback_origin":        config.IdentifierFromProvider,
	"cloudflare_custom_ssl":                             config.IdentifierFromProvider,
	"cloudflare_origin_ca_certificate":                  config.IdentifierFromProvider,
	"cloudflare_certificate_pack":                       config.IdentifierFromProvider,
	"cloudflare_keyless_certificate":                    config.IdentifierFromProvider,
	"cloudflare_mtls_certificate":                       config.IdentifierFromProvider,
	"cloudflare_total_tls":                              config.IdentifierFromProvider,
	"cloudflare_hostname_tls_setting":                   config.IdentifierFromProvider,
	"cloudflare_authenticated_origin_pulls":             config.IdentifierFromProvider,
	"cloudflare_authenticated_origin_pulls_certificate": config.IdentifierFromProvider,
	"cloudflare_authenticated_origin_pulls_settings":    config.IdentifierFromProvider,
	"cloudflare_universal_ssl_setting":                  config.IdentifierFromProvider,

	// Firewall & Security Resources
	"cloudflare_filter":                    config.IdentifierFromProvider,
	"cloudflare_firewall_rule":             config.IdentifierFromProvider,
	"cloudflare_rate_limit":                config.IdentifierFromProvider,
	"cloudflare_ruleset":                   config.IdentifierFromProvider,
	"cloudflare_url_normalization_settings": config.IdentifierFromProvider,
	"cloudflare_user_agent_blocking_rule":  config.IdentifierFromProvider,
	"cloudflare_page_rule":                 config.IdentifierFromProvider,
	"cloudflare_bot_management":            config.IdentifierFromProvider,
	"cloudflare_leaked_credential_check":      config.IdentifierFromProvider,
	"cloudflare_leaked_credential_check_rule": config.IdentifierFromProvider,
	"cloudflare_managed_transforms":        config.IdentifierFromProvider,
	"cloudflare_page_shield_policy":        config.IdentifierFromProvider,
	"cloudflare_custom_pages":              config.IdentifierFromProvider,

	// Access Rule (standalone, not part of zero_trust_access_* rename)
	"cloudflare_access_rule": config.IdentifierFromProvider,

	// Zero Trust Access Resources
	"cloudflare_zero_trust_access_application":            config.IdentifierFromProvider,
	"cloudflare_zero_trust_access_policy":                 config.IdentifierFromProvider,
	"cloudflare_zero_trust_access_group":                  config.IdentifierFromProvider,
	"cloudflare_zero_trust_access_identity_provider":      config.IdentifierFromProvider,
	"cloudflare_zero_trust_access_mtls_certificate":       config.IdentifierFromProvider,
	"cloudflare_zero_trust_access_mtls_hostname_settings": config.IdentifierFromProvider,
	"cloudflare_zero_trust_access_service_token":          config.IdentifierFromProvider,
	"cloudflare_zero_trust_access_key_configuration":      config.IdentifierFromProvider,
	"cloudflare_zero_trust_access_custom_page":            config.IdentifierFromProvider,
	"cloudflare_zero_trust_access_tag":                    config.IdentifierFromProvider,
	"cloudflare_zero_trust_access_short_lived_certificate": config.IdentifierFromProvider,
	"cloudflare_zero_trust_access_infrastructure_target":  config.IdentifierFromProvider,
	"cloudflare_zero_trust_access_ai_controls_mcp_portal": config.IdentifierFromProvider,
	"cloudflare_zero_trust_access_ai_controls_mcp_server": config.IdentifierFromProvider,
	"cloudflare_zero_trust_organization":                  config.IdentifierFromProvider,

	// Zero Trust Gateway & Network Resources
	"cloudflare_zero_trust_list":                      config.IdentifierFromProvider,
	"cloudflare_zero_trust_gateway_policy":            config.IdentifierFromProvider,
	"cloudflare_zero_trust_gateway_settings":          config.IdentifierFromProvider,
	"cloudflare_zero_trust_gateway_certificate":       config.IdentifierFromProvider,
	"cloudflare_zero_trust_gateway_logging":           config.IdentifierFromProvider,
	"cloudflare_zero_trust_gateway_proxy_endpoint":    config.IdentifierFromProvider,
	"cloudflare_zero_trust_dns_location":              config.IdentifierFromProvider,
	"cloudflare_zero_trust_tunnel_cloudflared":        config.IdentifierFromProvider,
	"cloudflare_zero_trust_tunnel_cloudflared_config": config.IdentifierFromProvider,
	"cloudflare_zero_trust_tunnel_cloudflared_route":  config.IdentifierFromProvider,
	"cloudflare_zero_trust_tunnel_cloudflared_virtual_network": config.IdentifierFromProvider,
	"cloudflare_zero_trust_tunnel_warp_connector":     config.IdentifierFromProvider,
	"cloudflare_zero_trust_network_hostname_route":    config.IdentifierFromProvider,

	// Zero Trust Device Resources
	"cloudflare_zero_trust_device_posture_rule":        config.IdentifierFromProvider,
	"cloudflare_zero_trust_device_posture_integration": config.IdentifierFromProvider,
	"cloudflare_zero_trust_device_managed_networks":    config.IdentifierFromProvider,
	"cloudflare_zero_trust_device_settings":            config.IdentifierFromProvider,
	"cloudflare_zero_trust_device_custom_profile":      config.IdentifierFromProvider,
	"cloudflare_zero_trust_device_custom_profile_local_domain_fallback": config.IdentifierFromProvider,
	"cloudflare_zero_trust_device_default_profile":                      config.IdentifierFromProvider,
	"cloudflare_zero_trust_device_default_profile_certificates":         config.IdentifierFromProvider,
	"cloudflare_zero_trust_device_default_profile_local_domain_fallback": config.IdentifierFromProvider,
	"cloudflare_zero_trust_dex_test": config.IdentifierFromProvider,

	// Zero Trust DLP Resources
	"cloudflare_zero_trust_dlp_custom_profile":    config.IdentifierFromProvider,
	"cloudflare_zero_trust_dlp_entry":             config.IdentifierFromProvider,
	"cloudflare_zero_trust_dlp_custom_entry":      config.IdentifierFromProvider,
	"cloudflare_zero_trust_dlp_dataset":           config.IdentifierFromProvider,
	"cloudflare_zero_trust_dlp_integration_entry": config.IdentifierFromProvider,
	"cloudflare_zero_trust_dlp_predefined_entry":  config.IdentifierFromProvider,
	"cloudflare_zero_trust_dlp_predefined_profile": config.IdentifierFromProvider,

	// Zero Trust Risk Resources
	"cloudflare_zero_trust_risk_behavior":           config.IdentifierFromProvider,
	"cloudflare_zero_trust_risk_scoring_integration": config.IdentifierFromProvider,

	// Load Balancing Resources
	"cloudflare_load_balancer":         config.IdentifierFromProvider,
	"cloudflare_load_balancer_pool":    config.IdentifierFromProvider,
	"cloudflare_load_balancer_monitor": config.IdentifierFromProvider,

	// Workers & Pages Resources
	"cloudflare_worker":                                config.IdentifierFromProvider,
	"cloudflare_worker_version":                        config.IdentifierFromProvider,
	"cloudflare_workers_script":                        config.NameAsIdentifier,
	"cloudflare_workers_cron_trigger":                  config.IdentifierFromProvider,
	"cloudflare_workers_custom_domain":                 config.IdentifierFromProvider,
	"cloudflare_workers_deployment":                    config.IdentifierFromProvider,
	"cloudflare_workers_for_platforms_dispatch_namespace": config.IdentifierFromProvider,
	"cloudflare_workers_kv":                            config.IdentifierFromProvider,
	"cloudflare_workers_kv_namespace":                  config.NameAsIdentifier,
	"cloudflare_workers_route":                         config.IdentifierFromProvider,
	"cloudflare_workers_script_subdomain":              config.IdentifierFromProvider,
	"cloudflare_workflow":                              config.IdentifierFromProvider,
	"cloudflare_pages_project":                         config.NameAsIdentifier,
	"cloudflare_pages_domain":                          config.IdentifierFromProvider,

	// R2 Storage Resources
	"cloudflare_r2_bucket":                  config.NameAsIdentifier,
	"cloudflare_r2_custom_domain":           config.IdentifierFromProvider,
	"cloudflare_r2_managed_domain":          config.IdentifierFromProvider,
	"cloudflare_r2_bucket_lock":             config.IdentifierFromProvider,
	"cloudflare_r2_bucket_lifecycle":        config.IdentifierFromProvider,
	"cloudflare_r2_bucket_cors":             config.IdentifierFromProvider,
	"cloudflare_r2_bucket_sippy":            config.IdentifierFromProvider,
	"cloudflare_r2_bucket_event_notification": config.IdentifierFromProvider,

	// Argo & Routing Resources
	"cloudflare_argo_smart_routing":    config.IdentifierFromProvider,
	"cloudflare_argo_tiered_caching":   config.IdentifierFromProvider,
	"cloudflare_tiered_cache":          config.IdentifierFromProvider,
	"cloudflare_regional_tiered_cache": config.IdentifierFromProvider,

	// Magic WAN & Transit Resources
	"cloudflare_magic_wan_gre_tunnel":                   config.IdentifierFromProvider,
	"cloudflare_magic_wan_ipsec_tunnel":                  config.IdentifierFromProvider,
	"cloudflare_magic_wan_static_route":                  config.IdentifierFromProvider,
	"cloudflare_magic_network_monitoring_configuration":  config.IdentifierFromProvider,
	"cloudflare_magic_network_monitoring_rule":           config.IdentifierFromProvider,
	"cloudflare_magic_transit_connector":                 config.IdentifierFromProvider,
	"cloudflare_magic_transit_site":                      config.IdentifierFromProvider,
	"cloudflare_magic_transit_site_acl":                  config.IdentifierFromProvider,
	"cloudflare_magic_transit_site_lan":                  config.IdentifierFromProvider,
	"cloudflare_magic_transit_site_wan":                  config.IdentifierFromProvider,

	// Spectrum Resources
	"cloudflare_spectrum_application": config.IdentifierFromProvider,

	// Logging Resources
	"cloudflare_logpush_job":                 config.IdentifierFromProvider,
	"cloudflare_logpush_ownership_challenge": config.IdentifierFromProvider,

	// Notification Resources
	"cloudflare_notification_policy":          config.IdentifierFromProvider,
	"cloudflare_notification_policy_webhooks": config.IdentifierFromProvider,

	// Healthcheck Resources
	"cloudflare_healthcheck": config.IdentifierFromProvider,

	// Waiting Room Resources
	"cloudflare_waiting_room":          config.IdentifierFromProvider,
	"cloudflare_waiting_room_event":    config.IdentifierFromProvider,
	"cloudflare_waiting_room_rules":    config.IdentifierFromProvider,
	"cloudflare_waiting_room_settings": config.IdentifierFromProvider,

	// Web3 Resources
	"cloudflare_web3_hostname": config.IdentifierFromProvider,

	// Stream Resources
	"cloudflare_stream":                  config.IdentifierFromProvider,
	"cloudflare_stream_key":              config.IdentifierFromProvider,
	"cloudflare_stream_live_input":       config.IdentifierFromProvider,
	"cloudflare_stream_webhook":          config.IdentifierFromProvider,
	"cloudflare_stream_audio_track":      config.IdentifierFromProvider,
	"cloudflare_stream_caption_language": config.IdentifierFromProvider,
	"cloudflare_stream_download":         config.IdentifierFromProvider,
	"cloudflare_stream_watermark":        config.IdentifierFromProvider,

	// Email Routing Resources
	"cloudflare_email_routing_address":   config.IdentifierFromProvider,
	"cloudflare_email_routing_catch_all": config.IdentifierFromProvider,
	"cloudflare_email_routing_dns":       config.IdentifierFromProvider,
	"cloudflare_email_routing_rule":      config.IdentifierFromProvider,
	"cloudflare_email_routing_settings":  config.IdentifierFromProvider,

	// Email Security Resources
	"cloudflare_email_security_block_sender":            config.IdentifierFromProvider,
	"cloudflare_email_security_impersonation_registry":  config.IdentifierFromProvider,
	"cloudflare_email_security_trusted_domains":         config.IdentifierFromProvider,

	// List Resources
	"cloudflare_list":      config.IdentifierFromProvider,
	"cloudflare_list_item": config.IdentifierFromProvider,

	// API Shield Resources
	"cloudflare_api_shield":                                      config.IdentifierFromProvider,
	"cloudflare_api_shield_operation":                            config.IdentifierFromProvider,
	"cloudflare_api_shield_schema":                               config.IdentifierFromProvider,
	"cloudflare_api_shield_operation_schema_validation_settings": config.IdentifierFromProvider,
	"cloudflare_api_shield_schema_validation_settings":           config.IdentifierFromProvider,
	"cloudflare_api_shield_discovery_operation":                  config.IdentifierFromProvider,
	"cloudflare_schema_validation_operation_settings":            config.IdentifierFromProvider,
	"cloudflare_schema_validation_schemas":                       config.IdentifierFromProvider,
	"cloudflare_schema_validation_settings":                      config.IdentifierFromProvider,

	// Address Map & BYO IP Resources
	"cloudflare_address_map":   config.IdentifierFromProvider,
	"cloudflare_byo_ip_prefix": config.IdentifierFromProvider,

	// Regional Hostname
	"cloudflare_regional_hostname": config.IdentifierFromProvider,

	// Snippet Resources
	"cloudflare_snippet":       config.IdentifierFromProvider,
	"cloudflare_snippet_rules": config.IdentifierFromProvider,
	"cloudflare_snippets":      config.IdentifierFromProvider,

	// Turnstile Widget
	"cloudflare_turnstile_widget": config.IdentifierFromProvider,

	// Web Analytics Resources
	"cloudflare_web_analytics_site": config.IdentifierFromProvider,
	"cloudflare_web_analytics_rule": config.IdentifierFromProvider,

	// Queue Resources
	"cloudflare_queue":          config.NameAsIdentifier,
	"cloudflare_queue_consumer": config.IdentifierFromProvider,

	// Hyperdrive Resources
	"cloudflare_hyperdrive_config": config.NameAsIdentifier,

	// D1 Database Resources
	"cloudflare_d1_database": config.NameAsIdentifier,

	// Observatory Resources
	"cloudflare_observatory_scheduled_test": config.IdentifierFromProvider,

	// Cloudforce One Resources
	"cloudflare_cloudforce_one_request":          config.IdentifierFromProvider,
	"cloudflare_cloudforce_one_request_asset":    config.IdentifierFromProvider,
	"cloudflare_cloudforce_one_request_message":  config.IdentifierFromProvider,
	"cloudflare_cloudforce_one_request_priority": config.IdentifierFromProvider,

	// Image Resources
	"cloudflare_image":         config.IdentifierFromProvider,
	"cloudflare_image_variant": config.IdentifierFromProvider,

	// Calls Resources
	"cloudflare_calls_sfu_app":  config.IdentifierFromProvider,
	"cloudflare_calls_turn_app": config.IdentifierFromProvider,

	// Content Scanning Resources
	"cloudflare_content_scanning":            config.IdentifierFromProvider,
	"cloudflare_content_scanning_expression": config.IdentifierFromProvider,

	// Cloud Connector Resources
	"cloudflare_cloud_connector_rules": config.IdentifierFromProvider,

	// Token Validation Resources
	"cloudflare_token_validation_config": config.IdentifierFromProvider,
	"cloudflare_token_validation_rules":  config.IdentifierFromProvider,

	// Connectivity Resources
	"cloudflare_connectivity_directory_service": config.IdentifierFromProvider,

	// Logpull Retention (removed but keeping for now)
	"cloudflare_logpull_retention": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
