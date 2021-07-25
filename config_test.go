package main

import (
	"reflect"
	"testing"
)

func TestParseConfig(t *testing.T) {
	input := `
  start: 0800
  end: 1700
  interval: 60
  messages:
    - Message One
    - Message Two
`
	expected := Config{"0800", "1700", 60, []string{"Message One", "Message Two"}}
	result := ParseConfig([]byte(input))

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
