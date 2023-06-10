FROM ubuntu:22.04 as builder

WORKDIR /root
RUN apt-get update && apt-get install -y wget

ARG SPIN_VERSION=canary

## spin
RUN wget https://github.com/fermyon/spin/releases/download/${SPIN_VERSION}/spin-${SPIN_VERSION}-linux-amd64.tar.gz &&         \
    tar -xvf spin-${SPIN_VERSION}-linux-amd64.tar.gz &&                                                                       \
    ls -ltr &&                                                                                                                \
    mv spin /usr/local/bin/spin;

RUN spin --version


ENV RUST_LOG='spin=trace'

COPY main.wasm main.wasm
COPY spin.toml spin.toml

ENTRYPOINT ["spin"]