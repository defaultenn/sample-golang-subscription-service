# Subscription Service

## Configuration

Configuration priority order: YAML -> Enviroment. Each subsequent configuration step overwrites the previous one if specified.

### Enviroment

Список возможных переменных окружения/настроек сервиса:

| Название переменной | Тип           | Значение по умолчанию |
| ------------------------------------- | ---------------- | ---------------------------------------- |
| DATABASE_DSN                          | string           |                                          |
| LOG_LEVEL                             | string?          | error                                    |
| HTTP_PORT                             | uint?            | 8000                                     |
| HTTP_READ_TIMEOUT                     | duration string? | 120s                                     |
| HTTP_WRITE_TIMEOUT                    | duration string? | 120s                                     |
| HTTP_SHUTDOWN_TIMEOUT                 | duration string? | 3s                                       |
| HTTP_CORS_ORIGINS                     | list of strings? | ["*"]                                    |

### YAML

```yaml
database:
	dsn: "string, required"

logging:
	level: "string, default: error"

http:
	port: number, default: 8000
	read_timeout: "duration string, default: 120s"
	write_timeout: "duration string, default: 120s"
	shutdown_timeout: "duration string, default: 3s"

	cors_origins: ["list of strings, separated by ',', default ["*"]"]
```
### Migrating

Make database in postgres.
Service migrate database automatically if needed.
But if you want to migrate it by yourself:

```bash
PGPASSWORD="gorm" psql -U gorm -d gorm -h localhost -p 5432 -f ./migrations/20260214095130_initial_setup.sql
```

## Run

### In shell

```bash
# Run 
DATABASE_DSN="host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Europe/Moscow" go run ./cmd/app/main.go

# OR

# Build
go build ./cmd/app/main.go

# Make it executable
chmod +x ./main

# Run
# Make sure that database is running on the machine
DATABASE_DSN="host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Europe/Moscow" ./main 
```

### By docker

```bash
# Run PSQL database
docker run -e POSTGRES_USER=gorm -e POSTGRES_DB=gorm -p 5432:5432 -v postgres_data:/var/lib/postgresql/data -e TZ=Europe/Moscow pgsql_service -e POSTGRES_PASSWORD=gorm --name subscription_service_database --restart unless-stopped -d postgres:15-alpine

# Build
docker build -t subscription_service .

# Run
docker run -e DATABASE_DSN="host=postgres user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Europe/Moscow" -p 8000:8000 --name subscription_service --restart always -d subscription_service
```

### By docker-compose
```bash
docker-compose up
# or
docker-compose up -d # in deamon
```

## Task

**Задача:** спроектировать и реализовать REST-сервис для агрегации данных об онлайн
подписках пользователей.

**Требования:**

1. Выставить HTTP-ручки для CRUDL-операций над записями о подписках. Каждая
   запись содержит:

   1. Название сервиса, предоставляющего подписку
   2. Стоимость месячной подписки в рублях
   3. ID пользователя в формате UUID
   4. Дата начала подписки (месяц и год)
   5. Опционально дата окончания подписки
2. Выставить HTTP-ручку для подсчета суммарной стоимости всех подписок за
   выбранный период с фильтрацией по id пользователя и названию подписки
3. СУБД – PostgreSQL. Должны быть миграции для инициализации базы данных
4. Покрыть код логами
5. Вынести конфигурационные данные в .env/.yaml-файл
6. Предоставить swagger-документацию к реализованному API
7. Запуск сервиса с помощью docker compose

**Примечания:**

* Проверка существования пользователя не требуется. Управление пользователями
* вне зоны ответственности вашего сервиса
* Стоимость любой подписки – целое число рублей, копейки не учитываются
  Пример тела запроса на создание записи о подписке:

```json
{
“service_name”: “Yandex Plus”,
“price”: 400,
“user_id”: “60601fee-2bf1-4721-ae6f-7636e79a0cba”,
“start_date”: “07-2025”
}
```
