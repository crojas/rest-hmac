## REST HMAC

Ejemplo de una API REST utilizando [GO] y HMAC, la idea es que se haga un post al endpoint _signature_ con un payload que se desea firmar, esto retorna una llave que se debe enviar como cabecera _X-CB-Signature_ al momento de crear el cliente. El proyecto se encuentra compilado por lo que basta con correr el .exe en windows o el binatio en linux. La llave secreta de prueba es _superawesomesecketkey_

###Endpoins disponibles

- `GET /clients` lista los clientes
- `GET /clients/:id` liste el cliente por id
- `POST /signature` crea una firma con el json para enviarlo como cabecera "X-CB-Signature" al momento del post del cliente.
- `POST /clients` crea el cliente solo si viene la cabecera "X-CB-Signature" y es valida con el secret.

###Ejemplo de JSON

```json
{
  "rut": "1-9",
  "firstName": "Juan Ricardo",
  "lastName": "Perez",
  "secondLastName": "Paredes",
  "pep": true,
  "gender": "male",
  "dateOfBirth": "1980/10/15",
  "nationality": "CL",
  "phone": 569789222,
  "residenceCountry": "CL",
  "address": "nowhere 1234",
  "city": "Puerto Varas",
  "commune": "Puerto Varas",
  "postalCode": 4132132,
  "maritalStatus": "married",
  "occupation": "employee",
  "degree": "engineer"
}
```

### Generar binario para windows en desarrollo

```sh
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./build/rest-hmac.exe
```

### Generar binario para linux en producción

```sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/rest-hmac
```
