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

const VertexAIIndexEndpointAssetType string = "{{region}}-aiplatform.googleapis.com/IndexEndpoint"

func resourceConverterVertexAIIndexEndpoint() ResourceConverter {
	return ResourceConverter{
		AssetType: VertexAIIndexEndpointAssetType,
		Convert:   GetVertexAIIndexEndpointCaiObject,
	}
}

func GetVertexAIIndexEndpointCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//{{region}}-aiplatform.googleapis.com/projects/{{project}}/locations/{{region}}/indexEndpoints/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetVertexAIIndexEndpointApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: VertexAIIndexEndpointAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/{{region}}-aiplatform/v1/rest",
				DiscoveryName:        "IndexEndpoint",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetVertexAIIndexEndpointApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandVertexAIIndexEndpointDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandVertexAIIndexEndpointDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandVertexAIIndexEndpointLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	networkProp, err := expandVertexAIIndexEndpointNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}

	return obj, nil
}

func expandVertexAIIndexEndpointDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexEndpointDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexEndpointLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandVertexAIIndexEndpointNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
