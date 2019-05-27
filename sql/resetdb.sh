kill -9 $(lsof -t -i:5433)
kill -9 $(lsof -t -i:5434)

# orbital db
pg_ctl -D .orbitaldb_dev -o '-p 5433' start
dropdb --if-exists -h localhost -p 5433 orbitaldb_dev
pg_ctl -D .orbitaldb_dev -o '-p 5433' stop

rm -rf .orbitaldb_dev && initdb .orbitaldb_dev

pg_ctl -D .orbitaldb_dev -o '-p 5433' start
createdb -h localhost -p 5433 orbitaldb_dev

psql 'postgres://bokwoon@localhost:5433/orbitaldb_dev?sslmode=disable' -f sql/init.sql
psql 'postgres://bokwoon@localhost:5433/orbitaldb_dev?sslmode=disable' -f sql/data.sql
psql 'postgres://bokwoon@localhost:5433/orbitaldb_dev?sslmode=disable' -f sql/views.sql

# session db
pg_ctl -D .sessiondb_dev -o '-p 5434' start
dropdb --if-exists -h localhost -p 5434 sessiondb_dev
pg_ctl -D .sessiondb_dev -o '-p 5434' stop

rm -rf .sessiondb_dev && initdb .sessiondb_dev

pg_ctl -D .sessiondb_dev -o '-p 5434' start
createdb -h localhost -p 5434 sessiondb_dev

psql 'postgres://bokwoon@localhost:5434/sessiondb_dev?sslmode=disable' -f sql/session.sql
