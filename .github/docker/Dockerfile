FROM ghcr.io/loongarchlinux/archlinux:latest
ARG TARGETARCH
RUN rm -f /root/.bash_history /var/log/pacman.log
RUN --mount=type=cache,target=/var/cache/pacmant,sharing=locked \
    --mount=type=cache,target=/var/lib/pacman/sync,sharing=locked \
    pacman --noconfirm -Syu
