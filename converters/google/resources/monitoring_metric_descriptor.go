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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const MonitoringMetricDescriptorAssetType string = "monitoring.googleapis.com/MetricDescriptor"

func resourceConverterMonitoringMetricDescriptor() ResourceConverter {
	return ResourceConverter{
		AssetType: MonitoringMetricDescriptorAssetType,
		Convert:   GetMonitoringMetricDescriptorCaiObject,
	}
}

func GetMonitoringMetricDescriptorCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//monitoring.googleapis.com/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetMonitoringMetricDescriptorApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: MonitoringMetricDescriptorAssetType,
			Resource: &AssetResource{
				Version:              "v3",
				DiscoveryDocumentURI: assetName("https://www.googleapis.com/discovery/v1/apis/monitoring/v3/rest"),
				DiscoveryName:        "MetricDescriptor",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetMonitoringMetricDescriptorApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	typeProp, err := expandMonitoringMetricDescriptorType(d.Get("type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	labelsProp, err := expandMonitoringMetricDescriptorLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	metricKindProp, err := expandMonitoringMetricDescriptorMetricKind(d.Get("metric_kind"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("metric_kind"); !isEmptyValue(reflect.ValueOf(metricKindProp)) && (ok || !reflect.DeepEqual(v, metricKindProp)) {
		obj["metricKind"] = metricKindProp
	}
	valueTypeProp, err := expandMonitoringMetricDescriptorValueType(d.Get("value_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("value_type"); !isEmptyValue(reflect.ValueOf(valueTypeProp)) && (ok || !reflect.DeepEqual(v, valueTypeProp)) {
		obj["valueType"] = valueTypeProp
	}
	unitProp, err := expandMonitoringMetricDescriptorUnit(d.Get("unit"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("unit"); !isEmptyValue(reflect.ValueOf(unitProp)) && (ok || !reflect.DeepEqual(v, unitProp)) {
		obj["unit"] = unitProp
	}
	descriptionProp, err := expandMonitoringMetricDescriptorDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	displayNameProp, err := expandMonitoringMetricDescriptorDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	metadataProp, err := expandMonitoringMetricDescriptorMetadata(d.Get("metadata"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("metadata"); !isEmptyValue(reflect.ValueOf(metadataProp)) && (ok || !reflect.DeepEqual(v, metadataProp)) {
		obj["metadata"] = metadataProp
	}
	launchStageProp, err := expandMonitoringMetricDescriptorLaunchStage(d.Get("launch_stage"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("launch_stage"); !isEmptyValue(reflect.ValueOf(launchStageProp)) && (ok || !reflect.DeepEqual(v, launchStageProp)) {
		obj["launchStage"] = launchStageProp
	}

	return obj, nil
}

func expandMonitoringMetricDescriptorType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorLabels(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedKey, err := expandMonitoringMetricDescriptorLabelsKey(original["key"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKey); val.IsValid() && !isEmptyValue(val) {
			transformed["key"] = transformedKey
		}

		transformedValueType, err := expandMonitoringMetricDescriptorLabelsValueType(original["value_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValueType); val.IsValid() && !isEmptyValue(val) {
			transformed["valueType"] = transformedValueType
		}

		transformedDescription, err := expandMonitoringMetricDescriptorLabelsDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandMonitoringMetricDescriptorLabelsKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorLabelsValueType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorLabelsDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorMetricKind(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorValueType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorUnit(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorMetadata(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSamplePeriod, err := expandMonitoringMetricDescriptorMetadataSamplePeriod(original["sample_period"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSamplePeriod); val.IsValid() && !isEmptyValue(val) {
		transformed["samplePeriod"] = transformedSamplePeriod
	}

	transformedIngestDelay, err := expandMonitoringMetricDescriptorMetadataIngestDelay(original["ingest_delay"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIngestDelay); val.IsValid() && !isEmptyValue(val) {
		transformed["ingestDelay"] = transformedIngestDelay
	}

	return transformed, nil
}

func expandMonitoringMetricDescriptorMetadataSamplePeriod(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorMetadataIngestDelay(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorLaunchStage(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
