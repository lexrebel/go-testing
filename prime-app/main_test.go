package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
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

// apesar de parecer pequeno o prox teste explica como testar saidas para console.
func TestPrompt(t *testing.T) {
	// salvar o std out antes do teste
	oldOut := os.Stdout

	// criar um write e um read pipe
	r, w, _ := os.Pipe()

	// stear o std out para o pipe criado onde consigo ler
	os.Stdout = w

	// chamar a funcao a ser testada
	prompt()

	// fechar o recurso para evitar problemas
	_ = w.Close()

	// voltar o std out ao seu modo antes do teste
	os.Stdout = oldOut

	// ler o resultado escrito no pipe
	out, _ := io.ReadAll(r)

	// agora sim testar o resultado
	if string(out) != "-> " {
		t.Errorf("Incorrect promt: expected \"-> \" but got \"%s\"", string(out))
	}
}

// Similar ao anterior porem com um check diferente
func TestIntro(t *testing.T) {
	// salvar o std out antes do teste
	oldOut := os.Stdout

	// criar um write e um read pipe
	r, w, _ := os.Pipe()

	// stear o std out para o pipe criado onde consigo ler
	os.Stdout = w

	// chamar a funcao a ser testada
	intro()

	// fechar o recurso para evitar problemas
	_ = w.Close()

	// voltar o std out ao seu modo antes do teste
	os.Stdout = oldOut

	// ler o resultado escrito no pipe
	out, _ := io.ReadAll(r)

	// agora sim testar o resultado
	if !strings.Contains(string(out), "Digite um numero inteiro") {
		t.Errorf("Intro text incorrect got: \"%s\"", string(out))
	}
}

func TestCheckNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "Empty", input: "", expected: "Digite um numero inteiro."},
		{name: "Zero", input: "0", expected: "0 não é um numero primo por definição."},
		{name: "One", input: "1", expected: "1 não é um numero primo por definição."},
		{name: "Two", input: "2", expected: "2 é um numero primo."},
		{name: "Negative", input: "-2", expected: "Por definição numeros negativos não são numeros primos."},
		{name: "Typed", input: "three", expected: "Digite um numero inteiro."},
		{name: "Decimal", input: "1.1", expected: "Digite um numero inteiro."},
		{name: "Greek", input: "ελπίδα", expected: "Digite um numero inteiro."},
		{name: "Quit lowercase", input: "q", expected: ""},
		{name: "Quit uppercase", input: "Q", expected: ""},
	}

	for _, e := range tests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: expected %s, but got %s", e.name, e.expected, res)
		}
	}
}
func TestReadUserInput(t *testing.T) {
	// um canal para usar como parametro
	doneChan := make(chan bool)

	// bytes buffer satisfaz a interface e possibilita escrita
	var stdin bytes.Buffer

	// simulando o usuario digitanto 1 + enter + q + enter
	stdin.Write([]byte("1\nq\n"))

	go readUserImput(&stdin, doneChan)

	// esperar o canal finalizar
	<-doneChan

	// fechar o canal
	close(doneChan)
}
