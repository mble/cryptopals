package set1

import (
	"encoding/base64"
	"io/ioutil"
	"strings"
	"testing"
)

func TestDetectAes128Ecb(t *testing.T) {
	all, _ := ioutil.ReadFile("testdata/8.txt")
	for i, hexString := range strings.Split(string(all), "\n") {
		decodedData, _ := base64.StdEncoding.DecodeString(hexString)
		if DetectAes128Ecb(decodedData, 16) {
			t.Logf("Line %d is encrypted with ECB", i+1)
		}
	}
}
