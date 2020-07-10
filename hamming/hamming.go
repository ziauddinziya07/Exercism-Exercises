package hamming

import "errors"

// Distance returns the differences between the two Strads and error as nil when stards are of Equal length ....
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		
		return 0, errors.New("Unequal Strards")
	}
	c := 0
	for i := 0; i < len(a); i ++ {
		if a[i] != b[i] {
			c ++
		}
	}
	return c, nil
}
