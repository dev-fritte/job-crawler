package domain

import "testing"

func TestHasTag(t *testing.T) {
	var j Job
	j.AddTag("Go")
	if !j.HasTag("go") {
		t.Errorf("HasTag(%q) = false, want true", "go")
	}
}
