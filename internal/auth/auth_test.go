package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type testCase struct {
		headerKey   string
		headerVal   string
		expectedErr bool
	}

	runCases := []testCase{
		{"Authorization", "ApiKey 95174985719", true},
		{"Authorization", "  ApiKey 73849173", true},
		{"Authorization", "ApiKey ApiKey", false},
		{"Authorization", "  ", true},
		{"Authorization", "", true},
	}

	headers := http.Header{}
	for _, v := range runCases {
		headers.Set(v.headerKey, v.headerVal)
		_, err := GetAPIKey(headers)
		if err != nil && !v.expectedErr || err == nil && v.expectedErr {
			t.Errorf(`
Header: %s %s
Expected error: %v
Got error: %v
				`, v.headerVal, v.headerKey, v.expectedErr, err)
		}
	}
}
