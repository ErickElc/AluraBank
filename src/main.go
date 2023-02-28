package main

import (
	conta "AluraBank/pkg/BankAccounts"
	cliente "AluraBank/pkg/Client"
	"fmt"
)

func main() {

	clienteErick := cliente.Titular{
		Nome:      "Erick Lucas",
		CPF:       "167676877-74",
		Profissao: "Desenvolvedor IOS",
	}

	contaDoErick := conta.ContaCorrente{
		Titular:       clienteErick,
		NumeroAgencia: 1234,
		NumeroConta:   123456,
		Saldo:         2356.30,
	}
	contaDaBruna := conta.ContaPoupanca{
		Titular: cliente.Titular{
			Nome:      "Bruna Rodrigues",
			CPF:       "167676877-74",
			Profissao: "Desenvolvedora web full-stack",
		},
		NumeroAgencia: 1234,
		NumeroConta:   654321,
		Saldo:         20341.31,
	}
	fmt.Println(contaDoErick.Saldo)
	fmt.Printf(contaDaBruna.Depositar(10000))

	fmt.Printf(contaDoErick.VerSaldo())
}
