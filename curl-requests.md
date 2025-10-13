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
