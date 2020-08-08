FROM alpine:latest
LABEL version="v1.0"
LABEL description="ma-novel-crawler-api"

WORKDIR /data/app
COPY ./build/ma-novel-crawler-api ./ma-novel-crawler-api

EXPOSE 8086
ENTRYPOINT ["./ma-novel-crawler-api"]