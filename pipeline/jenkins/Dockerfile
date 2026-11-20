# Jenkins controller that can run the repo's make targets itself (the pipeline uses the
# controller as its agent; the stock image has neither make nor Go)
FROM jenkins/jenkins:2.580.2-jdk25

USER root
RUN apt-get update \
    && apt-get install -y --no-install-recommends make \
    && rm -rf /var/lib/apt/lists/*

COPY --from=quay.io/projectquay/golang:1.25 /usr/local/go /usr/local/go
ENV PATH=/usr/local/go/bin:$PATH

USER jenkins
