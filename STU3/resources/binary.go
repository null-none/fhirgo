package resources

import (
	d "github.com/null-none/fhirgo/STU3/datatypes"
	"github.com/null-none/fhirgo/schema"
)

// Binary resource
type Binary struct {
	Base
	ContentType     *d.Code         `json:"contentType,omitempty"`
	SecurityContext *d.Reference    `json:"securityContext,omitempty"`
	Content         *d.Base64Binary `json:"content,omitempty"`
}

// Validate returns a check against schema
func (a *Binary) Validate() (bool, []error) {
	return schema.ValidateResource(*a, "3")
}
