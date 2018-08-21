FROM debian:stretch

RUN apt-get update
RUN apt-get upgrade -y
RUN apt-get install -y git gcc gawk curl make texinfo autoconf automake autotools-dev libmpc-dev libmpfr-dev libgmp-dev build-essential bison flex gperf libtool patchutils bc zlib1g-dev libexpat-dev
RUN useradd -m -s /sbin/nologin builder
RUN mkdir /home/builder/toolchain
RUN git clone --recursive https://github.com/riscv/riscv-gnu-toolchain
RUN mkdir /build
RUN cd /build; /riscv-gnu-toolchain/configure --prefix=/home/builder/toolchain --with-guile=no
RUN cd /build; make -j $(shell nproc)
RUN rm -rf /build /riscv-gnu-toolchain
