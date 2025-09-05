# Todo Go App

Простое приложение для управления задачами на Go с PostgreSQL.

## Запуск с Docker Compose

```bash
# Запуск всех сервисов
docker-compose up -d

# Просмотр логов
docker-compose logs -f

# Остановка сервисов
docker-compose down
```

## API Endpoints

- `GET /` - Информация о API
- `GET /health` - Проверка состояния
- `GET /api/v1/todos` - Получить все задачи
- `GET /api/v1/todos/{id}` - Получить задачу по ID
- `POST /api/v1/todos` - Создать новую задачу
- `PUT /api/v1/todos/{id}` - Обновить задачу
- `DELETE /api/v1/todos/{id}` - Удалить задачу

## Примеры запросов

### Создание задачи
```bash
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{"title": "Новая задача", "description": "Описание задачи", "completed": false}'
```

### Получение всех задач
```bash
curl http://localhost:8080/api/v1/todos
```

## Порты

- Приложение: 8080
- PostgreSQL: 5434 (внешний порт)

## Переменные окружения

Скопируйте `env.example` в `.env` и настройте по необходимости.
