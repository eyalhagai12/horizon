export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING="user=horizon-dev password=horizon-dev host=localhost port=5432 sslmode=disable"
export GOOSE_MIGRATION_DIR=./migrations

goose $@