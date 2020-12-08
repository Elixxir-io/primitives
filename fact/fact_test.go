package fact

import (
	"reflect"
	"testing"
)

// Test NewFact() function returns a correctly formatted Fact
func TestNewFact(t *testing.T) {
	// Expected result
	e := Fact{
		Fact: "testing",
		T:    1,
	}

	g, err := NewFact(Email, "testing")
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(e, g) {
		t.Errorf("The returned Fact did not match the expected Fact")
	}
}

// Test Stringify() creates a string of the Fact
// The output is verified to work in the test below
func TestFact_Stringify(t *testing.T) {
	f := Fact{
		Fact: "testing",
		T:    1,
	}

	expected := "Etesting"
	got := f.Stringify()
	t.Log(got)

	if got != expected {
		t.Errorf("Marshalled object from Got did not match Expected.\n\tGot: %v\n\tExpected: %v", got, expected)
	}
}

// Test the UnstringifyFact function creates a Fact from a string
// NOTE: this test does not pass, with error "Unknown Fact FactType: Etesting"
func TestFact_UnstringifyFact(t *testing.T) {
	// Expected fact from above test
	e := Fact{
		Fact: "testing",
		T:    Email,
	}

	// Stringify-ed Fact from above test
	m := "Etesting"
	f, err := UnstringifyFact(m)
	if err != nil {
		t.Error(err)
	}

	t.Log(f.Fact)
	t.Log(f.T)

	if !reflect.DeepEqual(e, f) {
		t.Errorf("The returned Fact did not match the expected Fact")
	}
}

func TestFact_ValidateFact(t *testing.T)  {
	// Expected fact from above test
	e := Fact{
		Fact: "xxxxxxxxxxxx123873j7djd741jrfhoiajdfhoewnuflkjvauirfhvkjdsafqyuusakjcg@carrere.cc",
		T:    Email,
	}

	err := ValidateFact(e, "")
	if err != nil {
		t.Errorf("Unexpected error in happy path: %v", err)
	}
}