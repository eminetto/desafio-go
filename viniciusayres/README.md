## Informações pessoais

Vinicius Ayres  
GitHub: [@ViniciusAyres](https://github.com/ViniciusAyres)

## Informações sobre o desafio

Inicialmente, o foco da otimização foi a ordenação do vetor de entrada, visando evitar múltiplas ordenações desnecessárias (nas funções `calculateMedian` e `calculatePercentile`). Após a ordenação do vetor, foi possível melhorar a complexidade assintótica do algoritmo de cálculo da moda de `O(n^2)` para `O(n)`, sem a necessidade de utilizar memória adicional, através de um mapa. Em seguida, foi implementada uma busca binária para encontrar o elemento 47, reduzindo a complexidade assintótica do algoritmo de `O(n)` para `O(log n)`. Por fim, foi adotada uma abordagem de concorrência/paralelismo com goroutines, visando reduzir ainda mais o tempo de execução.

## Especificação da Máquina:
- Sistema Operacional: macOS (darwin)
- Arquitetura: amd64
- CPU: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz

## Resultados

### Tempo de execução, 10.000:
- Versão Base: 0.09343 ns/op
- Versão Implementada: 0.0001070 ns/op

### Tempo de execução, 100.000:
- Versão Base: 9010434692 ns/op
- Versão Implementada: 0.001341 ns/op

### Tempo de execução, 1.000.000:
- ~~Versão Base: ?? ns/op~~
- Versão Implementada: 0.008819 ns/op

### Tempo de execução, 10.000.000:
- ~~Versão Base: ?? ns/op~~
- Versão Implementada: 1227605347 ns/op

## Conclusão

Os resultados demonstram um aumento significativo de desempenho, sendo a versão implementada 873 vezes mais rápida que a versão base para uma entrada de 10 mil elementos. Este ganho é ainda mais expressivo quando a entrada aumenta para 100 mil elementos, alcançando uma melhoria de desempenho de aproximadamente 6,71 trilhões de vezes em relação à versão base. Infelizmente, não foi possível comparar com entradas maiores devido ao longo tempo de execução do algoritmo base.
