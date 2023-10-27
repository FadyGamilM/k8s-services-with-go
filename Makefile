# i gave my postgres container name = db 
# then i used @db in the connection string to represent the host of this database or the ip of the host of this database
DATABASE_URL:=postgres://postgres:pswd@db:5432/postgres


# now i set this database url as env variable 
# ➜ docker-with-golang  $env:DATABASE_URL= "postgres://postgres:pswd@db:5432/postgres"
# ➜ docker-with-golang  echo $env:DATABASE_URL
#    postgres://postgres:pswd@db:5432/postgres

docker-run-postgres:
	docker run -d --name db --network trial-network -e POSTGRES_PASSWORD=pswd -v pgdata:/var/lib/postgresql/data -p 5432:5432 --restart unless-stopped postgres:14

docker-run-golang-api:
	docker run -d --name api-golang --network trial-network -e DATABASE_URL=${DATABASE_URL} -p 8080:8080 --restart unless-stopped --link=db golang-api-image:2.0

docker-compose-up:
	docker-compose up