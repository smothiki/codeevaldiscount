package main

import (
	"testing"
)

func Testutils(t *testing.T) {
	if Hasgcd(8, 9) == true {
		t.Fatal("8 and 9 has no common GCD other than 1")
	}
	if Hasgcd(8, 4) == false {
		t.Fatal("8 and 4 has 2 as a common GCD other than 1")
	}

	if Wordcount("zenreach", vowels) != 3 {
		t.Fatal("expected 3")
	}
	if Wordcount("zenreach", consos) != 3 {
		t.Fatal("expected 5")
	}
	if Wordcount("zenreach", alpha) != 3 {
		t.Fatal("expected 8")
	}
}

func TestSSmatrix(t *testing.T) {
	products := []string{"iPad 2 - 4-pack"}
	customers := []string{"Jack Abraham", "John Evans"}
	matrix := SSmatrix(products, customers)
	test := []float64{6, 4.5}
	for i, mat := range matrix {
		if test[i] != mat.ss {
			t.Fatal("wrong ss product calculated")
		}
	}
}

func TestMaxSS(t *testing.T) {
	products := []string{"iPad 2 - 4-pack"}
	customers := []string{"Jack Abraham", "John Evans"}
	if MaxSS(products, customers) != float64(6) {
		t.Fatal("wrinog maxss product returned")
	}
}
