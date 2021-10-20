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

const AccessContextManagerAccessLevelAssetType string = "accesscontextmanager.googleapis.com/AccessLevel"

func resourceConverterAccessContextManagerAccessLevel() ResourceConverter {
	return ResourceConverter{
		AssetType: AccessContextManagerAccessLevelAssetType,
		Convert:   GetAccessContextManagerAccessLevelCaiObject,
	}
}

func GetAccessContextManagerAccessLevelCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//accesscontextmanager.googleapis.com/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetAccessContextManagerAccessLevelApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: AccessContextManagerAccessLevelAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/accesscontextmanager/v1/rest",
				DiscoveryName:        "AccessLevel",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetAccessContextManagerAccessLevelApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	titleProp, err := expandAccessContextManagerAccessLevelTitle(d.Get("title"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("title"); !isEmptyValue(reflect.ValueOf(titleProp)) && (ok || !reflect.DeepEqual(v, titleProp)) {
		obj["title"] = titleProp
	}
	descriptionProp, err := expandAccessContextManagerAccessLevelDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	basicProp, err := expandAccessContextManagerAccessLevelBasic(d.Get("basic"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("basic"); !isEmptyValue(reflect.ValueOf(basicProp)) && (ok || !reflect.DeepEqual(v, basicProp)) {
		obj["basic"] = basicProp
	}
	customProp, err := expandAccessContextManagerAccessLevelCustom(d.Get("custom"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("custom"); !isEmptyValue(reflect.ValueOf(customProp)) && (ok || !reflect.DeepEqual(v, customProp)) {
		obj["custom"] = customProp
	}
	parentProp, err := expandAccessContextManagerAccessLevelParent(d.Get("parent"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("parent"); !isEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}
	nameProp, err := expandAccessContextManagerAccessLevelName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	return resourceAccessContextManagerAccessLevelEncoder(d, config, obj)
}

func resourceAccessContextManagerAccessLevelEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	delete(obj, "parent")
	return obj, nil
}

func expandAccessContextManagerAccessLevelTitle(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasic(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCombiningFunction, err := expandAccessContextManagerAccessLevelBasicCombiningFunction(original["combining_function"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCombiningFunction); val.IsValid() && !isEmptyValue(val) {
		transformed["combiningFunction"] = transformedCombiningFunction
	}

	transformedConditions, err := expandAccessContextManagerAccessLevelBasicConditions(original["conditions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConditions); val.IsValid() && !isEmptyValue(val) {
		transformed["conditions"] = transformedConditions
	}

	return transformed, nil
}

func expandAccessContextManagerAccessLevelBasicCombiningFunction(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedIpSubnetworks, err := expandAccessContextManagerAccessLevelBasicConditionsIpSubnetworks(original["ip_subnetworks"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpSubnetworks); val.IsValid() && !isEmptyValue(val) {
			transformed["ipSubnetworks"] = transformedIpSubnetworks
		}

		transformedRequiredAccessLevels, err := expandAccessContextManagerAccessLevelBasicConditionsRequiredAccessLevels(original["required_access_levels"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRequiredAccessLevels); val.IsValid() && !isEmptyValue(val) {
			transformed["requiredAccessLevels"] = transformedRequiredAccessLevels
		}

		transformedMembers, err := expandAccessContextManagerAccessLevelBasicConditionsMembers(original["members"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMembers); val.IsValid() && !isEmptyValue(val) {
			transformed["members"] = transformedMembers
		}

		transformedNegate, err := expandAccessContextManagerAccessLevelBasicConditionsNegate(original["negate"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNegate); val.IsValid() && !isEmptyValue(val) {
			transformed["negate"] = transformedNegate
		}

		transformedDevicePolicy, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicy(original["device_policy"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDevicePolicy); val.IsValid() && !isEmptyValue(val) {
			transformed["devicePolicy"] = transformedDevicePolicy
		}

		transformedRegions, err := expandAccessContextManagerAccessLevelBasicConditionsRegions(original["regions"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRegions); val.IsValid() && !isEmptyValue(val) {
			transformed["regions"] = transformedRegions
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsIpSubnetworks(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsRequiredAccessLevels(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsMembers(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsNegate(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRequireScreenLock, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyRequireScreenLock(original["require_screen_lock"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequireScreenLock); val.IsValid() && !isEmptyValue(val) {
		transformed["requireScreenlock"] = transformedRequireScreenLock
	}

	transformedAllowedEncryptionStatuses, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatuses(original["allowed_encryption_statuses"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedEncryptionStatuses); val.IsValid() && !isEmptyValue(val) {
		transformed["allowedEncryptionStatuses"] = transformedAllowedEncryptionStatuses
	}

	transformedAllowedDeviceManagementLevels, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevels(original["allowed_device_management_levels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedDeviceManagementLevels); val.IsValid() && !isEmptyValue(val) {
		transformed["allowedDeviceManagementLevels"] = transformedAllowedDeviceManagementLevels
	}

	transformedOsConstraints, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraints(original["os_constraints"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOsConstraints); val.IsValid() && !isEmptyValue(val) {
		transformed["osConstraints"] = transformedOsConstraints
	}

	transformedRequireAdminApproval, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyRequireAdminApproval(original["require_admin_approval"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequireAdminApproval); val.IsValid() && !isEmptyValue(val) {
		transformed["requireAdminApproval"] = transformedRequireAdminApproval
	}

	transformedRequireCorpOwned, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyRequireCorpOwned(original["require_corp_owned"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequireCorpOwned); val.IsValid() && !isEmptyValue(val) {
		transformed["requireCorpOwned"] = transformedRequireCorpOwned
	}

	return transformed, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyRequireScreenLock(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatuses(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevels(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraints(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedMinimumVersion, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsMinimumVersion(original["minimum_version"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMinimumVersion); val.IsValid() && !isEmptyValue(val) {
			transformed["minimumVersion"] = transformedMinimumVersion
		}

		transformedRequireVerifiedChromeOs, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsRequireVerifiedChromeOs(original["require_verified_chrome_os"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRequireVerifiedChromeOs); val.IsValid() && !isEmptyValue(val) {
			transformed["requireVerifiedChromeOs"] = transformedRequireVerifiedChromeOs
		}

		transformedOsType, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOsType(original["os_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedOsType); val.IsValid() && !isEmptyValue(val) {
			transformed["osType"] = transformedOsType
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsMinimumVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsRequireVerifiedChromeOs(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOsType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyRequireAdminApproval(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyRequireCorpOwned(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsRegions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelCustom(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExpr, err := expandAccessContextManagerAccessLevelCustomExpr(original["expr"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpr); val.IsValid() && !isEmptyValue(val) {
		transformed["expr"] = transformedExpr
	}

	return transformed, nil
}

func expandAccessContextManagerAccessLevelCustomExpr(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExpression, err := expandAccessContextManagerAccessLevelCustomExprExpression(original["expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpression); val.IsValid() && !isEmptyValue(val) {
		transformed["expression"] = transformedExpression
	}

	transformedTitle, err := expandAccessContextManagerAccessLevelCustomExprTitle(original["title"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTitle); val.IsValid() && !isEmptyValue(val) {
		transformed["title"] = transformedTitle
	}

	transformedDescription, err := expandAccessContextManagerAccessLevelCustomExprDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedLocation, err := expandAccessContextManagerAccessLevelCustomExprLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !isEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandAccessContextManagerAccessLevelCustomExprExpression(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelCustomExprTitle(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelCustomExprDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelCustomExprLocation(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelParent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
