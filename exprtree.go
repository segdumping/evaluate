package evaluate

import (
	"container/list"
	"errors"
	"github.com/segdumping/shared/stack"
)

var (
	expressError = errors.New("express syntax error")
)

type ExpressToken struct {
	token *Token
	Left  *ExpressToken
	Right *ExpressToken
}

//后缀转表达式树
func PostfixToBinaryTree(element *list.Element) (*ExpressToken, error) {
	var current *ExpressToken
	var left, right *ExpressToken
	stack := stack.Stack(make([]interface{}, 0))

	for ; element != nil; element = element.Next() {
		v := element.Value.(string)
		token, err := toToken(v)
		if err != nil {
			return nil, err
		}

		if IsOperator(v) { //运算符
			current = &ExpressToken{}
			current.token = token

			if e := stack.Pop(); e != nil {
				left = e.(*ExpressToken)
			} else {
				return nil, expressError
			}

			if !IsUnary(v) {
				if e := stack.Pop(); e != nil {
					right = e.(*ExpressToken)
				} else {
					return nil, expressError
				}
			} else {
				right = &ExpressToken{token: &Token{}}
			}

			current.Left = right
			current.Right = left
			stack.Push(current)
		} else {
			current = &ExpressToken{}
			current.token = token
			stack.Push(current)
		}
	}

	return stack.Top().(*ExpressToken), nil
}

//中缀转后缀
func infixToPostfix(stream *stream) *list.List {
	l := list.New()
	stacks := make([]stack.Stack, 6)

	for stream.CanRead() {
		element := stream.Read().(string)
		if element == ")" {
			break
		} else if element == "(" { //括号中的表达式递归处理
			l.PushBackList(infixToPostfix(stream))
		} else {
			if ident, ok := OpIdentifier(element); !ok { //普通字符串
				l.PushBack(element)
			} else { //高优先级的运算符全部弹出
				for i := 1; i <= ident.Precedence(); i++ {
					for !stacks[i].Empty() {
						l.PushBack(stacks[i].Pop().(string))
					}
				}

				stacks[ident.Precedence()].Push(element)
			}
		}
	}

	for i := range stacks {
		for !stacks[i].Empty() {
			l.PushBack(stacks[i].Pop().(string))
		}
	}

	return l
}
