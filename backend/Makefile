
MIGRATION_FOLDER = ./migrations

local_migration_up:
	goose -dir $(MIGRATION_FOLDER) postgres "user=postgres password=admin database=homeworkmanager host=localhost port=5431" up

local_migration_down:
	goose -dir $(MIGRATION_FOLDER) postgres "user=postgres password=admin database=homeworkmanager host=localhost port=5431" down




