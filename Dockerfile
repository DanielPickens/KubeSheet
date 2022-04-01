FROM golang  
WORKDIR /work
ADD . .
# RUN go test ./...
RUN go build -o /kubesheet .
WORKDIR /
RUN rm -r /work
CMD ["kubesheet"] 
