package util

import (
	"testing"
)

func TestGenToken(t *testing.T) {
	t.Setenv("JWT_KEY", "somekey")
	id := "1"
	token, err := GenToken(id)

	if err != nil {
		t.Fatalf("Should have returned a nil error instead got: %v\n", err)
	}

	if len(token) == 0 {
		t.Fatalf("Should have returned a token %v\n", len(token))
	}
}

func TestGenTokenErrorNoId(t *testing.T) {
	t.Setenv("JWT_KEY", "somekey")

	id := ""
	token, err := GenToken(id)
	if err == nil {
		t.Fatal("Expected an error returned, got nil instead\n")
	}

	if len(token) > 0 {
		t.Fatalf("Token should be empty: %v\n", token)
	}
}

func TestVerifyToken(t *testing.T) {
	t.Setenv("JWT_KEY", "somekey")

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6IjEiLCJleHAiOjE2NDQ2MjY5NzEsImlzcyI6ImdvcG1zZXJ2ZXIifQ.FUfGmYbv1GT6Kr_sbaJNWgX8MRtqbK43KyjQXCW7jfk"
	userId, err := VerifyToken(token)
	expectedId := "1"
	if err != nil {
		t.Fatalf("Nil should have been nil instead got: %v\n", err)
	}

	if userId != expectedId {
		t.Fatalf("Wrong id returned. Expect: %v got: %v\n", expectedId, userId)
	}
}

func TestVerifyTokenError(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6IjEiLCJleHAiOjE2NDQ2MjY5NzEsImlzcyI6ImdvcG1zZXJ2ZXIifQ.FUfGmYbv1GT6Kr_sbaJNWgX8MRtqbK43KyjQXCW7jfk"
	userId, err := VerifyToken(token)
	if err == nil {
		t.Fatal("Nil should not be nil")
	}

	if len(userId) > 0 {
		t.Fatalf("Expected an empty id instead got: %v\n", userId)
	}

}

func TestGetTokenFromBearerString(t *testing.T) {
	bearer := "Bearer token"
	token, err := GetTokenFromBearerString(bearer)
	if err != nil {
		t.Fatalf("Should have returned a nil err instead got: %v\n", err)
	}
	if token != "token" {
		t.Fatalf("Wrong token returned. Expect: %v got: %v\n", "token", token)
	}
}

func TestGetTokenFromBearerStringErr(t *testing.T) {
	bearer := "Bearer"
	token, err := GetTokenFromBearerString(bearer)
	if err == nil {
		t.Fatalf("Error expect instead got nil")
	}
	if err.Error() != "No token provided" {
		t.Fatalf("Wrong error returned. Expect: %v got: %v\n", "No token provided", err)
	}

	if len(token) > 0 {
		t.Fatalf("Token lenght should be 0 instead got: %v\n", len(token))
	}
}
