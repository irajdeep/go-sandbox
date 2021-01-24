package performance

// Avoid
func copyIntoSliceAndMap1(biggy []string) (a []string, b map[string]struct{}) {
	b = make(map[string]struct{})

	for _, item := range biggy {
		a = append(a, item)
		b[item] = struct{}{}
	}

	return a, b
}

// Prefer
func copyIntoSliceAndMap2(biggy []string) (a []string, b map[string]struct{}) {
	b = make(map[string]struct{}, len(biggy))
	a = make([]string, len(biggy))

	copy(a, biggy)
	for _, item := range biggy {
		b[item] = struct{}{}
	}
	return a, b
}
