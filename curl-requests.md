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
