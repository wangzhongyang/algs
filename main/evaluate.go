package main

import (
	"flag"

	"math"

	"fmt"

	"stack"
	"strconv"
	"typeChange"
)

func main() {
	var evaluate string
	flag.StringVar(&evaluate, "s", "(1+1)", "输入计算式：	")
	flag.Parse()

	opt, vals := stack.New(), stack.New()
	for _, v := range evaluate {
		vChar := string(v)
		if vChar == "(" || vChar == " " {

		} else if vChar == "+" {
			opt.Push(vChar)
		} else if vChar == "-" {
			opt.Push(vChar)
		} else if vChar == "*" {
			opt.Push(vChar)
		} else if vChar == "/" {
			opt.Push(vChar)
		} else if vChar == "sqrt" {
			opt.Push(vChar)
		} else if vChar == ")" {
			// 弹出上一个运算符，和前两个数值进行计算，结果进栈
			bol := opt.Pop().(string)
			val1, val := toFloat(vals.Pop()), 0.0
			switch bol {
			case "+":
				val = val1 + toFloat(vals.Pop())
			case "-":
				val = toFloat(vals.Pop()) - val1
			case "*":
				val = val1 * toFloat(vals.Pop())
			case "/":
				val = toFloat(vals.Pop()) / val1
			case "sqrt":
				val = math.Sqrt(val1)
			}
			fmt.Println(val, val1)
			vals.Push(toString(val))
		} else { // 不为字符，进vals栈
			vals.Push(vChar)
		}
	}
	fmt.Println(vals.Pop())
}

func toFloat(s interface{}) float64 {
	tc := typeChange.New()
	return tc.ObjToFloat(s)
}

func toString(f float64) string {
	return strconv.FormatFloat(f, 'f', 2, 64)
}
