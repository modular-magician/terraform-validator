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

const SQLSourceRepresentationInstanceAssetType string = "sqladmin.googleapis.com/SourceRepresentationInstance"

func resourceConverterSQLSourceRepresentationInstance() ResourceConverter {
	return ResourceConverter{
		AssetType: SQLSourceRepresentationInstanceAssetType,
		Convert:   GetSQLSourceRepresentationInstanceCaiObject,
	}
}

func GetSQLSourceRepresentationInstanceCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//sqladmin.googleapis.com/projects/{{project}}/instances/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetSQLSourceRepresentationInstanceApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: SQLSourceRepresentationInstanceAssetType,
			Resource: &AssetResource{
				Version:              "v1beta4",
				DiscoveryDocumentURI: assetName("https://www.googleapis.com/discovery/v1/apis/sqladmin/v1beta4/rest"),
				DiscoveryName:        "SourceRepresentationInstance",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetSQLSourceRepresentationInstanceApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandSQLSourceRepresentationInstanceName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	regionProp, err := expandSQLSourceRepresentationInstanceRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}
	databaseVersionProp, err := expandSQLSourceRepresentationInstanceDatabaseVersion(d.Get("database_version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("database_version"); !isEmptyValue(reflect.ValueOf(databaseVersionProp)) && (ok || !reflect.DeepEqual(v, databaseVersionProp)) {
		obj["databaseVersion"] = databaseVersionProp
	}
	onPremisesConfigurationProp, err := expandSQLSourceRepresentationInstanceOnPremisesConfiguration(nil, d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("on_premises_configuration"); !isEmptyValue(reflect.ValueOf(onPremisesConfigurationProp)) && (ok || !reflect.DeepEqual(v, onPremisesConfigurationProp)) {
		obj["onPremisesConfiguration"] = onPremisesConfigurationProp
	}

	return resourceSQLSourceRepresentationInstanceEncoder(d, config, obj)
}

func resourceSQLSourceRepresentationInstanceEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	opc := obj["onPremisesConfiguration"].(map[string]interface{})
	opc["hostPort"] = fmt.Sprintf("%v:%v", opc["host"], opc["port"])
	delete(opc, "host")
	delete(opc, "port")
	return obj, nil
}

func expandSQLSourceRepresentationInstanceName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandSQLSourceRepresentationInstanceRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandSQLSourceRepresentationInstanceDatabaseVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandSQLSourceRepresentationInstanceOnPremisesConfiguration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	transformedHost, err := expandSQLSourceRepresentationInstanceOnPremisesConfigurationHost(d.Get("host"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHost); val.IsValid() && !isEmptyValue(val) {
		transformed["host"] = transformedHost
	}

	transformedPort, err := expandSQLSourceRepresentationInstanceOnPremisesConfigurationPort(d.Get("port"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !isEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	return transformed, nil
}

func expandSQLSourceRepresentationInstanceOnPremisesConfigurationHost(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandSQLSourceRepresentationInstanceOnPremisesConfigurationPort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
