FROM golang:1.7.1

COPY . /go/src/github.com/docker/go
WORKDIR /go/src/github.com/docker/go

EXPOSE 8000
EXPOSE 8443
EXPOSE 3000

CMD ["./appbin", "kubesheet"]
