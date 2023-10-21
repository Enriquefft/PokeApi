-- Test data
BEGIN TRANSACTION;

-- Clean previous data
DELETE FROM pokemons;

-- Insert Ditto
INSERT INTO pokemons VALUES (1, 'Ditto', 48, 48, 48, 48, 48, 48, 'Normal', NULL, 0.3, 4.0, 'Limber', NULL, 'Imposter', NULL);

-- Insert Pikachu
INSERT INTO pokemons VALUES (2, 'Pikachu', 90, 50, 50, 40, 55, 35, 'Electric', NULL, 0.4, 6.0, 'Static', NULL, 'Lightning Rod', NULL);

-- Insert Charizard
INSERT INTO pokemons VALUES (3, 'Charizard', 100, 85, 109, 78, 84, 78, 'Fire', 'Flying', 1.7, 90.5, 'Blaze', NULL, 'Solar Power', NULL);

-- Insert Bulbasaur
INSERT INTO pokemons VALUES (4, 'Bulbasaur', 45, 65, 65, 49, 49, 45, 'Grass', 'Poison', 0.7, 6.9, 'Overgrow', NULL, 'Chlorophyll', NULL);

-- Insert Squirtle
INSERT INTO pokemons VALUES (5, 'Squirtle', 43, 64, 50, 65, 48, 44, 'Water', NULL, 0.5, 9.0, 'Torrent', NULL, 'Rain Dish', NULL);


-- Gardevoir family

-- Insert Ralts
INSERT INTO pokemons VALUES (280, 'Ralts', 28, 25, 25, 45, 35, 40, 'Psychic', 'Fairy', 0.4, 6.6, 'Synchronize', 'Trace', 'Telepathy', NULL);

-- Insert Kirlia
INSERT INTO pokemons VALUES (281, 'Kirlia', 50, 65, 65, 35, 45, 38, 'Psychic', 'Fairy', 0.8, 20.2, 'Synchronize', 'Trace', 'Telepathy', 280);

-- Insert Gardevoir
INSERT INTO pokemons VALUES (282, 'Gardevoir', 80, 115, 125, 65, 65, 68, 'Psychic', 'Fairy', 1.6, 48.4, 'Synchronize', 'Trace', 'Telepathy', 281);

-- Insert Gallade
INSERT INTO pokemons VALUES (475, 'Gallade', 80, 95, 125, 65, 125, 68, 'Psychic', 'Fighting', 1.6, 52.0, 'Steadfast', 'Sharpness', 'Justified', 281);


-- Evolutions
INSERT INTO pokemon_evolutions VALUES (280, 281);
INSERT INTO pokemon_evolutions VALUES (281, 282);
INSERT INTO pokemon_evolutions VALUES (281, 475);

COMMIT;
