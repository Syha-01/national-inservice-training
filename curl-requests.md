//------------------------------------ MIGRATIONS ------------------------------------//

migrate -path=./migrations -database="postgres://nits:bananaforscale@localhost/nits?sslmode=disable" up

//------------------------------------ LOGIN DB CLI ------------------------------------//

sudo -u postgres psql

psql --host=localhost --dbname=nits --username=nits

//------------------------------------ OFFICERS ------------------------------------//

// Get a specific officer
curl -i localhost:4000/v1/officers/1

// Update an officer's first name
curl -X PATCH -H "Content-Type: application/json" -d '{"first_name": "NewFirstName"}' localhost:4000/v1/officers/1

// Update multiple fields
curl -X PATCH -H "Content-Type: application/json" -d '{"rank_id": 2, "is_active": false}' localhost:4000/v1/officers/1

// Failed validation
curl -X PATCH -H "Content-Type: application/json" -d '{"first_name": ""}' localhost:4000/v1/officers/1

//Delete
curl -X DELETE localhost:4000/v1/officers/1

//Get all officers
curl -i localhost:4000/v1/officers

// Add a new officer
curl -X POST -H "Content-Type: application/json" -d '
{
    "regulation_number": "12345",
    "first_name": "John",
    "last_name": "Doe",
    "sex": "Male",
    "rank_id": 1,
    "formation_id": 1,
    "posting_id": 1,
    "is_active": true
}
' localhost:4000/v1/officers

/*
### Fields for Adding a New Officer

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

*/

//------------------------------------ UNIT TESTING ------------------------------------//


//run validator tests
go test -v ./internal/data


//------------------------------------ FACILITATORS ------------------------------------//

// Get a specific facilitator
curl -i localhost:4000/v1/facilitators/1

// Update a facilitator's information
curl -X PATCH -H "Content-Type: application/json" -d '{
    "first_name": "NewFirstName",
    "last_name": "NewLastName",
    "email": "newemail@example.com",
    "personnel_id": 1
}' localhost:4000/v1/facilitators/1

// Update a facilitator's email
curl -X PATCH -H "Content-Type: application/json" -d '{"email": "anothernewemail@example.com"}' localhost:4000/v1/facilitators/1




upcomming routes:

Facilitator Management
You have a facilitators table, and these routes would manage that data.

GET /v1/facilitators: List all available facilitators.
POST /v1/facilitators: Add a new facilitator.
GET /v1/facilitators/:id: Get the details of a specific facilitator.
PATCH /v1/facilitators/:id: Update a facilitator's information.
DELETE /v1/facilitators/:id: Remove a facilitator.
POST /v1/sessions/:id/facilitators: Assign a facilitator to a specific training session.
DELETE /v1/sessions/:id/facilitators/:facilitator_id: Remove a facilitator from a session.
User and Authentication Management
Your users table and the middleware.go file suggest that authentication is planned.

POST /v1/users: Register a new user for the application.
POST /v1/tokens: Authenticate a user and receive a token for accessing protected routes.
Ratings and Feedback
To use your course_ratings and facilitator_ratings tables:

POST /v1/enrollments/:id/courserating: Submit a rating for a course, linked to a specific enrollment.
POST /v1/enrollments/:id/facilitatorrating: Submit a rating for a facilitator.
Lookup Data
To make it easier for a front-end application to populate dropdown menus, you could add routes to expose your lookup tables:

GET /v1/ranks
GET /v1/regions
GET /v1/formations
GET /v1/postings
GET /v1/roles

