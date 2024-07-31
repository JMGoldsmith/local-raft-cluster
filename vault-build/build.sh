#/bin/sh
chmod +x vault
docker buildx build --platform=linux/amd64 --no-cache -t vault-enterprise:dev .