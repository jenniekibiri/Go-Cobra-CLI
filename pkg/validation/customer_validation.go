// pkg/validation/customer_validation.go
package validation

import (
    "regexp"
)

// ValidateEmail validates the format of an email address.
func ValidateEmail(email string) bool {
    // Simple regular expression for email validation.
    // You can use a more comprehensive regex for email validation.
    emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
    matched, _ := regexp.MatchString(emailPattern, email)
    return matched
}


