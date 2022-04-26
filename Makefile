pgName := try_bank-db
pgUser := postgres
pgPassword := secret
databaseName := bank

installpg:
	docker run -d --name ${pgName} -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=${pgPassword} -e POSTGRES_DB=${databaseName} postgres:12-alpine3.14

uninstallpg:
	docker container rm ${pgName}

startpg:
	docker start ${pgName}

stoppg:
	docker stop ${pgName}

execpg:
	docker exec -it ${pgName} psql -U postgres

createmigrate:
	migrate create -ext sql -dir database/migration -seq init_schema

migrateup:
	migrate -path database/postgresql/migration/ -database "postgresql://${pgUser}:${pgPassword}@localhost:5432/${databaseName}?sslmode=disable" -verbose up

migratedown:
	migrate -path database/postgresql/migration/ -database "postgresql://${pgUser}:${pgPassword}@localhost:5432/${databaseName}?sslmode=disable" -verbose down
  
.PHONY: installpg uninstallpg startpg stoppg execdb createmigrate migrateup migratedown