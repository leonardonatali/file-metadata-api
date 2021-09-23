default: run

run:
	@docker-compose -f docker-compose.yml up --build

test: 
	@docker-compose -f docker-compose-test.yml run --rm api go test -v /app/pkg/test/...

migrations:
	@go-bindata \
	-ignore=\\.go \
	-pkg=migrations \
	-prefix "pkg/database/migrations/" \
	-o $(PWD)/pkg/database/migrations/migrations.go \
	pkg/database/migrations/... 
