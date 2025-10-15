# cURL Requests

## Migrations

To run database migrations:

```bash
migrate -path=./migrations -database="postgres://nits:bananaforscale@localhost/nits?sslmode=disable" up
```

## Database CLI

To log in to the PostgreSQL CLI:

As the `postgres` user:
```bash
sudo -u postgres psql
```

With a specific user and database:
```bash
psql --host=localhost --dbname=nits --username=nits
```

## Officers

### Get a specific officer

```bash
curl -i localhost:4000/v1/officers/1
```

### Update an officer's first name

```bash
curl -X PATCH -H "Content-Type: application/json" -d '{"first_name": "NewFirstName"}' localhost:4000/v1/officers/1
```

### Update multiple fields

```bash
curl -X PATCH -H "Content-Type: application/json" -d '{"rank_id": 2, "is_active": false}' localhost:4000/v1/officers/1
```

### Failed validation

This example will fail validation because the `first_name` is empty.

```bash
curl -X PATCH -H "Content-Type: application/json" -d '{"first_name": ""}' localhost:4000/v1/officers/1
```

### Delete an officer

```bash
curl -X DELETE localhost:4000/v1/officers/1
```

### Get all officers

```bash
curl -i localhost:4000/v1/officers
```

### Add a new officer

When adding a new officer, you need to provide the following fields in the JSON body of your request:

- `regulation_number` (string, required): The unique regulation number of the officer. Must not be more than 50 characters long.
- `first_name` (string, required): The first name of the officer. Must not be more than 100 characters long.
- `last_name` (string, required): The last name of the officer. Must not be more than 100 characters long.
- `sex` (string, required): The sex of the officer. Must be either "Male" or "Female".
- `rank_id` (integer, required): The ID of the officer's rank. This should correspond to an existing ID in the `ranks` table.
- `formation_id` (integer, required): The ID of the officer's formation. This should correspond to an existing ID in the `formations` table.
- `posting_id` (integer, required): The ID of the officer's posting. This should correspond to an existing ID in the `postings` table.
- `is_active` (boolean, required): A boolean value indicating whether the officer is currently active.

**Example `curl` command:**

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "regulation_number": "12345",
    "first_name": "John",
    "last_name": "Doe",
    "sex": "Male",
    "rank_id": 1,
    "formation_id": 1,
    "posting_id": 1,
    "is_active": true
}' localhost:4000/v1/officers
```

## Unit Testing

To run validator tests:

```bash
go test -v ./internal/data
```

## Facilitators

### Get a specific facilitator

```bash
curl -i localhost:4000/v1/facilitators/1
```

### Update a facilitator's information

```bash
curl -X PATCH -H "Content-Type: application/json" -d '{
    "first_name": "NewFirstName",
    "last_name": "NewLastName",
    "email": "newemail@example.com",
    "personnel_id": 1
}' localhost:4000/v1/facilitators/1
```

### Update a facilitator's email

```bash
curl -X PATCH -H "Content-Type: application/json" -d '{"email": "anothernewemail@example.com"}' localhost:4000/v1/facilitators/1
```

### Delete a facilitator

```bash
curl -X DELETE localhost:4000/v1/facilitators/1
```

### Get all facilitators

```bash
curl -i localhost:4000/v1/facilitators
```

### Add a new facilitator

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "first_name": "John",
    "last_name": "Doe",
    "email": "johndoe@example.com"
}' localhost:4000/v1/facilitators
```

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "first_name": "Jane",
    "last_name": "Doe",
    "email": "janedoe@example.com",
    "personnel_id": 1
}' localhost:4000/v1/facilitators
```

## Feedback

### Facilitator Feedback

#### Create Facilitator Feedback

```bash
curl -X POST -H "Content-Type: application/json" -d '{
  "user_id": 1,
  "rating": 5,
  "comment": "Excellent facilitator!"
}' http://localhost:4000/v1/facilitators/1/feedback
```

#### List Facilitator Feedback

```bash
curl http://localhost:4000/v1/facilitators/1/feedback
```

### Course Feedback

#### Create Course Feedback

```bash
curl -X POST -H "Content-Type: application/json" -d '{
  "user_id": 1,
  "rating": 4,
  "comment": "The course was very informative."
}' http://localhost:4000/v1/courses/1/feedback
```

#### List Course Feedback

```bash
curl http://localhost:4000/v1/courses/1/feedback
```

## Upcoming Routes

### Facilitator Management
- `GET /v1/facilitators`: List all available facilitators.
- `POST /v1/facilitators`: Add a new facilitator.
- `GET /v1/facilitators/:id`: Get the details of a specific facilitator.
- `PATCH /v1/facilitators/:id`: Update a facilitator's information.
- `DELETE /v1/facilitators/:id`: Remove a facilitator.
- `POST /v1/sessions/:id/facilitators`: Assign a facilitator to a specific training session.
- `DELETE /v1/sessions/:id/facilitators/:facilitator_id`: Remove a facilitator from a session.

### User and Authentication Management
- `POST /v1/users`: Register a new user for the application.
- `POST /v1/tokens`: Authenticate a user and receive a token for accessing protected routes.

### Ratings and Feedback
- `POST /v1/enrollments/:id/courserating`: Submit a rating for a course, linked to a specific enrollment.
- `POST /v1/enrollments/:id/facilitatorrating`: Submit a rating for a facilitator.

### Lookup Data
- `GET /v1/ranks`
- `GET /v1/regions`
- `GET /v1/formations`
- `GET /v1/postings`
- `GET /v1/roles`