## Informações pessoais
https://www.linkedin.com/in/rodrigo-santorato-dev/

## Informações sobre o desafio
- Passo 1: refatorar o algoritmo de encontrar o Mode e o n47 que loopava os dados 1 milhão de vezes em O(n²) e transformar em 1 loop só O(n).
- Passo 2: colocar o processo citado acima mais as 4 chamadas de funções obrigatórias em go routines para execução concorrente.

### Benchmarks

#### Tempo de execução com o payload de tamanho padrão, 10.000:
- Solução base: 0.08300 ns/op
- Solução implementada: 0.0009988 ns/op

#### Tempo de execução com um payload de 100.000:
- Solução base: 8001216000 ns/op
- Solução implementada: 0.01500 ns/op

#### Tempo de execução com um payload de 1.000.000:
- Solução base: 799522986500 ns/op
- Solução implementada: 0.1795 ns/op