## Database
```bash
sudo docker run -d \
  --name ocserv \
  -e POSTGRES_USER=ocserv \
  -e POSTGRES_PASSWORD=ocserv \
  -e POSTGRES_DB=ocserv \
  -p 5435:5432 \
  postgres:latest
```

## Atlas
### install:
```bash
curl  https://atlasgo.sh -o /tmp/atlasgo.sh
chmod +x /tmp/atlasgo.sh
/tmp/atlasgo.sh --yes
```

### check migration files
```bash
atlas migrate diff --env development --dev-url "postgres://ocserv:ocserv@:5435/ocserv?sslmode=disable"
```
