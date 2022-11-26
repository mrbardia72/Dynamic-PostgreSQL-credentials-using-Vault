# user endpoint service

## login
```http request
POST {{baseUrl}}/login
Content-Type: application/json
```
##### Request
```json
{
    "identity": "email or mobile",
    "password": "password"
}
```
##### Response
```json
{
  "token": "eyJhbGciOiJIUzi8uo0ytnR5cCI6IkpXVCJ9.eyJleHAiOjE2MzU2ODkwODmmmnbzZXJuYW1lIjoib3JnMWFkbWluIn0.PhRbKxuhbvcf5XFfVRGDEMtD06cRajrEVRmx6AXbVy8",
  "identity": "email or mobile",
  "role": "writer"
}
```