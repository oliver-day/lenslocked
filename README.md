# lenslocked

## Command to run postgresql queries

docker compose exec -it db psql -U baloo -d lenslocked

## Command to view production logs when ssh into the server

cd ~/apps/lenslocked.com/
docker compose \
 -f docker-compose.yml \
 -f docker-compose.production.yml \
 logs
