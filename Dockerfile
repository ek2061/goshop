FROM golang

MAINTAINER github.com/Milesfeng

RUN apt-get update

RUN apt-get -y install git 

RUN git clone https://github.com/ek2061/goshop.git

#RUN cd goshop

WORKDIR /go/goshop

EXPOSE 5000

CMD go build main.go && ./main




