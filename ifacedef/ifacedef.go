package ifacedef

type I interface {
	F()
	G() string
	H(x string) error
}

type J interface {
	F()
}

type K interface {
	J
	G() string
}

type L interface {
	F() T
}

type T string

type LImpl struct{}

func (_ LImpl) F() T {
	return "qux"
}

var limpl L = LImpl{}
