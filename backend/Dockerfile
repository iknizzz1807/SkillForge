# Sử dụng Golang image để build
FROM golang:1.24-alpine

# Thiết lập thư mục làm việc
WORKDIR /app

# Cài đặt các dependencies cần thiết
RUN apk add --no-cache git

# Copy go.mod và go.sum (nếu có)
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy toàn bộ source code
COPY . .

# Thiết lập port
EXPOSE 8080

# Command để chạy ứng dụng trong development
CMD ["go", "run", "cmd/main.go"]