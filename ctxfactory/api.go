package ctxfactory

var contextFactory ContextFactory

// Factory() returns the currently set ContextFactory
// All calls to create contexts show go through this function.
func Factory() ContextFactory {
	return contextFactory
}

func UseAeFactory() {
	contextFactory = AppEngineFactory
}

func UseBaseFactory() {
	contextFactory = BaseFactory
}