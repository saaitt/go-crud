#Starting DB
```
DB_PASSWORD=<some pass> docker-compose up
```
#Dropping DB
```bigquery
DB_PASSWORD=<some pass> docker-compose down
```
#Starting App
```bigquery
GO111MODULE=on DB_PASSWORD=<some pass> go run main.go
```