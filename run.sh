echo "kaiclairvoyant" > result.txt 
cp kaiclairvoyant/run.go run.go
go test -benchmem -run=^$ -bench ^BenchmarkRun$ github.com/eminetto/desafio-go >> result.txt
git checkout run.go

echo "\n\nleogues" >> result.txt
cp leogues/run.go run.go
go test -benchmem -run=^$ -bench ^BenchmarkRun$ github.com/eminetto/desafio-go >> result.txt
git checkout run.go

echo "\n\nVassopoli" >> result.txt
cp Vassopoli/run.go run.go
go test -benchmem -run=^$ -bench ^BenchmarkRun$ github.com/eminetto/desafio-go >> result.txt
git checkout run.go

echo "\n\nviniciusayres" >> result.txt
cp viniciusayres/run.go run.go
go test -benchmem -run=^$ -bench ^BenchmarkRun$ github.com/eminetto/desafio-go >> result.txt
git checkout run.go
