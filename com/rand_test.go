package com

import "testing"

func TestRandLetter(t *testing.T) {
	t.Log(RandLetter(4))
}

func TestRandDigital(t *testing.T) {
	t.Log(RandDigital(4))
}
