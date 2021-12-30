FROM node:15.11.0 as builder
ARG VUE=/usr/src/vue
COPY ./dashboard $VUE
WORKDIR $VUE
RUN npm config set registry https://registry.npm.taobao.org
RUN npm config set ignore-engines true
RUN npm install && npm run build:prod

FROM golang:alpine AS development
WORKDIR $GOPATH/src
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"
ENV CGO_ENABLED=1
COPY ./admin .

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add build-base
RUN go mod tidy & go mod vendor
RUN go build -a -ldflags '-extldflags "-static"' -o ./bin/admin  ./cmd/main.go


FROM alpine:latest AS production
WORKDIR /opt/shenshu
COPY --from=development /go/src/bin/admin .
RUN mkdir etc
COPY --from=development /go/src/etc/config.yaml etc/
COPY --from=development /go/src/etc/auth.json etc/
COPY --from=development /go/src/etc/basic_model.conf etc/
COPY --from=development /go/src/etc/basic_policy.csv etc/

RUN mkdir -p dashboard/dist
COPY --from=builder /usr/src/vue/dist dashboard/dist

EXPOSE 8080
ENTRYPOINT ["./admin", "-config", "etc/config.yaml"]

