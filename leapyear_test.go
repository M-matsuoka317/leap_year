package main

import (
    "testing"
)

func TestLeapYear01(t *testing.T) {
    expected := Leapyear(1)
    actual := false
    if expected != actual {
        t.Errorf("Error")
    }
}

func TestLeapYear02(t *testing.T) {
    expected := Leapyear(4)
    actual := true
    if expected != actual {
        t.Errorf("Error")
    }
}

// sample