package domain

import "testing"

func TestParseRemoteMode(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    RemoteMode
		wantErr bool
	}{
		{"canonical", "hybrid", RemoteModeHybrid, false},
		{"trimmed and upper", " REMOTE ", RemoteModeRemote, false},
		{"german alias remote", "Homeoffice", RemoteModeRemote, false},
		{"german alias onsite", "Vor Ort", RemoteModeOnsite, false},
		{"unknown", "xyz", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseRemoteMode(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestValidEmploymentType(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"canonical full-time", "full_time", true},
		{"canonical part-time", "part_time", true},
		{"canonical working_student", "working_student", true},
		{"unknown", "xyz", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EmploymentType(tt.input).Valid()
			if got != tt.want {
				t.Fatalf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestValidSenorityLevel(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"canonical junior", "junior", true},
		{"canonical mid", "mid", true},
		{"canonical senior", "senior", true},
		{"canonical lead", "lead", true},
		{"canonical principal", "principal", true},
		{"unknown", "xyz", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SeniorityLevel(tt.input).Valid()
			if got != tt.want {
				t.Fatalf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}
