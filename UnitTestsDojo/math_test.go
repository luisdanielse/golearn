package math

import (
	"testing"
)

func TestAbs(t *testing.T) {
	/* Pruebas para las sig combinaciones -1, 0 ,1 */
	if Abs(a:-1) < 0 {
		t.Error("Negative value found in abs() with %d", -1)
	}

	if Abs(a:0) < 0 {
		t.Error("Negative value found in abs() with %d", 0)
	}

	if Abs(a:1) < 0 {
		t.Error("Negative value found in abs() with %d", 1)
	}
}
