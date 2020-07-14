include .env

.PHONY: new-migration run-migration

VARS:=$(shell sed -ne 's/ *\#.*$$//; /./ s/=.*$$// p' .env )
$(foreach v,$(VARS),$(eval $(shell echo export $(v)="$($(v))")))

ndef = $(if $(value $(1)),,$(error $(1) not set))

run:
	wire && go run main.go wire_gen.go

run-migration:
	docker-compose run --rm migrate -database ${DATABASE_URL} -path /migrations up

new-migration:
	$(call ndef,MIGRATION_NAME)
	docker-compose run --rm migrate create -ext sql -dir /migrations -seq ${MIGRATION_NAME}
