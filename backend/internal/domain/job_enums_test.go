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
