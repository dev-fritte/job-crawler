package domain

import (
	"slices"
	"testing"
)

func TestHasTag(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		present []string
		want    bool
	}{
		{"lowercased", "NormaLized", []string{"normalized"}, true},
		{"trimmed", "  go  ", []string{"go"}, true},
		{"unavailable", "golang", []string{"go", "lang", "typescript"}, false},
		{"empty ignored", "", []string{"tag1", "tag2"}, false},
		{"whitespace only ignored", "   ", []string{"tag1", "tag2"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			job := Job{
				Tags: tt.present,
			}

			if got := job.HasTag(tt.input); got != tt.want {
				t.Errorf("HasTag(%q) with Tags = %q, got %v, want %v", tt.input, tt.present, got, tt.want)
			}

		})
	}
}

func TestAddTag(t *testing.T) {
	tests := []struct {
		name   string
		inputs []string
		want   []string
	}{
		{"lowercased", []string{"NormaLized"}, []string{"normalized"}},
		{"trimmed", []string{"  go  "}, []string{"go"}},
		{"empty ignored", []string{""}, nil},
		{"whitespace only ignored", []string{"   "}, nil},
		{"duplicate ignored", []string{"go", "go"}, []string{"go"}},
		{"duplicate ignoring case and spaces", []string{"go", " GO "}, []string{"go"}},
		{"keeps order", []string{"go", "docker"}, []string{"go", "docker"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var job Job
			for _, in := range tt.inputs {
				job.AddTag(in)

			}
			if !slices.Equal(job.Tags, tt.want) {
				t.Errorf("AddTag(%q) -> Tags = %q, want %q", tt.inputs, job.Tags, tt.want)
			}
		})
	}
}
