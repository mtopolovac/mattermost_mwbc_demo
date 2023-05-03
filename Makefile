build:
	rm -rf plugin.exe
	rm -rf plugin.tar.gz
	go build -o plugin.exe plugin.go
	tar -czvf plugin.tar.gz plugin.exe plugin.json