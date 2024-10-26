# panel


## SSL Certificate Generation

You can edit the `san.cnf` file to add more domains or IP addresses as needed for the SSL/TLS configuration for your site


### Generate the certificates
```bash
openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 \
-keyout server.key -out server.crt \
-subj "/CN=example.com" \
-addext "subjectAltName=DNS:example.com,DNS:*.example.com,DNS:anotherdomain.com,DNS:*.anotherdomain.com,DNS:localhost"

```

Once you have generated the files, run the following command to get a base64 representation of the data so you can add them as environment variables
```bash
base64 -i server.crt -o server.crt.txt
base64 -i server.key -o server.key.txt
```
