#### Base Image: Debian 12, Codename "Bookworm"
FROM debian:bookworm-slim

LABEL maintainer="Michael Mayer <hello@photoprism.app>"

ARG TARGETARCH
ARG BUILD_TAG

# set environment variables
ENV PHOTOPRISM_ARCH=$TARGETARCH \
    DOCKER_TAG=$BUILD_TAG \
    DOCKER_ENV="develop" \
    NODE_ENV="production" \
    DEBIAN_FRONTEND="noninteractive" \
    TMPDIR="/tmp" \
    LD_LIBRARY_PATH="/root/.local/lib:/usr/local/lib:/usr/lib" \
    TF_CPP_MIN_LOG_LEVEL="0" \
    GOPATH="/go" \
    GOBIN="/go/bin" \
    PATH="/usr/local/sbin:/usr/sbin:/sbin:/usr/local/bin:/usr/bin:/bin:/scripts:/usr/local/go/bin:/go/bin" \
    GO111MODULE="on" \
    CGO_CFLAGS="-g -O2 -Wno-return-local-addr"

# copy scripts and debian backports sources list
COPY --chown=root:root --chmod=755 /scripts/dist/* /scripts/
COPY --chown=root:root --chmod=644 /docker/develop/bookworm/sources.list /etc/apt/sources.list.d/bookworm.list
COPY --chown=root:root --chmod=644 /.my.cnf /etc/my.cnf

# update image and install build dependencies
RUN echo 'APT::Acquire::Retries "3";' > /etc/apt/apt.conf.d/80retries && \
    echo 'APT::Install-Recommends "false";' > /etc/apt/apt.conf.d/80recommends && \
    echo 'APT::Install-Suggests "false";' > /etc/apt/apt.conf.d/80suggests && \
    echo 'APT::Get::Assume-Yes "true";' > /etc/apt/apt.conf.d/80forceyes && \
    echo 'APT::Get::Fix-Missing "true";' > /etc/apt/apt.conf.d/80fixmissing && \
    groupadd -f -r -g 44 video && \
    groupadd -f -r -g 109 render && \
    groupadd -f -g 1000 photoprism && \
    useradd -m -g 1000 -u 1000 -d /photoprism -G video,render photoprism && \
    chmod 777 /photoprism && \
    apt-get update && apt-get -qq dist-upgrade && apt-get -qq install --no-install-recommends \
      libc6 ca-certificates sudo bash tzdata \
      gpg zip unzip wget curl rsync make nano \
      jq lsof lshw sqlite3 mariadb-client \
      exiftool darktable rawtherapee libheif-examples librsvg2-bin \
      ffmpeg ffmpegthumbnailer libavcodec-extra libwebm1 \
      libmatroska7 libdvdread8 libebml5 libgav1-0 libatomic1 \
      libx264-163 libx265-199 && \
    apt-get -qq install --no-install-recommends \
      apt-utils pkg-config software-properties-common \
      build-essential gcc g++ git gettext davfs2 chrpath apache2-utils \
      chromium chromium-driver chromium-sandbox firefox-esr \
      libx264-dev libx265-dev libpng-dev libxft-dev \
      libc6-dev libhdf5-serial-dev libzmq3-dev libssl-dev libnss3 \
      libfreetype6 libfreetype6-dev libfontconfig1 libfontconfig1-dev fonts-roboto \
    && \
    /scripts/install-nodejs.sh && \
    /scripts/install-tensorflow.sh && \
    /scripts/install-go.sh && \
    /scripts/cleanup.sh && \
    mkdir -p "/go/src" "/go/bin" && \
    chmod -R 777 "/go" && \
    install -d -m 0777 -o 1000 -g 1000 \
      /var/lib/photoprism \
      /tmp/photoprism \
      /photoprism/originals \
      /photoprism/import \
      /photoprism/storage \
      /photoprism/storage/sidecar \
      /photoprism/storage/albums \
      /photoprism/storage/backups \
      /photoprism/storage/config \
      /photoprism/storage/cache \
    && \
    wget "https://dl.photoprism.app/tensorflow/nsfw.zip?${BUILD_TAG}" -O /tmp/photoprism/nsfw.zip && \
    wget "https://dl.photoprism.app/tensorflow/nasnet.zip?${BUILD_TAG}" -O /tmp/photoprism/nasnet.zip && \
    wget "https://dl.photoprism.app/tensorflow/facenet.zip?${BUILD_TAG}" -O /tmp/photoprism/facenet.zip && \
    wget "https://dl.photoprism.app/qa/testdata.zip?${BUILD_TAG}" -O /tmp/photoprism/testdata.zip

# install Go tools
RUN /usr/local/go/bin/go install github.com/tianon/gosu@latest && \
    /usr/local/go/bin/go install golang.org/x/tools/cmd/goimports@latest && \
    /usr/local/go/bin/go install github.com/kyoh86/richgo@latest && \
    /usr/local/go/bin/go install github.com/psampaz/go-mod-outdated@latest && \
    /usr/local/go/bin/go install github.com/dsoprea/go-exif/v3/command/exif-read-tool@latest; \
    cp /go/bin/gosu /sbin/gosu && \
    chown root:root /sbin/gosu && \
    chmod 755 /sbin/gosu && \
    echo "alias go=richgo ll='ls -alh'" > /photoprism/.bash_aliases && \
    echo "alias go=richgo ll='ls -alh'" > /root/.bash_aliases && \
    echo "ALL ALL=(ALL) NOPASSWD:SETENV: ALL" >> /etc/sudoers.d/all && \
    cp /scripts/heif-convert.sh /usr/local/bin/heif-convert && \
    chmod -R a+rwX /go

# set up project directory
WORKDIR "/go/src/github.com/photoprism/photoprism"

# expose the following container ports:
# - 2342 (HTTP)
# - 2343 (Acceptance Tests)
# - 9515 (Chromedriver)
# - 40000 (Go Debugger)
EXPOSE 2342 2343 9515 40000

# set container entrypoint script
ENTRYPOINT ["/scripts/entrypoint.sh"]

# keep container running
CMD ["tail", "-f", "/dev/null"]
