FROM alpine:3.12

RUN mkdir "/app"
WORKDIR "/app"

COPY docker-demo /app/app

ENTRYPOINT ["./app"]