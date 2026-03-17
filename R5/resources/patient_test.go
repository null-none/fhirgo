package resources

import (
	"encoding/json"
	"testing"

	d "github.com/null-none/fhirgo/R5/datatypes"
)

func TestPatientMarshal(t *testing.T) {
	family := d.String("Doe")
	given := []d.String{"John"}
	use := d.Code("official")
	gender := d.Code("male")
	birth := d.Date("1990-05-15")

	p := Patient{}
	p.ResourceType = "Patient"
	p.Name = []d.HumanName{{Use: &use, Family: &family, Given: given}}
	p.Gender = &gender
	p.BirthDate = &birth

	data, err := json.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}

	var got Patient
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if got.ResourceType != "Patient" {
		t.Errorf("ResourceType: got %v, want Patient", got.ResourceType)
	}
	if *got.Gender != gender {
		t.Errorf("Gender: got %v, want %v", *got.Gender, gender)
	}
	if *got.BirthDate != birth {
		t.Errorf("BirthDate: got %v, want %v", *got.BirthDate, birth)
	}
	if len(got.Name) != 1 || *got.Name[0].Family != family {
		t.Errorf("Name.Family: got %v, want %v", got.Name, family)
	}
}

func TestPatientWithIdentifier(t *testing.T) {
	idUse := d.Code("official")
	idSystem := d.URI("http://hospital.org/patients")
	idValue := d.String("PAT-001")

	p := Patient{}
	p.ResourceType = "Patient"
	p.Identifier = []d.Identifier{{Use: &idUse, System: &idSystem, Value: &idValue}}

	data, err := json.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}

	var got Patient
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if len(got.Identifier) != 1 {
		t.Fatalf("Identifier length: got %d, want 1", len(got.Identifier))
	}
	if *got.Identifier[0].Value != idValue {
		t.Errorf("Identifier.Value: got %v, want %v", *got.Identifier[0].Value, idValue)
	}
}

func TestPatientActive(t *testing.T) {
	active := d.Boolean(true)
	p := Patient{}
	p.ResourceType = "Patient"
	p.Active = &active

	data, err := json.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if m["active"] != true {
		t.Errorf("active: got %v, want true", m["active"])
	}
}

func TestPatientWithContact(t *testing.T) {
	family := d.String("Smith")
	contact := d.PatientContact{
		Name: &d.HumanName{Family: &family},
	}

	p := Patient{}
	p.ResourceType = "Patient"
	p.Contact = []d.PatientContact{contact}

	data, err := json.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}

	var got Patient
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if len(got.Contact) != 1 {
		t.Fatalf("Contact length: got %d, want 1", len(got.Contact))
	}
	if *got.Contact[0].Name.Family != family {
		t.Errorf("Contact.Name.Family: got %v, want %v", *got.Contact[0].Name.Family, family)
	}
}

func TestPatientEmptyMarshal(t *testing.T) {
	p := Patient{}
	p.ResourceType = "Patient"

	data, err := json.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}

	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}

	if m["resourceType"] != "Patient" {
		t.Errorf("resourceType: got %v, want Patient", m["resourceType"])
	}
}
