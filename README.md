![hieda-social-card](https://github.com/pinkphantasm/hieda/assets/110753839/b11cfc4d-f2f0-412c-be1c-c9cb1a6249c6)

<div align="center"><strong>Hieda</strong> - это решение для электронного документооборота с открытым исходным кодом</div>

## Требования

- [Docker](https://www.docker.com/) (`docker-compose`)
- Сторонние зависимости
- RSA-сертификаты

Проверить сторонние зависимости:

```shell
cd src/static_service
./scripts/ensure_vendor.sh
```

Далее, следуйте инструкциям, предоставленным скриптом.

Чтобы сгенерировать приватный RSA-ключ (требуется для работы удостоверяющего сервиса), запустите следующую команду:

```shell
openssl genrsa 2048 | openssl pkcs8 -topk8 -nocrypt > ./src/cert_service/key.pem
```

## Запуск

Используя `docker-compose`:

```shell
docker-compose up
```

Или, в некоторых случаях:

```shell
docker compose up
```

Теперь Hieda доступна по адресу [localhost](http://localhost).

## Конфигурация

Вы можете настраивать Hieda через `docker-compose.yml`.

Сами сервисы настраиваются через переменные окружения. См. `README.md` каждого сервиса для более подробной информации.
