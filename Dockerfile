FROM golang:1.17.9-alpine3.14 as builder

WORKDIR /app

COPY . .

# Install go-swagger and update the API docs
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN swag init -g ./cmd/api/main.go

# Set environment variables for build
RUN go env -w CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Finally, do the build
RUN go build -a -tags netgo \
    -ldflags '-w -extldflags "-static"'\
    -o /app/api.bin ./cmd/api/main.go

RUN go build -a -tags netgo \
    -ldflags '-w -extldflags "-static"'\
    -o /app/migration.bin ./cmd/migration/main.go


FROM alpine:3.15 as release

RUN apk add --no-cache bash \
    && adduser --disabled-password --gecos "" --no-create-home app

COPY --from=builder --chown=app /app/*.bin /app/

WORKDIR /app

USER app

CMD ./migration.bin && ./api.bin
