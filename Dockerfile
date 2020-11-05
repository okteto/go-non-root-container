FROM golang:buster as builder
RUN addgroup --gid 1000 okteto && adduser --uid 1100 --gid 1000 okteto 

WORKDIR /usr/src/app
ADD . .
RUN go build -o /usr/local/bin/non-root

EXPOSE 8080
CMD ["/usr/local/bin/non-root"]