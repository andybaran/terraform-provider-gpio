package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// testAccProtoV6ProviderFactories is used by acceptance tests.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"gpio": providerserver.NewProtocol6WithError(New("dev")()),
}

func TestProviderBuilds(t *testing.T) {
	// Simply ensure the provider can be constructed for tests
	factory := providerserver.NewProtocol6WithError(New("dev")())
	if _, err := factory(); err != nil {
		t.Fatalf("provider failed to build: %s", err)
	}
}

func testAccPreCheck(t *testing.T) { /* no-op for now */ }
