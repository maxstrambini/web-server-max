title buildAPP

REM del /q *.log

rem prompt $p$g
set GOPATH=%~dp0
set GOBIN=%~dp0bin

del /q webservermax.exe

go build -o webservermax.exe main.go config.go 

dir webservermax.exe


