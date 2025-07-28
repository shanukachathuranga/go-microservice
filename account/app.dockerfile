FROM golang:1.22-alpine AS build

WORKDIR /go/src/github.com/shanukachathuranga/go-microservice

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /go/bin/app ./account/cmd/account

FROM scratch

COPY --from=build /go/app/bin /app

EXPOSE 8080

CMD ["app"]