FROM golang as builder

WORKDIR /usr/src/go-app

COPY . ./
RUN go mod download && go mod verify

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /usr/local/bin/go-app ./cmd/main.go

FROM alpine

COPY --from=builder /usr/local/bin/go-app /usr/local/bin/go-app

CMD ["go-app"]
