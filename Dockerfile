####    Build stage    ####
FROM golang:1.23.4-alpine AS builder

RUN apk add --no-cache \
	gcc \
	musl-dev \
	make

WORKDIR /workspace

COPY . .

RUN go mod download

RUN make build


####    Deploy stage    ####
FROM alpine

WORKDIR /app

# install a libc compatibility layer
RUN apk add --no-cache \
	gcompat

COPY --from=builder \
	/workspace/build/api-server \
	/app/api-server

ENTRYPOINT ["/app/api-server"]
