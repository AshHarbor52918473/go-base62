package base62

import "testing"

func TestEncodeDecodeRoundTrip(t *testing.T) {
	for _, n := range []uint64{0, 1, 61, 62, 12345, 1 << 40} {
		if got := Decode(Encode(n)); got != n {
			t.Errorf("round trip of %d gave %d (encoded %q)", n, got, Encode(n))
		}
	}
}

func TestEncodeKnownValues(t *testing.T) {
	cases := map[uint64]string{0: "0", 10: "a", 61: "Z", 62: "10"}
	for n, want := range cases {
		if got := Encode(n); got != want {
			t.Errorf("Encode(%d) = %q, want %q", n, got, want)
		}
	}
}
