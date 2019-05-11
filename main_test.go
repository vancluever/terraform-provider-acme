package main

import (
	"reflect"
	"sort"
	"testing"

	"github.com/hashicorp/terraform/plugin"
)

// protocolVersionsExpected provides the expected protocol versions
// that this release of the plugin is built with.
var protocolVersionsExpected = []int{4, 5}

// TestPluginVersion tests the protocol versions of the vendored
// Terraform libraries.
func TestPluginVersion(t *testing.T) {
	// We need to source the legacy plugin version from
	// DefaultProtocolVersion as there's no way to get it from the
	// plugin package otherwise.
	protocolVersionsActual := []int{plugin.DefaultProtocolVersion}

	// Add the rest of the plugins from plugin.VersionedPlugins.
	for ver := range plugin.VersionedPlugins {
		protocolVersionsActual = append(protocolVersionsActual, ver)
	}

	sort.Ints(protocolVersionsActual)
	if !reflect.DeepEqual(protocolVersionsExpected, protocolVersionsActual) {
		t.Fatalf("expected versions %v, got %v", protocolVersionsExpected, protocolVersionsActual)
	}
}
