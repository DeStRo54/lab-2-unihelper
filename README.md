# Запуск UniHelper

Из корня репозитория (Bash и OpenSSL):

```bash
bash scripts/setup.sh
docker compose config --quiet
docker compose up -d --build
bash scripts/check.sh
```

Открыть https://localhost:8443. Браузер покажет предупреждение о self-signed сертификате.

`setup.sh` создаёт локальные файлы в игнорируемом каталоге `secrets/`. Compose передаёт ключ и сертификат nginx через `build.secrets` в BuildKit secret mounts. Стадия `cert-builder` экспортирует их в runtime-образ в `/etc/nginx/ssl/`; runtime secrets для TLS не нужны. Это учебный self-signed сертификат: его приватный ключ находится в финальном образе nginx, поэтому такой образ не следует публиковать с ключом реального сервера.

После замены сертификата или ключа пересобрать nginx без кэша: содержимое BuildKit secrets само по себе не сбрасывает кэш сборки.

```bash
docker compose build --no-cache nginx
docker compose up -d nginx
```

Пароль PostgreSQL передаётся только через runtime secret `postgres_password`. Старый `backend/docker-compose.dev.yml` тоже читает `secrets/postgres_password.txt` через secret; для него дополнительно нужен локальный `backend/backend.env` с остальными настройками приложения. `backend/db.env` содержит только несекретные настройки и путь к secret. Для уже существующей БД изменение файла пароля не меняет пароль роли PostgreSQL — его нужно согласовать с существующей БД.

# Как добавить группу БСБО-01-23 в UniHelper

## 1. Назначить пользователя администратором

Из корня проекта открыть `psql`:

```bash
docker compose exec db psql -U unihelper -d unihelper
```

В `psql` посмотреть список пользователей:

```sql
SELECT user_id, name, email, role FROM users ORDER BY user_id;
```

Назначить администратором пользователя с `user_id = 1`:

```sql
UPDATE users
SET role = 3
WHERE user_id = 1
RETURNING user_id, name, email, role;
```

Выйти из `psql`:

```text
\q
```

## 2. Добавить группу через приложение

Войти заново на [https://localhost:8443](https://localhost:8443) и добавить группу через раздел администратора:

- **Название:** `БСБО-01-23`
- **Курс:** `4`
- **Ссылка iCal:**

```text
http://english.mirea.ru/schedule/api/ical/1/698
```

> Группу нужно добавлять через приложение. Прямой `INSERT INTO groups` не загружает предметы и занятия.

## 3. Получить код группы и проверить загрузку расписания

Снова открыть `psql` из корня проекта:

```bash
docker compose exec db psql -U unihelper -d unihelper
```

Выполнить:

```sql
SELECT g.group_id, g.group_name, g.course, g.ical_link, g.register_key,
       (SELECT count(*) FROM subjects s WHERE s.group_id = g.group_id) AS subjects,
       (SELECT count(*) FROM classes c WHERE c.group_id = g.group_id) AS classes
FROM groups g
WHERE g.group_name = 'БСБО-01-23';
```

`register_key` — код регистрации. Для загруженного расписания `subjects` и `classes` должны быть больше нуля.

> Если группа уже существует, повторно добавлять её не нужно.
>
> **Не использовать `/admin/refreshAllData` для этой проверки:** она удаляет заметки и домашние задания.
