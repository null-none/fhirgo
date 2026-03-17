package resources

import (
	"encoding/json"
	"testing"

	d "github.com/null-none/fhirgo/R4/datatypes"
)

func TestObservationMarshal(t *testing.T) {
	status := d.Code("final")
	system := d.URI("http://loinc.org")
	code := d.Code("8867-4")
	display := d.String("Heart rate")
	codeCC := d.CodeableConcept{Coding: []d.Coding{{System: &system, Code: &code, Display: &display}}}
	value := d.Decimal(72)
	unit := d.String("bpm")
	valueQ := d.Quantity{Value: &value, Unit: &unit}

	o := Observation{}
	o.ResourceType = "Observation"
	o.Status = &status
	o.Code = &codeCC
	o.ValueQuantity = &valueQ

	data, err := json.Marshal(o)
	if err != nil {
		t.Fatal(err)
	}

	var got Observation
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if got.ResourceType != "Observation" {
		t.Errorf("ResourceType: got %v, want Observation", got.ResourceType)
	}
	if *got.Status != status {
		t.Errorf("Status: got %v, want %v", *got.Status, status)
	}
	if *got.ValueQuantity.Value != value {
		t.Errorf("ValueQuantity.Value: got %v, want %v", *got.ValueQuantity.Value, value)
	}
}

func TestObservationWithSubject(t *testing.T) {
	status := d.Code("final")
	ref := d.String("Patient/123")
	subject := d.Reference{Reference: &ref}

	o := Observation{}
	o.ResourceType = "Observation"
	o.Status = &status
	o.Subject = &subject

	data, err := json.Marshal(o)
	if err != nil {
		t.Fatal(err)
	}

	var got Observation
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if *got.Subject.Reference != ref {
		t.Errorf("Subject.Reference: got %v, want %v", *got.Subject.Reference, ref)
	}
}

func TestObservationWithComponent(t *testing.T) {
	system := d.URI("http://loinc.org")
	code := d.Code("8480-6")
	display := d.String("Systolic blood pressure")
	codeCC := d.CodeableConcept{Coding: []d.Coding{{System: &system, Code: &code, Display: &display}}}
	val := d.Decimal(120)
	unit := d.String("mmHg")
	comp := d.ObservationComponent{
		Code:          &codeCC,
		ValueQuantity: &d.Quantity{Value: &val, Unit: &unit},
	}

	o := Observation{}
	o.ResourceType = "Observation"
	o.Component = []d.ObservationComponent{comp}

	data, err := json.Marshal(o)
	if err != nil {
		t.Fatal(err)
	}

	var got Observation
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if len(got.Component) != 1 {
		t.Fatalf("Component length: got %d, want 1", len(got.Component))
	}
	if *got.Component[0].ValueQuantity.Value != val {
		t.Errorf("Component.ValueQuantity.Value: got %v, want %v",
			*got.Component[0].ValueQuantity.Value, val)
	}
}

func TestObservationWithReferenceRange(t *testing.T) {
	low := d.Decimal(60)
	high := d.Decimal(100)
	rr := d.ObservationReferenceRange{
		Low:  &d.SimpleQuantity{Value: &low},
		High: &d.SimpleQuantity{Value: &high},
	}

	o := Observation{}
	o.ResourceType = "Observation"
	o.ReferenceRange = []d.ObservationReferenceRange{rr}

	data, err := json.Marshal(o)
	if err != nil {
		t.Fatal(err)
	}

	var got Observation
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if len(got.ReferenceRange) != 1 {
		t.Fatalf("ReferenceRange length: got %d, want 1", len(got.ReferenceRange))
	}
	if *got.ReferenceRange[0].Low.Value != low {
		t.Errorf("ReferenceRange.Low: got %v, want %v", *got.ReferenceRange[0].Low.Value, low)
	}
}

func TestObservationValidate(t *testing.T) {
	status := d.Code("final")
	system := d.URI("http://loinc.org")
	code := d.Code("8867-4")
	codeCC := d.CodeableConcept{Coding: []d.Coding{{System: &system, Code: &code}}}
	effectiveDateTime := d.DateTime("2024-01-15T10:30:00Z")

	o := Observation{}
	o.ResourceType = "Observation"
	o.Status = &status
	o.Code = &codeCC
	o.EffectiveDateTime = &effectiveDateTime

	valid, errs := o.Validate()
	if !valid {
		t.Errorf("expected valid Observation, got errors: %v", errs)
	}
}
