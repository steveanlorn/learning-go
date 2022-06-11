package equal

// Slice ...
func SliceByte(a []byte, b []byte) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
