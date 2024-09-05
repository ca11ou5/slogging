FROM registry.sovcombank.group/web-ecom/golang-1.22-alpine-certs:latest AS builder
WORKDIR /usr/local/src
#RUN apk --no-cache add curl tar

COPY ["go.mod", "go.sum", "/"]
RUN go mod download

COPY ./ ./
RUN go build -o ./bin/main ./cmd/main.go


FROM alpine:latest AS runner

COPY --from=builder /usr/local/src/bin/main /main

CMD ["./main"]