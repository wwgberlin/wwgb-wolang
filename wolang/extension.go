package wolang

type ExtFuncDef struct {
	FuncName       string
	Implementation func(args []interface{}) (interface{}, error)
}

func (fd ExtFuncDef) Name() string {
	return fd.FuncName
}

func (fd ExtFuncDef) Call(args []interface{}) (interface{}, error) {
	return fd.Implementation(args)
}

type Callable interface {
	Name() string
	Call(args []interface{}) (interface{}, error)
}

var extendedFunctions map[string]Callable = map[string]Callable{}

func RegExtFunc(ef Callable) {
	extendedFunctions[ef.Name()] = ef
}
