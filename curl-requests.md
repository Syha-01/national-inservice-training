
//------------------------------------ MIGRATIONS ------------------------------------//

migrate -path=./migrations -database="postgres://nits:bananaforscale@localhost/nits?sslmode=disable" up

//------------------------------------ LOGIN DB CLI ------------------------------------//

sudo -u postgres psql

psql --host=localhost --dbname=nits --username=nits