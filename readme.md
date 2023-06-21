# 2022-2023 Project-5 Algorithmics Team-3

## Quick start

Install [Go 1.20.4](https://go.dev/dl/), then run the following commands to move in the correct directory and run the program:

```bash
cd src
```

```go
go run blend
```

## Change the CSV file

The code is preconfigured to use a specific CSV file. If you wish to use another CSV file, you will need to modify the line ```records := csvutils.OpenCSV("UseCase.csv")``` in **blend.go** to specify the path to the desired CSV file, like this:

```go
records := csvutils.OpenCSV("MyUseCase.csv")
```
