# AliothCenter Foreseen

![Go Version](https://img.shields.io/github/go-mod/go-version/alioth-center/foreseen)
![Release](https://img.shields.io/github/v/release/alioth-center/foreseen)
![Go Report Card](https://goreportcard.com/badge/github.com/alioth-center/foreseen)
![GitHub Actions](https://img.shields.io/github/actions/workflow/status/alioth-center/foreseen/build.yml?branch=main)
![License](https://img.shields.io/github/license/alioth-center/foreseen)

Foreseen is a notification service that pushes messages to users/groups via various channels. It is designed to be a robust and flexible service that can be easily integrated into any application.

> AliothCenter Foreseen is an application powered by [alioth-center/infrastructure](github.com/alioth-center/infrastructure).

## How to use

### Prepare the configuration file

```yaml
# config.yaml
app_id: "YOUR_LARK_APP_ID"
app_secret: "YOUR_LARK_APP_SECRET"
token: "RANDOM_REQUEST_TOKEN"
log_dir: "LOG_DIRECTORY"

database:
  host: 'postgres'
  port: 5432
  username: 'postgres'
  password: 'password'
  database: 'database_name'
```

### Run the service

Use the following command to run the service via Docker:

```shell
docker pull ghcr.io/alioth-center/foreseen:nightly

docker run -d \
  -p 8881:8881 \
  -v /path/to/config.yaml:/app/config.yaml \
  -v /path/to/log:/app/log \
  ghcr.io/alioth-center/foreseen:nightly
```

Or you can use the following command to run the service via Docker-compose:

```yaml
# docker-compose.yaml
version: '3.8'

services:
  foreseen:
    image: ghcr.io/alioth-center/foreseen:nightly
    ports:
      - "8881:8881"
    volumes:
      - /path/to/config.yaml:/app/config.yaml
      - /path/to/log:/app/log
    environment:
      - AC_SERVICE=foreseen
      - AC_EXTRA_FIELDS=app_env,app_version
      - app_env=pre-release
      - app_version=0.0.1
```


## Integrations

### Lark/Feishu

Foreseen supports sending messages to Lark/Feishu users and groups. To use this feature, you need to create a Lark app and get the `app_id` and `app_secret`.

For further information, please refer to the [Foreseen API Document](https://docs.alioth.center/foreseen-api.html).

## Contributors

![Contributors](https://contrib.rocks/image?repo=alioth-center/foreseen&max=1000)
