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
const HealthcareConsentStoreIAMAssetType string = "healthcare.googleapis.com/ConsentStore"

func resourceConverterHealthcareConsentStoreIamPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType:         HealthcareConsentStoreIAMAssetType,
		Convert:           GetHealthcareConsentStoreIamPolicyCaiObject,
		MergeCreateUpdate: MergeHealthcareConsentStoreIamPolicy,
	}
}

func resourceConverterHealthcareConsentStoreIamBinding() ResourceConverter {
	return ResourceConverter{
		AssetType:         HealthcareConsentStoreIAMAssetType,
		Convert:           GetHealthcareConsentStoreIamBindingCaiObject,
		FetchFullResource: FetchHealthcareConsentStoreIamPolicy,
		MergeCreateUpdate: MergeHealthcareConsentStoreIamBinding,
		MergeDelete:       MergeHealthcareConsentStoreIamBindingDelete,
	}
}

func resourceConverterHealthcareConsentStoreIamMember() ResourceConverter {
	return ResourceConverter{
		AssetType:         HealthcareConsentStoreIAMAssetType,
		Convert:           GetHealthcareConsentStoreIamMemberCaiObject,
		FetchFullResource: FetchHealthcareConsentStoreIamPolicy,
		MergeCreateUpdate: MergeHealthcareConsentStoreIamMember,
		MergeDelete:       MergeHealthcareConsentStoreIamMemberDelete,
	}
}

func GetHealthcareConsentStoreIamPolicyCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newHealthcareConsentStoreIamAsset(d, config, expandIamPolicyBindings)
}

func GetHealthcareConsentStoreIamBindingCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newHealthcareConsentStoreIamAsset(d, config, expandIamRoleBindings)
}

func GetHealthcareConsentStoreIamMemberCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newHealthcareConsentStoreIamAsset(d, config, expandIamMemberBindings)
}

func MergeHealthcareConsentStoreIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeHealthcareConsentStoreIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeHealthcareConsentStoreIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeHealthcareConsentStoreIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeHealthcareConsentStoreIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newHealthcareConsentStoreIamAsset(
	d TerraformResourceData,
	config *Config,
	expandBindings func(d TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//healthcare.googleapis.com/{{dataset}}/consentStores/{{consent_store_id}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: HealthcareConsentStoreIAMAssetType,
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchHealthcareConsentStoreIamPolicy(d TerraformResourceData, config *Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("dataset"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("consent_store_id"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		HealthcareConsentStoreIamUpdaterProducer,
		d,
		config,
		"//healthcare.googleapis.com/{{dataset}}/consentStores/{{consent_store_id}}",
		HealthcareConsentStoreIAMAssetType,
	)
}
