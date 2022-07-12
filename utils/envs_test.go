package utils

import "testing"

func TestGetEnv(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name     string
		env      string
		fallback string
		want     string
	}{
		{
			name:     "GET-PORT",
			env:      "PORT",
			fallback: "123",
			want:     "3000",
		},
		{
			name:     "ENV-NOT-FOUND-GET-DEFAULT",
			env:      "XPTO",
			fallback: "123",
			want:     "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnv(tt.env, tt.fallback); got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
