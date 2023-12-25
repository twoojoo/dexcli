```bash
git clone https://github.com/twoojoo/dexctl.git
cd dexctl
make install
```

```bash
dexctl version
```

```bash
dexctl signin 
```

```bash
dexctl verify <token> --client-id <client_id>
```

**clients**

```bash
dexctl client create my-client-id-123 --name MyClientName --secret s3Cr3TsTr1nG --redirect-uris https://redirect/auth/here
```

```bash
dexctl client update my-client-id-123 --name MyNewClientName
```

```bash
dexctl client update my-client-id-123
```