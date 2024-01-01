package integers

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	desiredRes := 4

	if sum != desiredRes {
		t.Errorf("desired %d sum %d", sum, desiredRes)
	}
}
