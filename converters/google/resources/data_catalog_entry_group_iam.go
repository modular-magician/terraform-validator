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

import "fmt"

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const DataCatalogEntryGroupIAMAssetType string = "datacatalog.googleapis.com/EntryGroup"

func resourceConverterDataCatalogEntryGroupIamPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType:         DataCatalogEntryGroupIAMAssetType,
		Convert:           GetDataCatalogEntryGroupIamPolicyCaiObject,
		MergeCreateUpdate: MergeDataCatalogEntryGroupIamPolicy,
	}
}

func resourceConverterDataCatalogEntryGroupIamBinding() ResourceConverter {
	return ResourceConverter{
		AssetType:         DataCatalogEntryGroupIAMAssetType,
		Convert:           GetDataCatalogEntryGroupIamBindingCaiObject,
		FetchFullResource: FetchDataCatalogEntryGroupIamPolicy,
		MergeCreateUpdate: MergeDataCatalogEntryGroupIamBinding,
		MergeDelete:       MergeDataCatalogEntryGroupIamBindingDelete,
	}
}

func resourceConverterDataCatalogEntryGroupIamMember() ResourceConverter {
	return ResourceConverter{
		AssetType:         DataCatalogEntryGroupIAMAssetType,
		Convert:           GetDataCatalogEntryGroupIamMemberCaiObject,
		FetchFullResource: FetchDataCatalogEntryGroupIamPolicy,
		MergeCreateUpdate: MergeDataCatalogEntryGroupIamMember,
		MergeDelete:       MergeDataCatalogEntryGroupIamMemberDelete,
	}
}

func GetDataCatalogEntryGroupIamPolicyCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newDataCatalogEntryGroupIamAsset(d, config, expandIamPolicyBindings)
}

func GetDataCatalogEntryGroupIamBindingCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newDataCatalogEntryGroupIamAsset(d, config, expandIamRoleBindings)
}

func GetDataCatalogEntryGroupIamMemberCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newDataCatalogEntryGroupIamAsset(d, config, expandIamMemberBindings)
}

func MergeDataCatalogEntryGroupIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeDataCatalogEntryGroupIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeDataCatalogEntryGroupIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeDataCatalogEntryGroupIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeDataCatalogEntryGroupIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newDataCatalogEntryGroupIamAsset(
	d TerraformResourceData,
	config *Config,
	expandBindings func(d TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//datacatalog.googleapis.com/projects/{{project}}/locations/{{region}}/entryGroups/{{entry_group}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: DataCatalogEntryGroupIAMAssetType,
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchDataCatalogEntryGroupIamPolicy(d TerraformResourceData, config *Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("region"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("entry_group"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		DataCatalogEntryGroupIamUpdaterProducer,
		d,
		config,
		"//datacatalog.googleapis.com/projects/{{project}}/locations/{{region}}/entryGroups/{{entry_group}}",
		DataCatalogEntryGroupIAMAssetType,
	)
}
