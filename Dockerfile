FROM golang:1.20 AS build

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY main.go main.go
COPY client/ client/
COPY model/ model/
COPY psql/ psql/
COPY securities/ securities/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api-identities main.go

FROM build AS run

WORKDIR /app

COPY --from=build /build/api-identities .

ENTRYPOINT [ "/app/api-identities" ]
