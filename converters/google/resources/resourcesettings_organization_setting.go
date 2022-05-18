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

const ResourceSettingsOrganizationSettingAssetType string = "resourcesettings.googleapis.com/OrganizationSetting"

func resourceConverterResourceSettingsOrganizationSetting() ResourceConverter {
	return ResourceConverter{
		AssetType: ResourceSettingsOrganizationSettingAssetType,
		Convert:   GetResourceSettingsOrganizationSettingCaiObject,
	}
}

func GetResourceSettingsOrganizationSettingCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//resourcesettings.googleapis.com/organizations/{{organization_id}}/settings/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetResourceSettingsOrganizationSettingApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: ResourceSettingsOrganizationSettingAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/resourcesettings/v1/rest",
				DiscoveryName:        "OrganizationSetting",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetResourceSettingsOrganizationSettingApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	localValueProp, err := expandResourceSettingsOrganizationSettingLocalValue(d.Get("local_value"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("local_value"); !isEmptyValue(reflect.ValueOf(localValueProp)) && (ok || !reflect.DeepEqual(v, localValueProp)) {
		obj["localValue"] = localValueProp
	}

	return obj, nil
}

func expandResourceSettingsOrganizationSettingLocalValue(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedStringValue, err := expandResourceSettingsOrganizationSettingLocalValueStringValue(original["string_value"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStringValue); val.IsValid() && !isEmptyValue(val) {
		transformed["stringValue"] = transformedStringValue
	}

	return transformed, nil
}

func expandResourceSettingsOrganizationSettingLocalValueStringValue(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
