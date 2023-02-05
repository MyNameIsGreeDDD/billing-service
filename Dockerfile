FROM golang:latest

RUN go version
ENV GOPATH=/

ADD ./ ./

RUN go mod download
RUN go build -o avito-test-case ./cmd/main.go


CMD ["./billingService"]