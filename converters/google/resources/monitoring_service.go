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

const MonitoringServiceAssetType string = "monitoring.googleapis.com/Service"

func resourceConverterMonitoringService() ResourceConverter {
	return ResourceConverter{
		AssetType: MonitoringServiceAssetType,
		Convert:   GetMonitoringServiceCaiObject,
	}
}

func GetMonitoringServiceCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//monitoring.googleapis.com/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetMonitoringServiceApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: MonitoringServiceAssetType,
			Resource: &AssetResource{
				Version:              "v3",
				DiscoveryDocumentURI: assetName("https://www.googleapis.com/discovery/v1/apis/monitoring/v3/rest"),
				DiscoveryName:        "Service",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetMonitoringServiceApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandMonitoringServiceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	telemetryProp, err := expandMonitoringServiceTelemetry(d.Get("telemetry"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("telemetry"); !isEmptyValue(reflect.ValueOf(telemetryProp)) && (ok || !reflect.DeepEqual(v, telemetryProp)) {
		obj["telemetry"] = telemetryProp
	}
	nameProp, err := expandMonitoringServiceServiceId(d.Get("service_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("service_id"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	return resourceMonitoringServiceEncoder(d, config, obj)
}

func resourceMonitoringServiceEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Currently only CUSTOM service types can be created, but the
	// custom identifier block does not actually have fields right now.
	// Set to empty to indicate manually-created service type is CUSTOM.
	if _, ok := obj["custom"]; !ok {
		obj["custom"] = map[string]interface{}{}
	}
	// Name/Service ID is a query parameter only
	delete(obj, "name")

	return obj, nil
}

func expandMonitoringServiceDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringServiceTelemetry(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedResourceName, err := expandMonitoringServiceTelemetryResourceName(original["resource_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResourceName); val.IsValid() && !isEmptyValue(val) {
		transformed["resourceName"] = transformedResourceName
	}

	return transformed, nil
}

func expandMonitoringServiceTelemetryResourceName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringServiceServiceId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
