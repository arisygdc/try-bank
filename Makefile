pgName := try_bank-db
pgUser := postgres
pgPassword := secret
databaseName := bank
timeZone := Asia/Jakarta
migrate := migrate -path database/postgresql/migration/ -database "postgresql://${pgUser}:${pgPassword}@localhost:5432/${databaseName}?sslmode=disable" -verbose up

testPgname := try_bank-testDB
testPgUser := postgresTest
testmigrate := migrate -path database/postgresql/migration/ -database "postgresql://${testPgUser}:${pgPassword}@localhost:5432/${databaseName}?sslmode=disable" -verbose up
testDB := docker run -d --name ${testPgname} -p 5432:5432 \
	-e POSTGRES_USER=${testPgUser} -e POSTGRES_PASSWORD=${pgPassword} \
	-e POSTGRES_DB=${databaseName} \
	-e TZ=${timeZone} -e PGTZ=${timeZone} \
	postgres:12-alpine3.14



installpg:
	docker run -d --name ${pgName} -p 5432:5432 \
	-e POSTGRES_USER=${pgUser} -e POSTGRES_PASSWORD=${pgPassword} \
	-e POSTGRES_DB=${databaseName} \
	-e TZ=${timeZone} -e PGTZ=${timeZone} \
	postgres:12-alpine3.14

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
	${migrate}

migratedown:
	migrate -path database/postgresql/migration/ -database "postgresql://${pgUser}:${pgPassword}@localhost:5432/${databaseName}?sslmode=disable" -verbose down

testup:
	${testDB} && \
	sleep 3 && \
	${testmigrate}

testdown:
	docker stop ${testPgname} && \
	docker container rm ${testPgname}

.PHONY: installpg uninstallpg startpg stoppg execdb createmigrate migrateup migratedown testup testdown