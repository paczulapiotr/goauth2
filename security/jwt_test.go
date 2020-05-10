package security

import (
	"testing"
	"time"
)

func TestCreateAuthoizationCode(t *testing.T) {
	codeBase := "auth code base stringz"
	code, validUntil := CreateAuthorizationCode(codeBase)

	expectedCodeStringLength := 44
	actualLength := len(code)
	if actualLength != expectedCodeStringLength {
		t.Fatalf("Authorization code has invalid lenght")
	}

	if validUntil.Before(time.Now()) {
		t.Fatalf("Auth code valid until date in invalid")
	}
}
