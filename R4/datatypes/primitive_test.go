package datatypes

import (
	"encoding/json"
	"testing"
)

func TestBooleanType(t *testing.T) {
	b := Boolean(true)
	data, err := json.Marshal(b)
	if err != nil {
		t.Fatal(err)
	}
	var got Boolean
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}
	if got != b {
		t.Errorf("got %v, want %v", got, b)
	}
}

func TestIntegerType(t *testing.T) {
	i := Integer(42)
	data, _ := json.Marshal(i)
	var got Integer
	json.Unmarshal(data, &got)
	if got != i {
		t.Errorf("got %v, want %v", got, i)
	}
}

func TestStringType(t *testing.T) {
	s := String("hello")
	data, _ := json.Marshal(s)
	var got String
	json.Unmarshal(data, &got)
	if got != s {
		t.Errorf("got %v, want %v", got, s)
	}
}

func TestDecimalType(t *testing.T) {
	d := Decimal(3.14)
	data, _ := json.Marshal(d)
	var got Decimal
	json.Unmarshal(data, &got)
	if got != d {
		t.Errorf("got %v, want %v", got, d)
	}
}

func TestURIType(t *testing.T) {
	u := URI("http://hl7.org/fhir")
	data, _ := json.Marshal(u)
	var got URI
	json.Unmarshal(data, &got)
	if got != u {
		t.Errorf("got %v, want %v", got, u)
	}
}

func TestCodeType(t *testing.T) {
	c := Code("active")
	data, _ := json.Marshal(c)
	var got Code
	json.Unmarshal(data, &got)
	if got != c {
		t.Errorf("got %v, want %v", got, c)
	}
}

func TestDateType(t *testing.T) {
	d := Date("1990-01-01")
	data, _ := json.Marshal(d)
	var got Date
	json.Unmarshal(data, &got)
	if got != d {
		t.Errorf("got %v, want %v", got, d)
	}
}

func TestDateTimeType(t *testing.T) {
	dt := DateTime("2024-01-15T10:30:00Z")
	data, _ := json.Marshal(dt)
	var got DateTime
	json.Unmarshal(data, &got)
	if got != dt {
		t.Errorf("got %v, want %v", got, dt)
	}
}

func TestIDType(t *testing.T) {
	id := ID("patient-123")
	data, _ := json.Marshal(id)
	var got ID
	json.Unmarshal(data, &got)
	if got != id {
		t.Errorf("got %v, want %v", got, id)
	}
}

func TestPositiveIntType(t *testing.T) {
	p := PositiveInt(1)
	data, _ := json.Marshal(p)
	var got PositiveInt
	json.Unmarshal(data, &got)
	if got != p {
		t.Errorf("got %v, want %v", got, p)
	}
}

func TestUUIDType(t *testing.T) {
	u := UUID("urn:uuid:c757873d-ec9a-4326-a141-0a95f1b1e5f9")
	data, _ := json.Marshal(u)
	var got UUID
	json.Unmarshal(data, &got)
	if got != u {
		t.Errorf("got %v, want %v", got, u)
	}
}

func TestMarkdownType(t *testing.T) {
	m := Markdown("**bold** text")
	data, _ := json.Marshal(m)
	var got Markdown
	json.Unmarshal(data, &got)
	if got != m {
		t.Errorf("got %v, want %v", got, m)
	}
}
