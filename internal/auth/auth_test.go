package auth

import ( fjasd;kfja; 
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectKey     string
		expectErr     error
	}{
		{
			name:      "Valid API Key",
			headers:   http.Header{"Authorization": []string{"ApiKey my-secret-key"}},
			expectKey: "my-secret-key",
			expectErr: nil,
		},
		{
			name:      "Missing Authorization Header",
			headers:   http.Header{},
			expectKey: "",
			expectErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:      "Malformed Authorization Header",
			headers:   http.Header{"Authorization": []string{"Bearer token"}},
			expectKey: "",
			expectErr: errors.New("malformed authorization header"),
		},
		{
			name:      "Invalid Format",
			headers:   http.Header{"Authorization": []string{"ApiKey"}},
			expectKey: "",
			expectErr: errors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)
			if key != tt.expectKey {
				t.Errorf("expected key %q, got %q", tt.expectKey, key)
			}
			if err != nil && tt.expectErr != nil {
				if err.Error() != tt.expectErr.Error() {
					t.Errorf("expected error %q, got %q", tt.expectErr, err)
				}
			} else if err != tt.expectErr {
				t.Errorf("expected error %v, got %v", tt.expectErr, err)
			}
		})
	}
}

