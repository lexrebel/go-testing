package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	intro()
	doneChan := make(chan bool)
	go readUserImput(os.Stdin, doneChan)
	<-doneChan
	close(doneChan)
	fmt.Println("Até logo.")
}

func readUserImput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)

	for {
		res, done := checkNumbers(scanner)
		if done {
			doneChan <- true
			return
		}
		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {

	scanner.Scan()
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Digite um numero inteiro.", false
	}

	_, msg := isPrime(numToCheck)

	return msg, false
}

func intro() {
	fmt.Println("Numeros primos")
	fmt.Println("--------------")
	fmt.Println("Digite um numero inteiro e este programa irá lhe responder se é um numero primo ou não. Digite q para sair.")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
}

func isPrime(number int) (bool, string) {
	// 0 e 1 não são primos por definição
	if number == 0 || number == 1 {
		return false, fmt.Sprintf("%d não é um numero primo por definição.", number)
	}

	// numeros negativos não são primos
	if number < 0 {
		return false, "Por definição numeros negativos não são numeros primos."
	}

	// usar operador de modulo repedidamente para ver se temos um numero primo
	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			// não é um numero primo
			return false, fmt.Sprintf("%d não é um numero primo pois é divisivel por %d.", number, i)
		}
	}

	// provavelmente um numero primo
	return true, fmt.Sprintf("%d é um numero primo.", number)
}
