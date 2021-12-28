NAME := gostocks

build:
	go get
	go build -v -o ${NAME}

install:
	go install -v

clearn:
	rm -f ${GOPATH}/bin/gostocks