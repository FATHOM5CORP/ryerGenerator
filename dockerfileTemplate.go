package main

type dockerfileDefiner struct {
	Name	string
	Port	int
}

var dockerfileTemplate = &Dockerfile {
	`FROM golang:1.13

	COPY ./{{.Name}} /go/src/app/
	
	EXPOSE {{.Port}}
	
	WORKDIR /go/src/app/
	
	# -d flag downloads package dependencies but does not
	# build or install them in the container
	RUN go get -d -v .
	RUN go install -v .
	
	CMD [./{{.Name}}]
	`,
}