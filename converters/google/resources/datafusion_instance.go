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

var instanceAcceleratorOptions = []string{
	"delta.default.checkpoint.directory",
	"ui.feature.cdc",
}

func instanceOptionsDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	// Suppress diffs for the options generated by adding an accelerator to a data fusion instance
	for _, option := range instanceAcceleratorOptions {
		if strings.Contains(k, option) && new == "" {
			return true
		}
	}

	// Let diff be determined by options (above)
	if strings.Contains(k, "options.%") {
		return true
	}

	// For other keys, don't suppress diff.
	return false
}

const DataFusionInstanceAssetType string = "datafusion.googleapis.com/Instance"

func resourceConverterDataFusionInstance() ResourceConverter {
	return ResourceConverter{
		AssetType: DataFusionInstanceAssetType,
		Convert:   GetDataFusionInstanceCaiObject,
	}
}

func GetDataFusionInstanceCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//datafusion.googleapis.com/projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetDataFusionInstanceApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: DataFusionInstanceAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/datafusion/v1/rest",
				DiscoveryName:        "Instance",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetDataFusionInstanceApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandDataFusionInstanceName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandDataFusionInstanceDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	typeProp, err := expandDataFusionInstanceType(d.Get("type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	enableStackdriverLoggingProp, err := expandDataFusionInstanceEnableStackdriverLogging(d.Get("enable_stackdriver_logging"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_stackdriver_logging"); !isEmptyValue(reflect.ValueOf(enableStackdriverLoggingProp)) && (ok || !reflect.DeepEqual(v, enableStackdriverLoggingProp)) {
		obj["enableStackdriverLogging"] = enableStackdriverLoggingProp
	}
	enableStackdriverMonitoringProp, err := expandDataFusionInstanceEnableStackdriverMonitoring(d.Get("enable_stackdriver_monitoring"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_stackdriver_monitoring"); !isEmptyValue(reflect.ValueOf(enableStackdriverMonitoringProp)) && (ok || !reflect.DeepEqual(v, enableStackdriverMonitoringProp)) {
		obj["enableStackdriverMonitoring"] = enableStackdriverMonitoringProp
	}
	enableRbacProp, err := expandDataFusionInstanceEnableRbac(d.Get("enable_rbac"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_rbac"); !isEmptyValue(reflect.ValueOf(enableRbacProp)) && (ok || !reflect.DeepEqual(v, enableRbacProp)) {
		obj["enableRbac"] = enableRbacProp
	}
	labelsProp, err := expandDataFusionInstanceLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	optionsProp, err := expandDataFusionInstanceOptions(d.Get("options"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("options"); !isEmptyValue(reflect.ValueOf(optionsProp)) && (ok || !reflect.DeepEqual(v, optionsProp)) {
		obj["options"] = optionsProp
	}
	versionProp, err := expandDataFusionInstanceVersion(d.Get("version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("version"); !isEmptyValue(reflect.ValueOf(versionProp)) && (ok || !reflect.DeepEqual(v, versionProp)) {
		obj["version"] = versionProp
	}
	privateInstanceProp, err := expandDataFusionInstancePrivateInstance(d.Get("private_instance"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("private_instance"); !isEmptyValue(reflect.ValueOf(privateInstanceProp)) && (ok || !reflect.DeepEqual(v, privateInstanceProp)) {
		obj["privateInstance"] = privateInstanceProp
	}
	dataprocServiceAccountProp, err := expandDataFusionInstanceDataprocServiceAccount(d.Get("dataproc_service_account"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("dataproc_service_account"); !isEmptyValue(reflect.ValueOf(dataprocServiceAccountProp)) && (ok || !reflect.DeepEqual(v, dataprocServiceAccountProp)) {
		obj["dataprocServiceAccount"] = dataprocServiceAccountProp
	}
	networkConfigProp, err := expandDataFusionInstanceNetworkConfig(d.Get("network_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network_config"); !isEmptyValue(reflect.ValueOf(networkConfigProp)) && (ok || !reflect.DeepEqual(v, networkConfigProp)) {
		obj["networkConfig"] = networkConfigProp
	}
	zoneProp, err := expandDataFusionInstanceZone(d.Get("zone"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("zone"); !isEmptyValue(reflect.ValueOf(zoneProp)) && (ok || !reflect.DeepEqual(v, zoneProp)) {
		obj["zone"] = zoneProp
	}
	displayNameProp, err := expandDataFusionInstanceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	cryptoKeyConfigProp, err := expandDataFusionInstanceCryptoKeyConfig(d.Get("crypto_key_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("crypto_key_config"); !isEmptyValue(reflect.ValueOf(cryptoKeyConfigProp)) && (ok || !reflect.DeepEqual(v, cryptoKeyConfigProp)) {
		obj["cryptoKeyConfig"] = cryptoKeyConfigProp
	}
	eventPublishConfigProp, err := expandDataFusionInstanceEventPublishConfig(d.Get("event_publish_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("event_publish_config"); !isEmptyValue(reflect.ValueOf(eventPublishConfigProp)) && (ok || !reflect.DeepEqual(v, eventPublishConfigProp)) {
		obj["eventPublishConfig"] = eventPublishConfigProp
	}
	acceleratorsProp, err := expandDataFusionInstanceAccelerators(d.Get("accelerators"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("accelerators"); !isEmptyValue(reflect.ValueOf(acceleratorsProp)) && (ok || !reflect.DeepEqual(v, acceleratorsProp)) {
		obj["accelerators"] = acceleratorsProp
	}

	return obj, nil
}

func expandDataFusionInstanceName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return replaceVars(d, config, "projects/{{project}}/locations/{{region}}/instances/{{name}}")
}

func expandDataFusionInstanceDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceEnableStackdriverLogging(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceEnableStackdriverMonitoring(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceEnableRbac(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDataFusionInstanceOptions(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDataFusionInstanceVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstancePrivateInstance(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceDataprocServiceAccount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceNetworkConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIpAllocation, err := expandDataFusionInstanceNetworkConfigIpAllocation(original["ip_allocation"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIpAllocation); val.IsValid() && !isEmptyValue(val) {
		transformed["ipAllocation"] = transformedIpAllocation
	}

	transformedNetwork, err := expandDataFusionInstanceNetworkConfigNetwork(original["network"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetwork); val.IsValid() && !isEmptyValue(val) {
		transformed["network"] = transformedNetwork
	}

	return transformed, nil
}

func expandDataFusionInstanceNetworkConfigIpAllocation(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceNetworkConfigNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceZone(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceCryptoKeyConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKeyReference, err := expandDataFusionInstanceCryptoKeyConfigKeyReference(original["key_reference"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKeyReference); val.IsValid() && !isEmptyValue(val) {
		transformed["keyReference"] = transformedKeyReference
	}

	return transformed, nil
}

func expandDataFusionInstanceCryptoKeyConfigKeyReference(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceEventPublishConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandDataFusionInstanceEventPublishConfigEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !isEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	transformedTopic, err := expandDataFusionInstanceEventPublishConfigTopic(original["topic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTopic); val.IsValid() && !isEmptyValue(val) {
		transformed["topic"] = transformedTopic
	}

	return transformed, nil
}

func expandDataFusionInstanceEventPublishConfigEnabled(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceEventPublishConfigTopic(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceAccelerators(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedAcceleratorType, err := expandDataFusionInstanceAcceleratorsAcceleratorType(original["accelerator_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAcceleratorType); val.IsValid() && !isEmptyValue(val) {
			transformed["acceleratorType"] = transformedAcceleratorType
		}

		transformedState, err := expandDataFusionInstanceAcceleratorsState(original["state"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedState); val.IsValid() && !isEmptyValue(val) {
			transformed["state"] = transformedState
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDataFusionInstanceAcceleratorsAcceleratorType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceAcceleratorsState(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
