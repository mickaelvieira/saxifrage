FROM golang:1.16.9-alpine AS build

WORKDIR /src

ENV GO111MODULE=on
ENV CGO_ENABLED=0

COPY cmd/ cmd/
COPY config/ config/
COPY keys/ keys/
COPY lexer/ lexer/
COPY parser/ parser/
COPY prompt/ prompt/
COPY template/ template/
COPY upgrade/ upgrade/
COPY sax.go .
COPY go.sum .
COPY go.mod .

RUN go mod download

ARG APP_VERSION
ARG TARGETOS
ARG TARGETARCH

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags "-s -w -X main.version=${APP_VERSION}" -o saxifrage .

FROM scratch AS bin-unix

COPY --from=build /src/saxifrage /

ENTRYPOINT [ "/saxifrage" ]

CMD ["version"]
