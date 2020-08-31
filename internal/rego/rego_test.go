package rego

import (
	"testing"
)

func TestKind(t *testing.T) {
	policy := Rego{
		path: "some/path/my-policy/src.rego",
	}

	actual := policy.Kind()

	const expected = "MyPolicy"
	if actual != expected {
		t.Errorf("unexpected Kind. expected %v, actual %v", expected, actual)
	}
}

func TestName(t *testing.T) {
	policy := Rego{
		path: "some/path/my-policy/src.rego",
	}

	actual := policy.Name()

	const expected = "mypolicy"
	if actual != expected {
		t.Errorf("unexpected Name. expected %v, actual %v", expected, actual)
	}
}

func TestTitle(t *testing.T) {
	comments := []string{
		"@title The title",
	}

	rego := Rego{
		comments: comments,
	}

	actual := rego.Title()

	const expected = "The title"
	if actual != expected {
		t.Errorf("unexpected Title. expected %v, actual %v", expected, actual)
	}
}

func TestDescription(t *testing.T) {
	comments := []string{
		"@title The title",
		"The description",
		"@kinds The kinds",
		"Extra comment",
	}

	rego := Rego{
		comments: comments,
	}

	actual := rego.Description()

	const expected = "The description"
	if actual != expected {
		t.Errorf("unexpected Description. expected %v, actual %v", expected, actual)
	}
}

func TestSeverity(t *testing.T) {
	rules := []string{
		"violation",
		"warn",
	}

	rego := Rego{
		rules: rules,
	}

	actual := rego.Severity()

	const expected = Violation
	if actual != expected {
		t.Errorf("unexpected Severity. expected %v, actual %v", expected, actual)
	}
}

func TestSource(t *testing.T) {
	raw := `first
# second
third
# fourth
`

	rego := Rego{
		raw: raw,
	}

	actual := rego.Source()

	const expected = `first
third`

	if actual != expected {
		t.Errorf("unexpected Source. expected %v, actual %v", expected, actual)
	}
}
