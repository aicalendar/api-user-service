## API Reference

### Register new user

```http
  POST /api/users/register
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `name` | `string` | **Required**. Name of the user |
| `password` | `string` | **Required**. Password of the user to hash |

#### Response

```
{
  "status": "success",
  "error": null,
  "user": {
    "id": "345fd9bf-c510-498d-b1dd-ad97274a317a",
    "name": "John",
    "password": "QNmA+SusycmMOhRWtUk+o84c7Vpbq69/9b7tsINKqbw=",
    "salt": "zOC6StaWJ1XLYf0iI7e6RA==",
    "createdAt": "2025-02-18T05:02:18.469190426Z"
  }
}
```
