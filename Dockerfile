FROM golang:1.14.9-alpine 
RUN mkdir /build
ADD go.mod go.sum hello.go /build/
WORKDIR /build
RUN go build

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=0 /build/hello /app/
COPY views/ /app/views
WORKDIR /app
EXPOSE 3000
CMD ["./hello"]
