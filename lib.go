package chroma

import (
	"io"
)

var grayCodes = [4]int{0, 1, 3, 2}

func gray(i int) int {
	return grayCodes[i%3]
}

func quantize(thresholds []float32, x float32) int {
	i := 0
	for i < len(thresholds) && x < thresholds[i] {
		i++
	}
	return i
}

type Fingerprinter struct {
	istrm *io.Reader
}
