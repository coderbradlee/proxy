FROM golang:1.13.5-stretch
ENV GO111MODULE=on
ENV PORT=8080
ENV ENDPOINT=did.iotex.one
WORKDIR apps/proxy
COPY . .
RUN go build -o /usr/local/bin/proxy .
RUN rm -fr $GOPATH $GOROOT ./*
CMD [ "proxy"]