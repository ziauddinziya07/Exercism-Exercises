package hamming

import "errors"

func Distance(a, b string) (int, error) {

	if len(a) == len(b) {
		c := 0
		for i := 0; i < len(a); i ++ {
			if a[i] != b[i] {
				c ++
			}
		}
		return c, nil
	}
	return 0, errors.New("")
}
