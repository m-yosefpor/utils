package utils

// isIn checks if a string array contains an element
func IsIn[T comparable](s T, list []T) bool {
	for _, l := range list {
		if s == l {
			return true
		}
	}
	return false
}

// difference returns the elements in `a` that aren't in `b`.
func ListDifference[T comparable](a, b []T) []T {
	mb := make(map[T]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []T
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

// IsErrorMessage check two error if two errors have the same err.Error() and avoiding nil pointer derefrence issue.
func IsErrorMessage(err, target error) bool {
	if target == nil || err == nil {
		return err == target
	}
	return err.Error() == target.Error()
}
