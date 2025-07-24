# 🛒 GoLang - Exercício Mercado

Este repositório contém um exercício prático em Go (Golang), com o objetivo de consolidar conceitos de modelagem de structs, modularização de pacotes, uso de ponteiros e tratamento de erros.

---

## 🎯 Proposta do Exercício

1. Criar um **modelo** que represente uma compra mensal, contendo:
   - Data da compra
   - Nome do mercado
   - Lista de itens para comprar

2. Mover esse modelo para um pacote separado chamado `model`.

3. Criar uma **função de inicialização** no pacote `model`, que retorne a struct como ponteiro.

4. Implementar **validação**:
   - Nome do mercado não pode ser vazio.
   - Lista de itens não pode estar vazia.
