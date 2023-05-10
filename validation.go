package govalidation

func Lambda(name string) bool {
	if name != "" {
		return true
	} else {
		return false
	}
}
