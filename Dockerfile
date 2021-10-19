FROM golang:alpine as builder
WORKDIR /app
COPY go.* *.go config.yaml ./
RUN go build
RUN ls -hal /app/redirects

FROM alpine
COPY --from=builder /app/redirects /app/redirects
EXPOSE 80
ENTRYPOINT [ "/app/redirects" ]
