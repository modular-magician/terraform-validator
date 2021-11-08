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

const DialogflowCXEnvironmentAssetType string = "dialogflow.googleapis.com/Environment"

func resourceConverterDialogflowCXEnvironment() ResourceConverter {
	return ResourceConverter{
		AssetType: DialogflowCXEnvironmentAssetType,
		Convert:   GetDialogflowCXEnvironmentCaiObject,
	}
}

func GetDialogflowCXEnvironmentCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//dialogflow.googleapis.com/{{parent}}/environments/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetDialogflowCXEnvironmentApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: DialogflowCXEnvironmentAssetType,
			Resource: &AssetResource{
				Version:              "v3",
				DiscoveryDocumentURI: assetName("https://www.googleapis.com/discovery/v1/apis/dialogflow/v3/rest"),
				DiscoveryName:        "Environment",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetDialogflowCXEnvironmentApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowCXEnvironmentDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandDialogflowCXEnvironmentDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	versionConfigsProp, err := expandDialogflowCXEnvironmentVersionConfigs(d.Get("version_configs"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("version_configs"); !isEmptyValue(reflect.ValueOf(versionConfigsProp)) && (ok || !reflect.DeepEqual(v, versionConfigsProp)) {
		obj["versionConfigs"] = versionConfigsProp
	}

	return obj, nil
}

func expandDialogflowCXEnvironmentDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXEnvironmentDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXEnvironmentVersionConfigs(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedVersion, err := expandDialogflowCXEnvironmentVersionConfigsVersion(original["version"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !isEmptyValue(val) {
			transformed["version"] = transformedVersion
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowCXEnvironmentVersionConfigsVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
