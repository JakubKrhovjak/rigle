# Rigle Multiprojekt

Multiprojekt obsahující dvě REST API aplikace:
- **rigle-java**: Quarkus REST API s PostgreSQL
- **rigle-go**: Gin REST API s PostgreSQL

## Struktura projektu

```
rigle/
├── rigle-java/          # Quarkus aplikace (port 8080)
├── rigle-go/            # Go Gin aplikace (port 8081)
└── docker-compose.yml   # PostgreSQL databáze
```

## Předpoklady

- Java 17+ (pro rigle-java)
- Maven (pro rigle-java)
- Go 1.21+ (pro rigle-go)
- Docker a Docker Compose

## Spuštění

### 1. Spuštění databází

```bash
docker-compose up -d
```

To spustí:
- PostgreSQL pro Java aplikaci na portu 5432 (databáze: `rigle_java`)
- PostgreSQL pro Go aplikaci na portu 5433 (databáze: `rigle_go`)

### 2. Spuštění Java aplikace

```bash
cd rigle-java
./mvnw quarkus:dev
```

Aplikace běží na `http://localhost:8080`

### 3. Spuštění Go aplikace

```bash
cd rigle-go
go run main.go
```

Aplikace běží na `http://localhost:8081`

## API Endpointy

Obě aplikace mají stejné endpointy:

### Java aplikace (port 8080)

- `GET /api/items` - Získat všechny položky
- `GET /api/items/{id}` - Získat položku podle ID
- `POST /api/items` - Vytvořit novou položku
- `PUT /api/items/{id}` - Aktualizovat položku
- `DELETE /api/items/{id}` - Smazat položku

### Go aplikace (port 8081)

- `GET /api/items` - Získat všechny položky
- `GET /api/items/{id}` - Získat položku podle ID
- `POST /api/items` - Vytvořit novou položku
- `PUT /api/items/{id}` - Aktualizovat položku
- `DELETE /api/items/{id}` - Smazat položku

## Příklady požadavků

### Vytvoření položky

```bash
# Java API
curl -X POST http://localhost:8080/api/items \
  -H "Content-Type: application/json" \
  -d '{"name":"Test Item","description":"Test Description"}'

# Go API
curl -X POST http://localhost:8081/api/items \
  -H "Content-Type: application/json" \
  -d '{"name":"Test Item","description":"Test Description"}'
```

### Získání všech položek

```bash
# Java API
curl http://localhost:8080/api/items

# Go API
curl http://localhost:8081/api/items
```

### Získání položky podle ID

```bash
# Java API
curl http://localhost:8080/api/items/1

# Go API
curl http://localhost:8081/api/items/1
```

### Aktualizace položky

```bash
# Java API
curl -X PUT http://localhost:8080/api/items/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated Item","description":"Updated Description"}'

# Go API
curl -X PUT http://localhost:8081/api/items/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated Item","description":"Updated Description"}'
```

### Smazání položky

```bash
# Java API
curl -X DELETE http://localhost:8080/api/items/1

# Go API
curl -X DELETE http://localhost:8081/api/items/1
```

## Model Item

```json
{
  "id": 1,
  "name": "Item Name",
  "description": "Item Description"
}
```

## Zastavení aplikací

```bash
# Zastavit databáze
docker-compose down

# Zastavit Java aplikaci: Ctrl+C
# Zastavit Go aplikaci: Ctrl+C
```
