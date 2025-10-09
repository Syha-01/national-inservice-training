
//------------------------------------ MIGRATIONS ------------------------------------//

migrate -path=./migrations -database="$TRAINING_DB_DSN" up

//------------------------------------ LOGIN DB CLI ------------------------------------//

sudo -u postgres psql