```bash
git clone https://github.com/twoojoo/dexctl.git
cd dexctl
make install
```

```bash
dexctl version
```

```bash
dexctl signin [--browser <browser>]
```

```bash
dexctl verify <token> --client-id <client_id>
```

**clients**

```bash
dexctl client create <client_id> --name <name> --secret <secret>  --redirect-uris <uri1,uri2...> 
```

```bash
dexctl client update <client_id> --name <name> 
```

```bash
dexctl client update <client_id>
```