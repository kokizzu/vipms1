package model

import (
	"strings"
	"testing"
)

func TestVipValidationErrorAcceptsValidVip(t *testing.T) {
	vip := Vip{
		Name:     "Jane Visitor",
		Country:  "Indonesia",
		ETA:      "2026-07-14 10:30:00",
		PhotoUrl: "https://example.com/jane.jpg",
	}

	if got := vip.ValidationError(); got != "" {
		t.Fatalf("ValidationError() = %q, want empty", got)
	}
}

func TestVipValidationErrorRejectsMissingETA(t *testing.T) {
	vip := Vip{
		Name:    "Jane Visitor",
		Country: "Indonesia",
	}

	if got := vip.ValidationError(); got != "eta must use format YYYY-MM-DD HH:MM:SS" {
		t.Fatalf("ValidationError() = %q", got)
	}
}

func TestVipValidationErrorRejectsInvalidETA(t *testing.T) {
	vip := Vip{
		Name:    "Jane Visitor",
		Country: "Indonesia",
		ETA:     "tomorrow",
	}

	if got := vip.ValidationError(); got != "eta must use format YYYY-MM-DD HH:MM:SS" {
		t.Fatalf("ValidationError() = %q", got)
	}
}

func TestVipValidationErrorRejectsOverlongPhotoURL(t *testing.T) {
	vip := Vip{
		Name:     "Jane Visitor",
		Country:  "Indonesia",
		ETA:      "2026-07-14 10:30:00",
		PhotoUrl: "https://example.com/" + strings.Repeat("x", 121),
	}

	if got := vip.ValidationError(); got != "photo_url must be maximum 120 characters" {
		t.Fatalf("ValidationError() = %q", got)
	}
}
