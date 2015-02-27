package money

type Options map[string]interface{}

func defaults() Options {
	return Options{
		"currency":                 "usd",
		"with_cents":               true,
		"with_currency":            false,
		"with_symbol":              true,
		"with_symbol_space":        false,
		"with_thousands_separator": true,
	}
}

func override(original, override Options) Options {
	for k, v := range override {
		if override[k] != nil {
			original[k] = v
		}
	}

	return original
}
