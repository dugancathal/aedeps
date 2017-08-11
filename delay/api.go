package delay

var funcFactory func(key string, callable Callable) Function

// The primary interface into the delay package. It has
// the same requirements as the underlying appengine Func function -
// you MUST instantiate all "Functions" at application init.
func Func(key string, callable Callable) Function {
	return funcFactory(key, callable)
}

func UseAe() {
	funcFactory = appengineFuncFactory
}

func UseBase() {
	funcFactory = baseFuncFactory
}
