package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "HelloWorld(rafli)",
			request: "rafli",
		},
		{
			name:    "HelloWorld(muhammad)",
			request: "muhammad",
		},
		{
			name:    "HelloWorld(x)",
			request: "x",
		},
		{
			name:    "HelloWorld(z)",
			request: "z",
		},
		{
			name:    "HelloWorld(w)",
			request: "w",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.name)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("rafli", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("rafli")
		}
	})
	b.Run("muhammad", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("muhammad")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; 0 < b.N; i++ {
		HelloWorld("rafli")
	}
}

func BenchmarkHelloWorldMuhammad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("muhammad")
	}
}

func TestHelloWorldTable(t *testing.T) {

	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(rafli)",
			request:  "rafli",
			expected: "hello rafli",
		},
		{
			name:     "HelloWorld(muhammad)",
			request:  "muhammad",
			expected: "hello muhammad",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {

	t.Run("rafli", func(t *testing.T) {
		result := HelloWorld("rafli")
		require.Equal(t, "hello rafli", result, "must be 'hello rafli'")
	})

	t.Run("muhammad", func(t *testing.T) {
		result := HelloWorld("muhammad")
		require.Equal(t, "muhammad", result, "must be 'hello muhammad'")
	})
}

func TestMain(m *testing.M) {

	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {

	if runtime.GOOS == "linux" {
		t.Skip("can not run on mac OS")
	}

	result := HelloWorld("ibrahimx")
	require.Equal(t, "hello ibrahim", result, "must be 'ibrahim'")

	fmt.Println("TestSkip done")
}

func TestHelloWorldRequire(t *testing.T) {

	result := HelloWorld("ibrahimx")
	require.Equal(t, "hello ibrahim", result, "must be 'ibrahim'")

	fmt.Println("TestHelloWorldRequire done")
}

func TestHelloWorldAssert(t *testing.T) {

	result := HelloWorld("ibrahimx")
	assert.Equal(t, "hello ibrahim", result, "must be 'ibrahim'")

	fmt.Println("TestHelloWorldAssert done")
}

func TestHelloWorld(t *testing.T) {

	result := HelloWorld("rafli")

	if result != "hello rafli" {
		// error
		// panic("result is not 'rafli' ")
		// t.Fail()
		t.Error("result is not 'rafli'")
	}

	fmt.Println("TestHelloWorld done")
}

func TestHelloDunia(t *testing.T) {

	result := HelloWorld("muhammad")

	if result != "hello 'daud'" {

		// t.FailNow()
		t.Fatal("must be hello muhammad")
	}

	fmt.Println("TestHelloDunia done")
}
