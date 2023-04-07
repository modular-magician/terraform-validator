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
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const ComputeForwardingRuleAssetType string = "compute.googleapis.com/ForwardingRule"

func resourceConverterComputeForwardingRule() ResourceConverter {
	return ResourceConverter{
		AssetType: ComputeForwardingRuleAssetType,
		Convert:   GetComputeForwardingRuleCaiObject,
	}
}

func GetComputeForwardingRuleCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/forwardingRules/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetComputeForwardingRuleApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: ComputeForwardingRuleAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "ForwardingRule",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetComputeForwardingRuleApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	isMirroringCollectorProp, err := expandComputeForwardingRuleIsMirroringCollector(d.Get("is_mirroring_collector"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("is_mirroring_collector"); !isEmptyValue(reflect.ValueOf(isMirroringCollectorProp)) && (ok || !reflect.DeepEqual(v, isMirroringCollectorProp)) {
		obj["isMirroringCollector"] = isMirroringCollectorProp
	}
	descriptionProp, err := expandComputeForwardingRuleDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	IPAddressProp, err := expandComputeForwardingRuleIPAddress(d.Get("ip_address"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ip_address"); !isEmptyValue(reflect.ValueOf(IPAddressProp)) && (ok || !reflect.DeepEqual(v, IPAddressProp)) {
		obj["IPAddress"] = IPAddressProp
	}
	IPProtocolProp, err := expandComputeForwardingRuleIPProtocol(d.Get("ip_protocol"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ip_protocol"); !isEmptyValue(reflect.ValueOf(IPProtocolProp)) && (ok || !reflect.DeepEqual(v, IPProtocolProp)) {
		obj["IPProtocol"] = IPProtocolProp
	}
	backendServiceProp, err := expandComputeForwardingRuleBackendService(d.Get("backend_service"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("backend_service"); !isEmptyValue(reflect.ValueOf(backendServiceProp)) && (ok || !reflect.DeepEqual(v, backendServiceProp)) {
		obj["backendService"] = backendServiceProp
	}
	loadBalancingSchemeProp, err := expandComputeForwardingRuleLoadBalancingScheme(d.Get("load_balancing_scheme"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("load_balancing_scheme"); !isEmptyValue(reflect.ValueOf(loadBalancingSchemeProp)) && (ok || !reflect.DeepEqual(v, loadBalancingSchemeProp)) {
		obj["loadBalancingScheme"] = loadBalancingSchemeProp
	}
	nameProp, err := expandComputeForwardingRuleName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networkProp, err := expandComputeForwardingRuleNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	portRangeProp, err := expandComputeForwardingRulePortRange(d.Get("port_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("port_range"); !isEmptyValue(reflect.ValueOf(portRangeProp)) && (ok || !reflect.DeepEqual(v, portRangeProp)) {
		obj["portRange"] = portRangeProp
	}
	portsProp, err := expandComputeForwardingRulePorts(d.Get("ports"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ports"); !isEmptyValue(reflect.ValueOf(portsProp)) && (ok || !reflect.DeepEqual(v, portsProp)) {
		obj["ports"] = portsProp
	}
	subnetworkProp, err := expandComputeForwardingRuleSubnetwork(d.Get("subnetwork"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("subnetwork"); !isEmptyValue(reflect.ValueOf(subnetworkProp)) && (ok || !reflect.DeepEqual(v, subnetworkProp)) {
		obj["subnetwork"] = subnetworkProp
	}
	targetProp, err := expandComputeForwardingRuleTarget(d.Get("target"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("target"); !isEmptyValue(reflect.ValueOf(targetProp)) && (ok || !reflect.DeepEqual(v, targetProp)) {
		obj["target"] = targetProp
	}
	allowGlobalAccessProp, err := expandComputeForwardingRuleAllowGlobalAccess(d.Get("allow_global_access"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("allow_global_access"); ok || !reflect.DeepEqual(v, allowGlobalAccessProp) {
		obj["allowGlobalAccess"] = allowGlobalAccessProp
	}
	allPortsProp, err := expandComputeForwardingRuleAllPorts(d.Get("all_ports"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("all_ports"); !isEmptyValue(reflect.ValueOf(allPortsProp)) && (ok || !reflect.DeepEqual(v, allPortsProp)) {
		obj["allPorts"] = allPortsProp
	}
	networkTierProp, err := expandComputeForwardingRuleNetworkTier(d.Get("network_tier"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network_tier"); !isEmptyValue(reflect.ValueOf(networkTierProp)) && (ok || !reflect.DeepEqual(v, networkTierProp)) {
		obj["networkTier"] = networkTierProp
	}
	serviceDirectoryRegistrationsProp, err := expandComputeForwardingRuleServiceDirectoryRegistrations(d.Get("service_directory_registrations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("service_directory_registrations"); !isEmptyValue(reflect.ValueOf(serviceDirectoryRegistrationsProp)) && (ok || !reflect.DeepEqual(v, serviceDirectoryRegistrationsProp)) {
		obj["serviceDirectoryRegistrations"] = serviceDirectoryRegistrationsProp
	}
	serviceLabelProp, err := expandComputeForwardingRuleServiceLabel(d.Get("service_label"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("service_label"); !isEmptyValue(reflect.ValueOf(serviceLabelProp)) && (ok || !reflect.DeepEqual(v, serviceLabelProp)) {
		obj["serviceLabel"] = serviceLabelProp
	}
	regionProp, err := expandComputeForwardingRuleRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	return obj, nil
}

func expandComputeForwardingRuleIsMirroringCollector(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleIPAddress(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleIPProtocol(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleBackendService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	// This method returns a full self link from a partial self link.
	if v == nil || v.(string) == "" {
		// It does not try to construct anything from empty.
		return "", nil
	} else if strings.HasPrefix(v.(string), "https://") {
		// Anything that starts with a URL scheme is assumed to be a self link worth using.
		return v, nil
	} else if strings.HasPrefix(v.(string), "projects/") {
		// If the self link references a project, we'll just stuck the compute prefix on it
		url, err := ReplaceVars(d, config, "{{ComputeBasePath}}"+v.(string))
		if err != nil {
			return "", err
		}
		return url, nil
	} else if strings.HasPrefix(v.(string), "regions/") || strings.HasPrefix(v.(string), "zones/") {
		// For regional or zonal resources which include their region or zone, just put the project in front.
		url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/")
		if err != nil {
			return nil, err
		}
		return url + v.(string), nil
	}
	// Anything else is assumed to be a regional resource, with a partial link that begins with the resource name.
	// This isn't very likely - it's a last-ditch effort to extract something useful here.  We can do a better job
	// as soon as MultiResourceRefs are working since we'll know the types that this field is supposed to point to.
	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/")
	if err != nil {
		return nil, err
	}
	return url + v.(string), nil
}

func expandComputeForwardingRuleLoadBalancingScheme(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeForwardingRulePortRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRulePorts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v.(*schema.Set).List(), nil
}

func expandComputeForwardingRuleSubnetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("subnetworks", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for subnetwork: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeForwardingRuleTarget(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	// This method returns a full self link from a partial self link.
	if v == nil || v.(string) == "" {
		// It does not try to construct anything from empty.
		return "", nil
	} else if strings.HasPrefix(v.(string), "https://") {
		// Anything that starts with a URL scheme is assumed to be a self link worth using.
		return v, nil
	} else if strings.HasPrefix(v.(string), "projects/") {
		// If the self link references a project, we'll just stuck the compute prefix on it
		url, err := ReplaceVars(d, config, "{{ComputeBasePath}}"+v.(string))
		if err != nil {
			return "", err
		}
		return url, nil
	} else if strings.HasPrefix(v.(string), "regions/") || strings.HasPrefix(v.(string), "zones/") {
		// For regional or zonal resources which include their region or zone, just put the project in front.
		url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/")
		if err != nil {
			return nil, err
		}
		return url + v.(string), nil
	}
	// Anything else is assumed to be a regional resource, with a partial link that begins with the resource name.
	// This isn't very likely - it's a last-ditch effort to extract something useful here.  We can do a better job
	// as soon as MultiResourceRefs are working since we'll know the types that this field is supposed to point to.
	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/")
	if err != nil {
		return nil, err
	}
	return url + v.(string), nil
}

func expandComputeForwardingRuleAllowGlobalAccess(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleAllPorts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleNetworkTier(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleServiceDirectoryRegistrations(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNamespace, err := expandComputeForwardingRuleServiceDirectoryRegistrationsNamespace(original["namespace"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNamespace); val.IsValid() && !isEmptyValue(val) {
			transformed["namespace"] = transformedNamespace
		}

		transformedService, err := expandComputeForwardingRuleServiceDirectoryRegistrationsService(original["service"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedService); val.IsValid() && !isEmptyValue(val) {
			transformed["service"] = transformedService
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeForwardingRuleServiceDirectoryRegistrationsNamespace(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleServiceDirectoryRegistrationsService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleServiceLabel(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
