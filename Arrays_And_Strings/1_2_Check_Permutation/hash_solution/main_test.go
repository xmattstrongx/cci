package main

import "testing"

const (
	TRUE  = true
	FALSE = false
)

func TestValidPermutationCaseInsensitive(t *testing.T) {
	isPermutation := IsPermutation("dog", "god", false)
	if have, want := isPermutation, TRUE; have != want {
		t.Errorf("Have %t but wanted %t", isPermutation, TRUE)
	}
}

func TestValidPermutationCaseSensitive(t *testing.T) {
	isPermutation := IsPermutation("dog", "GOD", true)
	if have, want := isPermutation, FALSE; have != want {
		t.Errorf("Have %t but wanted %t", isPermutation, FALSE)
	}
}

func TestInvalidPermutation(t *testing.T) {
	isPermutation := IsPermutation("dog", "godd", false)
	if have, want := isPermutation, FALSE; have != want {
		t.Errorf("Have %t but wanted %t", isPermutation, FALSE)
	}
}

func TestStringAgainstEmptyString(t *testing.T) {
	isPermutation := IsPermutation("dog", "", false)
	if have, want := isPermutation, FALSE; have != want {
		t.Errorf("Have %t but wanted %t", isPermutation, FALSE)
	}
}

func TestAgainstEmptyStrings(t *testing.T) {
	isPermutation := IsPermutation("", "", false)
	if have, want := isPermutation, TRUE; have != want {
		t.Errorf("Have %t but wanted %t", isPermutation, TRUE)
	}
}

func BenchmarkIsPermutation1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchmarkIsPermutation("dog", "god", false, b)
	}
}

func BenchmarkIsPermutation2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchmarkIsPermutation("aassddffgghhjjkkll", "llkkjjhhggffddssaa", false, b)
	}
}

func BenchmarkIsPermutation3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchmarkIsPermutation("aassddffgghhjjkklls", "llkkjjhhggffddssaa", false, b)
	}
}

func BenchmarkIsPermutation4(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchmarkIsPermutation("aassddffgghhjjkkll", "llkkjjhhggffddssAA", true, b)
	}
}

func benchmarkIsPermutation(string1 string, string2 string, caseSentive bool, b *testing.B) {
	IsPermutation(string1, string2, caseSentive)
}

// go test -bench=.
// BenchmarkIsPermutation1-8   	  500000	      2193 ns/op
// BenchmarkIsPermutation2-8   	  200000	      7821 ns/op
// BenchmarkIsPermutation3-8   	 5000000	       336 ns/op
// BenchmarkIsPermutation4-8   	  300000	      5332 ns/op