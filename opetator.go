package evaluate

type Operator int

const (
	ILLEGAL Operator = iota
	NEG              // -
	NOT              // !
	LAND             // &&
	LOR              // ||
	EQL              // ==
	LSS              // <
	GTR              // >
	NEQ              // !=
	LEQ              // <=
	GEQ              // >=
	LPAREN           // (
	RPAREN           // )
)

//运算符前缀
var opPrefix = map[rune]bool{
	'&': true,
	'|': true,
	'>': true,
	'<': true,
	'-': true,
	'(': true,
	')': true,
	'=': true,
	'!': true,
}

//运算符
var operators = map[string]Operator{
	"-":  NEG,
	">":  GTR,
	"<":  LSS,
	">=": GEQ,
	"<=": LEQ,
	"!=": NEQ,
	"==": EQL,
	"&&": LAND,
	"||": LOR,
	"!":  NOT,
	"(":  LPAREN,
	")":  RPAREN,
}

//忽略字符
var skips = map[rune]bool{
	' ': true,
}

//标点
var puncts = map[rune]bool{
	'\'': true,
	'.':  true,
}

//运算符优先级，数字越小优先级越高
func (o Operator) Precedence() int {
	switch o {
	case NEG, NOT:
		return 1
	case GTR, LSS, GEQ, LEQ, NEQ, EQL:
		return 3
	case LAND:
		return 4
	case LOR:
		return 5
	}

	return 0
}

func IsOpPrefix(character rune) bool {
	_, ok := opPrefix[character]
	return ok
}

func IsOperator(word string) bool {
	_, ok := operators[word]
	return ok
}

func IsSkip(character rune) bool {
	_, ok := skips[character]
	return ok
}

func IsPunct(character rune) bool {
	_, ok := puncts[character]
	return ok
}

//一元运算符
func IsUnary(word string) bool {
	switch word {
	case "-", "!":
		return true
	}

	return false
}

func OpIdentifier(word string) (Operator, bool) {
	v, ok := operators[word]

	return v, ok
}
