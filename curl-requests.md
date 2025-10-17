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

## Healthcheck

### Check the health of the API

```bash
curl -i localhost:4000/v1/healthcheck
```

## Nits

### Create a new nit

**Format:** `localhost:4000/v1/nits`

**Important:** This request will fail if the `course_id` does not exist in the database.

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "course_id": [course_id],
    "start_date": "2025-01-01T00:00:00Z",
    "end_date": "2025-01-02T00:00:00Z",
    "location": "Belize City"
}' localhost:4000/v1/nits
```

## Officers

### Get a specific officer

**Format:** `localhost:4000/v1/officers/[officer_id]`

- `[officer_id]`: The ID of the officer.

```bash
curl -i localhost:4000/v1/officers/[officer_id]
```

### Update an officer's first name

**Format:** `localhost:4000/v1/officers/[officer_id]`

- `[officer_id]`: The ID of the officer.

```bash
curl -X PATCH -H "Content-Type: application/json" -d '{"first_name": "NewFirstName"}' localhost:4000/v1/officers/[officer_id]
```

### Update multiple fields

**Format:** `localhost:4000/v1/officers/[officer_id]`

- `[officer_id]`: The ID of the officer.

```bash
curl -X PATCH -H "Content-Type: application/json" -d '{"rank_id": 2, "is_active": false}' localhost:4000/v1/officers/[officer_id]
```

### Failed validation

This example will fail validation because the `first_name` is empty.

```bash
curl -X PATCH -H "Content-Type: application/json" -d '{"first_name": ""}' localhost:4000/v1/officers/[officer_id]
```

### Delete an officer

**Format:** `localhost:4000/v1/officers/[officer_id]`

- `[officer_id]`: The ID of the officer.

```bash
curl -X DELETE localhost:4000/v1/officers/[officer_id]
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

**Sample `curl` command:**

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

**Format:** `localhost:4000/v1/facilitators/[facilitator_id]`

- `[facilitator_id]`: The ID of the facilitator.

```bash
curl -i localhost:4000/v1/facilitators/[facilitator_id]
```

### Update a facilitator's information

**Format:** `localhost:4000/v1/facilitators/[facilitator_id]`

- `[facilitator_id]`: The ID of the facilitator.

```bash
curl -X PATCH -H "Content-Type: application/json" -d '{
    "first_name": "NewFirstName",
    "last_name": "NewLastName",
    "email": "newemail@example.com",
    "personnel_id": 1
}' localhost:4000/v1/facilitators/[facilitator_id]
```

### Update a facilitator's email

**Format:** `localhost:4000/v1/facilitators/[facilitator_id]`

- `[facilitator_id]`: The ID of the facilitator.

```bash
curl -X PATCH -H "Content-Type: application/json" -d '{"email": "anothernewemail@example.com"}' localhost:4000/v1/facilitators/[facilitator_id]
```

### Delete a facilitator

**Format:** `localhost:4000/v1/facilitators/[facilitator_id]`

- `[facilitator_id]`: The ID of the facilitator.

```bash
curl -X DELETE localhost:4000/v1/facilitators/[facilitator_id]
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

## Courses

### Get all courses

```bash
curl -i localhost:4000/v1/courses
```

### Create a new course

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "title": "New Course",
    "description": "This is a new course.",
    "category": "Mandatory",
    "credit_hours": 3
}' localhost:4000/v1/courses
```

### Get a specific course

**Format:** `localhost:4000/v1/courses/[course_id]`

- `[course_id]`: The ID of the course.

```bash
curl -i localhost:4000/v1/courses/[course_id]
```

### Update a course

**Format:** `localhost:4000/v1/courses/[course_id]`

- `[course_id]`: The ID of the course.

```bash
curl -X PATCH -H "Content-Type: application/json" -d '{
    "title": "Updated Course Title"
}' localhost:4000/v1/courses/[course_id]
```

### Delete a course

**Format:** `localhost:4000/v1/courses/[course_id]`

- `[course_id]`: The ID of the course.

```bash
curl -X DELETE localhost:4000/v1/courses/[course_id]
```

## Feedback

### Facilitator Feedback

#### Create Facilitator Feedback

**Format:** `http://localhost:4000/v1/facilitators/[facilitator_id]/feedback`

- `[facilitator_id]`: The ID of the facilitator.

**Important:** This request will fail if the `facilitator_id` or the `session_enrollment_id` in the request body does not exist in the database.

```bash
curl -X POST -H "Content-Type: application/json" -d '{
  "session_enrollment_id": [session_enrollment_id],
  "score": 5,
  "comment": "Excellent facilitator!"
}' http://localhost:4000/v1/facilitators/[facilitator_id]/feedback
```

#### List Facilitator Feedback

**Format:** `http://localhost:4000/v1/facilitators/[facilitator_id]/feedback`

- `[facilitator_id]`: The ID of the facilitator.

```bash
curl http://localhost:4000/v1/facilitators/[facilitator_id]/feedback
```

### Course Feedback

#### Create Course Feedback

**Format:** `http://localhost:4000/v1/courses/[course_id]/feedback`

- `[course_id]`: The ID of the course.

**Important:** This request will fail if the `course_id` or the `session_enrollment_id` in the request body does not exist in the database.

```bash
curl -X POST -H "Content-Type: application/json" -d '{
  "session_enrollment_id": [session_enrollment_id],
  "score": 4,
  "comment": "The course was very informative."
}' http://localhost:4000/v1/courses/[course_id]/feedback
```

#### List Course Feedback

**Format:** `http://localhost:4000/v1/courses/[course_id]/feedback`

- `[course_id]`: The ID of the course.

```bash
curl http://localhost:4000/v1/courses/[course_id]/feedback
```

## New Endpoints

### Facilitator Session Management

#### Assign a facilitator to a session

**Format:** `localhost:4000/v1/sessions/[session_id]/facilitators`

- `[session_id]`: The ID of the training session.

**Important:** This request will fail if the `session_id` or the `facilitator_id` in the request body does not exist in the database.

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "facilitator_id": [facilitator_id]
}' localhost:4000/v1/sessions/[session_id]/facilitators
```

#### Remove a facilitator from a session

**Format:** `localhost:4000/v1/sessions/[session_id]/facilitators/[facilitator_id]`

- `[session_id]`: The ID of the training session.
- `[facilitator_id]`: The ID of the facilitator to be removed.

**Important:** This request will fail if the `session_id` or `facilitator_id` does not exist in the database.

```bash
curl -X DELETE localhost:4000/v1/sessions/[session_id]/facilitators/[facilitator_id]
```

### Ratings and Feedback

#### Submit a rating for a course, linked to a specific enrollment

**Format:** `localhost:4000/v1/enrollments/[enrollment_id]/courserating`

- `[enrollment_id]`: The ID of the session enrollment.

**Important:** This request will fail if the `enrollment_id` does not exist in the database.

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "score": 5,
    "comment": "Great course!"
}' localhost:4000/v1/enrollments/[enrollment_id]/courserating
```

#### Submit a rating for a facilitator, linked to a specific enrollment

**Format:** `localhost:4000/v1/enrollments/[enrollment_id]/facilitatorrating`

- `[enrollment_id]`: The ID of the session enrollment.

**Important:** This request will fail if the `enrollment_id` or the `facilitator_id` in the request body does not exist in the database.

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "facilitator_id": [facilitator_id],
    "score": 5,
    "comment": "Excellent facilitator!"
}' localhost:4000/v1/enrollments/[enrollment_id]/facilitatorrating
