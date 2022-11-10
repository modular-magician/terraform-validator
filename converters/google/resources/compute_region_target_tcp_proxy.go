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

const ComputeRegionTargetTcpProxyAssetType string = "compute.googleapis.com/RegionTargetTcpProxy"

func resourceConverterComputeRegionTargetTcpProxy() ResourceConverter {
	return ResourceConverter{
		AssetType: ComputeRegionTargetTcpProxyAssetType,
		Convert:   GetComputeRegionTargetTcpProxyCaiObject,
	}
}

func GetComputeRegionTargetTcpProxyCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/targetTcpProxies/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetComputeRegionTargetTcpProxyApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: ComputeRegionTargetTcpProxyAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "RegionTargetTcpProxy",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetComputeRegionTargetTcpProxyApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeRegionTargetTcpProxyDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeRegionTargetTcpProxyName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	proxyHeaderProp, err := expandComputeRegionTargetTcpProxyProxyHeader(d.Get("proxy_header"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("proxy_header"); !isEmptyValue(reflect.ValueOf(proxyHeaderProp)) && (ok || !reflect.DeepEqual(v, proxyHeaderProp)) {
		obj["proxyHeader"] = proxyHeaderProp
	}
	serviceProp, err := expandComputeRegionTargetTcpProxyBackendService(d.Get("backend_service"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("backend_service"); !isEmptyValue(reflect.ValueOf(serviceProp)) && (ok || !reflect.DeepEqual(v, serviceProp)) {
		obj["service"] = serviceProp
	}
	proxyBindProp, err := expandComputeRegionTargetTcpProxyProxyBind(d.Get("proxy_bind"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("proxy_bind"); !isEmptyValue(reflect.ValueOf(proxyBindProp)) && (ok || !reflect.DeepEqual(v, proxyBindProp)) {
		obj["proxyBind"] = proxyBindProp
	}
	regionProp, err := expandComputeRegionTargetTcpProxyRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	return obj, nil
}

func expandComputeRegionTargetTcpProxyDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionTargetTcpProxyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionTargetTcpProxyProxyHeader(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionTargetTcpProxyBackendService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("backendServices", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for backend_service: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRegionTargetTcpProxyProxyBind(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionTargetTcpProxyRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
