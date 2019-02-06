package currency

import (
	"testing"
)

func TestInit(t *testing.T) {
	if currencies == nil {
		t.Error("Expected currencies to be initialized.")
	}
}
