package stream

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "stream"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_stream", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Stream"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_stream_key", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "StreamKey"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if v, ok := attr["jwk"]; ok {
				conn["jwk"] = []byte(v.(string))
			}
			if v, ok := attr["pem"]; ok {
				conn["pem"] = []byte(v.(string))
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("cloudflare_stream_live_input", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "StreamLiveInput"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if v, ok := attr["stream_key"]; ok {
				conn["stream_key"] = []byte(v.(string))
			}
			if v, ok := attr["url"]; ok {
				conn["url"] = []byte(v.(string))
			}
			if v, ok := attr["passphrase"]; ok {
				conn["passphrase"] = []byte(v.(string))
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("cloudflare_stream_webhook", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "StreamWebhook"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_stream_audio_track", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "StreamAudioTrack"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_stream_caption_language", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "StreamCaptionLanguage"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_stream_download", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "StreamDownload"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_stream_watermark", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "StreamWatermark"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})
}
