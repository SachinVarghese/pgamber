build_image:
	docker build -t sachinmv31/postgres-pgamber:latest builder

test_setup_up:
	docker compose -f builder/docker-compose.yaml up -d

test_setup_down:
	docker compose -f builder/docker-compose.yaml down

purge_test_volume:
	docker volume rm builder_pgdata

ingest_test_data:
	cd setup && go run main.go && cd ..