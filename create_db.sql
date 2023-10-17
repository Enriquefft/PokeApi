BEGIN TRANSACTION;

DROP TABLE IF EXISTS pokemons;
CREATE TABLE IF NOT EXISTS pokemons (
  national_id SMALLINT PRIMARY KEY,
  name VARCHAR(8),
  speed SMALLINT,
  s_def SMALLINT,
  s_attack SMALLINT,
  def SMALLINT,
  attack SMALLINT,
  hp SMALLINT,
  type1 VARCHAR(8),
  type2 VARCHAR(8),
  height REAL,
  weight REAL,
  ability1 VARCHAR(14),
  ability2 VARCHAR(14),
  hidden_ability VARCHAR(14),
  pre_evo SMALLINT
);

DROP TABLE IF EXISTS pokemon_evolutions;
CREATE TABLE IF NOT EXISTS pokemon_evolutions (
  pokemon_id SMALLINT,
  evolution_id SMALLINT,
  PRIMARY KEY (pokemon_id, evolution_id),
  FOREIGN KEY (pokemon_id) REFERENCES pokemons(national_id),
  FOREIGN KEY (evolution_id) REFERENCES pokemons(national_id)
);

CREATE INDEX IF NOT EXISTS pokemon_name_idx ON pokemons (name);

-- CREATE INDEX IF NOT EXISTS pre_evolution_name ON pokemon_evolutions (pokemon_id);
-- This might not be needed based on the performance of the composite PK

COMMIT;
