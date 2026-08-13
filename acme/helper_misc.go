package acme

// stringSlice converts an interface slice to a string slice.
func stringSlice(src []any) []string {
	var dst []string
	for _, v := range src {
		dst = append(dst, v.(string))
	}
	return dst
}
