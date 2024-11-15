package usecases

import (
	"context"
	"hashtracker/internal/usecases"
	"regexp"
	"strings"
)

type ensValidator struct {
}

func NewENSValidator() usecases.ENSValidator {
	return &ensValidator{}
}

func (v *ensValidator) IsValid(ctx context.Context, domain string) bool {
	// Check if domain ends with .eth
	if !strings.HasSuffix(domain, ".eth") {
		return false
	}

	// Remove .eth suffix for further validation
	name := strings.TrimSuffix(domain, ".eth")

	// Check minimum length (3 characters)
	if len(name) < 3 {
		return false
	}

	// Check for invalid characters and patterns
	// Only alphanumeric and hyphens allowed
	validChars := regexp.MustCompile(`^[a-zA-Z0-9]+(-[a-zA-Z0-9]+)*$`)
	if !validChars.MatchString(name) {
		return false
	}

	return true
}
