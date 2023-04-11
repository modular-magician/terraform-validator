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

const DNSPolicyAssetType string = "dns.googleapis.com/Policy"

func resourceConverterDNSPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType: DNSPolicyAssetType,
		Convert:   GetDNSPolicyCaiObject,
	}
}

func GetDNSPolicyCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//dns.googleapis.com/projects/{{project}}/policies/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetDNSPolicyApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: DNSPolicyAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/dns/v1/rest",
				DiscoveryName:        "Policy",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetDNSPolicyApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	alternativeNameServerConfigProp, err := expandDNSPolicyAlternativeNameServerConfig(d.Get("alternative_name_server_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("alternative_name_server_config"); !isEmptyValue(reflect.ValueOf(alternativeNameServerConfigProp)) && (ok || !reflect.DeepEqual(v, alternativeNameServerConfigProp)) {
		obj["alternativeNameServerConfig"] = alternativeNameServerConfigProp
	}
	descriptionProp, err := expandDNSPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	enableInboundForwardingProp, err := expandDNSPolicyEnableInboundForwarding(d.Get("enable_inbound_forwarding"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_inbound_forwarding"); ok || !reflect.DeepEqual(v, enableInboundForwardingProp) {
		obj["enableInboundForwarding"] = enableInboundForwardingProp
	}
	enableLoggingProp, err := expandDNSPolicyEnableLogging(d.Get("enable_logging"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_logging"); ok || !reflect.DeepEqual(v, enableLoggingProp) {
		obj["enableLogging"] = enableLoggingProp
	}
	nameProp, err := expandDNSPolicyName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networksProp, err := expandDNSPolicyNetworks(d.Get("networks"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("networks"); !isEmptyValue(reflect.ValueOf(networksProp)) && (ok || !reflect.DeepEqual(v, networksProp)) {
		obj["networks"] = networksProp
	}

	return obj, nil
}

func expandDNSPolicyAlternativeNameServerConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTargetNameServers, err := expandDNSPolicyAlternativeNameServerConfigTargetNameServers(original["target_name_servers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTargetNameServers); val.IsValid() && !isEmptyValue(val) {
		transformed["targetNameServers"] = transformedTargetNameServers
	}

	return transformed, nil
}

func expandDNSPolicyAlternativeNameServerConfigTargetNameServers(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedIpv4Address, err := expandDNSPolicyAlternativeNameServerConfigTargetNameServersIpv4Address(original["ipv4_address"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpv4Address); val.IsValid() && !isEmptyValue(val) {
			transformed["ipv4Address"] = transformedIpv4Address
		}

		transformedForwardingPath, err := expandDNSPolicyAlternativeNameServerConfigTargetNameServersForwardingPath(original["forwarding_path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedForwardingPath); val.IsValid() && !isEmptyValue(val) {
			transformed["forwardingPath"] = transformedForwardingPath
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDNSPolicyAlternativeNameServerConfigTargetNameServersIpv4Address(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSPolicyAlternativeNameServerConfigTargetNameServersForwardingPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSPolicyDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSPolicyEnableInboundForwarding(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSPolicyEnableLogging(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSPolicyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSPolicyNetworks(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNetworkUrl, err := expandDNSPolicyNetworksNetworkUrl(original["network_url"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNetworkUrl); val.IsValid() && !isEmptyValue(val) {
			transformed["networkUrl"] = transformedNetworkUrl
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDNSPolicyNetworksNetworkUrl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	if v == nil || v.(string) == "" {
		return "", nil
	} else if strings.HasPrefix(v.(string), "https://") {
		return v, nil
	}
	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}"+v.(string))
	if err != nil {
		return "", err
	}
	return ConvertSelfLinkToV1(url), nil
}
