## REST HMAC

Ejemplo de una API REST utilizando [GO] y HMAC

####Endpoins disponibles

- `GET /clients`
- `GET /clients/:id`
- `POST /clients`

####Ejemplo de JSON

```json
{
  "id": "d290f1ee-6c54-4b01-90e6-d701748f0851",
  "rut": "15643345-6",
  "firstName": "Pedro Pablo",
  "lastName": "Perez",
  "secondLastName": "Quiroz",
  "pep": true,
  "gender": "male",
  "dateOfBirth": "1990/12/01",
  "nationality": "PE",
  "phone": 56955449932,
  "residenceCountry": "CL",
  "address": "Avenida Los Acacios 2343",
  "city": "Santiago",
  "commune": "Quinta Normal",
  "postalCode": 3059382,
  "maritalStatus": "married",
  "occupation": "employee",
  "degree": "nurse"
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
