FROM registry.access.redhat.com/ubi8/go-toolset AS builder
ADD . /tmp/src/
USER 0
RUN chown -R 1001:0 /tmp/src
USER 1001
RUN  /usr/libexec/s2i/assemble

FROM registry.access.redhat.com/ubi8/ubi-minimal AS runner
RUN mkdir /app
WORKDIR /app

COPY --from=builder /opt/app-root/gobinary /app/container-helper
ENTRYPOINT ["/app/container-helper"]
CMD ["serve"]