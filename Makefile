dev:
	GOVENV=dev docker compose up -d
	GOVENV=dev air
tests:
	GOVENV=test go run ./cmd/cli/db create > /dev/null
	GOVENV=test go test -p 1 ./test/... $(EXTRA_ARGS)
tests-v:
	make tests EXTRA_ARGS
db-create:
	go run ./cmd/cli/db create
db-drop:
	go run ./cmd/cli/db drop
db-reset:
	go run ./cmd/cli/db reset