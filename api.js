const express = require("express");
const app = express();

app.use(express.json());

// IMAGE_GETTER
app.get("/image/:pokemon_name", (req, res) => {
  // Get the query and number of results from the request parameters

  const pokemon_name = req.params.pokemon_name;
  const image_url = `https://github.com/Edgar5377/Pokedex/blob/main/Pokemon%20Dataset/${pokemon_name}.png?raw=true`;

  // Return the data and time taken as JSON response
  res.json({ image_url });
});

// Start the server
app.listen(3001, "10.100.226.35", () => {
  console.log("App is listening on port 3000");
});
