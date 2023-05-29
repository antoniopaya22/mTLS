# mTLS

Mutual Transport Layer Security (mTLS) is a mechanism that allows a server to authenticate itself to a client, and vice versa, using certificates.
This is a two-way authentication mechanism, where the client needs to authenticate itself to the server as well.

In this repository, we provide a simple example of how to use mTLS in a client-server application on different programming languages.

---

## How to run the examples

### 1. Generate the certificates

```bash
$ cd ~/; openssl rand -writerand .rnd
$ openssl genrsa -des3 -passout pass:antonio -out ca.key 4096
$ openssl req -new -x509 -days 365 -key ca.key -out ca.crt -passin pass:antonio -subj "/C=ES/ST=PA/L=A/O=Uniovi/OU=SE/CN=PhD/emailAddress=abc@xyz.com"
$ openssl genrsa -out server.key 4096
$ openssl req -new -key server.key -out server.csr -passin pass:antonio -subj "/C=ES/ST=PA/L=A/O=Uniovi/OU=SE/CN=PhD/emailAddress=abc@xyz.com"
$ openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365 -sha256 -passin pass:antonio
```

### 2. Python

```bash
$ cd python
$ python server.py
$ python client.py
```