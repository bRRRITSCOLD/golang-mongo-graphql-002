package utils

func PointerBool(boolean bool) *bool {
	b := boolean
	return &b
}
