# Cluster for Vault Enterprise reproductions.

## To build:

First, compile Vault Enterprise for `linux/amd64` and drop the binary into this directory:

```bash
# In Vault Enterprise dir
$ env GOOS=linux GOARCH=amd64 make bin
$ cp ./bin/vault THIS_DIR/vault-build/bin>
```

Run vault-build/build.sh

## To run:

./init.sh

## to teardown:

./teardown.sh