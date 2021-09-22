package main

import "fmt"

type Dicionario map[string]string

const (
	ErrNaoEncontrado      = ErrDicionario("Não foi possível encontrar a palavra procurada.")
	ErrPalavraExistente   = ErrDicionario("Não foi possível adicionar a palavra, pois ela já existe.")
	ErrPalavraInexistente = ErrDicionario("Não foi possível atualizar a palavra, pois ela não existe.")
)

type ErrDicionario string

func (e ErrDicionario) Error() string {
	return string(e)
}

func (d Dicionario) Busca(palavra string) (string, error) {
	definicao, existe := d[palavra]

	if !existe {
		return "", ErrNaoEncontrado
	}

	return definicao, nil
}

// func Busca(dicionario map[string]string, palavra string) string {
// 	return dicionario[palavra]
// }

func (d Dicionario) Adiciona(palavra, definicao string) error {

	_, err := d.Busca(palavra)

	switch err {
	case ErrNaoEncontrado:
		d[palavra] = definicao
	case nil:
		return ErrPalavraExistente
	default:
		return err
	}

	return nil
}

func (d Dicionario) Atualiza(palavra, definicao string) error {

	_, err := d.Busca(palavra)

	switch err {
	case ErrNaoEncontrado:
		return ErrPalavraInexistente
	case nil:
		d[palavra] = definicao
	default:
		return err
	}

	return nil
}

func (d Dicionario) Deleta(palavra string) {
	delete(d, palavra)
}

func main() {

	dic := Dicionario{}
	dic.Adiciona("dicionário", "conjunto de palavras")
	//fmt.Println(dic)
	dic.Adiciona("teste", "apenas um teste")
	//fmt.Println(dic)

	//fmt.Print(dic.Busca("dicionário"))

	dic.Atualiza("teste", "testinho")
	//fmt.Println(dic)

	dic.Deleta("teste")
	fmt.Println(dic)
}
