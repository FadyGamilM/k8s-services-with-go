FROM golang:1.21-alpine as build
# -> i will use buster because the static linking stuff
# FROM golang:1.19-buster as build
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# setup the non-root user
# Setup the non-root user using adduser
RUN adduser -D -u 1001 nonroot

COPY go.mod go.sum /app/

RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build go mod download

COPY . /app/

RUN go build -ldflags="-linkmode external -extldflags -static" -tags netgo -o api-golang

# This stage will be used in deployment
FROM scratch

# For production environment
ENV GIN_MODE release

EXPOSE 8080

# for security presepective
COPY --from=build /etc/passwd /etc/passwd

COPY --from=build /app/api-golang api-golang

COPY --from=build /app/api-golang api-golang

CMD ["/api-golang"]
