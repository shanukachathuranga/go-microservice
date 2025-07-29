FROM golang:1.22-alpine AS build

WORKDIR /go/src/github.com/shanukachathuranga/go-microservice

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLE=0 go build -ldflags="-w -s" -o /go/bin/app ./catalog/cmd/catalog

FROM scratch

COPY --from=build /go/bin/app /app

EXPOSE 8080

CMD ["app"]