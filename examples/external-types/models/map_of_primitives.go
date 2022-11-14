// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	timeext "time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
)

// MapOfPrimitives A map of NoValidatePrimitive external types.
//
// If the "noValidation" hint is omitted in the definition above, this code won't build because `time.Duration` has no `Validate` method.
//
// swagger:model MapOfPrimitives
type MapOfPrimitives map[string]timeext.Duration

// Validate validates this map of primitives
func (m MapOfPrimitives) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this map of primitives based on context it is used
func (m MapOfPrimitives) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
