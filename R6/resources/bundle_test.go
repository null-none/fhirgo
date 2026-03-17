package resources

import (
	"encoding/json"
	"testing"

	d "github.com/null-none/fhirgo/R6/datatypes"
)

func TestBundleMarshal(t *testing.T) {
	bundleType := d.Code("collection")
	total := d.UnsignedInt(0)

	b := Bundle{}
	b.ResourceType = "Bundle"
	b.Type = &bundleType
	b.Total = &total

	data, err := json.Marshal(b)
	if err != nil {
		t.Fatal(err)
	}

	var got Bundle
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if got.ResourceType != "Bundle" {
		t.Errorf("ResourceType: got %v, want Bundle", got.ResourceType)
	}
	if *got.Type != bundleType {
		t.Errorf("Type: got %v, want %v", *got.Type, bundleType)
	}
}

func TestBundleWithPatientEntry(t *testing.T) {
	bundleType := d.Code("searchset")
	fullURL := d.URI("http://example.org/Patient/1")
	family := d.String("Ivanov")
	gender := d.Code("male")

	p := Patient{}
	p.ResourceType = "Patient"
	p.Name = []d.HumanName{{Family: &family}}
	p.Gender = &gender

	entry := d.BundleEntry{
		FullURL:  &fullURL,
		Resource: p,
	}

	b := Bundle{}
	b.ResourceType = "Bundle"
	b.Type = &bundleType
	b.Entry = []d.BundleEntry{entry}
	total := d.UnsignedInt(1)
	b.Total = &total

	data, err := json.Marshal(b)
	if err != nil {
		t.Fatal(err)
	}

	var got Bundle
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if len(got.Entry) != 1 {
		t.Fatalf("Entry length: got %d, want 1", len(got.Entry))
	}
	if *got.Total != d.UnsignedInt(1) {
		t.Errorf("Total: got %v, want 1", *got.Total)
	}
}

func TestBundleTransformPatient(t *testing.T) {
	bundleType := d.Code("searchset")
	family := d.String("Petrov")
	gender := d.Code("female")

	p := Patient{}
	p.ResourceType = "Patient"
	p.Name = []d.HumanName{{Family: &family}}
	p.Gender = &gender

	entry := d.BundleEntry{Resource: p}

	b := Bundle{}
	b.ResourceType = "Bundle"
	b.Type = &bundleType
	b.Entry = []d.BundleEntry{entry}

	// Marshal and unmarshal to simulate receiving from API
	data, err := json.Marshal(b)
	if err != nil {
		t.Fatal(err)
	}

	var received Bundle
	if err := json.Unmarshal(data, &received); err != nil {
		t.Fatal(err)
	}

	result, err := received.Transform()
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Entry) != 1 {
		t.Fatalf("Entry length after Transform: got %d, want 1", len(result.Entry))
	}

	pat, ok := result.Entry[0].Resource.(Patient)
	if !ok {
		t.Fatalf("Entry.Resource is not a Patient, got %T", result.Entry[0].Resource)
	}
	if *pat.Gender != gender {
		t.Errorf("Patient.Gender: got %v, want %v", *pat.Gender, gender)
	}
}

func TestBundleTransformFilter(t *testing.T) {
	bundleType := d.Code("searchset")

	p := Patient{}
	p.ResourceType = "Patient"

	o := Observation{}
	o.ResourceType = "Observation"

	b := Bundle{}
	b.ResourceType = "Bundle"
	b.Type = &bundleType
	b.Entry = []d.BundleEntry{{Resource: p}, {Resource: o}}

	data, _ := json.Marshal(b)
	var received Bundle
	json.Unmarshal(data, &received)

	// Filter only Patient entries
	result, err := received.Transform("Patient")
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Entry) != 1 {
		t.Fatalf("filtered Entry length: got %d, want 1", len(result.Entry))
	}
	if *result.Total != d.UnsignedInt(1) {
		t.Errorf("filtered Total: got %v, want 1", *result.Total)
	}
}

func TestBundleWithLink(t *testing.T) {
	bundleType := d.Code("searchset")
	rel := d.String("self")
	url := d.URI("http://example.org/fhir/Patient?_count=10")

	b := Bundle{}
	b.ResourceType = "Bundle"
	b.Type = &bundleType
	b.Link = []d.BundleLink{{Relation: &rel, URL: &url}}

	data, err := json.Marshal(b)
	if err != nil {
		t.Fatal(err)
	}

	var got Bundle
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if len(got.Link) != 1 {
		t.Fatalf("Link length: got %d, want 1", len(got.Link))
	}
	if *got.Link[0].Relation != rel {
		t.Errorf("Link.Relation: got %v, want %v", *got.Link[0].Relation, rel)
	}
}
