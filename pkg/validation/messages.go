package validation

func ErrorMessages() map[string]string {
	return map[string]string{
		"required": "The field is required",
		"email":    "Must have a valid email address",
		"min":      "Should be more than the limit",
		"max":      "Should be less the the limit",
	}
}
