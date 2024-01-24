FROM alpine

WORKDIR /data/apps/gutiktok-service-bundle/static

RUN apk update --no-cache \
    && apk upgrade \
    && apk add yasm \
    && apk add ffmpeg \
    && rm -rf /var/cache/apk/*

COPY ./static .