# Aprimorando desempenho de programas

## Desafio

Você trabalha como Engenheira(o) de Software no time de Engenharia de Dados de uma grande corporação. Como sua próxima tarefa, você deve avaliar um código previamente desenvolvido pelo time de dados e otimizá-lo ao máximo em termos de tempo de execução. Em outras palavras, você já recebeu um código que funciona, e a sua *única* tarefa é torna-lo mais rápido!

Porém, há algumas regras que devem ser seguidas:

* Sua implementação deve continuar retornando um **resultado correto**. Para isso, certifique-se de utilizar as rotinas de testes previamente implementadas em **run_test.go** para asseguar que suas modificações não estão *quebrando* a corretude do programa.

* Você **NÃO PODE** alterar nenhum dos arquivos: **lib.go**, **baseline.go**, **run_test.go**. Você pode somente modificar a função ```Run()``` definida em **run.go**, e se desejar, você pode criar funções adicionais para serem invocadas em ```Run()```.

* Você **NÃO PODE** calcular os valores de média, mediana, e percentis sem utilizar as funções definidas em **lib.go**.

Como resultado final, é esperado que sua solução entregue o resultado correto em **um tempo de execução de no mínimo 4x menor** que a solução inicial, ao considerar **uma entrada de no mínimo 10 mil elementos** aleatórios gerados pela função ```getRandDataSlice()```. Para isso, certifique-se de utilizar a rotina de benchmark implementada em **run_test.go**.

Importante mencionar que o ganho de 4x no tempo de execução é o *mínimo* esperado. Quando mais rápido for sua solução, maior será o seu PLR no próximo semestre :)

## Instruções

- Faça um fork deste repositório
- Abra um PR contendo um diretório com a sua solução
- Dentro do diretório deve conter um README.md com os dados de quem criou a solução (nome e @ do Github)