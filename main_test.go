package main

import "testing"

func TestChangeCloseTags(t *testing.T) {
	xml := "<tag key=\"val\"></tag>"
	expect := "<tag key=\"val\"/>"
	got := changeCloseTags(xml)
	if got != expect {
		t.Errorf("Result = %s; want %s", got, expect)
	}
}
