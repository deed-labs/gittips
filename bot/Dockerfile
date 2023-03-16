FROM golang:1.19.3-alpine3.16 as builder

COPY go.mod go.sum app-private-key.pem /go/src/gittips-backend/
WORKDIR /go/src/gittips-backend
RUN go mod download
COPY . /go/src/gittips-backend
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/gittips-backend ./cmd/bot

FROM alpine

COPY --from=builder /go/src/gittips-backend/build/gittips-backend /usr/bin/gittips-backend
COPY --from=builder /go/src/gittips-backend/app-private-key.pem /usr/bin/app-private-key.pem

EXPOSE 8000 8000

ENTRYPOINT ["/usr/bin/gittips-backend"]