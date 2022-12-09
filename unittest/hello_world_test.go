package unittest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Randy")

	if result != "Selamat Datang Randy" {
		/**
		* Error
		**/

		// t.Fail()                                // Unit test error, proses akan tetap berjalan
		// t.FailNow()                             // Unit test error, proses unit test akan langsung berhenti
		// t.Error("Result is not 'Hello Randy' ") // Sama seperti Log Error, dan secara otomatis memanggil function t.Fail()
		t.Fatal("Result is not 'Selamat Datang Randy'") // Sama spt t.Error(), proses unit test akan langsung berhenti
	}
	fmt.Println("String ini bisa diakses setelah 't.Fail() dan t.Error()'")
	fmt.Println("String ini tidak akan bisa diakses setelah 't.FailNow() dan t.Fatal()'")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Randy")
	assert.Equal(t, "Selamat Datang Randy", result)
	fmt.Println("Dieksekusi")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Randy")
	require.Equal(t, "Selamat Datang Randy", result)
	fmt.Println("Dieksekusi")
}

/*
* Pada SubTest, kita bisa menjalankan 2 atau lebih test yang sama pada satu function
 */
func TestSubTest(t *testing.T) {
	t.Run("Randy", func(t *testing.T) {
		result := HelloWorld("Randy")
		require.Equal(t, "Selamat Datang Randy", result, "Result must be 'Selamat Datang Randy'")
	})
	t.Run("Wiratama", func(t *testing.T) {
		result := HelloWorld("Wiratama")
		require.Equal(t, "Selamat Datang Wiratama", result, "Result must be 'Selamat Datang Wiratama'")
	})
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Randy",
			request:  "Randy",
			expected: "Selamat Datang Randy",
		},
		{
			name:     "Wiratama",
			request:  "Wiratama",
			expected: "Selamat Datang Wiratama",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Randy")
	}
}

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Randy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Randy")
		}
	})
	b.Run("Wiratama", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Wiratama")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	// deklarasi variable
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Randy",
			request: "Randy",
		},
		{
			name:    "Wiratama",
			request: "Wiratama",
		},
	}

	for _, data := range benchmarks {
		b.Run(data.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(data.request)
			}
		})
	}
}
