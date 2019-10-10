package evaluate

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type TokenKind int

const (
	UNKNOWN  TokenKind = iota
	NUMERIC            //数字
	BOOLEAN            //布尔
	OPERATOR           //操作符
	STRING             //字符串
	VARIABLE           //变量
)

var tokens = [...]string{
	UNKNOWN:  "unknown",
	NUMERIC:  "numeric",
	BOOLEAN:  "boolean",
	OPERATOR: "operator",
	STRING:   "string",
	VARIABLE: "variable",
}

type Token struct {
	Kind  TokenKind
	Value interface{}
}

func (t *Token) String() string {
	return tokens[t.Kind]
}

func toToken(value string) (*Token, error) {
	switch {
	case IsOperator(value):
		v, _ := OpIdentifier(value)
		return &Token{
			Kind:  OPERATOR,
			Value: v,
		}, nil
	case isNumeric(value):
		v, err := strconv.ParseFloat(value, 64)
		return &Token{
			Kind:  NUMERIC,
			Value: v,
		}, err
	case isBoolean(value):
		v, err := strconv.ParseBool(toString(value))
		return &Token{
			Kind:  BOOLEAN,
			Value: v,
		}, err
	case isVariable(value):
		return &Token{
			Kind:  VARIABLE,
			Value: value,
		}, nil
	case isString(value):
		return &Token{
			Kind:  STRING,
			Value: toString(value),
		}, nil
	default:
		return nil, errors.New(fmt.Sprintf("value: %v cannot to token", value))
	}

	return nil, nil
}

func isNumeric(word string) bool {
	if len(word) == 0 {
		return false
	}

	if strings.HasPrefix(word, ".") {
		return false
	}

	for _, v := range word {
		switch {
		case v == '-', v == '.':
			continue
		case unicode.IsNumber(v):
			continue
		default:
			return false
		}
	}

	return true
}

func isVariable(word string) bool {
	for _, v := range word {
		switch {
		case unicode.IsLetter(v):
			continue
		case v == '_', v == '-', v == '.':
			continue
		default:
			return false
		}
	}

	return true
}

func isString(word string) bool {
	return strings.HasPrefix(word, "'") &&
		strings.HasSuffix(word, "'")
}

func isBoolean(word string) bool {
	switch word {
	case "true", "false":
		return true
	}

	return false
}

func toString(word string) string {
	return strings.TrimSuffix(strings.TrimPrefix(word, "'"), "'")
}
