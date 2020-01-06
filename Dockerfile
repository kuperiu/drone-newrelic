FROM alpine:latest
ADD build/drone-newrelic /bin/
ENTRYPOINT ["/bin/drone-newrelic"]