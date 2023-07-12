FROM golang:1.20 AS build

WORKDIR /build

COPY go.mod mo.mod
COPY go.sum go.sum
COPY main.go main.go
COPY client/ client/
COPY model/ model/
COPY psql/ psql/
COPY securities/ securities/

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api-identities main.go

FROM build AS run

WORKDIR /app

COPY --from=build /build/api-identities .

ENTRYPOINT [ "/app/api-identities" ]
