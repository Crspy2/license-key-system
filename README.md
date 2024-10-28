# Licensing System Panel

## Cloning the repository
Use one of the following commands to clone the repository
```bash
git clone https://github.com/Crspy2/license-key-system.git
```
```bash
gh repo clone Crspy2/license-key-system
```

## Setup

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

### Setting environment variables

You can now go into your `.env.*` file and set te `SSL_ENCRYPTION_CERT` and `SSL_ENCRYPTION_KEY` environment variables to the string inside the generated text files
```dotenv
DATABASE_URL=""
PORT=""

SSL_ENCRYPTION_CERT=""
SSL_ENCRYPTION_KEY=""

COOKIE_ENCRYPTION_KEY=""
```
Also go into the `panel/.env.*` file and set the `SSL_CERTIFICATE` environment variable to the ssl encryption certificate
```dotenv
SSL_CERTIFICATE=""
```

## Running the server
Once all the environment variables are set, you can run the following command to start the server

### Build the go dependencies
```bash
go mod download
```

You have two options for how you can run the server. The first is you can run the go code directly:
### Run the server
```bash
go run ./cmd/grpc
```
The second option is to compile an executable, and then execute it:
### Build an executable and run it
```bash
go build -o grpc_server ./cmd/grpc
./grpc_server
```