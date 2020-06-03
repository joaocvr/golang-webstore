createDB:
	@docker volume create postgres_storage_alura
	@docker run -d --name postgres-alura --volume postgres_storage_alura:/var/lib/postgresql/data -p 5432:5432 -e POSTGRES_PASSWORD=postgres00 postgres:9.6
	@docker run --rm -ti --link postgres-alura:db -e PGPASSWORD=postgres00 -v $PWD:/bkp -w /bkp postgres:9.6 psql -h db -U postgres -c "CREATE USER webstore WITH SUPERUSER PASSWORD 'webstore00'"
	@docker run --rm -ti --link postgres-alura:db -e PGPASSWORD=postgres00 -v $PWD:/bkp -w /bkp postgres:9.6 psql -h db -U postgres -c "CREATE DATABASE webstore WITH OWNER webstore"
