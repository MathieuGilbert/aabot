package helpers

import "testing"

type Case struct {
	In     string
	Expect string
}

func TestFullNum(t *testing.T) {
	success := []*Case{
		// int
		&Case{In: "5", Expect: "5"},
		// decimal
		&Case{In: "3.4567", Expect: "3"},
		// garbage decimals
		&Case{In: "1.66625559592783679668000772712315460334794e+21", Expect: "1666255595927836796680"},
		// int with exponent
		&Case{In: "1e+21", Expect: "1000000000000000000000"},
		// decimal with exponent
		&Case{In: "9.87654321e+21", Expect: "9876543210000000000000"},
		// decimal with exponent, no zeros needed
		&Case{In: "9.876543210123456789012e+21", Expect: "9876543210123456789012"},
		&Case{In: "1.5e+21", Expect: "1500000000000000000000"},
		&Case{In: "2.497388534418055446234e+21", Expect: "2497388534418055446234"},
	}

	for _, c := range success {
		if n, _ := FullNum(c.In); n != c.Expect {
			t.Errorf("Expected: %v\nGot: %v\n", c.Expect, n)
		}
	}

	fail := []*Case{
		// should never happen
		&Case{In: "5e+1.5"},
		&Case{In: "1.23456e-5"},
	}

	for _, c := range fail {
		if n, err := FullNum(c.In); err == nil {
			t.Errorf("Expected an error from %v\nGot: %v\n", c.In, n)
		}
	}
}
