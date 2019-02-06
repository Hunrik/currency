/*
package currency is a library to deal with money and currency representation.

Usage

    Format(10, "USD")                                   // "$10.00"
    Format(10, "EUR")                   								// "€10.00"
    Format(10, "USD", WithCents(false))                 // "$10"
    Format(10, "USD", WithCurrency(true))              	// "$10.00 USD"
    Format(10, "USD", WithSymbol(false)                	// "10.00"
    Format(10, "USD", WithSymbolSpace(true)            	// "$ 10.00"
    Format(1000, "USD")                                 // "$1,000.00"
    Format(1000, "USD", WithThousandsSeparator(false)	  // "$1000.00"
*/
package currency

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// OptionsFn function parameter
type OptionsFn func(*Options)

// WithCents Should display cents or not
// Default true
func WithCents(val bool) OptionsFn {
	return func(opts *Options) {
		opts.WithCents = val
	}
}

// WithCurrency Should add currency text or not
// Default false
func WithCurrency(val bool) OptionsFn {
	return func(opts *Options) {
		opts.WithCurrency = val
	}
}

// WithSymbol Should display symbol or not
// Default true
func WithSymbol(val bool) OptionsFn {
	return func(opts *Options) {
		opts.WithSymbol = val
	}
}

// WithSymbolSpace Should add space between symbol or not
// Default false
func WithSymbolSpace(val bool) OptionsFn {
	return func(opts *Options) {
		opts.WithSymbolSpace = val
	}
}

// WithThousandsSeparator Should use thousand separator or not
// Default true
func WithThousandsSeparator(val bool) OptionsFn {
	return func(opts *Options) {
		opts.WithCents = val
	}
}

// Options {
// 	WithCents:              true,
// 	WithCurrency:           false,
// 	WithSymbol:             true,
// 	WithSymbolSpace:        false,
// 	WithThousandsSeparator: true,
// }
type Options struct {
	WithCents              bool
	WithCurrency           bool
	WithSymbol             bool
	WithSymbolSpace        bool
	WithThousandsSeparator bool
}

func defaultOptions() *Options {
	return &Options{
		WithCents:              true,
		WithCurrency:           false,
		WithSymbol:             true,
		WithSymbolSpace:        false,
		WithThousandsSeparator: true,
	}
}

// Format returns a formatted price string according to currency rules and options
func Format(amount float64, currency string, opts ...OptionsFn) string {
	options := defaultOptions()

	for _, opt := range opts {
		opt(options)
	}

	c := currencies[currency]

	sign, integer, fractional := splitValue(amount)

	result := strconv.Itoa(integer)

	if options.WithThousandsSeparator {
		result = separateThousands(result, c.ThousandsSeparator)
	}

	if options.WithCents && c.SubUnit != "" {
		result = fmt.Sprintf("%s%s%s", result, c.DecimalMark, fractional)
	}

	if options.WithSymbol {
		result = addSymbol(result, c, options)
	}

	if options.WithCurrency {
		result = fmt.Sprintf("%s %s", result, currency)
	}

	return sign + result
}

func addSymbol(result string, c currency, options *Options) string {
	var space string

	if options.WithSymbolSpace {
		space = " "
	}

	if c.SymbolFirst {
		result = fmt.Sprintf("%s%s%s", c.Symbol, space, result)
	} else {
		result = fmt.Sprintf("%s%s%s", result, space, c.Symbol)
	}

	return result
}

func separateThousands(value, separator string) string {
	chunks := len(value) / 3

	if chunks == 0 {
		return value
	}

	if partial := math.Mod(float64(len(value)), 3); partial > 0 {
		chunks++
	}

	result := make([]string, chunks)

	for i := chunks - 1; i >= 0; i-- {
		if i == 0 {
			result[i] = value
			break
		}

		chunk := value[len(value)-3:]
		value = strings.TrimSuffix(value, chunk)
		result[i] = chunk
	}

	return strings.Join(result, separator)
}

func splitValue(val float64) (sign string, integer int, fractional string) {
	if val < 0 {
		val = math.Abs(val)
		sign = "-"
	}

	i, f := math.Modf(val)
	integer = int(i)
	fractional = fmt.Sprintf("%.2F", f)[2:]
	return
}
