//
// This file contains copies of important types from mattermost-plugin-workflow
// to facilitate client functionality.
//
package workflowclient

import (
	"encoding/json"
)

type BaseTrigger struct {
	BaseName string `json:"name"`
	BaseType string `json:"type"`
}

type SetupParams struct {
	BaseTrigger BaseTrigger     `json:"base_trigger"`
	Trigger     json.RawMessage `json:"trigger"`
	CallbackURL string          `json:"callback_url"`
}

type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type VarInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TriggerParams struct {
	// Unique name for the trigger. Automaticly prefixed with pluginID
	TypeName string `json:"typename" validate:"nonzero"`

	// Name for the trigger to be showed in the UI
	// Not implemented yet. Here for forwards compatibility.
	DisplayName string `json:"display_name" validate:"nonzero"`

	// The parameters to the trigger
	// Not implemented yet. Here for forwards compatibility.
	Fields []Field `json:"fields"`

	// The vars that are provided
	VarInfos []VarInfo `json:"var_infos"`

	// Called with trigger fields when workflow is created
	TriggerSetupURL string `json:"trigger_setup_url" validate:"nonzero"`
}

type ActionParams struct {
	// Unique name for the trigger. Automaticly prefixed with pluginID
	TypeName string `json:"typename" validate:"nonzero"`

	// Name for the trigger to be showed in the UI
	// Not implemented yet. Here for forwards compatibility.
	DisplayName string `json:"display_name" validate:"nonzero"`

	// The parameters to the trigger
	// Not implemented yet. Here for forwards compatibility.
	Fields []Field `json:"fields"`

	// The vars that are provided
	VarInfos []VarInfo `json:"var_infos"`

	// The URL that will be called when the action is fired with the action parameters.
	URL string `json:"url" validate:"nonzero"`
}

type RegisterParams struct {
	Triggers []TriggerParams `json:"triggers"`
	Actions  []ActionParams  `json:"actions"`
}

type ActivateParameters struct {
	// The vars the trigger was activated with.
	TriggerVars map[string]string
}

type Context struct {
	WorkflowName    string          `json:"workflow_name"`
	InstanceId      string          `json:"instance_id"`
	DefaultTeamId   string          `json:"default_team"`
	BotUserId       string          `json:"bot_user_id"`
	CurrentStepName string          `json:"current_step_name"`
	ActionVars      map[string]Vars `json:"action_vars"`
}

type ActionActivationParams struct {
	Action  json.RawMessage `json:"action"`
	Context Context         `json:"context"`
}

type Vars map[string]string
