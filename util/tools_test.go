package util

import (
	"testing"
)

func init() {
	Seed(0)
}
func TestTools(t *testing.T) {
	//t.Log(RandIntRangeBetween(1000, 10000))
	//for i := 0; i < 5; i++ {
	//	num := RandIntRangeRand(5)
	//	randomFloat := RandFloatRangeRand(4)
	//	t.Log(num)
	//	t.Log(randomFloat)
	//}
	//for i := 0; i < 10; i++ {
	//	t.Log(RandDate("2020-04-29", "2020-05-01", "datetime"))
	//	t.Log(RandDate("2020-04-29", "2020-05-01", "date"))
	//}
	t.Log(RandEnum("a","b","c"))
}
func BenchmarkRandDate(b *testing.B) {
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		RandDate("2010-04-29", "2020-05-01", "date")
	}
}