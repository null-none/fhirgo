package resources

import (
	d "github.com/null-none/fhirgo/STU3/datatypes"
	"github.com/null-none/fhirgo/schema"
)

// Condition resource
type Condition struct {
	Domain
	Identifier         []d.Identifier      `json:"identifier,omitempty"`
	ClinicalStatus     *d.Code             `json:"clinicalStatus,omitempty"`
	VerificationStatus *d.Code             `json:"verificationStatus,omitempty"`
	Category           []d.CodeableConcept `json:"category,omitempty"`
	Severity           *d.CodeableConcept  `json:"severity,omitempty"`
	Code               *d.CodeableConcept  `json:"code,omitempty"`
	BodySite           *d.CodeableConcept  `json:"bodySite,omitempty"`
	Subject            *d.Reference        `json:"subject,omitempty"`
	Context            *d.Reference        `json:"context,omitempty"`
	OnsetDateTime      *d.DateTime         `json:"onsetDateTime,omitempty"`
	OnsetAge           *d.Age              `json:"onsetAge,omitempty"`
	OnsetPeriod        *d.Period           `json:"onsetPeriod,omitempty"`
	OnsetRange         *d.Range            `json:"onsetRange,omitempty"`
	OnsetString        *d.String           `json:"onsetString,omitempty"`
	AbatementDateTime  *d.DateTime         `json:"abatementDateTime,omitempty"`
	AbatementAge       *d.Age              `json:"abatementAge,omitempty"`
	AbatementBoolean   *d.Boolean          `json:"abatementBoolean,omitempty"`
	AbatementPeriod    *d.Period           `json:"abatementPeriod,omitempty"`
	AbatementRange     *d.Range            `json:"abatementRange,omitempty"`
	AbatementString    *d.String           `json:"abatementString,omitempty"`
	AssertedDate       *d.DateTime         `json:"assertedDate,omitempty"`
	Asserter           *d.Reference        `json:"asserter,omitempty"`
	Stage              *ConditionStage     `json:"stage,omitempty"`
	Evidence           []ConditionEvidence `json:"evidence,omitempty"`
	Note               []d.Annotation      `json:"note,omitempty"`
}

// Validate returns a check against schema
func (c *Condition) Validate() (bool, []error) {
	return schema.ValidateResource(*c, "3")
}

// ConditionStage subResource
type ConditionStage struct {
	Summary    *d.CodeableConcept `json:"summary,omitempty"`
	Assessment []d.Reference      `json:"assessment,omitempty"`
}

// ConditionEvidence subResource
type ConditionEvidence struct {
	Code   []d.CodeableConcept `json:"code,omitempty"`
	Detail []d.Reference       `json:"detail,omitempty"`
}
