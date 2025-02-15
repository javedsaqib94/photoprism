#### Base Image: Debian 12, Codename "Bookworm"
FROM debian:bookworm-slim

LABEL maintainer="Michael Mayer <hello@photoprism.app>"

ARG TARGETARCH
ARG BUILD_TAG

# set environment variables, see https://docs.photoprism.app/getting-started/config-options/
ENV PHOTOPRISM_ARCH=$TARGETARCH \
    DOCKER_TAG=$BUILD_TAG \
    DOCKER_ENV="prod" \
    PATH="/usr/local/sbin:/usr/sbin:/sbin:/usr/local/bin:/usr/bin:/bin:/scripts:/opt/photoprism/bin" \
    TMPDIR="/tmp" \
    DEBIAN_FRONTEND="noninteractive" \
    TF_CPP_MIN_LOG_LEVEL="2"

# copy scripts and debian backports sources list
COPY --chown=root:root --chmod=755 /scripts/dist/* /scripts/
COPY --chown=root:root --chmod=644 /docker/develop/bookworm/sources.list /etc/apt/sources.list.d/bookworm.list

# install additional distribution packages
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
    echo "ALL ALL=(ALL) NOPASSWD:SETENV: /scripts/entrypoint-init.sh" >> /etc/sudoers.d/init && \
    cp /scripts/heif-convert.sh /usr/local/bin/heif-convert && \
    /scripts/cleanup.sh

# define default directory and user
WORKDIR /photoprism

# expose default http port 2342
EXPOSE 2342

# keep container running
CMD ["tail", "-f", "/dev/null"]