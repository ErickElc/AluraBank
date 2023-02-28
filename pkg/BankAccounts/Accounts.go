package pkg

import clientes "AluraBank/pkg/Client"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	valorDepositado := valor > 0
	if valorDepositado {
		c.Saldo += valor
		return "Deposito feito com sucesso. Seu Saldo: %.2f\n", c.Saldo
	}
	return "Não foi possível, fazer o depósito: %.2f\n", c.Saldo
}
func (c *ContaCorrente) Sacar(valor float64) (string, float64) {
	podeSacar := valor > 0 && valor <= c.Saldo
	if podeSacar {
		c.Saldo -= valor
		return "O Saque foi feito com sucesso. %.2f\n", c.Saldo
	}
	return "Não foi possível sacar %.2f\n", c.Saldo
}
func (c *ContaCorrente) VerSaldo() (string, float64) {
	return "Seu saldo é de %.2f", c.Saldo
}
func (c *ContaCorrente) Tranferir(valor float64, ContaCorrente *ContaCorrente) bool {
	podeTransferir := valor > 0 && valor <= c.Saldo
	if podeTransferir {
		c.Saldo -= valor
		ContaCorrente.Depositar(valor);
		return true
	}
	return false
}
