package tswitch

import "testing"

func TestTypeSwitchCheck(t *testing.T) {
	TypeSwitchCheck(21)
	TypeSwitchCheck("hello")
	TypeSwitchCheck(true)
}
