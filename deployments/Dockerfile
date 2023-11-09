# 使用 alpine 映像建立一個輕量級的容器
FROM golang:1.21 AS builder

# 設定工作目錄
WORKDIR /app

# 複製 go.mod 和 go.sum 文件
COPY go.mod .
COPY go.sum .

# 下載依賴
RUN go mod download

# 複製整個專案源代碼
COPY . .

# 構建應用程序
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/api

# 使用 alpine 作為運行時容器基礎
FROM alpine:latest  

# 在 alpine 映像中安裝 ca-certificates，這對於執行 HTTPS 請求是必要的
RUN apk --no-cache add ca-certificates

# 設定工作目錄到 /app
WORKDIR /app

# 從 builder 階段複製編譯好的應用程序和 web 目錄到 /app 目錄
COPY --from=builder /app/server /app/
COPY --from=builder /app/web /app/web

# 應用程序埠
EXPOSE 8080

# 運行應用程序
CMD ["./server"]
