package RowBinary

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func reverseBytesOriginal(target []byte) []byte {
	a := bytes.Split(target, []byte{'.'})

	l := len(a)
	for i := 0; i < l/2; i++ {
		a[i], a[l-i-1] = a[l-i-1], a[i]
	}

	return bytes.Join(a, []byte{'.'})
}

func TestReverseInplace(t *testing.T) {
	assert := assert.New(t)
	table := []string{
		"carbon.agents.carbon-clickhouse.graphite1.tcp.metricsReceived",
		"",
		".",
		"carbon..xx",
		".hello..world.",
	}

	for i := 0; i < len(table); i++ {
		x := []byte(table[i])
		y := []byte(table[i])
		z := reverseBytesOriginal(x)
		reverseMetricInplace(y)
		assert.Equal(string(z), string(y))
	}
}

func BenchmarkReverseOriginal(b *testing.B) {
	m := []byte("carbon.agents.carbon-clickhouse.graphite1.tcp.metricsReceived")
	var a []byte

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		a = reverseBytesOriginal(m)
		copy(m, a)
	}
}

func BenchmarkMetricInplace(b *testing.B) {
	m := []byte("carbon.agents.carbon-clickhouse.graphite1.tcp.metricsReceived")
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		reverseMetricInplace(m)
	}
}
