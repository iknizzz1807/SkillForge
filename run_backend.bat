@echo off
cd /d "%~dp0backend\cmd"

go mod tidy

go run main.go

pause