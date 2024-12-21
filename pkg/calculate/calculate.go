package calculate

import (
	"strconv"
)

func priority(elem string) int {
	res := 0
	switch {
	case elem == "(" || elem == ")":
		res = 1
	case elem == "+" || elem == "-":
		res = 2
	case elem == "*" || elem == "/":
		res = 3
	}
	return res
}

func functions(x, y float64, s string) float64 {
	var res float64
	switch s {
	case "+":
		res = x + y
	case "-":
		res = x - y
	case "*":
		res = x * y
	case "/":
		res = x / y
	}
	return res
}

func Calc(expression string) (float64, error) {

	if expression == "" {
		return 0, ErrInvalidExpression
	}

	out := []string{}
	stack := []string{}
	for _, r := range expression {
		s := string(r)
		if _, err := strconv.Atoi(s); err == nil {
			out = append(out, s)
		} else if s != "(" && s != ")" {
			for {
				if !(len(stack) > 0 && priority(stack[len(stack)-1]) >= priority(s)) {
					break
				}
				out = append(out, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, s)
		} else if s == "(" {
			stack = append(stack, s)
		} else {
			for {
				if stack[len(stack)-1] == "(" {
					break
				}
				out = append(out, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		}
	}

	for i := len(stack) - 1; i > -1; i-- {
		out = append(out, stack[i])
	}

	new_stack := []float64{}
	for _, elem := range out {
		if i, err := strconv.Atoi(elem); err == nil {
			new_stack = append(new_stack, float64(i))
		} else {
			if len(new_stack) < 2 {
				return 0, ErrInvalidExpression
			}
			x := new_stack[len(new_stack)-2]
			y := new_stack[len(new_stack)-1]
			new_stack = new_stack[:len(new_stack)-2]
			new_stack = append(new_stack, functions(x, y, elem))
		}
	}

	return new_stack[0], nil
}
