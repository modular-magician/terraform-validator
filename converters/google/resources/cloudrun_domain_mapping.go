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

import (
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var domainMappingGoogleProvidedLabels = []string{
	"cloud.googleapis.com/location",
	"run.googleapis.com/overrideAt",
}

func domainMappingLabelDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	// Suppress diffs for the labels provided by Google
	for _, label := range domainMappingGoogleProvidedLabels {
		if strings.Contains(k, label) && new == "" {
			return true
		}
	}

	// Let diff be determined by labels (above)
	if strings.Contains(k, "labels.%") {
		return true
	}

	// For other keys, don't suppress diff.
	return false
}

const CloudRunDomainMappingAssetType string = "{{location}}-run.googleapis.com/DomainMapping"

func resourceConverterCloudRunDomainMapping() ResourceConverter {
	return ResourceConverter{
		AssetType: CloudRunDomainMappingAssetType,
		Convert:   GetCloudRunDomainMappingCaiObject,
	}
}

func GetCloudRunDomainMappingCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//{{location}}-run.googleapis.com/apis/domains.cloudrun.com/namespaces/{{project}}/domainmappings/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetCloudRunDomainMappingApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: CloudRunDomainMappingAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: assetName("https://www.googleapis.com/discovery/v1/apis/{{location}}-run/v1/rest"),
				DiscoveryName:        "DomainMapping",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetCloudRunDomainMappingApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	specProp, err := expandCloudRunDomainMappingSpec(d.Get("spec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("spec"); !isEmptyValue(reflect.ValueOf(specProp)) && (ok || !reflect.DeepEqual(v, specProp)) {
		obj["spec"] = specProp
	}
	metadataProp, err := expandCloudRunDomainMappingMetadata(d.Get("metadata"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("metadata"); !isEmptyValue(reflect.ValueOf(metadataProp)) && (ok || !reflect.DeepEqual(v, metadataProp)) {
		obj["metadata"] = metadataProp
	}

	return resourceCloudRunDomainMappingEncoder(d, config, obj)
}

func resourceCloudRunDomainMappingEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	name := d.Get("name").(string)
	metadata := obj["metadata"].(map[string]interface{})
	metadata["name"] = name

	// The only acceptable version/kind right now
	obj["apiVersion"] = "domains.cloudrun.com/v1"
	obj["kind"] = "DomainMapping"
	return obj, nil
}

func expandCloudRunDomainMappingSpec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedForceOverride, err := expandCloudRunDomainMappingSpecForceOverride(original["force_override"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedForceOverride); val.IsValid() && !isEmptyValue(val) {
		transformed["forceOverride"] = transformedForceOverride
	}

	transformedRouteName, err := expandCloudRunDomainMappingSpecRouteName(original["route_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRouteName); val.IsValid() && !isEmptyValue(val) {
		transformed["routeName"] = transformedRouteName
	}

	transformedCertificateMode, err := expandCloudRunDomainMappingSpecCertificateMode(original["certificate_mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCertificateMode); val.IsValid() && !isEmptyValue(val) {
		transformed["certificateMode"] = transformedCertificateMode
	}

	return transformed, nil
}

func expandCloudRunDomainMappingSpecForceOverride(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunDomainMappingSpecRouteName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return GetResourceNameFromSelfLink(v.(string)), nil
}

func expandCloudRunDomainMappingSpecCertificateMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunDomainMappingMetadata(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLabels, err := expandCloudRunDomainMappingMetadataLabels(original["labels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !isEmptyValue(val) {
		transformed["labels"] = transformedLabels
	}

	transformedGeneration, err := expandCloudRunDomainMappingMetadataGeneration(original["generation"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGeneration); val.IsValid() && !isEmptyValue(val) {
		transformed["generation"] = transformedGeneration
	}

	transformedResourceVersion, err := expandCloudRunDomainMappingMetadataResourceVersion(original["resource_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResourceVersion); val.IsValid() && !isEmptyValue(val) {
		transformed["resourceVersion"] = transformedResourceVersion
	}

	transformedSelfLink, err := expandCloudRunDomainMappingMetadataSelfLink(original["self_link"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSelfLink); val.IsValid() && !isEmptyValue(val) {
		transformed["selfLink"] = transformedSelfLink
	}

	transformedUid, err := expandCloudRunDomainMappingMetadataUid(original["uid"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUid); val.IsValid() && !isEmptyValue(val) {
		transformed["uid"] = transformedUid
	}

	transformedNamespace, err := expandCloudRunDomainMappingMetadataNamespace(original["namespace"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNamespace); val.IsValid() && !isEmptyValue(val) {
		transformed["namespace"] = transformedNamespace
	}

	transformedAnnotations, err := expandCloudRunDomainMappingMetadataAnnotations(original["annotations"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAnnotations); val.IsValid() && !isEmptyValue(val) {
		transformed["annotations"] = transformedAnnotations
	}

	return transformed, nil
}

func expandCloudRunDomainMappingMetadataLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudRunDomainMappingMetadataGeneration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunDomainMappingMetadataResourceVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunDomainMappingMetadataSelfLink(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunDomainMappingMetadataUid(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunDomainMappingMetadataNamespace(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunDomainMappingMetadataAnnotations(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
