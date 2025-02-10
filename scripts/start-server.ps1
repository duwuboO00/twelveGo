# 確保 build 目錄存在
if (!(Test-Path -Path "./build")) {
    New-Item -ItemType Directory -Path "./build"
}

# 編譯 Go 伺服器
go build -o .\build\server.exe .\cmd\api

# 執行應用程序
.\build\server.exe
