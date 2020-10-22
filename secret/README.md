# secret

With secret tool, one can encrypt secret message using a key. Same key is needed when reading that message. This message is usually a string. _secret_ saves its data in home directory named `.secret`.

## Usage

```
$ go run cmd/cli.go
Secret is an API key and other secrets manager

Usage:
  secret [command]

Available Commands:
  get         Gets a secret from your secret storage
  help        Help about any command
  set         Sets a secret in your secret storage

Flags:
  -h, --help         help for secret
  -k, --key string   the key to use when encoding and decoding secrets
```

If no key is set, the file key is not encrypted and can be called by calling `get <key>`

```
# to encrypt a value/message
go run cmd/cli.go --key "secret string" set message 'some secret message'

# to get saved value
go run cmd/cli.go --key "secret string" get message 
```
