package wolang

type ExtFuncDef struct {
	FuncName       string
	Implementation func(args []DataType) (DataType, error)
}

func (fd ExtFuncDef) Name() string {
	return fd.FuncName
}

func (fd ExtFuncDef) Call(args []DataType) (DataType, error) {
	return fd.Implementation(args)
}

type Callable interface {
	Name() string
	Call(args []DataType) (DataType, error)
}

var extendedFunctions map[string]Callable = map[string]Callable{}

func RegExtFunc(ef Callable) {
	extendedFunctions[ef.Name()] = ef
}
