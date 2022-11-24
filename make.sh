GOOS=windows go build -o executables/windows-main.exe main.go
GOOS=linux  go build -o executables/linux-main main.go
GOOS=darwin go build -o executables/darwin-main main.go