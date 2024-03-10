FROM docker.qdreads.com/base/golang-1.20:v1 as builder 
WORKDIR  /go/compile
## 编译
COPY . .
RUN go build -o http ./cmd/http/main.go

## 二级构建
FROM docker.qdreads.com/base/http-golang:v1
WORKDIR  /go
RUN mkdir log

COPY --from=builder /go/compile/http .

EXPOSE 8080
CMD ["./http"]