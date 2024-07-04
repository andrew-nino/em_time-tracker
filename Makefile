THIS_FILE := $(lastword $(MAKEFILE_LIST))

.PHONY:  start stop 

help:
	make -pRrq  -f $(THIS_FILE) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'

start:
	docker run -d -e POSTGRES_PASSWORD=qwerty -p 5432:5432 --name postgres 69e765e8cdbe && go run cmd/app/main.go

stop:
	docker stop postgres 
