// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import "reflect"

const KMSCryptoKeyVersionAssetType string = "cloudkms.googleapis.com/CryptoKeyVersion"

func resourceConverterKMSCryptoKeyVersion() ResourceConverter {
	return ResourceConverter{
		AssetType: KMSCryptoKeyVersionAssetType,
		Convert:   GetKMSCryptoKeyVersionCaiObject,
	}
}

func GetKMSCryptoKeyVersionCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//cloudkms.googleapis.com/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetKMSCryptoKeyVersionApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: KMSCryptoKeyVersionAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudkms/v1/rest",
				DiscoveryName:        "CryptoKeyVersion",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetKMSCryptoKeyVersionApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	stateProp, err := expandKMSCryptoKeyVersionState(d.Get("state"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("state"); !isEmptyValue(reflect.ValueOf(stateProp)) && (ok || !reflect.DeepEqual(v, stateProp)) {
		obj["state"] = stateProp
	}
	externalProtectionLevelOptionsProp, err := expandKMSCryptoKeyVersionExternalProtectionLevelOptions(d.Get("external_protection_level_options"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("external_protection_level_options"); !isEmptyValue(reflect.ValueOf(externalProtectionLevelOptionsProp)) && (ok || !reflect.DeepEqual(v, externalProtectionLevelOptionsProp)) {
		obj["externalProtectionLevelOptions"] = externalProtectionLevelOptionsProp
	}

	return obj, nil
}

func expandKMSCryptoKeyVersionState(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandKMSCryptoKeyVersionExternalProtectionLevelOptions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExternalKeyUri, err := expandKMSCryptoKeyVersionExternalProtectionLevelOptionsExternalKeyUri(original["external_key_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExternalKeyUri); val.IsValid() && !isEmptyValue(val) {
		transformed["externalKeyUri"] = transformedExternalKeyUri
	}

	transformedEkmConnectionKeyPath, err := expandKMSCryptoKeyVersionExternalProtectionLevelOptionsEkmConnectionKeyPath(original["ekm_connection_key_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEkmConnectionKeyPath); val.IsValid() && !isEmptyValue(val) {
		transformed["ekmConnectionKeyPath"] = transformedEkmConnectionKeyPath
	}

	return transformed, nil
}

func expandKMSCryptoKeyVersionExternalProtectionLevelOptionsExternalKeyUri(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandKMSCryptoKeyVersionExternalProtectionLevelOptionsEkmConnectionKeyPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
