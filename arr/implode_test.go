package arr

import (
	"testing"
)

func TestImplode(t *testing.T) {
	//var s = []int{1, 2, 3, 9, 2, 5, 6}
	var b = []byte{1, 2, 3, 4, 5, 6}
	str, err := Implode(b, ",")
	t.Log(err)
	t.Log(str)
}

func BenchmarkImplode(b *testing.B) {
	var s = []int{1, 2, 1, 2, 3, 9, 2, 5, 6, 9, 1, 2, 3, 9, 2, 5, 6, 5, 6, 11, 2, 3, 9, 2, 5, 6, 2, 3, 9, 2, 5, 6}
	for i := 0; i < b.N; i++ {
		_, _ = Implode(s, ",")
	}
}
