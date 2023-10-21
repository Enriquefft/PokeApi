BEGIN TRANSACTION;

DROP TABLE IF EXISTS pokemons;
CREATE TABLE IF NOT EXISTS pokemons (
  national_id SMALLINT PRIMARY KEY,
  name VARCHAR(8) NOT NULL,
  speed SMALLINT NOT NULL,
  s_def SMALLINT NOT NULL,
  s_attack SMALLINT NOT NULL,
  def SMALLINT NOT NULL,
  attack SMALLINT NOT NULL,
  hp SMALLINT NOT NULL,
  type1 VARCHAR(8) NOT NULL,
  type2 VARCHAR(8),
  height REAL NOT NULL,
  weight REAL NOT NULL,
  ability1 VARCHAR(14) NOT NULL,
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
