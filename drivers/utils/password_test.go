package utils

import "testing"

func TestPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     bool
		wantErr  bool
	}{
		{
			name:     "Password too short",
			password: "short",
			want:     false,
			wantErr:  true,
		},
		{
			name:     "No uppercase letter",
			password: "password123!",
			want:     false,
			wantErr:  true,
		},
		{
			name:     "No lowercase letter",
			password: "PASSWORD123!",
			want:     false,
			wantErr:  true,
		},
		{
			name:     "No number",
			password: "Password!!",
			want:     false,
			wantErr:  true,
		},
		{
			name:     "No special character",
			password: "Password123",
			want:     false,
			wantErr:  true,
		},
		{
			name:     "Valid password",
			password: "Password!123",
			want:     true,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Password(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Password() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Password() = %v, want %v", got, tt.want)
			}
		})
	}
}
