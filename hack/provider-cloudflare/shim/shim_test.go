package shim

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/provider"
)

func TestNewProviderReturnsValidFactory(t *testing.T) {
	factory := NewProvider("test")
	if factory == nil {
		t.Fatal("NewProvider returned nil factory")
	}

	p := factory()
	if p == nil {
		t.Fatal("factory returned nil provider")
	}

	// Verify it satisfies the provider.Provider interface
	var _ provider.Provider = p
}
