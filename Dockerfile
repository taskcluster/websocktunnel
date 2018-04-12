FROM golang

RUN mkdir -p /go/src/github.com/
WORKDIR /go/src/github.com/taskcluster/
# clone and run webhooktunnel
RUN git clone http://github.com/taskcluster/webhooktunnel
WORKDIR /go/src/github.com/taskcluster/webhooktunnel

# set envs
ENV HOSTNAME=tcproxy.dev
ENV SECRET_A=example-secret
ENV SECRET_B=another-example-secret

RUN go get -v
ENTRYPOINT ["go", "run", "main.go"]
# expose ports when starting container
