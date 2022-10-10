# Принцип работы
Создать событие можно с помощью post запроса передав json:
```shell
localhost:8080/create_event
{
  "uuid"
  "user_id"
  "event"
  "date"
}
```

Обновить событие можно с помощью post запроса передав json:
```shell
localhost:8080/create_event
{
  "uuid"
  "user_id"
  "event"
  "date"
}
```

Удалить событие можно с помощью post запроса передав json:
```shell
localhost:8080/delete_event
{
  "uuid"
}
```

Получить данные по событиям на день, неделю, месяц вперед можно с помощью get запроса:
```shell
localhost:8080/events_for_day?user_id=543&date="2020-10-03"
localhost:8080/events_for_week?user_id=543&date="2020-10-03"
localhost:8080/events_for_month?user_id=543&date="2020-10-03"
```