package ponteiros

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

type Carteira struct {
	saldo Bitcoin //letra minúscula == atributo privado
}

func (c *Carteira) Depositar(quantidade Bitcoin) {
	//fmt.Printf("Endereço saldo no método Depositar: %v \n", &c.saldo)
	c.saldo += quantidade
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

var ErroSaldoInsuficiente = errors.New("Não é possível retirar: saldo insuficiente")

func (c *Carteira) Retirar(quantidade Bitcoin) error {

	if quantidade > c.saldo {
		return ErroSaldoInsuficiente
	}

	c.saldo -= quantidade
	return nil
}
