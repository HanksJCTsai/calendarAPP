# 使用 Go 官方的 Docker 鏡像作為基礎鏡像來建立你的應用鏡像
FROM golang:1.22 as builder

# 設置工作目錄
WORKDIR /app

# 將 go.mod 和 go.sum 文件複製到容器中
COPY ./calendarAPI/go.mod ./
COPY ./calendarAPI/go.sum ./

# 下載所有依賴
RUN go mod download

# 將源代碼複製到容器中
COPY ./calendarAPI/ .

# 編譯 Go 程式碼（適用於 Linux 平台）
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .

# 使用一個新的輕量級鏡像來運行你的 Go 應用
FROM alpine:latest

# 在新的鏡像中設置工作目錄
WORKDIR /app

# 將從上一階段 builder 中編譯的二進制執行檔複製到新的工作目錄中
COPY --from=builder /app/main .

# 執行應用程式
# CMD ["./main"]
