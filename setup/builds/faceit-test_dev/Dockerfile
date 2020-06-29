# Compile stage
FROM golang:1.14.4-alpine3.12
ENV CGO_ENABLED 0

RUN apk add --no-cache bash
RUN apk add --no-cache git
RUN go get github.com/derekparker/delve/cmd/dlv

RUN apk add --no-cache libc6-compat
RUN echo 'alias build="cd /faceit-test;go build -gcflags \"all=-N -l\" -o /faceit github.com/skyerus/faceit-test"' >> ~/.bashrc
RUN echo 'alias delve="go/bin/dlv --listen=:40000 --headless=true --api-version=2 exec /faceit"' >> ~/.bashrc
RUN echo 'alias run="build; /faceit"' >> ~/.bashrc
RUN echo 'alias run-tests="cd /faceit-test; go test -v"' >> ~/.bashrc
RUN echo 'alias debug="build; delve"' >> ~/.bashrc
WORKDIR /
