package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "juan"
	// creamos una regex
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("juan")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`hello("juan") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}
}
