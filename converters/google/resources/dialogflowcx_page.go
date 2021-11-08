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

const DialogflowCXPageAssetType string = "dialogflow.googleapis.com/Page"

func resourceConverterDialogflowCXPage() ResourceConverter {
	return ResourceConverter{
		AssetType: DialogflowCXPageAssetType,
		Convert:   GetDialogflowCXPageCaiObject,
	}
}

func GetDialogflowCXPageCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//dialogflow.googleapis.com/{{parent}}/pages/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetDialogflowCXPageApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: DialogflowCXPageAssetType,
			Resource: &AssetResource{
				Version:              "v3",
				DiscoveryDocumentURI: assetName("https://www.googleapis.com/discovery/v1/apis/dialogflow/v3/rest"),
				DiscoveryName:        "Page",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetDialogflowCXPageApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowCXPageDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	entryFulfillmentProp, err := expandDialogflowCXPageEntryFulfillment(d.Get("entry_fulfillment"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("entry_fulfillment"); !isEmptyValue(reflect.ValueOf(entryFulfillmentProp)) && (ok || !reflect.DeepEqual(v, entryFulfillmentProp)) {
		obj["entryFulfillment"] = entryFulfillmentProp
	}
	formProp, err := expandDialogflowCXPageForm(d.Get("form"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("form"); !isEmptyValue(reflect.ValueOf(formProp)) && (ok || !reflect.DeepEqual(v, formProp)) {
		obj["form"] = formProp
	}
	transitionRouteGroupsProp, err := expandDialogflowCXPageTransitionRouteGroups(d.Get("transition_route_groups"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("transition_route_groups"); !isEmptyValue(reflect.ValueOf(transitionRouteGroupsProp)) && (ok || !reflect.DeepEqual(v, transitionRouteGroupsProp)) {
		obj["transitionRouteGroups"] = transitionRouteGroupsProp
	}
	transitionRoutesProp, err := expandDialogflowCXPageTransitionRoutes(d.Get("transition_routes"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("transition_routes"); !isEmptyValue(reflect.ValueOf(transitionRoutesProp)) && (ok || !reflect.DeepEqual(v, transitionRoutesProp)) {
		obj["transitionRoutes"] = transitionRoutesProp
	}
	eventHandlersProp, err := expandDialogflowCXPageEventHandlers(d.Get("event_handlers"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("event_handlers"); !isEmptyValue(reflect.ValueOf(eventHandlersProp)) && (ok || !reflect.DeepEqual(v, eventHandlersProp)) {
		obj["eventHandlers"] = eventHandlersProp
	}
	languageCodeProp, err := expandDialogflowCXPageLanguageCode(d.Get("language_code"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("language_code"); !isEmptyValue(reflect.ValueOf(languageCodeProp)) && (ok || !reflect.DeepEqual(v, languageCodeProp)) {
		obj["languageCode"] = languageCodeProp
	}

	return obj, nil
}

func expandDialogflowCXPageDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageEntryFulfillment(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMessages, err := expandDialogflowCXPageEntryFulfillmentMessages(original["messages"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMessages); val.IsValid() && !isEmptyValue(val) {
		transformed["messages"] = transformedMessages
	}

	transformedWebhook, err := expandDialogflowCXPageEntryFulfillmentWebhook(original["webhook"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWebhook); val.IsValid() && !isEmptyValue(val) {
		transformed["webhook"] = transformedWebhook
	}

	transformedReturnPartialResponses, err := expandDialogflowCXPageEntryFulfillmentReturnPartialResponses(original["return_partial_responses"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedReturnPartialResponses); val.IsValid() && !isEmptyValue(val) {
		transformed["returnPartialResponses"] = transformedReturnPartialResponses
	}

	transformedTag, err := expandDialogflowCXPageEntryFulfillmentTag(original["tag"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTag); val.IsValid() && !isEmptyValue(val) {
		transformed["tag"] = transformedTag
	}

	return transformed, nil
}

func expandDialogflowCXPageEntryFulfillmentMessages(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedText, err := expandDialogflowCXPageEntryFulfillmentMessagesText(original["text"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedText); val.IsValid() && !isEmptyValue(val) {
			transformed["text"] = transformedText
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowCXPageEntryFulfillmentMessagesText(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedText, err := expandDialogflowCXPageEntryFulfillmentMessagesTextText(original["text"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedText); val.IsValid() && !isEmptyValue(val) {
		transformed["text"] = transformedText
	}

	transformedAllowPlaybackInterruption, err := expandDialogflowCXPageEntryFulfillmentMessagesTextAllowPlaybackInterruption(original["allow_playback_interruption"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowPlaybackInterruption); val.IsValid() && !isEmptyValue(val) {
		transformed["allowPlaybackInterruption"] = transformedAllowPlaybackInterruption
	}

	return transformed, nil
}

func expandDialogflowCXPageEntryFulfillmentMessagesTextText(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageEntryFulfillmentMessagesTextAllowPlaybackInterruption(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageEntryFulfillmentWebhook(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageEntryFulfillmentReturnPartialResponses(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageEntryFulfillmentTag(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageForm(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedParameters, err := expandDialogflowCXPageFormParameters(original["parameters"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedParameters); val.IsValid() && !isEmptyValue(val) {
		transformed["parameters"] = transformedParameters
	}

	return transformed, nil
}

func expandDialogflowCXPageFormParameters(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDisplayName, err := expandDialogflowCXPageFormParametersDisplayName(original["display_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDisplayName); val.IsValid() && !isEmptyValue(val) {
			transformed["displayName"] = transformedDisplayName
		}

		transformedRequired, err := expandDialogflowCXPageFormParametersRequired(original["required"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRequired); val.IsValid() && !isEmptyValue(val) {
			transformed["required"] = transformedRequired
		}

		transformedEntityType, err := expandDialogflowCXPageFormParametersEntityType(original["entity_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEntityType); val.IsValid() && !isEmptyValue(val) {
			transformed["entityType"] = transformedEntityType
		}

		transformedIsList, err := expandDialogflowCXPageFormParametersIsList(original["is_list"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIsList); val.IsValid() && !isEmptyValue(val) {
			transformed["isList"] = transformedIsList
		}

		transformedFillBehavior, err := expandDialogflowCXPageFormParametersFillBehavior(original["fill_behavior"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFillBehavior); val.IsValid() && !isEmptyValue(val) {
			transformed["fillBehavior"] = transformedFillBehavior
		}

		transformedRedact, err := expandDialogflowCXPageFormParametersRedact(original["redact"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRedact); val.IsValid() && !isEmptyValue(val) {
			transformed["redact"] = transformedRedact
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowCXPageFormParametersDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageFormParametersRequired(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageFormParametersEntityType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageFormParametersIsList(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageFormParametersFillBehavior(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedInitialPromptFulfillment, err := expandDialogflowCXPageFormParametersFillBehaviorInitialPromptFulfillment(original["initial_prompt_fulfillment"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInitialPromptFulfillment); val.IsValid() && !isEmptyValue(val) {
		transformed["initialPromptFulfillment"] = transformedInitialPromptFulfillment
	}

	return transformed, nil
}

func expandDialogflowCXPageFormParametersFillBehaviorInitialPromptFulfillment(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMessages, err := expandDialogflowCXPageFormParametersFillBehaviorInitialPromptFulfillmentMessages(original["messages"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMessages); val.IsValid() && !isEmptyValue(val) {
		transformed["messages"] = transformedMessages
	}

	transformedWebhook, err := expandDialogflowCXPageFormParametersFillBehaviorInitialPromptFulfillmentWebhook(original["webhook"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWebhook); val.IsValid() && !isEmptyValue(val) {
		transformed["webhook"] = transformedWebhook
	}

	transformedReturnPartialResponses, err := expandDialogflowCXPageFormParametersFillBehaviorInitialPromptFulfillmentReturnPartialResponses(original["return_partial_responses"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedReturnPartialResponses); val.IsValid() && !isEmptyValue(val) {
		transformed["returnPartialResponses"] = transformedReturnPartialResponses
	}

	transformedTag, err := expandDialogflowCXPageFormParametersFillBehaviorInitialPromptFulfillmentTag(original["tag"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTag); val.IsValid() && !isEmptyValue(val) {
		transformed["tag"] = transformedTag
	}

	return transformed, nil
}

func expandDialogflowCXPageFormParametersFillBehaviorInitialPromptFulfillmentMessages(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedText, err := expandDialogflowCXPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesText(original["text"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedText); val.IsValid() && !isEmptyValue(val) {
			transformed["text"] = transformedText
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowCXPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesText(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedText, err := expandDialogflowCXPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesTextText(original["text"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedText); val.IsValid() && !isEmptyValue(val) {
		transformed["text"] = transformedText
	}

	transformedAllowPlaybackInterruption, err := expandDialogflowCXPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesTextAllowPlaybackInterruption(original["allow_playback_interruption"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowPlaybackInterruption); val.IsValid() && !isEmptyValue(val) {
		transformed["allowPlaybackInterruption"] = transformedAllowPlaybackInterruption
	}

	return transformed, nil
}

func expandDialogflowCXPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesTextText(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesTextAllowPlaybackInterruption(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageFormParametersFillBehaviorInitialPromptFulfillmentWebhook(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageFormParametersFillBehaviorInitialPromptFulfillmentReturnPartialResponses(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageFormParametersFillBehaviorInitialPromptFulfillmentTag(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageFormParametersRedact(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageTransitionRouteGroups(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageTransitionRoutes(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandDialogflowCXPageTransitionRoutesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedIntent, err := expandDialogflowCXPageTransitionRoutesIntent(original["intent"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIntent); val.IsValid() && !isEmptyValue(val) {
			transformed["intent"] = transformedIntent
		}

		transformedCondition, err := expandDialogflowCXPageTransitionRoutesCondition(original["condition"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCondition); val.IsValid() && !isEmptyValue(val) {
			transformed["condition"] = transformedCondition
		}

		transformedTriggerFulfillment, err := expandDialogflowCXPageTransitionRoutesTriggerFulfillment(original["trigger_fulfillment"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTriggerFulfillment); val.IsValid() && !isEmptyValue(val) {
			transformed["triggerFulfillment"] = transformedTriggerFulfillment
		}

		transformedTargetPage, err := expandDialogflowCXPageTransitionRoutesTargetPage(original["target_page"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTargetPage); val.IsValid() && !isEmptyValue(val) {
			transformed["targetPage"] = transformedTargetPage
		}

		transformedTargetFlow, err := expandDialogflowCXPageTransitionRoutesTargetFlow(original["target_flow"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTargetFlow); val.IsValid() && !isEmptyValue(val) {
			transformed["targetFlow"] = transformedTargetFlow
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowCXPageTransitionRoutesName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageTransitionRoutesIntent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageTransitionRoutesCondition(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageTransitionRoutesTriggerFulfillment(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMessages, err := expandDialogflowCXPageTransitionRoutesTriggerFulfillmentMessages(original["messages"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMessages); val.IsValid() && !isEmptyValue(val) {
		transformed["messages"] = transformedMessages
	}

	transformedWebhook, err := expandDialogflowCXPageTransitionRoutesTriggerFulfillmentWebhook(original["webhook"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWebhook); val.IsValid() && !isEmptyValue(val) {
		transformed["webhook"] = transformedWebhook
	}

	transformedReturnPartialResponses, err := expandDialogflowCXPageTransitionRoutesTriggerFulfillmentReturnPartialResponses(original["return_partial_responses"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedReturnPartialResponses); val.IsValid() && !isEmptyValue(val) {
		transformed["returnPartialResponses"] = transformedReturnPartialResponses
	}

	transformedTag, err := expandDialogflowCXPageTransitionRoutesTriggerFulfillmentTag(original["tag"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTag); val.IsValid() && !isEmptyValue(val) {
		transformed["tag"] = transformedTag
	}

	return transformed, nil
}

func expandDialogflowCXPageTransitionRoutesTriggerFulfillmentMessages(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedText, err := expandDialogflowCXPageTransitionRoutesTriggerFulfillmentMessagesText(original["text"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedText); val.IsValid() && !isEmptyValue(val) {
			transformed["text"] = transformedText
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowCXPageTransitionRoutesTriggerFulfillmentMessagesText(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedText, err := expandDialogflowCXPageTransitionRoutesTriggerFulfillmentMessagesTextText(original["text"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedText); val.IsValid() && !isEmptyValue(val) {
		transformed["text"] = transformedText
	}

	transformedAllowPlaybackInterruption, err := expandDialogflowCXPageTransitionRoutesTriggerFulfillmentMessagesTextAllowPlaybackInterruption(original["allow_playback_interruption"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowPlaybackInterruption); val.IsValid() && !isEmptyValue(val) {
		transformed["allowPlaybackInterruption"] = transformedAllowPlaybackInterruption
	}

	return transformed, nil
}

func expandDialogflowCXPageTransitionRoutesTriggerFulfillmentMessagesTextText(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageTransitionRoutesTriggerFulfillmentMessagesTextAllowPlaybackInterruption(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageTransitionRoutesTriggerFulfillmentWebhook(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageTransitionRoutesTriggerFulfillmentReturnPartialResponses(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageTransitionRoutesTriggerFulfillmentTag(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageTransitionRoutesTargetPage(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageTransitionRoutesTargetFlow(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageEventHandlers(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandDialogflowCXPageEventHandlersName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedEvent, err := expandDialogflowCXPageEventHandlersEvent(original["event"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEvent); val.IsValid() && !isEmptyValue(val) {
			transformed["event"] = transformedEvent
		}

		transformedTriggerFulfillment, err := expandDialogflowCXPageEventHandlersTriggerFulfillment(original["trigger_fulfillment"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTriggerFulfillment); val.IsValid() && !isEmptyValue(val) {
			transformed["triggerFulfillment"] = transformedTriggerFulfillment
		}

		transformedTargetPage, err := expandDialogflowCXPageEventHandlersTargetPage(original["target_page"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTargetPage); val.IsValid() && !isEmptyValue(val) {
			transformed["targetPage"] = transformedTargetPage
		}

		transformedTargetFlow, err := expandDialogflowCXPageEventHandlersTargetFlow(original["target_flow"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTargetFlow); val.IsValid() && !isEmptyValue(val) {
			transformed["targetFlow"] = transformedTargetFlow
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowCXPageEventHandlersName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageEventHandlersEvent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageEventHandlersTriggerFulfillment(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMessages, err := expandDialogflowCXPageEventHandlersTriggerFulfillmentMessages(original["messages"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMessages); val.IsValid() && !isEmptyValue(val) {
		transformed["messages"] = transformedMessages
	}

	transformedWebhook, err := expandDialogflowCXPageEventHandlersTriggerFulfillmentWebhook(original["webhook"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWebhook); val.IsValid() && !isEmptyValue(val) {
		transformed["webhook"] = transformedWebhook
	}

	transformedReturnPartialResponses, err := expandDialogflowCXPageEventHandlersTriggerFulfillmentReturnPartialResponses(original["return_partial_responses"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedReturnPartialResponses); val.IsValid() && !isEmptyValue(val) {
		transformed["returnPartialResponses"] = transformedReturnPartialResponses
	}

	transformedTag, err := expandDialogflowCXPageEventHandlersTriggerFulfillmentTag(original["tag"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTag); val.IsValid() && !isEmptyValue(val) {
		transformed["tag"] = transformedTag
	}

	return transformed, nil
}

func expandDialogflowCXPageEventHandlersTriggerFulfillmentMessages(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedText, err := expandDialogflowCXPageEventHandlersTriggerFulfillmentMessagesText(original["text"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedText); val.IsValid() && !isEmptyValue(val) {
			transformed["text"] = transformedText
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowCXPageEventHandlersTriggerFulfillmentMessagesText(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedText, err := expandDialogflowCXPageEventHandlersTriggerFulfillmentMessagesTextText(original["text"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedText); val.IsValid() && !isEmptyValue(val) {
		transformed["text"] = transformedText
	}

	transformedAllowPlaybackInterruption, err := expandDialogflowCXPageEventHandlersTriggerFulfillmentMessagesTextAllowPlaybackInterruption(original["allow_playback_interruption"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowPlaybackInterruption); val.IsValid() && !isEmptyValue(val) {
		transformed["allowPlaybackInterruption"] = transformedAllowPlaybackInterruption
	}

	return transformed, nil
}

func expandDialogflowCXPageEventHandlersTriggerFulfillmentMessagesTextText(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageEventHandlersTriggerFulfillmentMessagesTextAllowPlaybackInterruption(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageEventHandlersTriggerFulfillmentWebhook(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageEventHandlersTriggerFulfillmentReturnPartialResponses(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageEventHandlersTriggerFulfillmentTag(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageEventHandlersTargetPage(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageEventHandlersTargetFlow(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXPageLanguageCode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
