package main

import "testing"

/**
测试数据，在遍历大数据量时，range 性能比 for 差很多

	goos: darwin
	goarch: arm64
	cpu: Apple M4
	BenchmarkRangeHiPerformance
	BenchmarkRangeHiPerformance-10     	  954213	      1238 ns/op
	BenchmarkRangeLowPerformance
	BenchmarkRangeLowPerformance-10    	     408	   2899801 ns/op
	PASS

	Process finished with the exit code 0
*/

func CreateABigSlice(count int) [][4096]int {

	ret := make([][4096]int, count)
	for i := 0; i < count; i++ {
		ret[i] = [4096]int{}
	}

	return ret
}

func BenchmarkRangeHiPerformance(b *testing.B) {

	v := CreateABigSlice(1 << 12)

	for i := 0; i < b.N; i++ {
		var tmp [4096]int
		for k := 0; k < len(v); k++ {
			tmp = v[k]
		}
		_ = tmp
	}
}

func BenchmarkRangeLowPerformance(b *testing.B) {

	v := CreateABigSlice(1 << 12)

	for i := 0; i < b.N; i++ {

		var tmp [4096]int
		for _, e := range v {
			tmp = e
		}
		_ = tmp
	}
}
