FROM golang:1.20 AS build

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY main.go main.go
COPY helper/ helper/
COPY model/ model/
COPY psql/ psql/
COPY securities/ securities/
COPY controllers/ controllers/
COPY external/ external/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api-identities main.go

FROM build AS run

WORKDIR /app

COPY --from=build /build/api-identities .

EXPOSE 9080
ENTRYPOINT [ "/app/api-identities" ]
