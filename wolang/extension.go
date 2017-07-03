package wolang

type extfuncdef struct {
	name           string
	implementation func(args []interface{}) (interface{}, error)
}

func (fd extfuncdef) Name() string {
	return fd.name
}

func (fd extfuncdef) Call(args []interface{}) (interface{}, error) {
	return fd.implementation(args)
}

type callable interface {
	Name() string
	Call(args []interface{}) (interface{}, error)
}

var extendedFunctions map[string]callable = map[string]callable{
	always99.Name(): always99,
}

var always99 extfuncdef = extfuncdef{
	"always99",
	func(terms []interface{}) (result interface{}, err error) {
		return 99, nil
	},
}
