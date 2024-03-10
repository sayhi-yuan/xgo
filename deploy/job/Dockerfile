FROM docker.qdreads.com/base/golang-1.20:v1 as builder 
WORKDIR  /go/compile
## 编译
COPY . .
RUN go build -o corntab cmd/http/main.go

## 二级构建
FROM docker.qdreads.com/base/gocron-node:v3
WORKDIR  /go

COPY --from=builder /go/compile/corntab .

EXPOSE 8080
CMD ["./corntab"]