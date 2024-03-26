# Vassa

## Informações pessoais
Github: [@Vassopoli](https://github.com/Vassopoli)

Blog: https://blog.vassopoli.com/

## Informações sobre o desafio

Foi feita uma refatoração para eliminar a complexidade quadrática na comparação da lista com ela mesma. A solução elaborada foi colocar num mapa os números gerados e quantas vezes repetem. Com este mapa em mãos, se torna possível verificar a existência do número 47 no array em tempo constante.

### Benchmarks

#### Tempo de execução com o payload de tamanho padrão, 10.000:
- Solução base: 0.0398 ns/op
- Solução implementada: 0.0034349 ns/op

#### Tempo de execução com um payload de 100.000:
- Solução base: 4263635833 ns/op
- Solução implementada: 0.04417 ns/op

#### Tempo de execução com um payload de 1.000.000:
- Solução base: 331853817 250 ns/op
- Solução implementada: 0.5399 ns/op
