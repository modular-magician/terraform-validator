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
const ComputeDiskIAMAssetType string = "compute.googleapis.com/Disk"

func resourceConverterComputeDiskIamPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType:         ComputeDiskIAMAssetType,
		Convert:           GetComputeDiskIamPolicyCaiObject,
		MergeCreateUpdate: MergeComputeDiskIamPolicy,
	}
}

func resourceConverterComputeDiskIamBinding() ResourceConverter {
	return ResourceConverter{
		AssetType:         ComputeDiskIAMAssetType,
		Convert:           GetComputeDiskIamBindingCaiObject,
		FetchFullResource: FetchComputeDiskIamPolicy,
		MergeCreateUpdate: MergeComputeDiskIamBinding,
		MergeDelete:       MergeComputeDiskIamBindingDelete,
	}
}

func resourceConverterComputeDiskIamMember() ResourceConverter {
	return ResourceConverter{
		AssetType:         ComputeDiskIAMAssetType,
		Convert:           GetComputeDiskIamMemberCaiObject,
		FetchFullResource: FetchComputeDiskIamPolicy,
		MergeCreateUpdate: MergeComputeDiskIamMember,
		MergeDelete:       MergeComputeDiskIamMemberDelete,
	}
}

func GetComputeDiskIamPolicyCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newComputeDiskIamAsset(d, config, expandIamPolicyBindings)
}

func GetComputeDiskIamBindingCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newComputeDiskIamAsset(d, config, expandIamRoleBindings)
}

func GetComputeDiskIamMemberCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newComputeDiskIamAsset(d, config, expandIamMemberBindings)
}

func MergeComputeDiskIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeComputeDiskIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeComputeDiskIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeComputeDiskIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeComputeDiskIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newComputeDiskIamAsset(
	d TerraformResourceData,
	config *Config,
	expandBindings func(d TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/zones/{{zone}}/disks/{{name}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: ComputeDiskIAMAssetType,
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchComputeDiskIamPolicy(d TerraformResourceData, config *Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("{{zone}}"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("{{name}}"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		ComputeDiskIamUpdaterProducer,
		d,
		config,
		"//compute.googleapis.com/projects/{{project}}/zones/{{zone}}/disks/{{name}}",
		ComputeDiskIAMAssetType,
	)
}
