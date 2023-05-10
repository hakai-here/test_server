```
docker-compose up -d 
go run main.go 
```
Both Redis and Postgres database are dockerized using `docker-compose` 

`/api/proceedingentries` - will list all the avaliable entries 

`/api/proceedingentries?id=XXXXX` - will list the specific entry 

