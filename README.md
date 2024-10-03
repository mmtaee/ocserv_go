### Database
```shell
sudo docker run -d \
  --name ocserv \
  -e POSTGRES_USER=ocserv \
  -e POSTGRES_PASSWORD=ocserv \
  -e POSTGRES_DB=ocserv \
  -v /home/masoud/volumes:/var/lib/postgresql/data \
  -p 5435:5432 \
  postgres:latest
```

### .env file
```dotenv
DEBUG=false
SECRET_KEY=1234
HOST=0.0.0.0
PORT=8080
ALLOW_ORIGINS=


POSTGRES_HOST=127.0.0.1
POSTGRES_PORT=5435
POSTGRES_NAME=ocserv
POSTGRES_USER=ocserv
POSTGRES_PASSWORD=ocserv
POSTGRES_SSL_MODE=disable
```

### Commands
```shell
go run main.go migrate

```

### Test
```shell
export $(cat .env | xargs) && go test -v ./...

# verbose mode
go run main.go test -v

# benchmark testing
go run main.go test -b 

GIN_MODE=release CGO_ENABLED=0 GOOS=linux go build -o api main.go

# minor or major
go run main.go bump -m/--minor -j/--major

# patch
go run main.go bump

```