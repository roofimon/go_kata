package captcha

import (
	"errors"
	"strconv"
)

type Captcha struct {
	leftOperand  int
	operator     int
	rightOperand int
}

func New(leftOperand int, operator int, rightOperand int) (*Captcha, error) {
	c := Captcha{}
	if leftOperand > 9 {
		return nil, errors.New("Left operand value could not greater than 9")
	} else if leftOperand < 0 {
		return nil, errors.New("Left operand value could not less than 0")
	}
	c.leftOperand = leftOperand
	c.rightOperand = rightOperand
	c.operator = operator
	return &c, nil
}

func (c *Captcha) LeftOperand() string {
	return strconv.Itoa(c.leftOperand)
}
