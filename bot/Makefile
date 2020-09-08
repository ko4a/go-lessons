DC=docker-compose
DCEXEC_GO=${DC} exec go

up:
	${DC} up -d --build

up_local:
	env $(cat .env.local) ${DC} up -d --build

down:
	${DC} down

bash:
	${DCEXEC_GO} bash

deps:
	${DCEXEC_GO} go get .

start:
	${DCEXEC_GO} go run *.go

csfix:
	${DCEXEC_GO} gofmt -w .