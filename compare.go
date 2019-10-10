package evaluate

import (
	"errors"
	"fmt"
	"reflect"
)

var cmpHandlers = map[Operator]func(left, right interface{}) (interface{}, error){
	NEG:  neg,
	GTR:  gtr,
	LSS:  lss,
	GEQ:  geq,
	LEQ:  leq,
	NEQ:  neq,
	EQL:  eql,
	LAND: land,
	LOR:  lor,
}

func compare(op Operator, left, right interface{}) (interface{}, error) {
	h, ok := cmpHandlers[op]
	if !ok {
		return nil, errors.New(fmt.Sprintf("invalid operator: %v", op))
	}

	return h(left, right)
}

func neg(left, right interface{}) (interface{}, error) {
	if v, ok := right.(float64); ok {
		return -1 * v, nil
	} else {
		return nil, errors.New(fmt.Sprintf("op (-) invalid param: %v", right))
	}
}

func gtr(left, right interface{}) (interface{}, error) {
	if isFloat64(left) && isFloat64(right) {
		return left.(float64) > right.(float64), nil
	}

	return left.(string) > right.(string), nil
}

func lss(left, right interface{}) (interface{}, error) {
	if isFloat64(left) && isFloat64(right) {
		return left.(float64) < right.(float64), nil
	}

	return left.(string) < right.(string), nil
}

func geq(left, right interface{}) (interface{}, error) {
	if isFloat64(left) && isFloat64(right) {
		return left.(float64) >= right.(float64), nil
	}

	return left.(string) >= right.(string), nil
}

func leq(left, right interface{}) (interface{}, error) {
	if isFloat64(left) && isFloat64(right) {
		return left.(float64) <= right.(float64), nil
	}

	return left.(string) <= right.(string), nil
}

func neq(left, right interface{}) (interface{}, error) {
	return !reflect.DeepEqual(left, right), nil
}

func eql(left, right interface{}) (interface{}, error) {
	return reflect.DeepEqual(left, right), nil
}

func land(left, right interface{}) (interface{}, error) {
	return left.(bool) && right.(bool), nil
}

func lor(left, right interface{}) (interface{}, error) {
	return left.(bool) || right.(bool), nil
}

func isFloat64(value interface{}) bool {
	switch value.(type) {
	case float64:
		return true
	}

	return false
}
