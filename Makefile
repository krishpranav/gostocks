NAME := gostocks

build:
	go mod tidy 
	go build -v -o ${NAME}

install:
	go install -v

clean:
	rm gostocks

clear:
	rm -f ${GOPATH}/bin/gostocks
