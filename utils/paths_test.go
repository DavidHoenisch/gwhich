package utils

import "testing"

func TestGetUserConfigPath(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name: "config exists"
			want:
		}

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserConfigPath()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserConfigPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUserConfigPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
