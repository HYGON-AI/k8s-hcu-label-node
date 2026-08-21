# SPDX-License-Identifier: Apache-2.0
# Copyright 2026 Hygon Information Technology Co., Ltd.

FROM ubuntu:22.04 AS label-node

COPY LICENSE  /licenses/

WORKDIR /root

RUN  apt update && apt install -y kmod pciutils

COPY hcu-label-node .

COPY internal/pkg/lib ./lib

RUN chmod +x /root/hcu-label-node \
    && ln -s /root/lib/librocm_smi64.so.2.8 /root/lib/librocm_smi64.so.2 \
    && ln -s /root/lib/librocm_smi64.so.2 /root/lib/librocm_smi64.so \
    && ln -s /root/lib/libhydmi.so.1.5 /root/lib/libhydmi.so.1 \
    && ln -s /root/lib/libhydmi.so.1 /root/lib/libhydmi.so \
    && ln -s /root/lib/libhydmi_mig.so.1.3 /root/lib/libhydmi_mig.so.1 \
    && ln -s /root/lib/libhydmi_mig.so.1   /root/lib/libhydmi_mig.so


ENV LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/opt/hyhal/lib::/root/lib

CMD ["./hcu-label-node", "-name", "-cu-count", "-vram", "-driver-version", "-logtostderr=true", "-stderrthreshold=INFO", "-v=0"]
