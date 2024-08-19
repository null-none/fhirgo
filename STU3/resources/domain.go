package resources

import (
	d "github.com/null-none/fhirgo/STU3/datatypes"
)

// Domain resource
type Domain struct {
	Base
	Text      *d.Narrative `json:"text,omitempty"`
	Contained []Base       `json:"contained,omitempty"`
}
