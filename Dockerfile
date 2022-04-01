FROM golang  
WORKDIR /work
ADD . .

RUN go build -o /bin/kubesheet
WORKDIR /
RUN rm -r /work
CMD ["/bin/kubesheet"] 
