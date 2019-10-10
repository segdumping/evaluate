package evaluate

import "container/list"

type stream struct {
	source   []interface{}
	position int
	length   int
}

func NewStream(source interface{}) *stream {
	ret := new(stream)
	switch i := source.(type) {
	case string:
		ret.fromString(i)
	case *list.List:
		ret.fromList(i)
	}

	ret.length = len(ret.source)
	return ret
}

func (s *stream) fromString(source string) {
	for _, c := range source {
		s.source = append(s.source, c)
	}
}

func (s *stream) fromList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		s.source = append(s.source, e.Value)
	}
}

func (s *stream) Read() interface{} {
	var character interface{}

	character = s.source[s.position]
	s.position += 1
	return character
}

func (s *stream) Rewind(amount int) {
	s.position -= amount
}

func (s stream) CanRead() bool {
	return s.position < s.length
}
