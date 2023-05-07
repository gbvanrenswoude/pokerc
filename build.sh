GOOS=darwin GOARCH=amd64 go build -o dist/pokerc_macos main.go
GOOS=linux GOARCH=amd64 go build -o dist/pokerc_linux main.go
GOOS=windows GOARCH=amd64 go build -o dist/pokerc_windows.exe main.go
GOOS=darwin GOARCH=arm64 go build -o dist/pokerc_macos_arm main.go
GOOS=linux GOARCH=arm64 go build -o dist/pokerc_linux_arm main.go
GOOS=windows GOARCH=arm64 go build -o dist/pokerc_windows_arm.exe main.go