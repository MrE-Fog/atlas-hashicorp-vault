package atlasvault

import (
	"strings"

	"github.com/hashicorp/vault/sdk/framework"
)

func normalizeSerial(serial string) string {
	return strings.Replace(strings.ToLower(serial), ":", "-", -1)
}

// addIssueAndSignCommonFields adds fields common to both CA and non-CA issuing
// and signing
func addIssueAndSignCommonFields(fields map[string]*framework.FieldSchema) map[string]*framework.FieldSchema {
	fields["exclude_cn_from_sans"] = &framework.FieldSchema{
		Type:    framework.TypeBool,
		Default: false,
		Description: `If true, the Common Name will not be
included in DNS or Email Subject Alternate Names.
Defaults to false (CN is included).`,
	}

	fields["format"] = &framework.FieldSchema{
		Type:    framework.TypeString,
		Default: "pem",
		Description: `Format for returned data. Can be "pem", "der",
or "pem_bundle". If "pem_bundle" any private
key and issuing cert will be appended to the
certificate pem. Defaults to "pem".`,
		AllowedValues: []interface{}{"pem", "der", "pem_bundle"},
	}

	fields["private_key_format"] = &framework.FieldSchema{
		Type:    framework.TypeString,
		Default: "der",
		Description: `Format for the returned private key.
Generally the default will be controlled by the "format"
parameter as either base64-encoded DER or PEM-encoded DER.
However, this can be set to "pkcs8" to have the returned
private key contain base64-encoded pkcs8 or PEM-encoded
pkcs8 instead. Defaults to "der".`,
		AllowedValues: []interface{}{"", "der", "pem", "pkcs8"},
	}

	fields["ip_sans"] = &framework.FieldSchema{
		Type: framework.TypeCommaStringSlice,
		Description: `The requested IP SANs, if any, in a
comma-delimited list`,
	}

	fields["uri_sans"] = &framework.FieldSchema{
		Type: framework.TypeCommaStringSlice,
		Description: `The requested URI SANs, if any, in a
comma-delimited list.`,
	}

	fields["other_sans"] = &framework.FieldSchema{
		Type: framework.TypeCommaStringSlice,
		Description: `Requested other SANs, in an array with the format
<oid>;UTF8:<utf8 string value> for each entry.`,
	}

	return fields
}

// addNonCACommonFields adds fields with help text specific to non-CA
// certificate issuing and signing
func addNonCACommonFields(fields map[string]*framework.FieldSchema) map[string]*framework.FieldSchema {
	fields = addIssueAndSignCommonFields(fields)

	fields["role"] = &framework.FieldSchema{
		Type: framework.TypeString,
		Description: `The desired role with configuration for this
request`,
	}

	fields["common_name"] = &framework.FieldSchema{
		Type: framework.TypeString,
		Description: `The requested common name; if you want more than
one, specify the alternative names in the
alt_names map. If email protection is enabled
in the role, this may be an email address.`,
	}

	fields["alt_names"] = &framework.FieldSchema{
		Type: framework.TypeString,
		Description: `The requested Subject Alternative Names, if any,
in a comma-delimited list. If email protection
is enabled for the role, this may contain
email addresses.`,
	}

	fields["serial_number"] = &framework.FieldSchema{
		Type: framework.TypeString,
		Description: `The requested serial number, if any. If you want
more than one, specify alternative names in
the alt_names map using OID 2.5.4.5.`,
	}

	fields["ttl"] = &framework.FieldSchema{
		Type: framework.TypeDurationSecond,
		Description: `The requested Time To Live for the certificate;
sets the expiration date. If not specified
the role default, backend default, or system
default TTL is used, in that order. Cannot
be larger than the role max TTL.`,
	}

	return fields
}
