# Nombre del binario de la aplicación
APP_NAME=my-app

# Target por defecto
all: build

# Levanta la base de datos usando docker-compose
up:
	@echo ">= Levantando base de datos..."
	@docker compose up -d

# Apaga la base de datos
down:
	@echo ">= Bajando base de datos..."
	@docker compose down

# Esperar a la base de datos
wait-db:
	@echo ">= Esperando a que Postgres esté listo..."
	@until docker exec postgresTP6 pg_isready -U user; do sleep 1; done

# Genera código SQLC
generate:
	@echo ">= Generando SQLC..."
	@sqlc generate

# Construye el binario
build: generate
	@echo ">= Compilando aplicación..."
	@go build -o $(APP_NAME) .

# Corre la app localmente (requiere que la base ya esté levantada)
run: build
	@echo ">= Ejecutando aplicación..."
	@./$(APP_NAME)

# Limpia binario
clean:
	@rm -f $(APP_NAME)

# Target completo: levanta DB, espera a que este lista, genera código y corre app
dev: up wait-db run
