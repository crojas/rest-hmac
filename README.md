## REST HMAC

Ejemplo de una API REST utilizando [GO] y HMAC

###Endpoins disponibles

- `GET /clients`
- `GET /clients/:id`
- `POST /clients`

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

### Generar binario para linux en desarrollo

```sh
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./build/rest-hmac.exe
```

### Generar binario para linux en producción

```sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/rest-hmac
```
