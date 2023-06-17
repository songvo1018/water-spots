после запуска окружения:
1. проверить созданную БД и таблицу в postgresql
2. создать очередь water-spots в rabbitmq
3. импортировать дашборд для Postgres в графане, создав источник Prometheus

Rabbitmq http://localhost:15672/#/

Postgres Exporter http://localhost:9090/targets

Grafana Dashboard http://localhost:3000
    Default
        User: admin
        Password: admin (refresh pasword like 'waterspot')




порядок поднятия:
- docker compose up
- go run /water-spot-enricher/main.go
- go run /water-spot-producer/main.go (каждый запуск создает новый срез данных)
- go run /water-spot-consumer/main.go