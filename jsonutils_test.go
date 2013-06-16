package jsonutils

import (
	"testing"
)

func BenchmarkPrintGo(b *testing.B) {
	b.StopTimer()
	j := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

	f, _ := ParseJson(j)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		PrintGo(f, "Test")
	}
}
