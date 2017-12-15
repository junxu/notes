package reverse_integer

import "testing"

func TestReverseInteger(t *testing.T) {
	tests := []int{12345, 67455, 33344}
	results := []int{54321, 55476, 44333}

	for i, v := range(tests) {
		if ret := ReverseInteger(v); results[i] != ret {
			t.Fatalf("case %d failed, actaul is %d, expect is %d\n", i, ret, results[i])
		}
	}
}
