package helpers

import (
	"errors"
	"strconv"
	"strings"
)

// FullNum turns potentially exponential string numbers into int strings
func FullNum(s string) (string, error) {
	if strings.Contains(s, "e-") {
		return "", errors.New("can't have negative exponents")
	}
	// get number and exponent
	parts := strings.Split(s, "e+")
	if len(parts) == 1 {
		parts = append(parts, "0")
	}

	num := parts[0]
	exp, err := strconv.Atoi(parts[1])
	if err != nil {
		// exponent wasn't an integer
		return "", err
	}

	// shift decimal
	nums := strings.Split(num, ".")
	var pre string
	if len(nums) == 1 {
		// no decimal
		pre = nums[0]
	} else {
		if len(nums[1]) > exp {
			// number had way too many decial places
			nums[1] = nums[1][:exp]
		}
		pre = strings.Join(nums, "")
		exp = exp - len(nums[1])
	}

	// add zeros from exponent
	return pre + strings.Repeat("0", exp), nil
}
