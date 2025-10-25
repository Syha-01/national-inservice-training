package data

import (
	"testing"

	"github.com/Syha-01/national-inservice-training/internal/validator"
)

func TestPassword_Set(t *testing.T) {
	t.Run("successfully hashes password", func(t *testing.T) {
		var p password
		plaintext := "password123"

		err := p.Set(plaintext)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}

		if p.hash == nil {
			t.Error("expected hash to be set, but it was nil")
		}

		if p.plaintext == nil {
			t.Error("expected plaintext to be stored, but it was nil")
		}

		if *p.plaintext != plaintext {
			t.Errorf("expected plaintext to be %q, got %q", plaintext, *p.plaintext)
		}
	})

	t.Run("generates different hashes for same password", func(t *testing.T) {
		var p1, p2 password
		plaintext := "password123"

		p1.Set(plaintext)
		p2.Set(plaintext)

		// Hashes should be different due to salt
		hash1 := string(p1.hash)
		hash2 := string(p2.hash)

		if hash1 == hash2 {
			t.Error("expected different hashes for same password, but they were identical")
		}
	})
}

func TestPassword_Matches(t *testing.T) {
	t.Run("correct password matches", func(t *testing.T) {
		var p password
		plaintext := "password123"
		p.Set(plaintext)

		matches, err := p.Matches(plaintext)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}

		if !matches {
			t.Error("expected password to match, but it didn't")
		}
	})

	t.Run("incorrect password does not match", func(t *testing.T) {
		var p password
		p.Set("password123")

		matches, err := p.Matches("wrongpassword")
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}

		if matches {
			t.Error("expected password not to match, but it did")
		}
	})
}

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name      string
		email     string
		wantValid bool
	}{
		{
			name:      "valid email",
			email:     "test@police.gov.bz",
			wantValid: true,
		},
		{
			name:      "empty email",
			email:     "",
			wantValid: false,
		},
		{
			name:      "invalid format - no @",
			email:     "notanemail",
			wantValid: false,
		},
		{
			name:      "invalid format - no domain",
			email:     "test@",
			wantValid: false,
		},
		{
			name:      "valid email with subdomain",
			email:     "admin@training.police.gov.bz",
			wantValid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validator.New()
			ValidateEmail(v, tt.email)

			if tt.wantValid && !v.IsEmpty() {
				t.Errorf("expected email %q to be valid, but got errors: %v", tt.email, v.Errors)
			}

			if !tt.wantValid && v.IsEmpty() {
				t.Errorf("expected email %q to be invalid, but it was valid", tt.email)
			}
		})
	}
}

func TestValidatePasswordPlaintext(t *testing.T) {
	tests := []struct {
		name      string
		password  string
		wantValid bool
	}{
		{
			name:      "valid password",
			password:  "password123",
			wantValid: true,
		},
		{
			name:      "empty password",
			password:  "",
			wantValid: false,
		},
		{
			name:      "too short",
			password:  "pass",
			wantValid: false,
		},
		{
			name:      "exactly 8 characters",
			password:  "pass1234",
			wantValid: true,
		},
		{
			name:      "very long password",
			password:  "verylongpasswordthatismorethan72characterslongandshouldfailvalidationbecausebcrypthaslimit",
			wantValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validator.New()
			ValidatePasswordPlaintext(v, tt.password)

			if tt.wantValid && !v.IsEmpty() {
				t.Errorf("expected password to be valid, but got errors: %v", v.Errors)
			}

			if !tt.wantValid && v.IsEmpty() {
				t.Error("expected password to be invalid, but it was valid")
			}
		})
	}
}

func TestValidateUser(t *testing.T) {
	t.Run("valid user", func(t *testing.T) {
		user := &User{
			Email: "test@police.gov.bz",
		}
		user.Password.Set("password123")

		v := validator.New()
		ValidateUser(v, user)

		if !v.IsEmpty() {
			t.Errorf("expected user to be valid, but got errors: %v", v.Errors)
		}
	})

	t.Run("invalid email", func(t *testing.T) {
		user := &User{
			Email: "invalidemail",
		}
		user.Password.Set("password123")

		v := validator.New()
		ValidateUser(v, user)

		if v.IsEmpty() {
			t.Error("expected user to be invalid, but it was valid")
		}

		if _, exists := v.Errors["email"]; !exists {
			t.Error("expected error on email field, but it was not found")
		}
	})
}