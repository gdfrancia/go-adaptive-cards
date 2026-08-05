package cards

// BoolPtr returns pointer to bool
func BoolPtr(b bool) *bool {
	return &b
}

// IntPtr retruns pointer to int
func IntPtr(i int) *int {
	return &i
}

// IntString returns pointer to string
func IntString(s string) *string {
	return &s
}

// TruePtr returns pointer to true
func TruePtr() *bool {
	return BoolPtr(true)
}

// FalsePtr returns pointer to false
func FalsePtr() *bool {
	return BoolPtr(false)
}
