## dexctl - a Command Line Interface for [Dex](https://dexidp.io/)

`dexctl` can be used both to perform CRUD operations on `Dex` resources and to test the sign-on process and verify tokens and passwords.

### install

```bash
git clone https://github.com/twoojoo/dexctl.git
cd dexctl
make install
```

### usage

A `Dex` instance must be running. `Dex` defaults match `dexctl` ones:

- host: 127.0.0.1 (won't match "localhost")
- App port: 5555
- Dex web port: 5556
- Dex grpc port: 5557

`help` flag available for every command
```bash
dexctl --help
```

```bash
dexctl version [--grpc-url <host:port>]
# gets Dex server version
```

```bash
dexctl signon <client_id> -s <client_secret> [--scope <scope1> --scope <scope2> --offine-access --browser <browser>]
# will open up a browser window (default browser is autodetected in X11)
```

#### token

```bash
dexctl token verify <client_id> <token>
```

#### clients

```bash
dexctl client create <id> --name <name> [--secret <secret> --redirect-uri <uri> --grpc-url <host:port>]
```

```bash
dexctl client update <id> [--name <new-name> --grpc-url <host:port>]
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
dexctl password create --email <email> --username <username> --hash <hash> [--grpc-url <host:port>]
```

```bash
dexctl password update <email> [--username <new-username> --hash <new-hash> --grpc-url <host:port>]
```

```bash
dexctl password delete <email> [--grpc-url <host:port>]
```

```bash
dexctl password verify <email> <password> [--grpc-url <host:port>]
```