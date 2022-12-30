db-migrate:
	goose -dir "./server/db/migrations" ${DB_DRIVER} "${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" up

swag-gen:
	swag init