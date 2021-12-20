installpg:
	docker run -d --name try_bank -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=bank postgres:12-alpine3.14

uninstallpg:
	docker container rm try_bank

startpg:
	docker start try_bank

stoppg:
	docker stop try_bank

execpg:
	docker exec -it try_bank psql -U postgres

createmigrate:
	migrate create -ext sql -dir database/migration -seq init_schema

migrateup:
	migrate -path database/migration/ -database "postgresql://postgres:secret@localhost:5432/bank?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migration/ -database "postgresql://postgres:secret@localhost:5432/bank?sslmode=disable" -verbose down
  
.PHONY: installpg uninstallpg startpg stoppg execdb createmigrate migrateup migratedown