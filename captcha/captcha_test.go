package captcha_test

import (
	"kata/captcha"
	"testing"
)

func TestLeftOperandShouldEquals_1_WhenInputIs_111(t *testing.T) {
	c, _ := captcha.New(1, 1, 1)
	lo := c.LeftOperand()
	if lo != "1" {
		t.Errorf("Expected %v but got %v", 1, lo)
	}
}

func TestLeftOperandShouldEquals_2_WhenInputIs_211(t *testing.T) {
	c, _ := captcha.New(2, 1, 1)
	lo := c.LeftOperand()
	if lo != "2" {
		t.Errorf("Expected %v but got %v", 2, lo)
	}
}

func TestReturnError_WhenInputIs_1011(t *testing.T) {
	message := "Left operand value could not greater than 9"
	_, err := captcha.New(10, 1, 1)
	if err.Error() != message {
		t.Errorf("Expected %v but got nothing", message)
	}
}

func TestReturnError_WhenInputIs_Minus111(t *testing.T) {
	message := "Left operand value could not less than 0"
	_, err := captcha.New(-1, 1, 1)
	if err.Error() != message {
		t.Errorf("Expected %v but got nothing", message)
	}
}
