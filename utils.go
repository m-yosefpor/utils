package utils

// isIn checks if a string array contains an element
func IsIn(s string, list []string) bool {
	for _, l := range list {
		if s == l {
			return true
		}
	}
	return false
}

// difference returns the elements in `a` that aren't in `b`.
func ListDifference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
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
