package util

import (
	"encoding/base64"
	"testing"
)

// ── StrOrBase64Encoded ──────────────────────────────────────────────────────

func TestStrOrBase64Encoded_ValidBase64(t *testing.T) {
	original := "hello world"
	encoded := base64.StdEncoding.EncodeToString([]byte(original))
	got := StrOrBase64Encoded(encoded)
	if got != original {
		t.Errorf("got %q, want %q", got, original)
	}
}

func TestStrOrBase64Encoded_PlainString(t *testing.T) {
	plain := "not base64!!!"
	got := StrOrBase64Encoded(plain)
	if got != plain {
		t.Errorf("got %q, want %q", got, plain)
	}
}

func TestStrOrBase64Encoded_EmptyString(t *testing.T) {
	got := StrOrBase64Encoded("")
	// empty string is valid base64 decoding to empty → returns ""
	if got != "" {
		t.Errorf("got %q, want empty string", got)
	}
}

// ── B64StrToByte ────────────────────────────────────────────────────────────

func TestB64StrToByte_ValidInput(t *testing.T) {
	original := []byte("zpanel rocks")
	encoded := base64.StdEncoding.EncodeToString(original)
	got, err := B64StrToByte(encoded)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(got) != string(original) {
		t.Errorf("got %q, want %q", got, original)
	}
}

func TestB64StrToByte_InvalidInput(t *testing.T) {
	_, err := B64StrToByte("not!valid!base64")
	if err == nil {
		t.Error("expected error for invalid base64 input, got nil")
	}
}

// ── ByteToB64Str ────────────────────────────────────────────────────────────

func TestByteToB64Str_RoundTrip(t *testing.T) {
	data := []byte("round-trip test 🎉")
	encoded := ByteToB64Str(data)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if string(decoded) != string(data) {
		t.Errorf("round-trip failed: got %q, want %q", decoded, data)
	}
}

func TestByteToB64Str_EmptyInput(t *testing.T) {
	got := ByteToB64Str([]byte{})
	if got != "" {
		t.Errorf("expected empty string, got %q", got)
	}
}
