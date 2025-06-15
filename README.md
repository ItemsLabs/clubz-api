# laliga-matchfantasy-api

## Requirements

Swagger - <https://goswagger.io/install.html> (install for your OS)
Sqlboiler:

```bash
go get -u -t github.com/volatiletech/sqlboiler@latest
go get -u -t github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest



```

and all the ENV variables in a `.env` file:

```env
DATABASE_USER=laliga
DATABASE_PASSWORD=laliga
DATABASE_NAME=laliga
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_SSLMODE=disable
ENV_NAME=development
API_PORT=8080
AMQP_GAMES_EXCHANGE=games
RMQ_HOST=localhost
RMQ_PORT=5673
RMQ_VHOST=
RMQ_USER=rabbitmq
RMQ_PASSWORD=rabbitmq
RMQ_GAME_UPDATES_EXCHANGE="game_updates"
RMQ_FCM_EXCHANGE="fcm"
RMQ_EXCHANGES=match_event,fcm,games,system,game_updates
RMQ_FCM_PUSHER_QUEUE="fcm:pusher"
WORKER_COUNT=20
PREFETCH_COUNT=100
```

## Instructions for Auto-generated Code

1. Algorithm for add or change db schema

    1. add new class or update columns in [models.py](..%2Flaliga-matchfantasy-admin%2Fcore%2Fmodels.py) in Admin app

        ```python
        class User
          ...
          wallet_address = models.TextField(null=True, blank=True)
        ```

    2. Next step is update db, create migration for it

        ```bash
          $> python manage.py makemigrations --name update_models
        ```

    3. rerun Django(or docker-container) for run migrations
    4. Update definitions in [swagger.yaml](swagger.yaml)
    5. Run

        ```bash
          $> make generate
        ```

### If you encounter an issue with the host address, add the following configuration to [sqlboiler.toml](database%2Fsqlboiler.toml)

```toml
# Replace with your credentials
[psql]
   dbname = "dbname"
   host   = "localhost"
   port   = 5432
   user   = "dbusername"
   pass   = "dbpassword"
   schema = "public"
```

or export some envs

```bash
$ export DATABASE_USER=laliga
export DATABASE_PASSWORD=laliga
export DATABASE_NAME=laliga
export DATABASE_HOST=localhost
export DATABASE_PORT=5433
export DATABASE_SSLMODE=disable
```

Then, run the generation command again.
# CI/CD Pipeline Test - Sun Jun 15 00:07:04 -05 2025
# Test commit to verify CI/CD - Sun Jun 15 00:41:53 -05 2025
