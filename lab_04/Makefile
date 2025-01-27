docker-run:
	docker compose up -d

docker-rm:
	docker compose down

run-e2e:
	go test --count=1 -v ./...

.PHONY: docker-run docker-rm run-e2e