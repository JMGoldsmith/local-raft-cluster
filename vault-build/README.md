# Custom Vault Enterprise Docker Image Builder

This is a small example of building a custom Vault Enterprise Docker image.

## Run

First, compile Vault Enterprise for `linux/amd64` and drop the binary into this directory:

```bash
# In Vault Enterprise dir
$ env GOOS=linux GOARCH=amd64 make bin
$ cp ./bin/vault <this directory>
```

Next, build the Docker image:

```bash
$ docker buildx build --platform=linux/amd64 --no-cache -t vault-enterprise:dev .
```
