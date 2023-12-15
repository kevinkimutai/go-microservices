mysql:
	docker run --name mysql8.2 -p 3307:3306 -e MYSQL_ROOT_PASSWORD=P@ssw0rd -d mysql:8.2.0

createdb:
	docker exec -it mysql8.2 mysql -uroot -p -e "CREATE DATABASE simple_bank;"

dropdb:
	docker exec -it mysql8.2 mysql -uroot -p -e "DROP DATABASE simple_bank;"

migrateup:
	migrate -path database/migrations -database "mysql://root:P@ssw0rd@tcp(localhost:3307)/simple_bank" -verbose up

migratedown:
	migrate -path database/migrations -database "mysql://root:P@ssw0rd@tcp(localhost:3307)/simple_bank" -verbose down

.PHONY: mysql createdb dropdb migrateup migratedown