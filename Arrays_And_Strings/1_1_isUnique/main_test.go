package main

import "testing"

const (
	TRUE  = true
	FALSE = false
)

func TestUniqueCharacterString(t *testing.T) {
	unique := IsUnique("Spiderman")
	if have, want := unique, TRUE; have != want {
		t.Errorf("Have %t but wanted %t", unique, TRUE)
	}
}

func TestNonUniqueCharacterString(t *testing.T) {
	unique := IsUnique("happy")
	if have, want := unique, FALSE; have != want {
		t.Errorf("Have %t but wanted %t", unique, FALSE)
	}
}

func TestSingleCharacterString(t *testing.T) {
	unique := IsUnique("a")
	if have, want := unique, TRUE; have != want {
		t.Errorf("Have %t but wanted %t", unique, TRUE)
	}
}

func TestEmptyString(t *testing.T) {
	unique := IsUnique("")
	if have, want := unique, TRUE; have != want {
		t.Errorf("Have %t but wanted %t", unique, TRUE)
	}
}

func TestWordLargerThan128Characters(t *testing.T) {
	unique := IsUnique("kQDtRPswt0vXTy7WY7gXNo6UP5OpxPRjvRIECoeyNkhpGkhRWOkKFIghRSRVglUBGpAYYVB6yYUtasFXJszlSgUDYXOXPs0uY81uvL7YCEK26TAfagGZVVkC22tpwTnBQ")
	if have, want := unique, FALSE; have != want {
		t.Errorf("Have %t but wanted %t", unique, FALSE)
	}
}
