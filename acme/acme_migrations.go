package acme

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/copystructure"
)

// resourceACMERegistrationStateUpgraderV1 returns the state upgrader
// that handles migrations from version 1 to version 2 for
// acme_registration.
func resourceACMERegistrationStateUpgraderV1() schema.StateUpgrader {
	return schema.StateUpgrader{
		Version: 1,
		Type:    resourceACMERegistrationV1().CoreConfigSchema().ImpliedType(),
		Upgrade: resourceACMERegistrationStateUpgraderV1Func,
	}
}

// resourceACMERegistrationStateUpgraderV1Func provides Terraform 0.12
// state upgrade functionality from schema version 1 to schema
// version 2 for acme_registration.
func resourceACMERegistrationStateUpgraderV1Func(
	_ context.Context,
	rawState map[string]any,
	meta any,
) (map[string]any, error) {
	z, err := copystructure.Copy(rawState)
	if err != nil {
		return nil, err
	}
	result := z.(map[string]any)
	result["account_key_algorithm"] = keyAlgorithmECDSA
	result["account_key_ecdsa_curve"] = keyECDSACurveP384
	result["account_key_rsa_bits"] = 4096
	return result, nil
}

// resourceACMECertificateStateUpgraderV4 returns the state upgrader
// that handles migrations from version 4 to version 5 for
// acme_certificate.
func resourceACMECertificateStateUpgraderV4() schema.StateUpgrader {
	return schema.StateUpgrader{
		Version: 4,
		Type:    resourceACMECertificateV4().CoreConfigSchema().ImpliedType(),
		Upgrade: resourceACMECertificateStateUpgraderV4Func,
	}
}

// resourceACMECertificateStateUpgraderV4Func provides Terraform 0.12
// state upgrade functionality from schema version 4 to schema
// version 5 for acme_certificate.
func resourceACMECertificateStateUpgraderV4Func(
	_ context.Context,
	rawState map[string]any,
	meta any,
) (map[string]any, error) {
	resourceUUID, err := uuid.GenerateUUID()
	if err != nil {
		return nil, fmt.Errorf("error generating new UUID for resource: %s", err)
	}

	z, err := copystructure.Copy(rawState)
	if err != nil {
		return nil, err
	}
	result := z.(map[string]any)
	result["id"] = resourceUUID
	return result, nil
}

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
	_ context.Context,
	rawState map[string]any,
	meta any,
) (map[string]any, error) {
	z, err := copystructure.Copy(rawState)
	if err != nil {
		return nil, err
	}
	result := z.(map[string]any)

	a, ok := rawState["dns_challenge"]
	if ok {
		b, ok := a.([]any)
		if ok && len(b) > 0 {
			c, ok := b[0].(map[string]any)
			if ok {
				d, ok := c["recursive_nameservers"]
				if ok {
					// Should be safe here to access this key directly.
					delete(result["dns_challenge"].([]any)[0].(map[string]any), "recursive_nameservers")
					result["recursive_nameservers"] = d
				}
			}
		}
	}

	return result, nil
}
