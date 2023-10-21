rm -f pokemon_test.db
sqlite3 Test.db -init Scripts/create_db.sql < Scripts/test_data.sql
