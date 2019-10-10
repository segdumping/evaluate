package evaluate

import (
	"errors"
	"reflect"
	"unicode"
)

type EvaluableExpression struct {
	expressTree *ExpressToken
	parameters  Parameters
}

func NewEvaluableExpression(express string) (*EvaluableExpression, error) {
	if !isExpressValid(express) {
		return nil, errors.New("express invalid")
	}

	words := ExpressLexer(express)
	infix := infixToPostfix(words, "")
	tree, err := PostfixToBinaryTree(infix.Front())

	if err != nil {
		return nil, err
	}

	return &EvaluableExpression{
		expressTree: tree,
	}, nil
}

func (e *EvaluableExpression) Evaluate(parameters Parameters) (bool, error) {
	e.parameters = parameters

	r, err := e.eval(e.expressTree)
	if err != nil {
		return false, err
	}

	return r.(bool), nil
}

func (e *EvaluableExpression) eval(express *ExpressToken) (interface{}, error) {
	if express == nil {
		return nil, nil
	}

	var err error
	var left, right interface{}
	if express.token.Kind == OPERATOR {
		if express.Left != nil {
			if left, err = e.eval(express.Left); err != nil {
				return nil, err
			}
		}

		if express.Right != nil {
			if right, err = e.eval(express.Right); err != nil {
				return nil, err
			}
		}
	} else {
		return express.token, nil
	}

	pl, err := e.getParam(left)
	if err != nil {
		return nil, err
	}

	pr, err := e.getParam(right)
	if err != nil {
		return nil, err
	}

	if !e.isSameType(pl, pr) {
		return nil, errors.New("not same type")
	}

	op := express.token.Value.(Operator)
	return compare(op, pl, pr)

}

func (e *EvaluableExpression) getParam(value interface{}) (interface{}, error) {
	switch v := value.(type) {
	case *Token:
		if v.Kind != VARIABLE {
			return v.Value, nil
		}

		p, err := e.parameters.Get(v.Value.(string))
		if err != nil {
			return nil, err
		}

		return e.convert(p), nil
	}

	return value, nil
}

func (e *EvaluableExpression) convert(value interface{}) interface{} {
	switch v := value.(type) {
	case int16:
		return float64(v)
	case uint16:
		return float64(v)
	case int32:
		return float64(v)
	case uint32:
		return float64(v)
	case int64:
		return float64(v)
	case uint64:
		return float64(v)
	case int:
		return float64(v)
	case uint:
		return float64(v)
	}

	return value
}

func (e *EvaluableExpression) isSameType(left, right interface{}) bool {
	if left == nil {
		return true
	}

	lt := reflect.TypeOf(left).Kind()
	rt := reflect.TypeOf(right).Kind()

	return lt == rt
}

//校验表达式合法性
func isExpressValid(exp string) bool {
	for _, v := range exp {
		switch {
		case unicode.IsNumber(v):
			continue
		case unicode.IsLetter(v):
			continue
		case IsOpPrefix(v):
			continue
		case IsSkip(v):
			continue
		case IsPunct(v):
			continue
		default:
			return false
		}
	}

	return true
}
