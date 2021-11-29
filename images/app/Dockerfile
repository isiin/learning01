ARG SRC_IMAGE
FROM ${SRC_IMAGE}

# install extension
RUN docker-php-ext-install -j$(nproc) \
    mysqli \
;

# install msmtp
RUN apk add --no-cache msmtp;
