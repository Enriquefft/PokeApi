const express = require("express");
const app = express();

var cors = require("cors");
app.use(cors());

const port = 3001;

let pokemon = {
  national_id: 281,
  name: "Kirlia",
  speed: 50,
  s_def: 65,
  s_attack: 65,
  def: 35,
  attack: 45,
  hp: 38,
  type1: "Psychic",
  type2: "Fairy",
  height: 0.8,
  weight: 20.2,
  ability1: "Synchronize",
  ability2: "Trace",
  hidden_ability: "Telepathy",
  pre_evo: {
    national_id: 280,
    name: "Ralts",
    img_url: "https://play.pokemonshowdown.com/sprites/dex/ralts.png",
  },
  evolutions: [
    {
      national_id: 282,
      name: "Gardevoir",
      img_url: "https://play.pokemonshowdown.com/sprites/dex/gardevoir.png",
    },
    {
      national_id: 475,
      name: "Gallade",
      img_url: "https://play.pokemonshowdown.com/sprites/dex/gallade.png",
    },
  ],
  img_url: "https://play.pokemonshowdown.com/sprites/dex/kirlia.png",
  sound_url: "https://play.pokemonshowdown.com/audio/cries/kirlia.mp3",

  matchups: {
    defensive: {
      x4: [
      ],
      x2: [
        "Poison",
        "Ghost",
        "Steel"
      ],
      x0_5: [
        "Psychic",
      ],
      x0_25: [
        "Fighting",
      ],
      x0: [
        "Dragon",
      ],
    },

    offensive: {
      "Psychic": {
        x4: [
        ],
        x2: [
          "Fighting",
          "Poison"
        ],
        x0_5: [ 
          "Steel",
          "Psychic",
        ],

        x0: [
          "Dark"
        ],

        x0_25: [
        ],

      },
      "Fairy": {
        x4: [
        ],
        x2: [
          "Fighting",
          "Poison",
          "Dragon",
          "Dark",
        ],
        x0_5: [ 
          "Steel",
        ],

        x0: [
        ],

        x0_25: [
        ],

      }

    }


  }

};

app.get("/pokemon/:name", (req, res) => {
  pokemon.name = req.params.name;
  res.json(pokemon);
});

app.listen(port, "10.100.232.143", () => {
  console.log(`Server running at http://10.100.232.143:${port}`);
});
