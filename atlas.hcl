data "external_schema" "development" {
  program = ["go", "run", "main.go", "make-migrations"]
}

env "development" {
  src = data.external_schema.development.url
  dev = "postgres://ocserv:ocserv@:5435/ocserv?sslmode=disable"
  migration {
    dir = "file://migrations"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

data "external_schema" "production" {
  program = ["/app", "make-migrations"]
}

env "production" {
  src = data.external_schema.production.url
  dev = "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_NAME}?sslmode=${POSTGRES_SSL_MODE}"
  migration {
    dir = "file://migrations"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}