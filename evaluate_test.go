package evaluate

import (
	"math/rand"
	"testing"
)

type EvaluateTest struct {
	Express    string
	Parameters MapParameters
	Expect     bool
}

func TestEvaluate(t *testing.T) {
	evaluateTests := []EvaluateTest{
		{
			Express: "!true",
			Expect:  true,
		},

		{
			Express: "10 > 5",
			Expect:  true,
		},

		{
			Express: "-1 < -5",
			Expect:  true,
		},

		{
			Express: "1.1 < 1.1",
			Expect:  true,
		},

		{
			Express: "1.1 == 1.1",
			Expect:  true,
		},

		{
			Express: "1.1 != 1.1",
			Expect:  true,
		},

		{
			Express:    "value > -5",
			Parameters: MapParameters{"value": -4},
			Expect:     true,
		},

		{
			Express:    "age > 5",
			Parameters: MapParameters{"age": 4},
			Expect:     true,
		},

		{
			Express:    "age >= 5",
			Parameters: MapParameters{"age": 5},
			Expect:     true,
		},

		{
			Express:    "age < 5",
			Parameters: MapParameters{"age": 5},
			Expect:     true,
		},

		{
			Express:    "age <= 5",
			Parameters: MapParameters{"age": 5},
			Expect:     true,
		},

		{
			Express:    "result == true",
			Parameters: MapParameters{"result": true},
			Expect:     true,
		},

		{
			Express:    "name == 'liyi'",
			Parameters: MapParameters{"name": "liyi"},
			Expect:     true,
		},

		{
			Express:    "name != 'liyi'",
			Parameters: MapParameters{"name": "liyi"},
			Expect:     true,
		},

		{
			Express:    "age > 5 && age < 20",
			Parameters: MapParameters{"age": 10},
			Expect:     true,
		},

		{
			Express:    "age > 5 || age == 3",
			Parameters: MapParameters{"age": 2},
			Expect:     true,
		},

		{
			Express:    "!(age > 5 || age == 3)",
			Parameters: MapParameters{"age": 2},
			Expect:     true,
		},

		{
			Express:    "(age > 5 || age == 3) && (name == 'liyi' || name == 'liming')",
			Parameters: MapParameters{"age": 3, "name": "liyi"},
			Expect:     true,
		},
	}

	for _, v := range evaluateTests {
		eval, err := NewEvaluableExpression(v.Express)
		if err != nil {
			t.Logf("new evaluate error: %v, test: %v", err, v)
			continue
		}

		r, err := eval.Evaluate(v.Parameters)
		t.Logf("result: %v, expect: %v, error: %v, express: %s, paramers: %v", r, v.Expect, err, v.Express, v.Parameters)
	}
}

func BenchmarkEvaluable(b *testing.B) {
	eval, err := NewEvaluableExpression("age > 5 && age < 10")
	if err != nil {
		b.Log(err)
		return
	}

	param := MapParameters{}
	for i := 0; i < b.N; i++ {
		param["age"] = rand.Intn(10)
		eval.Evaluate(param)
	}
}
