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

const IAM2AccessBoundaryPolicyAssetType string = "iam.googleapis.com/AccessBoundaryPolicy"

func resourceConverterIAM2AccessBoundaryPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType: IAM2AccessBoundaryPolicyAssetType,
		Convert:   GetIAM2AccessBoundaryPolicyCaiObject,
	}
}

func GetIAM2AccessBoundaryPolicyCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//iam.googleapis.com/policies/{{parent}}/accessboundarypolicies/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetIAM2AccessBoundaryPolicyApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: IAM2AccessBoundaryPolicyAssetType,
			Resource: &AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/iam/v2/rest",
				DiscoveryName:        "AccessBoundaryPolicy",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetIAM2AccessBoundaryPolicyApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandIAM2AccessBoundaryPolicyDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	etagProp, err := expandIAM2AccessBoundaryPolicyEtag(d.Get("etag"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("etag"); !isEmptyValue(reflect.ValueOf(etagProp)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}
	rulesProp, err := expandIAM2AccessBoundaryPolicyRules(d.Get("rules"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("rules"); !isEmptyValue(reflect.ValueOf(rulesProp)) && (ok || !reflect.DeepEqual(v, rulesProp)) {
		obj["rules"] = rulesProp
	}

	return obj, nil
}

func expandIAM2AccessBoundaryPolicyDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIAM2AccessBoundaryPolicyEtag(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIAM2AccessBoundaryPolicyRules(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDescription, err := expandIAM2AccessBoundaryPolicyRulesDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedAccessBoundaryRule, err := expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRule(original["access_boundary_rule"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAccessBoundaryRule); val.IsValid() && !isEmptyValue(val) {
			transformed["accessBoundaryRule"] = transformedAccessBoundaryRule
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandIAM2AccessBoundaryPolicyRulesDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAvailableResource, err := expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailableResource(original["available_resource"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAvailableResource); val.IsValid() && !isEmptyValue(val) {
		transformed["availableResource"] = transformedAvailableResource
	}

	transformedAvailablePermissions, err := expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailablePermissions(original["available_permissions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAvailablePermissions); val.IsValid() && !isEmptyValue(val) {
		transformed["availablePermissions"] = transformedAvailablePermissions
	}

	transformedAvailabilityCondition, err := expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityCondition(original["availability_condition"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAvailabilityCondition); val.IsValid() && !isEmptyValue(val) {
		transformed["availabilityCondition"] = transformedAvailabilityCondition
	}

	return transformed, nil
}

func expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailableResource(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailablePermissions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityCondition(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExpression, err := expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionExpression(original["expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpression); val.IsValid() && !isEmptyValue(val) {
		transformed["expression"] = transformedExpression
	}

	transformedTitle, err := expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionTitle(original["title"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTitle); val.IsValid() && !isEmptyValue(val) {
		transformed["title"] = transformedTitle
	}

	transformedDescription, err := expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedLocation, err := expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !isEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionExpression(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionTitle(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionLocation(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
