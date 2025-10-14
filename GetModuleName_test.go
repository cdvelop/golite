package golite

import "testing"

func TestGetModuleName(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
		wantErr  bool
	}{
		{
			name:     "valid module path",
			path:     "miproject/modules/user/model.go",
			expected: "user",
			wantErr:  false,
		},
		{
			name:     "valid module path \\",
			path:     "miproject\\modules\\user\\model.go",
			expected: "user",
			wantErr:  false,
		},
		{
			name:     "valid module with multiple levels",
			path:     "app/src/modules/auth/login.go",
			expected: "auth",
			wantErr:  false,
		},
		{
			name:     "invalid path without modules",
			path:     "miproject/web/module3.wasm",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "empty path",
			path:     "",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "path ends in modules",
			path:     "project/modules",
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetModuleName(tt.path)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetModuleName(%q) error = %v, wantErr %v", tt.path, err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Fatalf("GetModuleName(%q) = %v, want %v", tt.path, got, tt.expected)
			}
		})
	}
}
