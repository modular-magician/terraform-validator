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

const PubsubLiteTopicAssetType string = "{{region}}-pubsublite.googleapis.com/Topic"

func resourceConverterPubsubLiteTopic() ResourceConverter {
	return ResourceConverter{
		AssetType: PubsubLiteTopicAssetType,
		Convert:   GetPubsubLiteTopicCaiObject,
	}
}

func GetPubsubLiteTopicCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//{{region}}-pubsublite.googleapis.com/projects/{{project}}/locations/{{zone}}/topics/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetPubsubLiteTopicApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: PubsubLiteTopicAssetType,
			Resource: &AssetResource{
				Version:              "admin",
				DiscoveryDocumentURI: assetName("https://www.googleapis.com/discovery/v1/apis/{{region}}-pubsublite/admin/rest"),
				DiscoveryName:        "Topic",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetPubsubLiteTopicApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	partitionConfigProp, err := expandPubsubLiteTopicPartitionConfig(d.Get("partition_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("partition_config"); !isEmptyValue(reflect.ValueOf(partitionConfigProp)) && (ok || !reflect.DeepEqual(v, partitionConfigProp)) {
		obj["partitionConfig"] = partitionConfigProp
	}
	retentionConfigProp, err := expandPubsubLiteTopicRetentionConfig(d.Get("retention_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("retention_config"); !isEmptyValue(reflect.ValueOf(retentionConfigProp)) && (ok || !reflect.DeepEqual(v, retentionConfigProp)) {
		obj["retentionConfig"] = retentionConfigProp
	}
	reservationConfigProp, err := expandPubsubLiteTopicReservationConfig(d.Get("reservation_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("reservation_config"); !isEmptyValue(reflect.ValueOf(reservationConfigProp)) && (ok || !reflect.DeepEqual(v, reservationConfigProp)) {
		obj["reservationConfig"] = reservationConfigProp
	}

	return resourcePubsubLiteTopicEncoder(d, config, obj)
}

func resourcePubsubLiteTopicEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*Config)

	zone, err := getZone(d, config)
	if err != nil {
		return nil, err
	}

	if zone == "" {
		return nil, fmt.Errorf("zone must be non-empty - set in resource or at provider-level")
	}

	// API Endpoint requires region in the URL. We infer it from the zone.

	region := getRegionFromZone(zone)

	if region == "" {
		return nil, fmt.Errorf("invalid zone %q, unable to infer region from zone", zone)
	}

	return obj, nil
}

func expandPubsubLiteTopicPartitionConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCount, err := expandPubsubLiteTopicPartitionConfigCount(original["count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCount); val.IsValid() && !isEmptyValue(val) {
		transformed["count"] = transformedCount
	}

	transformedCapacity, err := expandPubsubLiteTopicPartitionConfigCapacity(original["capacity"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCapacity); val.IsValid() && !isEmptyValue(val) {
		transformed["capacity"] = transformedCapacity
	}

	return transformed, nil
}

func expandPubsubLiteTopicPartitionConfigCount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubLiteTopicPartitionConfigCapacity(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPublishMibPerSec, err := expandPubsubLiteTopicPartitionConfigCapacityPublishMibPerSec(original["publish_mib_per_sec"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPublishMibPerSec); val.IsValid() && !isEmptyValue(val) {
		transformed["publishMibPerSec"] = transformedPublishMibPerSec
	}

	transformedSubscribeMibPerSec, err := expandPubsubLiteTopicPartitionConfigCapacitySubscribeMibPerSec(original["subscribe_mib_per_sec"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubscribeMibPerSec); val.IsValid() && !isEmptyValue(val) {
		transformed["subscribeMibPerSec"] = transformedSubscribeMibPerSec
	}

	return transformed, nil
}

func expandPubsubLiteTopicPartitionConfigCapacityPublishMibPerSec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubLiteTopicPartitionConfigCapacitySubscribeMibPerSec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubLiteTopicRetentionConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPerPartitionBytes, err := expandPubsubLiteTopicRetentionConfigPerPartitionBytes(original["per_partition_bytes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPerPartitionBytes); val.IsValid() && !isEmptyValue(val) {
		transformed["perPartitionBytes"] = transformedPerPartitionBytes
	}

	transformedPeriod, err := expandPubsubLiteTopicRetentionConfigPeriod(original["period"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPeriod); val.IsValid() && !isEmptyValue(val) {
		transformed["period"] = transformedPeriod
	}

	return transformed, nil
}

func expandPubsubLiteTopicRetentionConfigPerPartitionBytes(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubLiteTopicRetentionConfigPeriod(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubLiteTopicReservationConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedThroughputReservation, err := expandPubsubLiteTopicReservationConfigThroughputReservation(original["throughput_reservation"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedThroughputReservation); val.IsValid() && !isEmptyValue(val) {
		transformed["throughputReservation"] = transformedThroughputReservation
	}

	return transformed, nil
}

func expandPubsubLiteTopicReservationConfigThroughputReservation(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("reservations", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for throughput_reservation: %s", err)
	}
	// Custom due to "locations" rather than "regions".
	return fmt.Sprintf("projects/%s/locations/%s/reservations/%s", f.Project, f.Region, f.Name), nil
}
