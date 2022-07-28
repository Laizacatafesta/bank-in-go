package conta

import "banco/cliente"

type ContaCorrente struct {
	Titular                    cliente.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {

	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito

		return "Depósito realizado com sucesso! Você possue um saldo de:", c.saldo
	} else {
		return "Não foi possível realizar depósito", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {

	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}

}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
