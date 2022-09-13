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

const Cloudfunctions2functionAssetType string = "cloudfunctions.googleapis.com/function"

func resourceConverterCloudfunctions2function() ResourceConverter {
	return ResourceConverter{
		AssetType: Cloudfunctions2functionAssetType,
		Convert:   GetCloudfunctions2functionCaiObject,
	}
}

func GetCloudfunctions2functionCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//cloudfunctions.googleapis.com/projects/{{project}}/locations/{{location}}/functions/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetCloudfunctions2functionApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: Cloudfunctions2functionAssetType,
			Resource: &AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudfunctions/v2/rest",
				DiscoveryName:        "function",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetCloudfunctions2functionApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandCloudfunctions2functionName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandCloudfunctions2functionDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	buildConfigProp, err := expandCloudfunctions2functionBuildConfig(d.Get("build_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("build_config"); !isEmptyValue(reflect.ValueOf(buildConfigProp)) && (ok || !reflect.DeepEqual(v, buildConfigProp)) {
		obj["buildConfig"] = buildConfigProp
	}
	serviceConfigProp, err := expandCloudfunctions2functionServiceConfig(d.Get("service_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("service_config"); !isEmptyValue(reflect.ValueOf(serviceConfigProp)) && (ok || !reflect.DeepEqual(v, serviceConfigProp)) {
		obj["serviceConfig"] = serviceConfigProp
	}
	eventTriggerProp, err := expandCloudfunctions2functionEventTrigger(d.Get("event_trigger"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("event_trigger"); !isEmptyValue(reflect.ValueOf(eventTriggerProp)) && (ok || !reflect.DeepEqual(v, eventTriggerProp)) {
		obj["eventTrigger"] = eventTriggerProp
	}
	labelsProp, err := expandCloudfunctions2functionLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandCloudfunctions2functionName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return replaceVars(d, config, "projects/{{project}}/locations/{{location}}/functions/{{name}}")
}

func expandCloudfunctions2functionDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBuild, err := expandCloudfunctions2functionBuildConfigBuild(original["build"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBuild); val.IsValid() && !isEmptyValue(val) {
		transformed["build"] = transformedBuild
	}

	transformedRuntime, err := expandCloudfunctions2functionBuildConfigRuntime(original["runtime"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRuntime); val.IsValid() && !isEmptyValue(val) {
		transformed["runtime"] = transformedRuntime
	}

	transformedEntryPoint, err := expandCloudfunctions2functionBuildConfigEntryPoint(original["entry_point"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEntryPoint); val.IsValid() && !isEmptyValue(val) {
		transformed["entryPoint"] = transformedEntryPoint
	}

	transformedSource, err := expandCloudfunctions2functionBuildConfigSource(original["source"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSource); val.IsValid() && !isEmptyValue(val) {
		transformed["source"] = transformedSource
	}

	transformedWorkerPool, err := expandCloudfunctions2functionBuildConfigWorkerPool(original["worker_pool"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWorkerPool); val.IsValid() && !isEmptyValue(val) {
		transformed["workerPool"] = transformedWorkerPool
	}

	transformedEnvironmentVariables, err := expandCloudfunctions2functionBuildConfigEnvironmentVariables(original["environment_variables"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnvironmentVariables); val.IsValid() && !isEmptyValue(val) {
		transformed["environmentVariables"] = transformedEnvironmentVariables
	}

	transformedDockerRepository, err := expandCloudfunctions2functionBuildConfigDockerRepository(original["docker_repository"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDockerRepository); val.IsValid() && !isEmptyValue(val) {
		transformed["dockerRepository"] = transformedDockerRepository
	}

	return transformed, nil
}

func expandCloudfunctions2functionBuildConfigBuild(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigRuntime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigEntryPoint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigSource(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedStorageSource, err := expandCloudfunctions2functionBuildConfigSourceStorageSource(original["storage_source"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStorageSource); val.IsValid() && !isEmptyValue(val) {
		transformed["storageSource"] = transformedStorageSource
	}

	transformedRepoSource, err := expandCloudfunctions2functionBuildConfigSourceRepoSource(original["repo_source"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRepoSource); val.IsValid() && !isEmptyValue(val) {
		transformed["repoSource"] = transformedRepoSource
	}

	return transformed, nil
}

func expandCloudfunctions2functionBuildConfigSourceStorageSource(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBucket, err := expandCloudfunctions2functionBuildConfigSourceStorageSourceBucket(original["bucket"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBucket); val.IsValid() && !isEmptyValue(val) {
		transformed["bucket"] = transformedBucket
	}

	transformedObject, err := expandCloudfunctions2functionBuildConfigSourceStorageSourceObject(original["object"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedObject); val.IsValid() && !isEmptyValue(val) {
		transformed["object"] = transformedObject
	}

	transformedGeneration, err := expandCloudfunctions2functionBuildConfigSourceStorageSourceGeneration(original["generation"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGeneration); val.IsValid() && !isEmptyValue(val) {
		transformed["generation"] = transformedGeneration
	}

	return transformed, nil
}

func expandCloudfunctions2functionBuildConfigSourceStorageSourceBucket(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigSourceStorageSourceObject(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigSourceStorageSourceGeneration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigSourceRepoSource(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProjectId, err := expandCloudfunctions2functionBuildConfigSourceRepoSourceProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !isEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	transformedRepoName, err := expandCloudfunctions2functionBuildConfigSourceRepoSourceRepoName(original["repo_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRepoName); val.IsValid() && !isEmptyValue(val) {
		transformed["repoName"] = transformedRepoName
	}

	transformedBranchName, err := expandCloudfunctions2functionBuildConfigSourceRepoSourceBranchName(original["branch_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBranchName); val.IsValid() && !isEmptyValue(val) {
		transformed["branchName"] = transformedBranchName
	}

	transformedTagName, err := expandCloudfunctions2functionBuildConfigSourceRepoSourceTagName(original["tag_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTagName); val.IsValid() && !isEmptyValue(val) {
		transformed["tagName"] = transformedTagName
	}

	transformedCommitSha, err := expandCloudfunctions2functionBuildConfigSourceRepoSourceCommitSha(original["commit_sha"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCommitSha); val.IsValid() && !isEmptyValue(val) {
		transformed["commitSha"] = transformedCommitSha
	}

	transformedDir, err := expandCloudfunctions2functionBuildConfigSourceRepoSourceDir(original["dir"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDir); val.IsValid() && !isEmptyValue(val) {
		transformed["dir"] = transformedDir
	}

	transformedInvertRegex, err := expandCloudfunctions2functionBuildConfigSourceRepoSourceInvertRegex(original["invert_regex"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInvertRegex); val.IsValid() && !isEmptyValue(val) {
		transformed["invertRegex"] = transformedInvertRegex
	}

	return transformed, nil
}

func expandCloudfunctions2functionBuildConfigSourceRepoSourceProjectId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigSourceRepoSourceRepoName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigSourceRepoSourceBranchName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigSourceRepoSourceTagName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigSourceRepoSourceCommitSha(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigSourceRepoSourceDir(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigSourceRepoSourceInvertRegex(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigWorkerPool(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigEnvironmentVariables(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudfunctions2functionBuildConfigDockerRepository(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedService, err := expandCloudfunctions2functionServiceConfigService(original["service"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedService); val.IsValid() && !isEmptyValue(val) {
		transformed["service"] = transformedService
	}

	transformedTimeoutSeconds, err := expandCloudfunctions2functionServiceConfigTimeoutSeconds(original["timeout_seconds"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTimeoutSeconds); val.IsValid() && !isEmptyValue(val) {
		transformed["timeoutSeconds"] = transformedTimeoutSeconds
	}

	transformedAvailableMemory, err := expandCloudfunctions2functionServiceConfigAvailableMemory(original["available_memory"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAvailableMemory); val.IsValid() && !isEmptyValue(val) {
		transformed["availableMemory"] = transformedAvailableMemory
	}

	transformedEnvironmentVariables, err := expandCloudfunctions2functionServiceConfigEnvironmentVariables(original["environment_variables"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnvironmentVariables); val.IsValid() && !isEmptyValue(val) {
		transformed["environmentVariables"] = transformedEnvironmentVariables
	}

	transformedMaxInstanceCount, err := expandCloudfunctions2functionServiceConfigMaxInstanceCount(original["max_instance_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxInstanceCount); val.IsValid() && !isEmptyValue(val) {
		transformed["maxInstanceCount"] = transformedMaxInstanceCount
	}

	transformedMinInstanceCount, err := expandCloudfunctions2functionServiceConfigMinInstanceCount(original["min_instance_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinInstanceCount); val.IsValid() && !isEmptyValue(val) {
		transformed["minInstanceCount"] = transformedMinInstanceCount
	}

	transformedConcurrency, err := expandCloudfunctions2functionServiceConfigConcurrency(original["concurrency"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConcurrency); val.IsValid() && !isEmptyValue(val) {
		transformed["concurrency"] = transformedConcurrency
	}

	transformedVPCConnector, err := expandCloudfunctions2functionServiceConfigVPCConnector(original["vpc_connector"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVPCConnector); val.IsValid() && !isEmptyValue(val) {
		transformed["vpcConnector"] = transformedVPCConnector
	}

	transformedVPCConnectorEgressSettings, err := expandCloudfunctions2functionServiceConfigVPCConnectorEgressSettings(original["vpc_connector_egress_settings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVPCConnectorEgressSettings); val.IsValid() && !isEmptyValue(val) {
		transformed["vpcConnectorEgressSettings"] = transformedVPCConnectorEgressSettings
	}

	transformedIngressSettings, err := expandCloudfunctions2functionServiceConfigIngressSettings(original["ingress_settings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIngressSettings); val.IsValid() && !isEmptyValue(val) {
		transformed["ingressSettings"] = transformedIngressSettings
	}

	transformedUri, err := expandCloudfunctions2functionServiceConfigUri(original["uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUri); val.IsValid() && !isEmptyValue(val) {
		transformed["uri"] = transformedUri
	}

	transformedGcfUri, err := expandCloudfunctions2functionServiceConfigGcfUri(original["gcf_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGcfUri); val.IsValid() && !isEmptyValue(val) {
		transformed["gcfUri"] = transformedGcfUri
	}

	transformedServiceAccountEmail, err := expandCloudfunctions2functionServiceConfigServiceAccountEmail(original["service_account_email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAccountEmail); val.IsValid() && !isEmptyValue(val) {
		transformed["serviceAccountEmail"] = transformedServiceAccountEmail
	}

	transformedAllTrafficOnLatestRevision, err := expandCloudfunctions2functionServiceConfigAllTrafficOnLatestRevision(original["all_traffic_on_latest_revision"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllTrafficOnLatestRevision); val.IsValid() && !isEmptyValue(val) {
		transformed["allTrafficOnLatestRevision"] = transformedAllTrafficOnLatestRevision
	}

	transformedSecretEnvironmentVariables, err := expandCloudfunctions2functionServiceConfigSecretEnvironmentVariables(original["secret_environment_variables"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSecretEnvironmentVariables); val.IsValid() && !isEmptyValue(val) {
		transformed["secretEnvironmentVariables"] = transformedSecretEnvironmentVariables
	}

	transformedSecretVolumes, err := expandCloudfunctions2functionServiceConfigSecretVolumes(original["secret_volumes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSecretVolumes); val.IsValid() && !isEmptyValue(val) {
		transformed["secretVolumes"] = transformedSecretVolumes
	}

	return transformed, nil
}

func expandCloudfunctions2functionServiceConfigService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigTimeoutSeconds(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigAvailableMemory(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigEnvironmentVariables(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudfunctions2functionServiceConfigMaxInstanceCount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigMinInstanceCount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigConcurrency(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigVPCConnector(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigVPCConnectorEgressSettings(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigIngressSettings(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigUri(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigGcfUri(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigServiceAccountEmail(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigAllTrafficOnLatestRevision(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigSecretEnvironmentVariables(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedKey, err := expandCloudfunctions2functionServiceConfigSecretEnvironmentVariablesKey(original["key"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKey); val.IsValid() && !isEmptyValue(val) {
			transformed["key"] = transformedKey
		}

		transformedProjectId, err := expandCloudfunctions2functionServiceConfigSecretEnvironmentVariablesProjectId(original["project_id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !isEmptyValue(val) {
			transformed["projectId"] = transformedProjectId
		}

		transformedSecret, err := expandCloudfunctions2functionServiceConfigSecretEnvironmentVariablesSecret(original["secret"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSecret); val.IsValid() && !isEmptyValue(val) {
			transformed["secret"] = transformedSecret
		}

		transformedVersion, err := expandCloudfunctions2functionServiceConfigSecretEnvironmentVariablesVersion(original["version"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !isEmptyValue(val) {
			transformed["version"] = transformedVersion
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudfunctions2functionServiceConfigSecretEnvironmentVariablesKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigSecretEnvironmentVariablesProjectId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigSecretEnvironmentVariablesSecret(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigSecretEnvironmentVariablesVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigSecretVolumes(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedMountPath, err := expandCloudfunctions2functionServiceConfigSecretVolumesMountPath(original["mount_path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMountPath); val.IsValid() && !isEmptyValue(val) {
			transformed["mountPath"] = transformedMountPath
		}

		transformedProjectId, err := expandCloudfunctions2functionServiceConfigSecretVolumesProjectId(original["project_id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !isEmptyValue(val) {
			transformed["projectId"] = transformedProjectId
		}

		transformedSecret, err := expandCloudfunctions2functionServiceConfigSecretVolumesSecret(original["secret"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSecret); val.IsValid() && !isEmptyValue(val) {
			transformed["secret"] = transformedSecret
		}

		transformedVersions, err := expandCloudfunctions2functionServiceConfigSecretVolumesVersions(original["versions"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedVersions); val.IsValid() && !isEmptyValue(val) {
			transformed["versions"] = transformedVersions
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudfunctions2functionServiceConfigSecretVolumesMountPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigSecretVolumesProjectId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigSecretVolumesSecret(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigSecretVolumesVersions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedVersion, err := expandCloudfunctions2functionServiceConfigSecretVolumesVersionsVersion(original["version"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !isEmptyValue(val) {
			transformed["version"] = transformedVersion
		}

		transformedPath, err := expandCloudfunctions2functionServiceConfigSecretVolumesVersionsPath(original["path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !isEmptyValue(val) {
			transformed["path"] = transformedPath
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudfunctions2functionServiceConfigSecretVolumesVersionsVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigSecretVolumesVersionsPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionEventTrigger(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTrigger, err := expandCloudfunctions2functionEventTriggerTrigger(original["trigger"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTrigger); val.IsValid() && !isEmptyValue(val) {
		transformed["trigger"] = transformedTrigger
	}

	transformedTriggerRegion, err := expandCloudfunctions2functionEventTriggerTriggerRegion(original["trigger_region"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTriggerRegion); val.IsValid() && !isEmptyValue(val) {
		transformed["triggerRegion"] = transformedTriggerRegion
	}

	transformedEventType, err := expandCloudfunctions2functionEventTriggerEventType(original["event_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEventType); val.IsValid() && !isEmptyValue(val) {
		transformed["eventType"] = transformedEventType
	}

	transformedEventFilters, err := expandCloudfunctions2functionEventTriggerEventFilters(original["event_filters"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEventFilters); val.IsValid() && !isEmptyValue(val) {
		transformed["eventFilters"] = transformedEventFilters
	}

	transformedPubsubTopic, err := expandCloudfunctions2functionEventTriggerPubsubTopic(original["pubsub_topic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPubsubTopic); val.IsValid() && !isEmptyValue(val) {
		transformed["pubsubTopic"] = transformedPubsubTopic
	}

	transformedServiceAccountEmail, err := expandCloudfunctions2functionEventTriggerServiceAccountEmail(original["service_account_email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAccountEmail); val.IsValid() && !isEmptyValue(val) {
		transformed["serviceAccountEmail"] = transformedServiceAccountEmail
	}

	transformedRetryPolicy, err := expandCloudfunctions2functionEventTriggerRetryPolicy(original["retry_policy"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRetryPolicy); val.IsValid() && !isEmptyValue(val) {
		transformed["retryPolicy"] = transformedRetryPolicy
	}

	return transformed, nil
}

func expandCloudfunctions2functionEventTriggerTrigger(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionEventTriggerTriggerRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionEventTriggerEventType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionEventTriggerEventFilters(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedAttribute, err := expandCloudfunctions2functionEventTriggerEventFiltersAttribute(original["attribute"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAttribute); val.IsValid() && !isEmptyValue(val) {
			transformed["attribute"] = transformedAttribute
		}

		transformedValue, err := expandCloudfunctions2functionEventTriggerEventFiltersValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !isEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		transformedOperator, err := expandCloudfunctions2functionEventTriggerEventFiltersOperator(original["operator"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedOperator); val.IsValid() && !isEmptyValue(val) {
			transformed["operator"] = transformedOperator
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudfunctions2functionEventTriggerEventFiltersAttribute(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionEventTriggerEventFiltersValue(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionEventTriggerEventFiltersOperator(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionEventTriggerPubsubTopic(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionEventTriggerServiceAccountEmail(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionEventTriggerRetryPolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
