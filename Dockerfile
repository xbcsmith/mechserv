ARG GO_VERSION=1.15

FROM golang:${GO_VERSION} as builder

WORKDIR /build
COPY . .
RUN make clean && make arm64

RUN mkdir etc
RUN echo "astromech:x:10001:10001:Polaris User:/home/astromech:/bin/false" > etc/passwd \
  && echo "astromech:x:10001:astromech" > etc/group

FROM scratch

COPY --from=builder /build/etc /etc/
COPY --chown=astromech:astromech --from=builder /build/bin/mechserv-arm64 /usr/local/bin/mechserv
COPY --chown=astromech:astromech --from=builder /build/resources /resources

EXPOSE 9423
USER astromech
CMD ["/usr/local/bin/mechserv"]
