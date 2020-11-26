package acme

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/mitchellh/copystructure"
)

// resourceACMECertificateStateUpgraderV3 returns the state upgrader
// that handles migrations from version 3 to version 4 for
// acme_certificate.
func resourceACMECertificateStateUpgraderV3() schema.StateUpgrader {
	return schema.StateUpgrader{
		Version: 3,
		Type:    resourceACMECertificateV3().CoreConfigSchema().ImpliedType(),
		Upgrade: resourceACMECertificateStateUpgraderV3Func,
	}
}

// resourceACMECertificateStateUpgraderV3Func provides Terraform 0.12
// state upgrade functionality from schema version 3 to schema
// version 4 for acme_certificate.
func resourceACMECertificateStateUpgraderV3Func(
	rawState map[string]interface{},
	meta interface{},
) (map[string]interface{}, error) {
	z, err := copystructure.Copy(rawState)
	if err != nil {
		return nil, err
	}
	result := z.(map[string]interface{})

	a, ok := rawState["dns_challenge"]
	if ok {
		b, ok := a.([]interface{})
		if ok && len(b) > 0 {
			c, ok := b[0].(map[string]interface{})
			if ok {
				d, ok := c["recursive_nameservers"]
				if ok {
					// Should be safe here to access this key directly.
					delete(result["dns_challenge"].([]interface{})[0].(map[string]interface{}), "recursive_nameservers")
					result["recursive_nameservers"] = d
				}
			}
		}
	}

	return result, nil
}
