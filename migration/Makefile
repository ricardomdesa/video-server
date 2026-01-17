createmigrations: 
	@echo "running create migration..."
	/opt/homebrew/bin/docker run --rm -v ./sql/migrations:/migrations --network host migrate/migrate create -ext sql -dir /migrations/migrations_sqlite 001_create_table.sql
# /opt/homebrew/bin/docker run -rm -v ./sql/migrations:/migrations --network host migrate/migrate create -ext=sql -dir=/migrations/ -seq init

	
migrateup:
	@echo "running migrate"
	go run sql/migrate/main.go up
# /opt/homebrew/bin/docker run -rm -v ./sql/migrations:/migrations --network host migrate/migrate --database postgres://postgres:bR632F9EYioo@localhost:5432/course\?sslmode=disable -path=/migrations/ up


migratedown:
	go run sql/migrate/main.go down
# 	docker run -v sql/migrations\:/migrations --network host migrate/migrate -path=sql/migrations -database "mysql\://root\:root@tcp(127.0.0.1\:3306)/courses" -verbose down

# .PHONY migrate createmigration migratedown

sqlcgen:
	/opt/homebrew/bin/docker run --rm -v ./sql:/src -w /src sqlc/sqlc generate


buildrun:
	go build -o app migration.go && ./app

gorun:
	go run migration.go

