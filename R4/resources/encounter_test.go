package resources

import (
	"encoding/json"
	"testing"

	d "github.com/null-none/fhirgo/R4/datatypes"
)

func TestEncounterMarshal(t *testing.T) {
	status := d.Code("finished")
	start := d.DateTime("2024-06-01T08:00:00Z")
	end := d.DateTime("2024-06-01T09:00:00Z")

	e := Encounter{}
	e.ResourceType = "Encounter"
	e.Status = &status
	e.Period = &d.Period{Start: &start, End: &end}

	data, err := json.Marshal(e)
	if err != nil {
		t.Fatal(err)
	}

	var got Encounter
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if got.ResourceType != "Encounter" {
		t.Errorf("ResourceType: got %v, want Encounter", got.ResourceType)
	}
	if *got.Status != status {
		t.Errorf("Status: got %v, want %v", *got.Status, status)
	}
	if *got.Period.Start != start {
		t.Errorf("Period.Start: got %v, want %v", *got.Period.Start, start)
	}
}

func TestEncounterWithSubject(t *testing.T) {
	status := d.Code("in-progress")
	ref := d.String("Patient/123")
	subject := d.Reference{Reference: &ref}

	e := Encounter{}
	e.ResourceType = "Encounter"
	e.Status = &status
	e.Subject = &subject

	data, err := json.Marshal(e)
	if err != nil {
		t.Fatal(err)
	}

	var got Encounter
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if got.Subject == nil {
		t.Fatal("Subject is nil")
	}
	if *got.Subject.Reference != ref {
		t.Errorf("Subject.Reference: got %v, want %v", *got.Subject.Reference, ref)
	}
}

func TestEncounterWithParticipant(t *testing.T) {
	ref := d.String("Practitioner/456")
	individual := d.Reference{Reference: &ref}
	participant := d.EncounterParticipant{Individual: &individual}

	e := Encounter{}
	e.ResourceType = "Encounter"
	e.Participant = []d.EncounterParticipant{participant}

	data, err := json.Marshal(e)
	if err != nil {
		t.Fatal(err)
	}

	var got Encounter
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if len(got.Participant) != 1 {
		t.Fatalf("Participant length: got %d, want 1", len(got.Participant))
	}
	if *got.Participant[0].Individual.Reference != ref {
		t.Errorf("Participant.Individual.Reference: got %v, want %v",
			*got.Participant[0].Individual.Reference, ref)
	}
}

func TestEncounterWithDiagnosis(t *testing.T) {
	ref := d.String("Condition/789")
	condition := d.Reference{Reference: &ref}
	rank := d.PositiveInt(1)
	diagnosis := d.EncounterDiagnosis{Condition: &condition, Rank: &rank}

	e := Encounter{}
	e.ResourceType = "Encounter"
	e.Diagnosis = []d.EncounterDiagnosis{diagnosis}

	data, err := json.Marshal(e)
	if err != nil {
		t.Fatal(err)
	}

	var got Encounter
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if len(got.Diagnosis) != 1 {
		t.Fatalf("Diagnosis length: got %d, want 1", len(got.Diagnosis))
	}
	if *got.Diagnosis[0].Rank != rank {
		t.Errorf("Diagnosis.Rank: got %v, want %v", *got.Diagnosis[0].Rank, rank)
	}
}

func TestEncounterValidate(t *testing.T) {
	status := d.Code("finished")
	classSystem := d.URI("http://terminology.hl7.org/CodeSystem/v3-ActCode")
	classCode := d.Code("AMB")
	class := d.Coding{System: &classSystem, Code: &classCode}

	e := Encounter{}
	e.ResourceType = "Encounter"
	e.Status = &status
	e.Class = &class

	valid, errs := e.Validate()
	if !valid {
		t.Errorf("expected valid Encounter, got errors: %v", errs)
	}
}
