set GOOS=linux
set GOARCH=amd64
go build main.go
del main.zip
tar.exe -a -cf main.zip main



@REM compilar lamba en mac
@REM GOOS=linux GOARCH=amd64 go build main.go
@REM zip myFunction.zip main