package evaluate

import (
	"bytes"
	"container/list"
)

//表达式分词
func ExpressLexer(express string) *stream {
	l := list.New()
	stream := NewStream(express)
	for stream.CanRead() {
		c := stream.Read().(rune)
		if IsSkip(c) {
			continue
		}

		if IsOpPrefix(c) {
			stream.Rewind(1)
			s := readOperate(stream)
			l.PushBack(s)
		} else {
			stream.Rewind(1)
			s := readNormal(stream)
			l.PushBack(s)
		}
	}

	return NewStream(l)
}

//普通字符
func readNormal(stream *stream) string {
	var buf bytes.Buffer
	for stream.CanRead() {
		c := stream.Read().(rune)
		if IsSkip(c) {
			break
		}

		if IsOpPrefix(c) {
			stream.Rewind(1)
			break
		}

		buf.WriteRune(c)
	}

	return buf.String()
}

//运算符，最多匹配原则
func readOperate(stream *stream) string {
	var buf bytes.Buffer
	for stream.CanRead() {
		c := stream.Read().(rune)
		if IsSkip(c) {
			break
		}

		if !IsOpPrefix(c) {
			stream.Rewind(1)
			break
		}

		buf.WriteRune(c)
	}

	word := buf.String()
	if IsOperator(word) {
		return word
	} else {
		stream.Rewind(1)
		for i := buf.Len() - 1; i > 0; i-- {
			buf.Truncate(i)
			word := buf.String()
			if IsOperator(word) {
				return word
			}

			stream.Rewind(1)
		}
	}

	return word
}
