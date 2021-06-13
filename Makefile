.PHONY: redis

redis:
	docker run --rm -ti --network host redis:alpine3.13

test:
	go test -v -cover ./...