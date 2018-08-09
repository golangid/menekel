BINARY=menekel

build: test
	go build -o ${BINARY} github.com/golangid/menekel/app

test:
	./test_cover.sh

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean install unittest