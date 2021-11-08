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

const ComputeNodeGroupAssetType string = "compute.googleapis.com/NodeGroup"

func resourceConverterComputeNodeGroup() ResourceConverter {
	return ResourceConverter{
		AssetType: ComputeNodeGroupAssetType,
		Convert:   GetComputeNodeGroupCaiObject,
	}
}

func GetComputeNodeGroupCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/zones/{{zone}}/nodeGroups/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetComputeNodeGroupApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: ComputeNodeGroupAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: assetName("https://www.googleapis.com/discovery/v1/apis/compute/v1/rest"),
				DiscoveryName:        "NodeGroup",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetComputeNodeGroupApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeNodeGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeNodeGroupName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	nodeTemplateProp, err := expandComputeNodeGroupNodeTemplate(d.Get("node_template"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("node_template"); !isEmptyValue(reflect.ValueOf(nodeTemplateProp)) && (ok || !reflect.DeepEqual(v, nodeTemplateProp)) {
		obj["nodeTemplate"] = nodeTemplateProp
	}
	sizeProp, err := expandComputeNodeGroupSize(d.Get("size"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("size"); ok || !reflect.DeepEqual(v, sizeProp) {
		obj["size"] = sizeProp
	}
	maintenancePolicyProp, err := expandComputeNodeGroupMaintenancePolicy(d.Get("maintenance_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("maintenance_policy"); !isEmptyValue(reflect.ValueOf(maintenancePolicyProp)) && (ok || !reflect.DeepEqual(v, maintenancePolicyProp)) {
		obj["maintenancePolicy"] = maintenancePolicyProp
	}
	maintenanceWindowProp, err := expandComputeNodeGroupMaintenanceWindow(d.Get("maintenance_window"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("maintenance_window"); !isEmptyValue(reflect.ValueOf(maintenanceWindowProp)) && (ok || !reflect.DeepEqual(v, maintenanceWindowProp)) {
		obj["maintenanceWindow"] = maintenanceWindowProp
	}
	autoscalingPolicyProp, err := expandComputeNodeGroupAutoscalingPolicy(d.Get("autoscaling_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("autoscaling_policy"); !isEmptyValue(reflect.ValueOf(autoscalingPolicyProp)) && (ok || !reflect.DeepEqual(v, autoscalingPolicyProp)) {
		obj["autoscalingPolicy"] = autoscalingPolicyProp
	}
	zoneProp, err := expandComputeNodeGroupZone(d.Get("zone"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("zone"); !isEmptyValue(reflect.ValueOf(zoneProp)) && (ok || !reflect.DeepEqual(v, zoneProp)) {
		obj["zone"] = zoneProp
	}

	return obj, nil
}

func expandComputeNodeGroupDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeGroupName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeGroupNodeTemplate(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("nodeTemplates", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for node_template: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeNodeGroupSize(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeGroupMaintenancePolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeGroupMaintenanceWindow(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedStartTime, err := expandComputeNodeGroupMaintenanceWindowStartTime(original["start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !isEmptyValue(val) {
		transformed["startTime"] = transformedStartTime
	}

	return transformed, nil
}

func expandComputeNodeGroupMaintenanceWindowStartTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeGroupAutoscalingPolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMode, err := expandComputeNodeGroupAutoscalingPolicyMode(original["mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMode); val.IsValid() && !isEmptyValue(val) {
		transformed["mode"] = transformedMode
	}

	transformedMinNodes, err := expandComputeNodeGroupAutoscalingPolicyMinNodes(original["min_nodes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinNodes); val.IsValid() && !isEmptyValue(val) {
		transformed["minNodes"] = transformedMinNodes
	}

	transformedMaxNodes, err := expandComputeNodeGroupAutoscalingPolicyMaxNodes(original["max_nodes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxNodes); val.IsValid() && !isEmptyValue(val) {
		transformed["maxNodes"] = transformedMaxNodes
	}

	return transformed, nil
}

func expandComputeNodeGroupAutoscalingPolicyMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeGroupAutoscalingPolicyMinNodes(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeGroupAutoscalingPolicyMaxNodes(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeGroupZone(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("zones", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for zone: %s", err)
	}
	return f.RelativeLink(), nil
}
