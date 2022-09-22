package main

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	result, msg := isPrime(0)

	if result {
		t.Error("with 0 as test parameter, got true, but expected false")
	}

	if msg != "0 não é um numero primo por definição." {
		t.Error("wrong message returned: ", msg)
	}
}

func TestIsprimeTable(t *testing.T) {
	primeTests := []struct {
		name   string
		number int
		result bool
		msg    string
	}{
		{"Numero primo 7", 7, true, "7 é um numero primo."},
		{"Zero", 0, false, "0 não é um numero primo por definição."},
		{"Um", 1, false, "1 não é um numero primo por definição."},
		{"Negativo", -7, false, "Por definição numeros negativos não são numeros primos."},
		{"Divisivel por dois", 4, false, "4 não é um numero primo pois é divisivel por 2."},
		{"Divisivel por tres", 9, false, "9 não é um numero primo pois é divisivel por 3."},
		{"Numero grande não primo", 34643642362626667, false, "34643642362626667 não é um numero primo pois é divisivel por 99943331."},
		{"Numero grande primo", 131071, true, "131071 é um numero primo."},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.number)

		if result != e.result {
			t.Errorf("\nScenario %s \nWith %d as test parameter.\nExpected: %v\nReturned: %v", e.name, e.number, e.result, result)
		}

		if msg != e.msg {
			t.Errorf("\nScenario %s: \nExpected message: %s\nReturned message: %s", e.name, e.msg, msg)
		}

	}
}
