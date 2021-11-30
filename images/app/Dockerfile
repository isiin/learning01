FROM php:8.1-fpm-alpine

# install extension
RUN docker-php-ext-install -j$(nproc) \
    mysqli \
;

# install msmtp
RUN apk add --no-cache msmtp;
