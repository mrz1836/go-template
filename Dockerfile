FROM scratch
COPY go-template /
ENTRYPOINT ["/go-template"]
