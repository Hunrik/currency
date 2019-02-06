# Currency

[![godoc](https://godoc.org/github.com/Hunrik/currency?status.svg)](https://godoc.org/github.com/Hunrik/currency)

A golang library to deal with money and currency representation.

## Installation

    $ go get github.com/Hunrik/currency

## Usage

First, import the package adding:

```go
import "github.com/Hunrik/currency"
```

Examples:

```go
currency.Format(10, "USD")                                  // "$10.00"
currency.Format(10, "EUR")                                  // "€10.00"
currency.Format(10, "USD", WithCents(false))                // "$10"
currency.Format(10, "USD", WithCurrency(true))              // "$10.00 USD"
currency.Format(10, "USD", WithSymbol(false)                // "10.00"
currency.Format(10, "USD", WithSymbolSpace(true)            // "$ 10.00"
currency.Format(1000, "USD")                                // "$1,000.00"
currency.Format(1000, "USD", WithThousandsSeparator(false)  // "$1000.00"
```

For more detailed documentation refer to [godoc](http://godoc.org/github.com/Hunrik/currency)

## Contributing

1. [Fork it](https://github.com/Hunrik/currency/fork)
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

## License

[The MIT licence](LICENSE.md)
