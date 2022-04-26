pgName := try_bank-db
pgUser := postgres
pgPassword := secret
databaseName := bank

installpg:
	docker run -d --name ${pgname} -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=${pgPassword} -e POSTGRES_DB=${databaseName} postgres:12-alpine3.14

uninstallpg:
	docker container rm ${pgname}

startpg:
	docker start ${pgname}

stoppg:
	docker stop ${pgname}

execpg:
	docker exec -it ${pgname} psql -U postgres

createmigrate:
	migrate create -ext sql -dir database/migration -seq init_schema

migrateup:
	migrate -path database/postgresql/migration/ -database "postgresql://${pgUser}:${pgPassword}@localhost:5432/${databaseName}?sslmode=disable" -verbose up

migratedown:
	migrate -path database/postgresql/migration/ -database "postgresql://${pgUser}:${pgPassword}@localhost:5432/${databaseName}?sslmode=disable" -verbose down
  
.PHONY: installpg uninstallpg startpg stoppg execdb createmigrate migrateup migratedown