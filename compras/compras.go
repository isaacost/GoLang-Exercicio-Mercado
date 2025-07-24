package Model

import (
	"errors"
	"time"
)

// letra b
type Mercado struct {
	Data  time.Time
	Local string
	Itens []ItensCompra
}

type ItensCompra struct {
	Nome string
}

// letra c
/*func NovaCompra(data time.Time, local string, novoItem []string) *Mercado {
	var itens []ItensCompra

	for _, nome := range novoItem {
		itens = append(itens, ItensCompra{Nome: nome})
	}

	return &Mercado{
		Data:  data,
		Local: local,
		Itens: itens,
	}
}
*/

// letra d
func NovaCompra(data time.Time, local string, novoItem []string) (*Mercado, error) {
	// valida se local(mercado) é definido
	if local == "" {
		return nil, errors.New("O nome do mercado precisa ser definido")
	}

	// valida se a lista de compras foi passada
	if len(novoItem) == 0 {
		return nil, errors.New("A lista de comprar precisa ter pelo menos 1 item")
	}

	var itens []ItensCompra

	for _, nome := range novoItem {
		itens = append(itens, ItensCompra{Nome: nome})
	}

	return &Mercado{
		Data:  data,
		Local: local,
		Itens: itens,
	}, nil
}
