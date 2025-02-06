:: filepath: /d:/work_space/twelveGo/scripts/start.bat
@echo off
:: 編譯應用程序
go build -o .\build\server.exe .\cmd\api

:: 運行應用程序
.\build\server