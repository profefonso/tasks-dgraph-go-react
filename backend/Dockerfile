FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /usr/src/

COPY go.mod .
COPY go.sum .
RUN go mod download

RUN apk add --update \
    curl \
    && rm -rf /var/cache/apk/*

COPY . .

# RUN cp /build/main .

EXPOSE 9075

COPY entrypoint.sh .

COPY schema.graphql .

RUN ["chmod", "+x", "/usr/src/entrypoint.sh"]

ENTRYPOINT ["/usr/src/entrypoint.sh"]
