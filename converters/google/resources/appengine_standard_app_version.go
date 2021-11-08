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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const AppEngineStandardAppVersionAssetType string = "appengine.googleapis.com/StandardAppVersion"

func resourceConverterAppEngineStandardAppVersion() ResourceConverter {
	return ResourceConverter{
		AssetType: AppEngineStandardAppVersionAssetType,
		Convert:   GetAppEngineStandardAppVersionCaiObject,
	}
}

func GetAppEngineStandardAppVersionCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//appengine.googleapis.com/apps/{{project}}/services/{{service}}/versions/{{version_id}}?view=FULL")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetAppEngineStandardAppVersionApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: AppEngineStandardAppVersionAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: assetName("https://www.googleapis.com/discovery/v1/apis/appengine/v1/rest"),
				DiscoveryName:        "StandardAppVersion",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetAppEngineStandardAppVersionApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	idProp, err := expandAppEngineStandardAppVersionVersionId(d.Get("version_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("version_id"); !isEmptyValue(reflect.ValueOf(idProp)) && (ok || !reflect.DeepEqual(v, idProp)) {
		obj["id"] = idProp
	}
	runtimeProp, err := expandAppEngineStandardAppVersionRuntime(d.Get("runtime"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("runtime"); !isEmptyValue(reflect.ValueOf(runtimeProp)) && (ok || !reflect.DeepEqual(v, runtimeProp)) {
		obj["runtime"] = runtimeProp
	}
	threadsafeProp, err := expandAppEngineStandardAppVersionThreadsafe(d.Get("threadsafe"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("threadsafe"); !isEmptyValue(reflect.ValueOf(threadsafeProp)) && (ok || !reflect.DeepEqual(v, threadsafeProp)) {
		obj["threadsafe"] = threadsafeProp
	}
	runtimeApiVersionProp, err := expandAppEngineStandardAppVersionRuntimeApiVersion(d.Get("runtime_api_version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("runtime_api_version"); !isEmptyValue(reflect.ValueOf(runtimeApiVersionProp)) && (ok || !reflect.DeepEqual(v, runtimeApiVersionProp)) {
		obj["runtimeApiVersion"] = runtimeApiVersionProp
	}
	handlersProp, err := expandAppEngineStandardAppVersionHandlers(d.Get("handlers"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("handlers"); !isEmptyValue(reflect.ValueOf(handlersProp)) && (ok || !reflect.DeepEqual(v, handlersProp)) {
		obj["handlers"] = handlersProp
	}
	librariesProp, err := expandAppEngineStandardAppVersionLibraries(d.Get("libraries"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("libraries"); !isEmptyValue(reflect.ValueOf(librariesProp)) && (ok || !reflect.DeepEqual(v, librariesProp)) {
		obj["libraries"] = librariesProp
	}
	envVariablesProp, err := expandAppEngineStandardAppVersionEnvVariables(d.Get("env_variables"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("env_variables"); !isEmptyValue(reflect.ValueOf(envVariablesProp)) && (ok || !reflect.DeepEqual(v, envVariablesProp)) {
		obj["envVariables"] = envVariablesProp
	}
	deploymentProp, err := expandAppEngineStandardAppVersionDeployment(d.Get("deployment"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("deployment"); !isEmptyValue(reflect.ValueOf(deploymentProp)) && (ok || !reflect.DeepEqual(v, deploymentProp)) {
		obj["deployment"] = deploymentProp
	}
	entrypointProp, err := expandAppEngineStandardAppVersionEntrypoint(d.Get("entrypoint"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("entrypoint"); !isEmptyValue(reflect.ValueOf(entrypointProp)) && (ok || !reflect.DeepEqual(v, entrypointProp)) {
		obj["entrypoint"] = entrypointProp
	}
	vpcAccessConnectorProp, err := expandAppEngineStandardAppVersionVPCAccessConnector(d.Get("vpc_access_connector"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("vpc_access_connector"); !isEmptyValue(reflect.ValueOf(vpcAccessConnectorProp)) && (ok || !reflect.DeepEqual(v, vpcAccessConnectorProp)) {
		obj["vpcAccessConnector"] = vpcAccessConnectorProp
	}
	inboundServicesProp, err := expandAppEngineStandardAppVersionInboundServices(d.Get("inbound_services"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("inbound_services"); !isEmptyValue(reflect.ValueOf(inboundServicesProp)) && (ok || !reflect.DeepEqual(v, inboundServicesProp)) {
		obj["inboundServices"] = inboundServicesProp
	}
	instanceClassProp, err := expandAppEngineStandardAppVersionInstanceClass(d.Get("instance_class"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("instance_class"); !isEmptyValue(reflect.ValueOf(instanceClassProp)) && (ok || !reflect.DeepEqual(v, instanceClassProp)) {
		obj["instanceClass"] = instanceClassProp
	}
	automaticScalingProp, err := expandAppEngineStandardAppVersionAutomaticScaling(d.Get("automatic_scaling"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("automatic_scaling"); !isEmptyValue(reflect.ValueOf(automaticScalingProp)) && (ok || !reflect.DeepEqual(v, automaticScalingProp)) {
		obj["automaticScaling"] = automaticScalingProp
	}
	basicScalingProp, err := expandAppEngineStandardAppVersionBasicScaling(d.Get("basic_scaling"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("basic_scaling"); !isEmptyValue(reflect.ValueOf(basicScalingProp)) && (ok || !reflect.DeepEqual(v, basicScalingProp)) {
		obj["basicScaling"] = basicScalingProp
	}
	manualScalingProp, err := expandAppEngineStandardAppVersionManualScaling(d.Get("manual_scaling"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("manual_scaling"); !isEmptyValue(reflect.ValueOf(manualScalingProp)) && (ok || !reflect.DeepEqual(v, manualScalingProp)) {
		obj["manualScaling"] = manualScalingProp
	}

	return obj, nil
}

func expandAppEngineStandardAppVersionVersionId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionRuntime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionThreadsafe(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionRuntimeApiVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlers(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedUrlRegex, err := expandAppEngineStandardAppVersionHandlersUrlRegex(original["url_regex"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedUrlRegex); val.IsValid() && !isEmptyValue(val) {
			transformed["urlRegex"] = transformedUrlRegex
		}

		transformedSecurityLevel, err := expandAppEngineStandardAppVersionHandlersSecurityLevel(original["security_level"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSecurityLevel); val.IsValid() && !isEmptyValue(val) {
			transformed["securityLevel"] = transformedSecurityLevel
		}

		transformedLogin, err := expandAppEngineStandardAppVersionHandlersLogin(original["login"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedLogin); val.IsValid() && !isEmptyValue(val) {
			transformed["login"] = transformedLogin
		}

		transformedAuthFailAction, err := expandAppEngineStandardAppVersionHandlersAuthFailAction(original["auth_fail_action"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAuthFailAction); val.IsValid() && !isEmptyValue(val) {
			transformed["authFailAction"] = transformedAuthFailAction
		}

		transformedRedirectHttpResponseCode, err := expandAppEngineStandardAppVersionHandlersRedirectHttpResponseCode(original["redirect_http_response_code"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRedirectHttpResponseCode); val.IsValid() && !isEmptyValue(val) {
			transformed["redirectHttpResponseCode"] = transformedRedirectHttpResponseCode
		}

		transformedScript, err := expandAppEngineStandardAppVersionHandlersScript(original["script"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedScript); val.IsValid() && !isEmptyValue(val) {
			transformed["script"] = transformedScript
		}

		transformedStaticFiles, err := expandAppEngineStandardAppVersionHandlersStaticFiles(original["static_files"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedStaticFiles); val.IsValid() && !isEmptyValue(val) {
			transformed["staticFiles"] = transformedStaticFiles
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAppEngineStandardAppVersionHandlersUrlRegex(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersSecurityLevel(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersLogin(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersAuthFailAction(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersRedirectHttpResponseCode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersScript(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedScriptPath, err := expandAppEngineStandardAppVersionHandlersScriptScriptPath(original["script_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScriptPath); val.IsValid() && !isEmptyValue(val) {
		transformed["scriptPath"] = transformedScriptPath
	}

	return transformed, nil
}

func expandAppEngineStandardAppVersionHandlersScriptScriptPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersStaticFiles(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPath, err := expandAppEngineStandardAppVersionHandlersStaticFilesPath(original["path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !isEmptyValue(val) {
		transformed["path"] = transformedPath
	}

	transformedUploadPathRegex, err := expandAppEngineStandardAppVersionHandlersStaticFilesUploadPathRegex(original["upload_path_regex"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUploadPathRegex); val.IsValid() && !isEmptyValue(val) {
		transformed["uploadPathRegex"] = transformedUploadPathRegex
	}

	transformedHttpHeaders, err := expandAppEngineStandardAppVersionHandlersStaticFilesHttpHeaders(original["http_headers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHttpHeaders); val.IsValid() && !isEmptyValue(val) {
		transformed["httpHeaders"] = transformedHttpHeaders
	}

	transformedMimeType, err := expandAppEngineStandardAppVersionHandlersStaticFilesMimeType(original["mime_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMimeType); val.IsValid() && !isEmptyValue(val) {
		transformed["mimeType"] = transformedMimeType
	}

	transformedExpiration, err := expandAppEngineStandardAppVersionHandlersStaticFilesExpiration(original["expiration"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpiration); val.IsValid() && !isEmptyValue(val) {
		transformed["expiration"] = transformedExpiration
	}

	transformedRequireMatchingFile, err := expandAppEngineStandardAppVersionHandlersStaticFilesRequireMatchingFile(original["require_matching_file"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequireMatchingFile); val.IsValid() && !isEmptyValue(val) {
		transformed["requireMatchingFile"] = transformedRequireMatchingFile
	}

	transformedApplicationReadable, err := expandAppEngineStandardAppVersionHandlersStaticFilesApplicationReadable(original["application_readable"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedApplicationReadable); val.IsValid() && !isEmptyValue(val) {
		transformed["applicationReadable"] = transformedApplicationReadable
	}

	return transformed, nil
}

func expandAppEngineStandardAppVersionHandlersStaticFilesPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersStaticFilesUploadPathRegex(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersStaticFilesHttpHeaders(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandAppEngineStandardAppVersionHandlersStaticFilesMimeType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersStaticFilesExpiration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersStaticFilesRequireMatchingFile(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionHandlersStaticFilesApplicationReadable(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionLibraries(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandAppEngineStandardAppVersionLibrariesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedVersion, err := expandAppEngineStandardAppVersionLibrariesVersion(original["version"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !isEmptyValue(val) {
			transformed["version"] = transformedVersion
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAppEngineStandardAppVersionLibrariesName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionLibrariesVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionEnvVariables(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandAppEngineStandardAppVersionDeployment(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedZip, err := expandAppEngineStandardAppVersionDeploymentZip(original["zip"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedZip); val.IsValid() && !isEmptyValue(val) {
		transformed["zip"] = transformedZip
	}

	transformedFiles, err := expandAppEngineStandardAppVersionDeploymentFiles(original["files"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFiles); val.IsValid() && !isEmptyValue(val) {
		transformed["files"] = transformedFiles
	}

	return transformed, nil
}

func expandAppEngineStandardAppVersionDeploymentZip(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSourceUrl, err := expandAppEngineStandardAppVersionDeploymentZipSourceUrl(original["source_url"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSourceUrl); val.IsValid() && !isEmptyValue(val) {
		transformed["sourceUrl"] = transformedSourceUrl
	}

	transformedFilesCount, err := expandAppEngineStandardAppVersionDeploymentZipFilesCount(original["files_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFilesCount); val.IsValid() && !isEmptyValue(val) {
		transformed["filesCount"] = transformedFilesCount
	}

	return transformed, nil
}

func expandAppEngineStandardAppVersionDeploymentZipSourceUrl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionDeploymentZipFilesCount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionDeploymentFiles(v interface{}, d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedSha1Sum, err := expandAppEngineStandardAppVersionDeploymentFilesSha1Sum(original["sha1_sum"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSha1Sum); val.IsValid() && !isEmptyValue(val) {
			transformed["sha1Sum"] = transformedSha1Sum
		}

		transformedSourceUrl, err := expandAppEngineStandardAppVersionDeploymentFilesSourceUrl(original["source_url"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSourceUrl); val.IsValid() && !isEmptyValue(val) {
			transformed["sourceUrl"] = transformedSourceUrl
		}

		transformedName, err := expandString(original["name"], d, config)
		if err != nil {
			return nil, err
		}
		m[transformedName] = transformed
	}
	return m, nil
}

func expandAppEngineStandardAppVersionDeploymentFilesSha1Sum(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionDeploymentFilesSourceUrl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionEntrypoint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedShell, err := expandAppEngineStandardAppVersionEntrypointShell(original["shell"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedShell); val.IsValid() && !isEmptyValue(val) {
		transformed["shell"] = transformedShell
	}

	return transformed, nil
}

func expandAppEngineStandardAppVersionEntrypointShell(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionVPCAccessConnector(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedName, err := expandAppEngineStandardAppVersionVPCAccessConnectorName(original["name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
		transformed["name"] = transformedName
	}

	return transformed, nil
}

func expandAppEngineStandardAppVersionVPCAccessConnectorName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionInboundServices(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandAppEngineStandardAppVersionInstanceClass(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionAutomaticScaling(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMaxConcurrentRequests, err := expandAppEngineStandardAppVersionAutomaticScalingMaxConcurrentRequests(original["max_concurrent_requests"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxConcurrentRequests); val.IsValid() && !isEmptyValue(val) {
		transformed["maxConcurrentRequests"] = transformedMaxConcurrentRequests
	}

	transformedMaxIdleInstances, err := expandAppEngineStandardAppVersionAutomaticScalingMaxIdleInstances(original["max_idle_instances"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxIdleInstances); val.IsValid() && !isEmptyValue(val) {
		transformed["maxIdleInstances"] = transformedMaxIdleInstances
	}

	transformedMaxPendingLatency, err := expandAppEngineStandardAppVersionAutomaticScalingMaxPendingLatency(original["max_pending_latency"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxPendingLatency); val.IsValid() && !isEmptyValue(val) {
		transformed["maxPendingLatency"] = transformedMaxPendingLatency
	}

	transformedMinIdleInstances, err := expandAppEngineStandardAppVersionAutomaticScalingMinIdleInstances(original["min_idle_instances"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinIdleInstances); val.IsValid() && !isEmptyValue(val) {
		transformed["minIdleInstances"] = transformedMinIdleInstances
	}

	transformedMinPendingLatency, err := expandAppEngineStandardAppVersionAutomaticScalingMinPendingLatency(original["min_pending_latency"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinPendingLatency); val.IsValid() && !isEmptyValue(val) {
		transformed["minPendingLatency"] = transformedMinPendingLatency
	}

	transformedStandardSchedulerSettings, err := expandAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettings(original["standard_scheduler_settings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStandardSchedulerSettings); val.IsValid() && !isEmptyValue(val) {
		transformed["standardSchedulerSettings"] = transformedStandardSchedulerSettings
	}

	return transformed, nil
}

func expandAppEngineStandardAppVersionAutomaticScalingMaxConcurrentRequests(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionAutomaticScalingMaxIdleInstances(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionAutomaticScalingMaxPendingLatency(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionAutomaticScalingMinIdleInstances(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionAutomaticScalingMinPendingLatency(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettings(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTargetCpuUtilization, err := expandAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsTargetCpuUtilization(original["target_cpu_utilization"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTargetCpuUtilization); val.IsValid() && !isEmptyValue(val) {
		transformed["targetCpuUtilization"] = transformedTargetCpuUtilization
	}

	transformedTargetThroughputUtilization, err := expandAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsTargetThroughputUtilization(original["target_throughput_utilization"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTargetThroughputUtilization); val.IsValid() && !isEmptyValue(val) {
		transformed["targetThroughputUtilization"] = transformedTargetThroughputUtilization
	}

	transformedMinInstances, err := expandAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsMinInstances(original["min_instances"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinInstances); val.IsValid() && !isEmptyValue(val) {
		transformed["minInstances"] = transformedMinInstances
	}

	transformedMaxInstances, err := expandAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsMaxInstances(original["max_instances"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxInstances); val.IsValid() && !isEmptyValue(val) {
		transformed["maxInstances"] = transformedMaxInstances
	}

	return transformed, nil
}

func expandAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsTargetCpuUtilization(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsTargetThroughputUtilization(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsMinInstances(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsMaxInstances(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionBasicScaling(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIdleTimeout, err := expandAppEngineStandardAppVersionBasicScalingIdleTimeout(original["idle_timeout"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIdleTimeout); val.IsValid() && !isEmptyValue(val) {
		transformed["idleTimeout"] = transformedIdleTimeout
	}

	transformedMaxInstances, err := expandAppEngineStandardAppVersionBasicScalingMaxInstances(original["max_instances"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxInstances); val.IsValid() && !isEmptyValue(val) {
		transformed["maxInstances"] = transformedMaxInstances
	}

	return transformed, nil
}

func expandAppEngineStandardAppVersionBasicScalingIdleTimeout(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionBasicScalingMaxInstances(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineStandardAppVersionManualScaling(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedInstances, err := expandAppEngineStandardAppVersionManualScalingInstances(original["instances"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInstances); val.IsValid() && !isEmptyValue(val) {
		transformed["instances"] = transformedInstances
	}

	return transformed, nil
}

func expandAppEngineStandardAppVersionManualScalingInstances(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
