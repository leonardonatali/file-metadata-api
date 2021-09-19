
default: 
	@export UID=$(UID); docker-compose up --build

migrations:
	@go-bindata \
	-ignore=\\.go \
	-pkg=migrations \
	-prefix "pkg/database/migrations/" \
	-o $(PWD)/pkg/database/migrations/migrations.go \
	pkg/database/migrations/... 
