```bash
git clone https://github.com/twoojoo/dexctl.git
cd dexctl
make install
```

```
dexctl version
```

```
dexctl signin 
```

```
dexctl verify <token> --client-id <client_id>
```

**clients**

```
dexctl client create my-client-id-123 --name MyClientName --secret s3Cr3TsTr1nG --redirect-uris <redirect-udi>
```

```
dexctl client update my-client-id-123 --name MyNewClientName
```

```
dexctl client update my-client-id-123
```