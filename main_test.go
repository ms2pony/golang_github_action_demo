package main

import "testing"

// 测试，猫是会喵的
func TestCat(t *testing.T) {
	saying := Cat()
	if saying != "Miao~~~" {
		t.Errorf("Cat say %s", saying)
	}
}
