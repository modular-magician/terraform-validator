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

const ComputeBackendBucketAssetType string = "compute.googleapis.com/BackendBucket"

func resourceConverterComputeBackendBucket() ResourceConverter {
	return ResourceConverter{
		AssetType: ComputeBackendBucketAssetType,
		Convert:   GetComputeBackendBucketCaiObject,
	}
}

func GetComputeBackendBucketCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/backendBuckets/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetComputeBackendBucketApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: ComputeBackendBucketAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: assetName("https://www.googleapis.com/discovery/v1/apis/compute/v1/rest"),
				DiscoveryName:        "BackendBucket",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetComputeBackendBucketApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	bucketNameProp, err := expandComputeBackendBucketBucketName(d.Get("bucket_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("bucket_name"); !isEmptyValue(reflect.ValueOf(bucketNameProp)) && (ok || !reflect.DeepEqual(v, bucketNameProp)) {
		obj["bucketName"] = bucketNameProp
	}
	cdnPolicyProp, err := expandComputeBackendBucketCdnPolicy(d.Get("cdn_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cdn_policy"); !isEmptyValue(reflect.ValueOf(cdnPolicyProp)) && (ok || !reflect.DeepEqual(v, cdnPolicyProp)) {
		obj["cdnPolicy"] = cdnPolicyProp
	}
	customResponseHeadersProp, err := expandComputeBackendBucketCustomResponseHeaders(d.Get("custom_response_headers"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("custom_response_headers"); !isEmptyValue(reflect.ValueOf(customResponseHeadersProp)) && (ok || !reflect.DeepEqual(v, customResponseHeadersProp)) {
		obj["customResponseHeaders"] = customResponseHeadersProp
	}
	descriptionProp, err := expandComputeBackendBucketDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	enableCdnProp, err := expandComputeBackendBucketEnableCdn(d.Get("enable_cdn"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_cdn"); !isEmptyValue(reflect.ValueOf(enableCdnProp)) && (ok || !reflect.DeepEqual(v, enableCdnProp)) {
		obj["enableCdn"] = enableCdnProp
	}
	nameProp, err := expandComputeBackendBucketName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	return obj, nil
}

func expandComputeBackendBucketBucketName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSignedUrlCacheMaxAgeSec, err := expandComputeBackendBucketCdnPolicySignedUrlCacheMaxAgeSec(original["signed_url_cache_max_age_sec"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["signedUrlCacheMaxAgeSec"] = transformedSignedUrlCacheMaxAgeSec
	}

	transformedDefaultTtl, err := expandComputeBackendBucketCdnPolicyDefaultTtl(original["default_ttl"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["defaultTtl"] = transformedDefaultTtl
	}

	transformedMaxTtl, err := expandComputeBackendBucketCdnPolicyMaxTtl(original["max_ttl"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["maxTtl"] = transformedMaxTtl
	}

	transformedClientTtl, err := expandComputeBackendBucketCdnPolicyClientTtl(original["client_ttl"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["clientTtl"] = transformedClientTtl
	}

	transformedNegativeCaching, err := expandComputeBackendBucketCdnPolicyNegativeCaching(original["negative_caching"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["negativeCaching"] = transformedNegativeCaching
	}

	transformedNegativeCachingPolicy, err := expandComputeBackendBucketCdnPolicyNegativeCachingPolicy(original["negative_caching_policy"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNegativeCachingPolicy); val.IsValid() && !isEmptyValue(val) {
		transformed["negativeCachingPolicy"] = transformedNegativeCachingPolicy
	}

	transformedCacheMode, err := expandComputeBackendBucketCdnPolicyCacheMode(original["cache_mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCacheMode); val.IsValid() && !isEmptyValue(val) {
		transformed["cacheMode"] = transformedCacheMode
	}

	transformedServeWhileStale, err := expandComputeBackendBucketCdnPolicyServeWhileStale(original["serve_while_stale"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["serveWhileStale"] = transformedServeWhileStale
	}

	return transformed, nil
}

func expandComputeBackendBucketCdnPolicySignedUrlCacheMaxAgeSec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyDefaultTtl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyMaxTtl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyClientTtl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyNegativeCaching(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyNegativeCachingPolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedCode, err := expandComputeBackendBucketCdnPolicyNegativeCachingPolicyCode(original["code"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCode); val.IsValid() && !isEmptyValue(val) {
			transformed["code"] = transformedCode
		}

		transformedTtl, err := expandComputeBackendBucketCdnPolicyNegativeCachingPolicyTtl(original["ttl"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["ttl"] = transformedTtl
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeBackendBucketCdnPolicyNegativeCachingPolicyCode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyNegativeCachingPolicyTtl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyCacheMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyServeWhileStale(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCustomResponseHeaders(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketEnableCdn(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
