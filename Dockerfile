FROM alpine

COPY gopath/bin/istio-test /go/bin/istio-test

ENTRYPOINT /go/bin/istio-test
