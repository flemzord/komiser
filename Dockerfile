FROM alpine
COPY komiser /
ENTRYPOINT ["/komiser", "start"]
