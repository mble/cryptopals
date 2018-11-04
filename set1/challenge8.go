package set1

// DetectAes128Ecb detects if given data is encrypted in ECB mode
func DetectAes128Ecb(data []byte, blockSize int) bool {
	seen := make(map[string]struct{})
	for i := 0; i < len(data); i += blockSize {
		val := string(data[i : i+blockSize])
		if _, ok := seen[val]; ok {
			return true
		}
		seen[val] = struct{}{}
	}
	return false
}
