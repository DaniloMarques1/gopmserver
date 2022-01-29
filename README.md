# Go Password Manager Server

this will be the server for the password manager  
checkout the cli [gopm](https://github.com/DaniloMarques1/gopm)

## Usage with curl

### Create a new master

```console
curl -X POST -d '{ "email": "fitz@gmail.com", "pwd": "12345" }' -H \
"Content-type: application/json" localhost:8080/master -v
```

### Log in as a master (get the token)

```console
curl -X POST -d '{ "email": "fitz@gmail.com", "pwd": "12345" }' -H \
"Content-type: application/json" localhost:8080/session -v
```
### Save a new password

```console
curl -X POST -d '{ "key": "your_password_key", "pwd": "some_password" }' -H \
"Content-type: application/json" -H "Authorization: Bearer {token_goes_here}" localhost:8080/password -v
```

### Get a password

```console
curl -X POST -d -H "Content-type: application/json" -H \
"Authorization: Bearer {token_goes_here}" localhost:8080/password/{passoword_key_goes_here} -v
```
