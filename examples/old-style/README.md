# old-style demo

- with external configuration files
  checkout `<project>/ci/etc/old-style/old-style.yml` and `<project>/ci/etc/old-style/conf.d/*.yml`.
- normalize app structure
-  





```bash

[ -d ci/certs ] || mkdir -p ci/certs
openssl req -newkey rsa:2048 -nodes -keyout ci/certs/server.key -x509 -days 3650 -out ci/certs/server.crt

```




