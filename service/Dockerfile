# Build
FROM golang:1.15-alpine AS build

RUN apk update && apk upgrade && \
  apk add --no-cache bash openssh openssl-dev build-base openssh-client hdf5 hdf5-dev

ADD . /src
WORKDIR /src

RUN go mod tidy && go get -u all
RUN go build -ldflags="-s -w" -o service ./main.go

# Runtime
FROM alpine

RUN apk update && apk add --no-cache ca-certificates openssl openssh openssl-dev hdf5
WORKDIR /app

COPY --from=build /src/service /app/

COPY --from=build /src/artifacts /app/artifacts

CMD ["./service"]