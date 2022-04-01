FROM golang  
WORKDIR /work
ADD . .
# RUN go test ./...
RUN go build 
WORKDIR /
RUN rm -r /work
CMD ["kubesheet"] 
