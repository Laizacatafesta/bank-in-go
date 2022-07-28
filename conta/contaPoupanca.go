package conta

import "banco/cliente"

type ContaPoupanca struct {
	Titular                              cliente.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {

	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito

		return "Depósito realizado com sucesso! Você possue um saldo de:", c.saldo
	} else {
		return "Não foi possível realizar depósito", c.saldo
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
