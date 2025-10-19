package parser

import "testing"

func TestNew(t *testing.T) {
	opts := map[string]interface{}{
		"ecmaVersion": 2020,
	}
	p := New(opts)
	if p == nil {
		t.Fatal("expected parser to be created")
	}
	if p.Options["ecmaVersion"] != 2020 {
		t.Errorf("expected ecmaVersion to be 2020, got %v", p.Options["ecmaVersion"])
	}
}

func TestParse(t *testing.T) {
	p := New(nil)
	_, err := p.Parse("const x = 1;")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
