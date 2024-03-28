## Informações pessoais

Leonardo Rodrigues
Github: [@leogues](https://github.com/leogues)

## Informações sobre desafio

A principio, busquei otimizar o cálculo da moda empregando um mapa para registrar a contagem de ocorrências de cada elemento. No entanto, ao aprender e depois fazer o profile da aplicação, identifiquei que o cálculo da moda ainda era a operação mais demorada. Observando que outros calculos utilizavam ordenação do array, optei por testar essa abordagem também para o cálculo da moda. E essa mudança acabou resultando em uma significativa redução de tempo de execução.

Obs: Apesar de não ver uma mudança no tempo ao executar os tests, quando eu fiz a ordenação no inicio da função Run tive ganho bem expressivo de tempo ao fazer o profile da aplicação.

Obs: Pelos tests que fiz até 100.000.000 a versão com sort era superior ao map em questão de tempo de execução, não consegui testar valores maiores porque a versão com map passava da quantidade de memoria que eu tinha disponivel.

#### Tempo de execução, 10.000:

- Versão Base: 0.09100 ns/op
- Versão Implementada: 0 ns/op

#### Tempo de execução, 100.000:

- Versão Base: 8951554500 ns/op 3214136 B/op
- Versão Implementada: 0.0009995 ns/op

#### Tempo de execução, 1.000.000:

- Versão Base: ??? ns/op
- Versão Implementada: 0.008499 ns/op

#### Tempo de execução, 10.000.000:

- Versão Base: ??? ns/op
- Versão Implementada: 1156666300 ns/op
