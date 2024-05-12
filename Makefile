postgres:
	docker run --name sweet_boyd -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.3-alpine3.19

createdb: 
	docker exec -it sweet_boyd createdb --username=root --owner=root vacation_planner

drobdb:
	docker exec -it sweet_boyd dropdb vacation_planner

migrateup:
	migrate -path db/migration --database "postgresql://root:secret@localhost:5431/vacation_planner?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration --database "postgresql://root:secret@localhost:5431/vacation_planner?sslmode=disable" -verbose up 1
	
migratedown:
	migrate -path db/migration --database "postgresql://root:secret@localhost:5431/vacation_planner?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration --database "postgresql://root:secret@localhost:5431/vacation_planner?sslmode=disable" -verbose down 1
	


	

