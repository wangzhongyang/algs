package myMath

type MyMath struct{}

func New() *MyMath {
	return &MyMath{}
}

func (m *MyMath) Add(args *interface{}, hincr interface{}) interface{} {
	switch v := args.(type) {
	case int:
		return v + hincr.(int)
	case float64:
		return v + hincr.(float64)
	default:
		return "error"
	}
}
