title runAPP

REM del /q *.log

rem prompt $p$g
set GOPATH=%~dp0
set GOBIN=%~dp0bin

go run main.go config.go 

