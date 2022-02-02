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

const DialogflowCXAgentAssetType string = "dialogflow.googleapis.com/Agent"

func resourceConverterDialogflowCXAgent() ResourceConverter {
	return ResourceConverter{
		AssetType: DialogflowCXAgentAssetType,
		Convert:   GetDialogflowCXAgentCaiObject,
	}
}

func GetDialogflowCXAgentCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//dialogflow.googleapis.com/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetDialogflowCXAgentApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: DialogflowCXAgentAssetType,
			Resource: &AssetResource{
				Version:              "v3",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/dialogflow/v3/rest",
				DiscoveryName:        "Agent",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetDialogflowCXAgentApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowCXAgentDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	defaultLanguageCodeProp, err := expandDialogflowCXAgentDefaultLanguageCode(d.Get("default_language_code"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("default_language_code"); !isEmptyValue(reflect.ValueOf(defaultLanguageCodeProp)) && (ok || !reflect.DeepEqual(v, defaultLanguageCodeProp)) {
		obj["defaultLanguageCode"] = defaultLanguageCodeProp
	}
	supportedLanguageCodesProp, err := expandDialogflowCXAgentSupportedLanguageCodes(d.Get("supported_language_codes"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("supported_language_codes"); !isEmptyValue(reflect.ValueOf(supportedLanguageCodesProp)) && (ok || !reflect.DeepEqual(v, supportedLanguageCodesProp)) {
		obj["supportedLanguageCodes"] = supportedLanguageCodesProp
	}
	timeZoneProp, err := expandDialogflowCXAgentTimeZone(d.Get("time_zone"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("time_zone"); !isEmptyValue(reflect.ValueOf(timeZoneProp)) && (ok || !reflect.DeepEqual(v, timeZoneProp)) {
		obj["timeZone"] = timeZoneProp
	}
	descriptionProp, err := expandDialogflowCXAgentDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	avatarUriProp, err := expandDialogflowCXAgentAvatarUri(d.Get("avatar_uri"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("avatar_uri"); !isEmptyValue(reflect.ValueOf(avatarUriProp)) && (ok || !reflect.DeepEqual(v, avatarUriProp)) {
		obj["avatarUri"] = avatarUriProp
	}
	speechToTextSettingsProp, err := expandDialogflowCXAgentSpeechToTextSettings(d.Get("speech_to_text_settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("speech_to_text_settings"); !isEmptyValue(reflect.ValueOf(speechToTextSettingsProp)) && (ok || !reflect.DeepEqual(v, speechToTextSettingsProp)) {
		obj["speechToTextSettings"] = speechToTextSettingsProp
	}
	securitySettingsProp, err := expandDialogflowCXAgentSecuritySettings(d.Get("security_settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("security_settings"); !isEmptyValue(reflect.ValueOf(securitySettingsProp)) && (ok || !reflect.DeepEqual(v, securitySettingsProp)) {
		obj["securitySettings"] = securitySettingsProp
	}
	enableStackdriverLoggingProp, err := expandDialogflowCXAgentEnableStackdriverLogging(d.Get("enable_stackdriver_logging"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_stackdriver_logging"); !isEmptyValue(reflect.ValueOf(enableStackdriverLoggingProp)) && (ok || !reflect.DeepEqual(v, enableStackdriverLoggingProp)) {
		obj["enableStackdriverLogging"] = enableStackdriverLoggingProp
	}
	enableSpellCorrectionProp, err := expandDialogflowCXAgentEnableSpellCorrection(d.Get("enable_spell_correction"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_spell_correction"); !isEmptyValue(reflect.ValueOf(enableSpellCorrectionProp)) && (ok || !reflect.DeepEqual(v, enableSpellCorrectionProp)) {
		obj["enableSpellCorrection"] = enableSpellCorrectionProp
	}

	return obj, nil
}

func expandDialogflowCXAgentDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentDefaultLanguageCode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentSupportedLanguageCodes(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentTimeZone(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentAvatarUri(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentSpeechToTextSettings(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnableSpeechAdaptation, err := expandDialogflowCXAgentSpeechToTextSettingsEnableSpeechAdaptation(original["enable_speech_adaptation"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableSpeechAdaptation); val.IsValid() && !isEmptyValue(val) {
		transformed["enableSpeechAdaptation"] = transformedEnableSpeechAdaptation
	}

	return transformed, nil
}

func expandDialogflowCXAgentSpeechToTextSettingsEnableSpeechAdaptation(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentSecuritySettings(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentEnableStackdriverLogging(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentEnableSpellCorrection(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
