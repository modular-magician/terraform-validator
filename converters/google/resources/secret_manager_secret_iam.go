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
const SecretManagerSecretIAMAssetType string = "secretmanager.googleapis.com/Secret"

func resourceConverterSecretManagerSecretIamPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType:         SecretManagerSecretIAMAssetType,
		Convert:           GetSecretManagerSecretIamPolicyCaiObject,
		MergeCreateUpdate: MergeSecretManagerSecretIamPolicy,
	}
}

func resourceConverterSecretManagerSecretIamBinding() ResourceConverter {
	return ResourceConverter{
		AssetType:         SecretManagerSecretIAMAssetType,
		Convert:           GetSecretManagerSecretIamBindingCaiObject,
		FetchFullResource: FetchSecretManagerSecretIamPolicy,
		MergeCreateUpdate: MergeSecretManagerSecretIamBinding,
		MergeDelete:       MergeSecretManagerSecretIamBindingDelete,
	}
}

func resourceConverterSecretManagerSecretIamMember() ResourceConverter {
	return ResourceConverter{
		AssetType:         SecretManagerSecretIAMAssetType,
		Convert:           GetSecretManagerSecretIamMemberCaiObject,
		FetchFullResource: FetchSecretManagerSecretIamPolicy,
		MergeCreateUpdate: MergeSecretManagerSecretIamMember,
		MergeDelete:       MergeSecretManagerSecretIamMemberDelete,
	}
}

func GetSecretManagerSecretIamPolicyCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newSecretManagerSecretIamAsset(d, config, expandIamPolicyBindings)
}

func GetSecretManagerSecretIamBindingCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newSecretManagerSecretIamAsset(d, config, expandIamRoleBindings)
}

func GetSecretManagerSecretIamMemberCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newSecretManagerSecretIamAsset(d, config, expandIamMemberBindings)
}

func MergeSecretManagerSecretIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeSecretManagerSecretIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeSecretManagerSecretIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeSecretManagerSecretIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeSecretManagerSecretIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newSecretManagerSecretIamAsset(
	d TerraformResourceData,
	config *Config,
	expandBindings func(d TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//secretmanager.googleapis.com/{{secret}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: SecretManagerSecretIAMAssetType,
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchSecretManagerSecretIamPolicy(d TerraformResourceData, config *Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("{{secret}}"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		SecretManagerSecretIamUpdaterProducer,
		d,
		config,
		assetName("//secretmanager.googleapis.com/{{secret}}"),
		SecretManagerSecretIAMAssetType,
	)
}
