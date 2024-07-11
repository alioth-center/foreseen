# AliothCenter Foreseen

![Go Version](https://img.shields.io/github/go-mod/go-version/alioth-center/foreseen)
![Release](https://img.shields.io/github/v/release/alioth-center/foreseen)
![Go Report Card](https://goreportcard.com/badge/github.com/alioth-center/foreseen)
![GitHub Actions](https://img.shields.io/github/actions/workflow/status/alioth-center/foreseen/codecov.yaml?branch=main)
![License](https://img.shields.io/github/license/alioth-center/foreseen)

Foreseen is a notification service that push messages to users/groups via various channels. It is designed to be a robust and flexible service that can be easily integrated into any application.

> AliothCenter Foreseen is an application powered by [alioth-center/infrastructure](github.com/alioth-center/infrastructure).

## How to use

### Prepare the configuration file

```yaml
# config.yaml
app_id: "YOUR_LARK_APP_ID"
app_secret: "YOUR_LARK_APP_SECRET"
token: "RANDOM_REQUEST_TOKEN"
log_dir: "LOG_DIRECTORY"
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
```


## Integrations

### Lark/Feishu

Foreseen supports sending messages to Lark/Feishu users and groups. To use this feature, you need to create a Lark app and get the `app_id` and `app_secret`.

For further information, please refer to the [[WIP]Foreseen API Document](https://docs.alioth.center/foreseen-apis.html).

## Contributors

![Contributors](https://contrib.rocks/image?repo=alioth-center/foreseen&max=1000)
