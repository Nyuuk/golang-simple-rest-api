package helpers

import "strings"

func IsDuplicateKeyError(err error) bool {
	return strings.Contains(err.Error(), "duplicate key value violates unique constraint")
}

func IsNullConstraintError(err error) bool {
	return strings.Contains(err.Error(), "null value in column")
}
