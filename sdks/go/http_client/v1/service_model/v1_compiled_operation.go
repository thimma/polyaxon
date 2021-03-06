// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CompiledOperation CompiledOperation specification
//
// swagger:model v1CompiledOperation
type V1CompiledOperation struct {

	// Optional flag to disable cache validation and force run this component
	Cache *V1Cache `json:"cache,omitempty"`

	// An optional list of condition to check before starting the run, entities should be a valid Condition
	Conditions []interface{} `json:"conditions"`

	// Optional graph dependencies of this op
	Dependencies []string `json:"dependencies"`

	// Optional component description
	Description string `json:"description,omitempty"`

	// Optional events section, must be a valid List of Event option (Git/Alert/Webhook/Dataset)
	Events []interface{} `json:"events"`

	// Optional inputs definition
	Inputs []*V1IO `json:"inputs"`

	// Optional component kind, should be equal to "operation"
	Kind string `json:"kind,omitempty"`

	// Optional matrix section, must be a valid matrix option (Random/Grid/BO/Hyperband/Hyperopt/Mapping/Iterative)
	Matrix interface{} `json:"matrix,omitempty"`

	// Optional component name, should a valid slug
	Name string `json:"name,omitempty"`

	// Optional outputs definition
	Outputs []*V1IO `json:"outputs"`

	// Optional plugins to enable
	Plugins *V1Plugins `json:"plugins,omitempty"`

	// Optional profile to use for running this component
	Profile string `json:"profile,omitempty"`

	// Optional queue to use for running this component
	Queue string `json:"queue,omitempty"`

	// Run definiton, should be one of run composition: Container/Spark/Flink/Kubeflow/Dask/Dag
	Run interface{} `json:"run,omitempty"`

	// Optional schedule section, must be a valid Schedule option (Cron/Interval/Repeatable/ExactTime)
	Schedule interface{} `json:"schedule,omitempty"`

	// Optional flag to skip this run if upstream was skipped
	SkipOnUpstreamSkip bool `json:"skip_on_upstream_skip,omitempty"`

	// Optional component tags
	Tags []string `json:"tags"`

	// optional termination section
	Termination *V1Termination `json:"termination,omitempty"`

	// Optional trigger policy
	Trigger V1TriggerPolicy `json:"trigger,omitempty"`

	// Spec version
	Version float32 `json:"version,omitempty"`
}

// Validate validates this v1 compiled operation
func (m *V1CompiledOperation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCache(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInputs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutputs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlugins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTermination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrigger(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CompiledOperation) validateCache(formats strfmt.Registry) error {

	if swag.IsZero(m.Cache) { // not required
		return nil
	}

	if m.Cache != nil {
		if err := m.Cache.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cache")
			}
			return err
		}
	}

	return nil
}

func (m *V1CompiledOperation) validateInputs(formats strfmt.Registry) error {

	if swag.IsZero(m.Inputs) { // not required
		return nil
	}

	for i := 0; i < len(m.Inputs); i++ {
		if swag.IsZero(m.Inputs[i]) { // not required
			continue
		}

		if m.Inputs[i] != nil {
			if err := m.Inputs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1CompiledOperation) validateOutputs(formats strfmt.Registry) error {

	if swag.IsZero(m.Outputs) { // not required
		return nil
	}

	for i := 0; i < len(m.Outputs); i++ {
		if swag.IsZero(m.Outputs[i]) { // not required
			continue
		}

		if m.Outputs[i] != nil {
			if err := m.Outputs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("outputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1CompiledOperation) validatePlugins(formats strfmt.Registry) error {

	if swag.IsZero(m.Plugins) { // not required
		return nil
	}

	if m.Plugins != nil {
		if err := m.Plugins.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plugins")
			}
			return err
		}
	}

	return nil
}

func (m *V1CompiledOperation) validateTermination(formats strfmt.Registry) error {

	if swag.IsZero(m.Termination) { // not required
		return nil
	}

	if m.Termination != nil {
		if err := m.Termination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("termination")
			}
			return err
		}
	}

	return nil
}

func (m *V1CompiledOperation) validateTrigger(formats strfmt.Registry) error {

	if swag.IsZero(m.Trigger) { // not required
		return nil
	}

	if err := m.Trigger.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("trigger")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CompiledOperation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CompiledOperation) UnmarshalBinary(b []byte) error {
	var res V1CompiledOperation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
