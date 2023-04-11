# ecommerce-api

Introduction to creating an e-commerce API for a clothing store, practicing and using the Fiber library, and also using Docker to create our containers and managing them using Docker Compose.

## Technologies used

- [Golang](https://go.dev/)

- [Fiber](https://gofiber.io/)

- [PostgreSQL](https://www.postgresql.org/)

- [Docker](https://www.docker.com/)

## Clonning and Run application

```
  git clone https://github.com/kimMichel/ecommerce-api
```
or fork the project.

### Create a .env file

Create a new .env file to create your environment variables.

Example:
```
  DB_USER="example"
  DB_PASSWORD="example"
  DB_NAME="example"
```

Run:

```
  docker compose build
```
and

```
  docker compose up
```

### Methods:

GET
```
http://localhost:3000
```

POST
```
http://localhost:3000/products
```

POST/JSON
```
{
  "id": "1",
  "name": "Floral Women's Blouse",
  "description": "Cotton blouse with floral print",
  "price": 49.99,
  "image": "https://i.postimg.cc/yWqRwhxg/Floral-Women-s-Blouse.jpg",
  "rating": 4.5,
  "quantity": 1
}
```
