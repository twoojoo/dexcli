### install

```bash
git clone https://github.com/twoojoo/dexctl.git
cd dexctl
make install
```

### usage

```bash
dexctl version [--grpc-url <host:port>]
```

```bash
dexctl signin [--browser <browser>]
```

```bash
dexctl verify <token> --client-id <id>
```

#### clients

```bash
dexctl client create <id> --name <name> --secret <secret>  --redirect-uris <uri1,uri2...> [--grpc-url <host:port>]
```

```bash
dexctl client update <id> --name <new-name> [--grpc-url <host:port>]
```

```bash
dexctl client delete <id> [--grpc-url <host:port>]
```

#### passwords

```bash
dexctl password get <id> [--grpc-url <host:port>]
```

```bash
dexctl password list [--grpc-url <host:port>]
```

```bash
dexctl password create <email> --username <username> --hash <hash> [--grpc-url <host:port>]
```

```bash
dexctl password update <email> --username <new-username> --hash <new-hash> [--grpc-url <host:port>]
```

```bash
dexctl password delete <email> [--grpc-url <host:port>]
```

```bash
dexctl password verify <email> --password <password> [--grpc-url <host:port>]
```