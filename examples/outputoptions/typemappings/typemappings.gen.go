// Package typemappings provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package typemappings

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Client defines model for Client.
type Client struct {
	// Email should have the overridden go type.
	Email *string `json:"email,omitempty"`

	// Uuid should have no pointer receiver, even though it's an optional field.
	Uuid openapi_types.UUID `json:"uuid,omitempty"`
}

// StringAlias should be defined via an alias.
type StringAlias = string

// StringType should not be defined via an alias.
type StringType string