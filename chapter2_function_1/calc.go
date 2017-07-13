// calc
package main

import "errors"

func Add(a, b int) (result int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("Should be non-negative numbers!")
		return
	}

	return a + b, nil
}
