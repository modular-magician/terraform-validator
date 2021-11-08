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
	"fmt"
	"reflect"
)

const NetworkManagementConnectivityTestAssetType string = "networkmanagement.googleapis.com/ConnectivityTest"

func resourceConverterNetworkManagementConnectivityTest() ResourceConverter {
	return ResourceConverter{
		AssetType: NetworkManagementConnectivityTestAssetType,
		Convert:   GetNetworkManagementConnectivityTestCaiObject,
	}
}

func GetNetworkManagementConnectivityTestCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//networkmanagement.googleapis.com/projects/{{project}}/locations/global/connectivityTests/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetNetworkManagementConnectivityTestApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: NetworkManagementConnectivityTestAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: assetName("https://www.googleapis.com/discovery/v1/apis/networkmanagement/v1/rest"),
				DiscoveryName:        "ConnectivityTest",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetNetworkManagementConnectivityTestApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandNetworkManagementConnectivityTestName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandNetworkManagementConnectivityTestDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	sourceProp, err := expandNetworkManagementConnectivityTestSource(d.Get("source"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source"); !isEmptyValue(reflect.ValueOf(sourceProp)) && (ok || !reflect.DeepEqual(v, sourceProp)) {
		obj["source"] = sourceProp
	}
	destinationProp, err := expandNetworkManagementConnectivityTestDestination(d.Get("destination"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("destination"); !isEmptyValue(reflect.ValueOf(destinationProp)) && (ok || !reflect.DeepEqual(v, destinationProp)) {
		obj["destination"] = destinationProp
	}
	protocolProp, err := expandNetworkManagementConnectivityTestProtocol(d.Get("protocol"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("protocol"); !isEmptyValue(reflect.ValueOf(protocolProp)) && (ok || !reflect.DeepEqual(v, protocolProp)) {
		obj["protocol"] = protocolProp
	}
	relatedProjectsProp, err := expandNetworkManagementConnectivityTestRelatedProjects(d.Get("related_projects"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("related_projects"); !isEmptyValue(reflect.ValueOf(relatedProjectsProp)) && (ok || !reflect.DeepEqual(v, relatedProjectsProp)) {
		obj["relatedProjects"] = relatedProjectsProp
	}
	labelsProp, err := expandNetworkManagementConnectivityTestLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandNetworkManagementConnectivityTestName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	// projects/X/tests/Y - note not "connectivityTests"
	f, err := parseGlobalFieldValue("tests", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for zone: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandNetworkManagementConnectivityTestDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestSource(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIpAddress, err := expandNetworkManagementConnectivityTestSourceIpAddress(original["ip_address"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIpAddress); val.IsValid() && !isEmptyValue(val) {
		transformed["ipAddress"] = transformedIpAddress
	}

	transformedPort, err := expandNetworkManagementConnectivityTestSourcePort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !isEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedInstance, err := expandNetworkManagementConnectivityTestSourceInstance(original["instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInstance); val.IsValid() && !isEmptyValue(val) {
		transformed["instance"] = transformedInstance
	}

	transformedNetwork, err := expandNetworkManagementConnectivityTestSourceNetwork(original["network"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetwork); val.IsValid() && !isEmptyValue(val) {
		transformed["network"] = transformedNetwork
	}

	transformedNetworkType, err := expandNetworkManagementConnectivityTestSourceNetworkType(original["network_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetworkType); val.IsValid() && !isEmptyValue(val) {
		transformed["networkType"] = transformedNetworkType
	}

	transformedProjectId, err := expandNetworkManagementConnectivityTestSourceProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !isEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	return transformed, nil
}

func expandNetworkManagementConnectivityTestSourceIpAddress(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestSourcePort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestSourceInstance(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestSourceNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestSourceNetworkType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestSourceProjectId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestDestination(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIpAddress, err := expandNetworkManagementConnectivityTestDestinationIpAddress(original["ip_address"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIpAddress); val.IsValid() && !isEmptyValue(val) {
		transformed["ipAddress"] = transformedIpAddress
	}

	transformedPort, err := expandNetworkManagementConnectivityTestDestinationPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !isEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedInstance, err := expandNetworkManagementConnectivityTestDestinationInstance(original["instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInstance); val.IsValid() && !isEmptyValue(val) {
		transformed["instance"] = transformedInstance
	}

	transformedNetwork, err := expandNetworkManagementConnectivityTestDestinationNetwork(original["network"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetwork); val.IsValid() && !isEmptyValue(val) {
		transformed["network"] = transformedNetwork
	}

	transformedProjectId, err := expandNetworkManagementConnectivityTestDestinationProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !isEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	return transformed, nil
}

func expandNetworkManagementConnectivityTestDestinationIpAddress(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestDestinationPort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestDestinationInstance(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestDestinationNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestDestinationProjectId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestProtocol(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestRelatedProjects(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
