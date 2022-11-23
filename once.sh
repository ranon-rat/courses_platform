rm -rf data/database.db
touch data/database.db
cat ./data/init.sql | sqlite3 ./data/database.db
