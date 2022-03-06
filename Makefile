all :: build

build ::
	CGO_ENABLED=0 go build

install :: build
	mv hex /usr/local/bin

