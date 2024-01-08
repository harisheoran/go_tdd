package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := repeat("a", 5)
	expexted := "aaaaa"
	if repeated != expexted {
		t.Errorf("expected %q repeated %q", expexted, repeated)
	}
}

func benchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeat("a", 5)
	}
}
