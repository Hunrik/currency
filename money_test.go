package currency

import (
	"testing"
)

func TestFormatWithDefaults(t *testing.T) {
	currency := Format(10, "USD")
	expected := "$10.00"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestFormatWithNegativeValues(t *testing.T) {
	currency := Format(-10, "USD")
	expected := "-$10.00"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestFormatWithCurrency(t *testing.T) {
	currency := Format(10, "USD", WithCurrency(true))
	expected := "$10.00 USD"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestFormatWithCurrencyWhenCanadian(t *testing.T) {
	currency := Format(10, "CAD", WithCurrency(true))
	expected := "$10.00 CAD"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestFormatWhenJapanese(t *testing.T) {
	currency := Format(10, "JPY")
	expected := "¥10"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestFormatWhenJapaneseOverThousand(t *testing.T) {
	currency := Format(1000, "JPY")
	expected := "¥1,000"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestFormatWhenJapaneseOverTenThousand(t *testing.T) {
	currency := Format(10000, "JPY")
	expected := "¥10,000"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestFormatWhenJapaneseOverHundredThousand(t *testing.T) {
	currency := Format(100000, "JPY")
	expected := "¥100,000"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestFormatWhenJapaneseOverMillion(t *testing.T) {
	currency := Format(1000000, "JPY")
	expected := "¥1,000,000"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestFormatWithoutCents(t *testing.T) {
	currency := Format(10, "USD", WithCents(false))
	expected := "$10"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestAddSymbol(t *testing.T) {
	q := "10.00"

	if afn := addSymbol(q, currencies["AFN"], defaultOptions()); afn != "10.00؋" {
		t.Errorf("Expected euro symbol to be placed after quantity, but got %s", afn)
	}

	if afn := addSymbol(q, currencies["AFN"], &Options{WithSymbolSpace: true}); afn != "10.00 ؋" {
		t.Errorf("Expected euro symbol to be placed after quantity and a space, but got %s", afn)
	}

	if usd := addSymbol(q, currencies["USD"], defaultOptions()); usd != "$10.00" {
		t.Errorf("Expected dollar symbol to be placed before quantity, but got %s", usd)
	}

	if usd := addSymbol(q, currencies["USD"], &Options{WithSymbolSpace: true}); usd != "$ 10.00" {
		t.Errorf("Expected dollar symbol to be placed before quantity and a space, but got %s", usd)
	}
}

func TestSeparateThousands(t *testing.T) {
	values := map[string]string{
		"1":          "1",
		"12":         "12",
		"123":        "123",
		"1234":       "1,234",
		"12345":      "12,345",
		"69310":      "69,310",
		"123456":     "123,456",
		"1234567":    "1,234,567",
		"12345678":   "12,345,678",
		"123456789":  "123,456,789",
		"1234567891": "1,234,567,891",
	}

	for value, expected := range values {
		if v := separateThousands(value, ","); v != expected {
			t.Errorf("Expected %s to be %s", v, expected)
		}
	}
}
