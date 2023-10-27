FROM golang:1.21-alpine AS build-base
# -> i will use buster because the static linking stuff
# FROM golang:1.19-buster as build
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum /app/

RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build go mod download

###
FROM build-base AS dev

# install two dependencies for dev only
# air => hot reload
# delve for debugging
RUN go install github.com/cosmtrek/air@latest && go install github.com/go-delve/delve/cmd/dlv@latest

COPY . /app/

CMD [ "air", "-c", ".air.toml" ]

###
# now build the production image and copy the result in the sctach layer (The final)
FROM build-base AS build-prod

# setup the non-root user
# Setup the non-root user using adduser
RUN adduser -D -u 1001 nonroot

COPY . /app/

RUN go build -ldflags="-linkmode external -extldflags -static" -tags netgo -o api-golang

# This stage will be used in deployment
FROM scratch

# For production environment
ENV GIN_MODE release

EXPOSE 5000

# for security presepective
COPY --from=build-prod /etc/passwd /etc/passwd

COPY --from=build-prod /app/api-golang api-golang

COPY --from=build-prod /app/api-golang api-golang

CMD ["/api-golang"]
