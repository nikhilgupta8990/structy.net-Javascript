package reverselist_test

import (
	"reflect"
	"testing"

	reverselist "github.com/ju-popov/structy.net/015-reverse-list"
)

func TestSimple(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := reverselist.Simple(tc.head)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for test '%v' is: '%v', but the actual result is: '%v'", tc.name, tc.expected, actual)
			}
		})
	}
}

func benchmarkSimple(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		reverselist.Simple(tc.head)
	}
}

func BenchmarkSimple000(b *testing.B) { benchmarkSimple(b, testCases[0]) }
func BenchmarkSimple001(b *testing.B) { benchmarkSimple(b, testCases[1]) }
func BenchmarkSimple002(b *testing.B) { benchmarkSimple(b, testCases[2]) }
