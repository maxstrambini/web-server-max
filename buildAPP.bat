title buildAPP

del /q webservermax.exe

go build -o webservermax.exe main.go config.go 

dir webservermax.exe


