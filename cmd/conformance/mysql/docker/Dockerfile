FROM golang:1.23.4-bookworm@sha256:2e838582004fab0931693a3a84743ceccfbfeeafa8187e87291a1afea457ff7a AS build

WORKDIR /build

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -v -o ./conformance-mysql ./cmd/conformance/mysql

FROM alpine:3.20@sha256:0a4eaa0eecf5f8c050e5bba433f58c052be7587ee8af3e8b3910ef9ab5fbe9f5

COPY --from=build /build/conformance-mysql /build/cmd/conformance/mysql/docker/testdata/key /build/cmd/conformance/mysql/docker/testdata/key.pub /build/storage/mysql/schema.sql /

ENTRYPOINT ["/conformance-mysql"]
