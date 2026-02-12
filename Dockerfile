### Build Go binary
ARG GOLANG_VERSION=1.24.11
ARG APP_NAME=app
ARG TZ=Europe/Moscow

FROM golang:${GOLANG_VERSION} AS golang_build

ARG APP_NAME
ARG TZ
ENV TZ=$TZ
ENV CGO_ENABLED=1


WORKDIR /${APP_NAME}

RUN echo "Run building..."
COPY ./go.mod /${APP_NAME}/go.mod
COPY ./go.sum /${APP_NAME}/go.sum
RUN go mod tidy
COPY . /${APP_NAME}
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN export PATH=$PATH:$(go env GOPATH)/bin

RUN swag init -g ./cmd/app/main.go -o /${APP_NAME}/docs --parseDependency --parseInternal || true && \
    swag init -g ./cmd/app/main.go -o /${APP_NAME}/docs --parseDependency --parseInternal

RUN go build -o /${APP_NAME}/bin/${APP_NAME} /${APP_NAME}/cmd/app/main.go

### Run binary in Alpine
FROM alpine:3.22
ARG APP_NAME
ENV APP_NAME=$APP_NAME

RUN apk -U upgrade \
    && apk add --no-cache tzdata libc6-compat ca-certificates \
    && rm -rf /var/cache/apk/ \
    && mkdir -p /${APP_NAME}/bin \
    && adduser -g "${APP_NAME}" -h "/${APP_NAME}/bin/" -u 10001 -D ${APP_NAME} \
    && chown -R ${APP_NAME}:${APP_NAME} /${APP_NAME}/bin \
    && update-ca-certificates

COPY --from=golang_build --chown=${APP_NAME}:${APP_NAME} --chmod=755 /${APP_NAME}/bin/${APP_NAME} /${APP_NAME}/bin/${APP_NAME}

### Естественно это плохо и делать так нельзя (хранить креды в репо), но это в качестве демонстрации
COPY ./config.yaml /${APP_NAME}/bin/config.yaml

WORKDIR /${APP_NAME}/bin/
USER 10001

ENTRYPOINT /${APP_NAME}/bin/${APP_NAME}