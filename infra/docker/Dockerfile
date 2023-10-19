FROM golang:1.21.1-alpine3.18 as build_sales-api
ENV CGO_ENABLED 0
ARG BUILD_REF

RUN mkdir /service
COPY go.* /service/
WORKDIR /service
RUN go mod download

# copy the source code
COPY main.go /service

# build the service binary
WORKDIR /service
RUN go build -ldflags "-X main.build=${BUILD_REF}"

# The Run stage
FROM alpine:3.18
ARG BUILD_DATE
ARG BUILD_REF
COPY --from=build_sales-api /service /service/service
WORKDIR /service

CMD [ "./service" ]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
    org.opencontainers.image.title="sales-api" \
    org.opencontainers.image.authors="fady gamil <gamilfady605@gmail.com>" \
    org.opencontainers.image.revision="${BUILD_REF}" \
    org.opencontainers.image.vendor="BACKEND"
