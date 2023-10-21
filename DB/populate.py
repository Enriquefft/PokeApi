"""Module for populating the pokemons database."""

import sqlite3

API_URL = "https://pokeapi.co/api/v2/pokemon-species/ivysaur"


def getDB():
    """Return a database connection."""
    con = sqlite3.connect("tutorial.db")
    cur = con.cursor()
    return cur


def verifyDB(db_cursor: sqlite3.Cursor) -> bool:
    """Verify that the database is populated."""
    res = db_cursor.execute("SELECT pokemons FROM sqlite_master")
    return res.fetchone() is not None


def main():
    """Execute the main function."""
    db = getDB()
    if not verifyDB(db):
        raise Exception("Database table not found.")


if __name__ == "__main__":
    main()
