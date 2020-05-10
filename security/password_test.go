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

func TestCheckPasswordStructure(t *testing.T) {
	tooShortPassword := "123"
	tooLongPassword := "12n0invfds9ovaafdsrngoi"
	validPassword := "Qwerty!1@"

	if CheckPasswordStructure(tooShortPassword) == nil {
		t.Fatalf("Should return error for too short password")
	}

	if CheckPasswordStructure(tooLongPassword) == nil {
		t.Fatalf("Should return error for too long password")
	}

	if CheckPasswordStructure(validPassword) != nil {
		t.Fatalf("Should not return error for valid password")
	}
}
