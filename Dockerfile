FROM golang:alpine as build_base
RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /app
# Force the go compiler to use modules
ENV GO111MODULE=on
# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
#COPY go.sum .
RUN go mod download

# This image builds the weavaite server
FROM build_base AS server_builder
# Here we copy the rest of the source code
COPY . .
ENV GOOS=linux
ENV GOARCH=amd64
#RUN go build -o /app  get_url_one_param.go -tags netgo -ldflags '-w -extldflags "-static"' .
RUN go build -o app get_url_one_param.go

### Put the binary onto base image
FROM plugins/base:linux-amd64
LABEL maintainer="Pgluffy <kuang7156@gmail.com>"
EXPOSE 18080
WORKDIR /app
COPY --from=server_builder /app /app
ENTRYPOINT ./app
#CMD ["/app"]