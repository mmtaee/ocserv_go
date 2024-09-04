## Database
```bash
sudo docker run -d \
  --name ocserv \
  -e POSTGRES_USER=ocserv \
  -e POSTGRES_PASSWORD=ocserv \
  -e POSTGRES_DB=ocserv \
  -v /home/masoud/volumes:/var/lib/postgresql/data \
  -p 5435:5432 \
  postgres:latest
```

