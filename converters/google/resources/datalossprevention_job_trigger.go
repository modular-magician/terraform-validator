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

const DataLossPreventionJobTriggerAssetType string = "dlp.googleapis.com/JobTrigger"

func resourceConverterDataLossPreventionJobTrigger() ResourceConverter {
	return ResourceConverter{
		AssetType: DataLossPreventionJobTriggerAssetType,
		Convert:   GetDataLossPreventionJobTriggerCaiObject,
	}
}

func GetDataLossPreventionJobTriggerCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//dlp.googleapis.com/{{parent}}/jobTriggers/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetDataLossPreventionJobTriggerApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: DataLossPreventionJobTriggerAssetType,
			Resource: &AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/dlp/v2/rest",
				DiscoveryName:        "JobTrigger",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetDataLossPreventionJobTriggerApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandDataLossPreventionJobTriggerDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	displayNameProp, err := expandDataLossPreventionJobTriggerDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	statusProp, err := expandDataLossPreventionJobTriggerStatus(d.Get("status"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("status"); !isEmptyValue(reflect.ValueOf(statusProp)) && (ok || !reflect.DeepEqual(v, statusProp)) {
		obj["status"] = statusProp
	}
	triggersProp, err := expandDataLossPreventionJobTriggerTriggers(d.Get("triggers"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("triggers"); !isEmptyValue(reflect.ValueOf(triggersProp)) && (ok || !reflect.DeepEqual(v, triggersProp)) {
		obj["triggers"] = triggersProp
	}
	inspectJobProp, err := expandDataLossPreventionJobTriggerInspectJob(d.Get("inspect_job"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("inspect_job"); !isEmptyValue(reflect.ValueOf(inspectJobProp)) && (ok || !reflect.DeepEqual(v, inspectJobProp)) {
		obj["inspectJob"] = inspectJobProp
	}

	return resourceDataLossPreventionJobTriggerEncoder(d, config, obj)
}

func resourceDataLossPreventionJobTriggerEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	newObj := make(map[string]interface{})
	newObj["jobTrigger"] = obj
	return newObj, nil
}

func expandDataLossPreventionJobTriggerDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerStatus(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerTriggers(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedSchedule, err := expandDataLossPreventionJobTriggerTriggersSchedule(original["schedule"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSchedule); val.IsValid() && !isEmptyValue(val) {
			transformed["schedule"] = transformedSchedule
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDataLossPreventionJobTriggerTriggersSchedule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRecurrencePeriodDuration, err := expandDataLossPreventionJobTriggerTriggersScheduleRecurrencePeriodDuration(original["recurrence_period_duration"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRecurrencePeriodDuration); val.IsValid() && !isEmptyValue(val) {
		transformed["recurrencePeriodDuration"] = transformedRecurrencePeriodDuration
	}

	return transformed, nil
}

func expandDataLossPreventionJobTriggerTriggersScheduleRecurrencePeriodDuration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJob(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedInspectTemplateName, err := expandDataLossPreventionJobTriggerInspectJobInspectTemplateName(original["inspect_template_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInspectTemplateName); val.IsValid() && !isEmptyValue(val) {
		transformed["inspectTemplateName"] = transformedInspectTemplateName
	}

	transformedStorageConfig, err := expandDataLossPreventionJobTriggerInspectJobStorageConfig(original["storage_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStorageConfig); val.IsValid() && !isEmptyValue(val) {
		transformed["storageConfig"] = transformedStorageConfig
	}

	transformedActions, err := expandDataLossPreventionJobTriggerInspectJobActions(original["actions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedActions); val.IsValid() && !isEmptyValue(val) {
		transformed["actions"] = transformedActions
	}

	return transformed, nil
}

func expandDataLossPreventionJobTriggerInspectJobInspectTemplateName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTimespanConfig, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfig(original["timespan_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTimespanConfig); val.IsValid() && !isEmptyValue(val) {
		transformed["timespanConfig"] = transformedTimespanConfig
	}

	transformedDatastoreOptions, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptions(original["datastore_options"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatastoreOptions); val.IsValid() && !isEmptyValue(val) {
		transformed["datastoreOptions"] = transformedDatastoreOptions
	}

	transformedCloudStorageOptions, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptions(original["cloud_storage_options"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCloudStorageOptions); val.IsValid() && !isEmptyValue(val) {
		transformed["cloudStorageOptions"] = transformedCloudStorageOptions
	}

	transformedBigQueryOptions, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions(original["big_query_options"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBigQueryOptions); val.IsValid() && !isEmptyValue(val) {
		transformed["bigQueryOptions"] = transformedBigQueryOptions
	}

	return transformed, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedStartTime, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfigStartTime(original["start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !isEmptyValue(val) {
		transformed["startTime"] = transformedStartTime
	}

	transformedEndTime, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfigEndTime(original["end_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEndTime); val.IsValid() && !isEmptyValue(val) {
		transformed["endTime"] = transformedEndTime
	}

	transformedEnableAutoPopulationOfTimespanConfig, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfigEnableAutoPopulationOfTimespanConfig(original["enable_auto_population_of_timespan_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableAutoPopulationOfTimespanConfig); val.IsValid() && !isEmptyValue(val) {
		transformed["enableAutoPopulationOfTimespanConfig"] = transformedEnableAutoPopulationOfTimespanConfig
	}

	transformedTimestampField, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfigTimestampField(original["timestamp_field"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTimestampField); val.IsValid() && !isEmptyValue(val) {
		transformed["timestampField"] = transformedTimestampField
	}

	return transformed, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfigStartTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfigEndTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfigEnableAutoPopulationOfTimespanConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfigTimestampField(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedName, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldName(original["name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
		transformed["name"] = transformedName
	}

	return transformed, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPartitionId, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(original["partition_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPartitionId); val.IsValid() && !isEmptyValue(val) {
		transformed["partitionId"] = transformedPartitionId
	}

	transformedKind, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptionsKind(original["kind"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKind); val.IsValid() && !isEmptyValue(val) {
		transformed["kind"] = transformedKind
	}

	return transformed, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProjectId, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !isEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	transformedNamespaceId, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdNamespaceId(original["namespace_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNamespaceId); val.IsValid() && !isEmptyValue(val) {
		transformed["namespaceId"] = transformedNamespaceId
	}

	return transformed, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdProjectId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdNamespaceId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptionsKind(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedName, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptionsKindName(original["name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
		transformed["name"] = transformedName
	}

	return transformed, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptionsKindName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFileSet, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(original["file_set"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFileSet); val.IsValid() && !isEmptyValue(val) {
		transformed["fileSet"] = transformedFileSet
	}

	transformedBytesLimitPerFile, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsBytesLimitPerFile(original["bytes_limit_per_file"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBytesLimitPerFile); val.IsValid() && !isEmptyValue(val) {
		transformed["bytesLimitPerFile"] = transformedBytesLimitPerFile
	}

	transformedBytesLimitPerFilePercent, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsBytesLimitPerFilePercent(original["bytes_limit_per_file_percent"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBytesLimitPerFilePercent); val.IsValid() && !isEmptyValue(val) {
		transformed["bytesLimitPerFilePercent"] = transformedBytesLimitPerFilePercent
	}

	transformedFilesLimitPercent, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFilesLimitPercent(original["files_limit_percent"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFilesLimitPercent); val.IsValid() && !isEmptyValue(val) {
		transformed["filesLimitPercent"] = transformedFilesLimitPercent
	}

	transformedFileTypes, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypes(original["file_types"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFileTypes); val.IsValid() && !isEmptyValue(val) {
		transformed["fileTypes"] = transformedFileTypes
	}

	transformedSampleMethod, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethod(original["sample_method"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSampleMethod); val.IsValid() && !isEmptyValue(val) {
		transformed["sampleMethod"] = transformedSampleMethod
	}

	return transformed, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUrl, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetUrl(original["url"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUrl); val.IsValid() && !isEmptyValue(val) {
		transformed["url"] = transformedUrl
	}

	transformedRegexFileSet, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(original["regex_file_set"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRegexFileSet); val.IsValid() && !isEmptyValue(val) {
		transformed["regexFileSet"] = transformedRegexFileSet
	}

	return transformed, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetUrl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBucketName, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetBucketName(original["bucket_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBucketName); val.IsValid() && !isEmptyValue(val) {
		transformed["bucketName"] = transformedBucketName
	}

	transformedIncludeRegex, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetIncludeRegex(original["include_regex"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIncludeRegex); val.IsValid() && !isEmptyValue(val) {
		transformed["includeRegex"] = transformedIncludeRegex
	}

	transformedExcludeRegex, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetExcludeRegex(original["exclude_regex"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExcludeRegex); val.IsValid() && !isEmptyValue(val) {
		transformed["excludeRegex"] = transformedExcludeRegex
	}

	return transformed, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetBucketName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetIncludeRegex(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetExcludeRegex(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsBytesLimitPerFile(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsBytesLimitPerFilePercent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFilesLimitPercent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypes(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethod(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTableReference, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(original["table_reference"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTableReference); val.IsValid() && !isEmptyValue(val) {
		transformed["tableReference"] = transformedTableReference
	}

	transformedRowsLimit, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsRowsLimit(original["rows_limit"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRowsLimit); val.IsValid() && !isEmptyValue(val) {
		transformed["rowsLimit"] = transformedRowsLimit
	}

	transformedRowsLimitPercent, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsRowsLimitPercent(original["rows_limit_percent"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRowsLimitPercent); val.IsValid() && !isEmptyValue(val) {
		transformed["rowsLimitPercent"] = transformedRowsLimitPercent
	}

	transformedSampleMethod, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethod(original["sample_method"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSampleMethod); val.IsValid() && !isEmptyValue(val) {
		transformed["sampleMethod"] = transformedSampleMethod
	}

	transformedIdentifyingFields, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields(original["identifying_fields"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIdentifyingFields); val.IsValid() && !isEmptyValue(val) {
		transformed["identifyingFields"] = transformedIdentifyingFields
	}

	return transformed, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProjectId, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !isEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	transformedDatasetId, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceDatasetId(original["dataset_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatasetId); val.IsValid() && !isEmptyValue(val) {
		transformed["datasetId"] = transformedDatasetId
	}

	transformedTableId, err := expandDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceTableId(original["table_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTableId); val.IsValid() && !isEmptyValue(val) {
		transformed["tableId"] = transformedTableId
	}

	return transformed, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceProjectId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceDatasetId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceTableId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsRowsLimit(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsRowsLimitPercent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethod(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobActions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedSaveFindings, err := expandDataLossPreventionJobTriggerInspectJobActionsSaveFindings(original["save_findings"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSaveFindings); val.IsValid() && !isEmptyValue(val) {
			transformed["saveFindings"] = transformedSaveFindings
		}

		transformedPubSub, err := expandDataLossPreventionJobTriggerInspectJobActionsPubSub(original["pub_sub"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPubSub); val.IsValid() && !isEmptyValue(val) {
			transformed["pubSub"] = transformedPubSub
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDataLossPreventionJobTriggerInspectJobActionsSaveFindings(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedOutputConfig, err := expandDataLossPreventionJobTriggerInspectJobActionsSaveFindingsOutputConfig(original["output_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOutputConfig); val.IsValid() && !isEmptyValue(val) {
		transformed["outputConfig"] = transformedOutputConfig
	}

	return transformed, nil
}

func expandDataLossPreventionJobTriggerInspectJobActionsSaveFindingsOutputConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTable, err := expandDataLossPreventionJobTriggerInspectJobActionsSaveFindingsOutputConfigTable(original["table"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTable); val.IsValid() && !isEmptyValue(val) {
		transformed["table"] = transformedTable
	}

	transformedOutputSchema, err := expandDataLossPreventionJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchema(original["output_schema"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOutputSchema); val.IsValid() && !isEmptyValue(val) {
		transformed["outputSchema"] = transformedOutputSchema
	}

	return transformed, nil
}

func expandDataLossPreventionJobTriggerInspectJobActionsSaveFindingsOutputConfigTable(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProjectId, err := expandDataLossPreventionJobTriggerInspectJobActionsSaveFindingsOutputConfigTableProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !isEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	transformedDatasetId, err := expandDataLossPreventionJobTriggerInspectJobActionsSaveFindingsOutputConfigTableDatasetId(original["dataset_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatasetId); val.IsValid() && !isEmptyValue(val) {
		transformed["datasetId"] = transformedDatasetId
	}

	transformedTableId, err := expandDataLossPreventionJobTriggerInspectJobActionsSaveFindingsOutputConfigTableTableId(original["table_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTableId); val.IsValid() && !isEmptyValue(val) {
		transformed["tableId"] = transformedTableId
	}

	return transformed, nil
}

func expandDataLossPreventionJobTriggerInspectJobActionsSaveFindingsOutputConfigTableProjectId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobActionsSaveFindingsOutputConfigTableDatasetId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobActionsSaveFindingsOutputConfigTableTableId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchema(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionJobTriggerInspectJobActionsPubSub(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTopic, err := expandDataLossPreventionJobTriggerInspectJobActionsPubSubTopic(original["topic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTopic); val.IsValid() && !isEmptyValue(val) {
		transformed["topic"] = transformedTopic
	}

	return transformed, nil
}

func expandDataLossPreventionJobTriggerInspectJobActionsPubSubTopic(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
