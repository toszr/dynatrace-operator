FROM gcr.io/distroless/base

LABEL name="Dynatrace Operator" \
      vendor="Dynatrace LLC" \
      maintainer="Dynatrace LLC" \
      version="1.x" \
      release="1" \
      url="https://www.dynatrace.com" \
      summary="Dynatrace is an all-in-one, zero-config monitoring platform designed by and for cloud natives. It is powered by artificial intelligence that identifies performance problems and pinpoints their root causes in seconds." \
      description="ActiveGate works as a proxy between Dynatrace OneAgent and Dynatrace Cluster"

ENV USER_UID=1001

COPY LICENSE /licenses/
COPY build/_output/bin /usr/local/bin

ENTRYPOINT ["/usr/local/bin/dynatrace-operator"]

USER ${USER_UID}