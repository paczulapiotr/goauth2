package security

import "testing"

func TestHashPassword(t *testing.T) {
	testPassword := "super_secure_password!2#"
	hash, err := HashPassword(testPassword)

	if err != nil || len(hash) == 0 {
		t.Fatal(err)
	}
}

func TestCheckPasswordHash(t *testing.T) {
	testPassword := "admin1!"
	hash, _ := HashPassword(testPassword)
	if !CheckPasswordHash(testPassword, hash) {
		t.Fatal("Hash is not matching password")
	}
}
