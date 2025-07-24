// a. Criar um modelo que irá receber os itens para compra do mês, nesse modelo teremos a data que a compra irá acontecer,
// o mercado (local) e os itens para comprar
// b. Dado o exercicio anterior, mover o modelo anterior para o pacote chamado "Model"
// c. Dado o exercicio anterior, criar uma função no pacote "Model" que inicializa a struct e retorna como ponteiro
// d. Fazer a tratativa de erros caso mercado não seja definido ou itens para comprar esteja vazio
package main

import (
	Model "ex3/compras"
	"fmt"
	"time"
)

// utiliza na letra a
/* type Mercado struct {
	Data  time.Time
	Local string
	Itens []ItensCompra
}

type ItensCompra struct {
	Nome string
}
*/

func main() {

	//utiliza na letra a e b
	/* var itens []Model.ItensCompra
	itens = append(itens, Model.ItensCompra{Nome: "arroz"})
	itens = append(itens, Model.ItensCompra{Nome: "feijão"})
	itens = append(itens, Model.ItensCompra{Nome: "alface"})
	itens = append(itens, Model.ItensCompra{Nome: "tomate"})
	itens = append(itens, Model.ItensCompra{Nome: "batata"})
	itens = append(itens, Model.ItensCompra{Nome: "carne"})

	mercado := Model.Mercado{
		Data:  time.Date(2025, 07, 25, 0, 0, 0, 0, time.Local),
		Local: "ABC",
		Itens: itens,
	} */

	// letra c
	/*var nomeItens []string
	nomeItens = append(nomeItens, "arroz")
	nomeItens = append(nomeItens, "feijão")
	nomeItens = append(nomeItens, "carne")

	mercado := Model.NovaCompra(time.Date(2025, 07, 25, 0, 0, 0, 0, time.Local), "Alvorada", nomeItens)

	fmt.Println("Data para ir ao mercado:", mercado.Data, "No mercado:", mercado.Local, "e comprar:", mercado.Itens)
	*/

	// letra d - validade para local não definido
	/*var nomeItens []string
	nomeItens = append(nomeItens, "arroz")
	nomeItens = append(nomeItens, "feijão")
	nomeItens = append(nomeItens, "carne")

	mercado, err := Model.NovaCompra(time.Date(2025, 07, 25, 0, 0, 0, 0, time.Local), "", nomeItens)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Data para ir ao mercado:", mercado.Data, "No mercado:", mercado.Local, "e comprar:", mercado.Itens)
	}
	*/

	// letra d - validade para lista de compras vazia
	/*	var nomeItens []string

		mercado, err := Model.NovaCompra(time.Date(2025, 07, 25, 0, 0, 0, 0, time.Local), "Geraldinho", nomeItens)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Data para ir ao mercado:", mercado.Data, "No mercado:", mercado.Local, "e comprar:", mercado.Itens)
		}
	*/

	// completo
	var nomeItens []string
	nomeItens = append(nomeItens, "arroz")
	nomeItens = append(nomeItens, "feijão")
	nomeItens = append(nomeItens, "carne")

	mercado, err := Model.NovaCompra(time.Date(2025, 07, 25, 0, 0, 0, 0, time.Local), "Alvorada", nomeItens)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Data para ir ao mercado:", mercado.Data, "No mercado:", mercado.Local, "e comprar:", mercado.Itens)
	}

}
