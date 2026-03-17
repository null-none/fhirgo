package datatypes

import (
	"encoding/json"
	"testing"
)

func TestHumanName(t *testing.T) {
	family := String("Doe")
	given := []String{"John"}
	use := Code("official")
	name := HumanName{
		Use:    &use,
		Family: &family,
		Given:  given,
	}

	data, err := json.Marshal(name)
	if err != nil {
		t.Fatal(err)
	}

	var got HumanName
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if *got.Family != family {
		t.Errorf("Family: got %v, want %v", *got.Family, family)
	}
	if *got.Use != use {
		t.Errorf("Use: got %v, want %v", *got.Use, use)
	}
	if len(got.Given) != 1 || got.Given[0] != "John" {
		t.Errorf("Given: got %v, want %v", got.Given, given)
	}
}

func TestAddress(t *testing.T) {
	city := String("Moscow")
	country := String("RU")
	use := Code("home")
	addr := Address{
		Use:     &use,
		City:    &city,
		Country: &country,
	}

	data, err := json.Marshal(addr)
	if err != nil {
		t.Fatal(err)
	}

	var got Address
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if *got.City != city {
		t.Errorf("City: got %v, want %v", *got.City, city)
	}
	if *got.Country != country {
		t.Errorf("Country: got %v, want %v", *got.Country, country)
	}
}

func TestContactPoint(t *testing.T) {
	system := Code("phone")
	value := String("+7-999-000-00-00")
	use := Code("work")
	cp := ContactPoint{
		System: &system,
		Value:  &value,
		Use:    &use,
	}

	data, err := json.Marshal(cp)
	if err != nil {
		t.Fatal(err)
	}

	var got ContactPoint
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if *got.System != system {
		t.Errorf("System: got %v, want %v", *got.System, system)
	}
	if *got.Value != value {
		t.Errorf("Value: got %v, want %v", *got.Value, value)
	}
}

func TestCoding(t *testing.T) {
	system := URI("http://loinc.org")
	code := Code("8867-4")
	display := String("Heart rate")
	coding := Coding{
		System:  &system,
		Code:    &code,
		Display: &display,
	}

	data, err := json.Marshal(coding)
	if err != nil {
		t.Fatal(err)
	}

	var got Coding
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if *got.Code != code {
		t.Errorf("Code: got %v, want %v", *got.Code, code)
	}
	if *got.Display != display {
		t.Errorf("Display: got %v, want %v", *got.Display, display)
	}
}

func TestCodeableConcept(t *testing.T) {
	system := URI("http://snomed.info/sct")
	code := Code("73211009")
	display := String("Diabetes mellitus")
	text := String("Diabetes mellitus")

	cc := CodeableConcept{
		Coding: []Coding{{System: &system, Code: &code, Display: &display}},
		Text:   &text,
	}

	data, err := json.Marshal(cc)
	if err != nil {
		t.Fatal(err)
	}

	var got CodeableConcept
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if *got.Text != text {
		t.Errorf("Text: got %v, want %v", *got.Text, text)
	}
	if len(got.Coding) != 1 {
		t.Fatalf("Coding length: got %d, want 1", len(got.Coding))
	}
	if *got.Coding[0].Code != code {
		t.Errorf("Coding.Code: got %v, want %v", *got.Coding[0].Code, code)
	}
}

func TestQuantity(t *testing.T) {
	value := Decimal(72.5)
	unit := String("bpm")
	system := URI("http://unitsofmeasure.org")
	code := Code("/min")

	q := Quantity{
		Value:  &value,
		Unit:   &unit,
		System: &system,
		Code:   &code,
	}

	data, err := json.Marshal(q)
	if err != nil {
		t.Fatal(err)
	}

	var got Quantity
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if *got.Value != value {
		t.Errorf("Value: got %v, want %v", *got.Value, value)
	}
	if *got.Unit != unit {
		t.Errorf("Unit: got %v, want %v", *got.Unit, unit)
	}
}

func TestPeriod(t *testing.T) {
	start := DateTime("2024-01-01T00:00:00Z")
	end := DateTime("2024-12-31T23:59:59Z")
	p := Period{Start: &start, End: &end}

	data, err := json.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}

	var got Period
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if *got.Start != start {
		t.Errorf("Start: got %v, want %v", *got.Start, start)
	}
	if *got.End != end {
		t.Errorf("End: got %v, want %v", *got.End, end)
	}
}

func TestIdentifier(t *testing.T) {
	use := Code("official")
	system := URI("http://hospital.org/patients")
	value := String("12345")

	id := Identifier{
		Use:    &use,
		System: &system,
		Value:  &value,
	}

	data, err := json.Marshal(id)
	if err != nil {
		t.Fatal(err)
	}

	var got Identifier
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if *got.Value != value {
		t.Errorf("Value: got %v, want %v", *got.Value, value)
	}
}

func TestMoney(t *testing.T) {
	val := Decimal(100.50)
	currency := Code("RUB")
	m := Money{Value: &val, Currency: &currency}

	data, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}

	var got Money
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if *got.Currency != currency {
		t.Errorf("Currency: got %v, want %v", *got.Currency, currency)
	}
}

func TestRange(t *testing.T) {
	lowVal := Decimal(60.0)
	highVal := Decimal(100.0)
	low := Quantity{Value: &lowVal}
	high := Quantity{Value: &highVal}
	r := Range{Low: &low, High: &high}

	data, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}

	var got Range
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	if *got.Low.Value != lowVal {
		t.Errorf("Low: got %v, want %v", *got.Low.Value, lowVal)
	}
	if *got.High.Value != highVal {
		t.Errorf("High: got %v, want %v", *got.High.Value, highVal)
	}
}
