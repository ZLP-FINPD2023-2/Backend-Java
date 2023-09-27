FROM golang as builder

WORKDIR /usr/src/go-app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /usr/local/bin/go-app ./cmd/main.go

FROM alpine

COPY --from=builder /usr/local/bin/go-app /usr/local/bin/go-app

CMD ["go-app"]
