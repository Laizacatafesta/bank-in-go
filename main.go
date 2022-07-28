package main

import (
	"banco/conta"
	"fmt"
)

// este progrgaminha servirá para gerenciar contas correntes
func PagarBoleto(conta veriricarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type veriricarConta interface {
	Sacar(valor float64) string
}

func main() {

	ContaGabe := conta.ContaPoupanca{}
	ContaGabe.Depositar(500)
	PagarBoleto(&ContaGabe, 50)
	fmt.Println(ContaGabe.ObterSaldo())

	contaLaiza := conta.ContaCorrente{}
	contaLaiza.Depositar(100)
	PagarBoleto(&contaLaiza, 250)
	fmt.Println(contaLaiza.ObterSaldo())
}
