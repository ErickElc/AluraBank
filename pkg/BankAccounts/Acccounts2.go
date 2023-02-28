package pkg

import clientes "AluraBank/pkg/Client"

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaPoupanca) Depositar(valor float64) (string, float64) {
	valorDepositado := valor > 0
	if valorDepositado {
		c.Saldo += valor
		return "Deposito feito com sucesso. Seu Saldo: %.2f\n", c.Saldo
	}
	return "Não foi possível, fazer o depósito: %.2f\n", c.Saldo
}
func (c *ContaPoupanca) Sacar(valor float64) (string, float64) {
	podeSacar := valor > 0 && valor <= c.Saldo
	if podeSacar {
		c.Saldo -= valor
		return "O Saque foi feito com sucesso. %.2f\n", c.Saldo
	}
	return "Não foi possível sacar %.2f\n", c.Saldo
}

func (c *ContaPoupanca) Tranferir(valor float64, Conta *ContaPoupanca) bool {
	podeTransferir := valor > 0 && valor <= c.Saldo
	if podeTransferir {
		c.Saldo -= valor
		Conta.Depositar(valor)
		return true
	}
	return false
}
