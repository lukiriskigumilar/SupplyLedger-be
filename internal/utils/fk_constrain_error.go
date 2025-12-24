package utils

import "strings"

func IsFKConstraintError(err error) bool {
	return strings.Contains(err.Error(), "foreign key") ||
		strings.Contains(err.Error(), "constraint fails")
}
