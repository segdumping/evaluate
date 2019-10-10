package evaluate

import "errors"

type Parameters interface {
	Get(string) (interface{}, error)
}

type MapParameters map[string]interface{}

func (m MapParameters) Get(name string) (interface{}, error) {
	v, ok := m[name]
	if !ok {
		return nil, errors.New("parameter " + name + "not found")
	}

	return v, nil
}
