const express = require("express");
const app = express();
import bodyParser from 'body-parser';
app.use(bodyParser.urlencoded({extended:true}));
app.use(bodyParser.json())

function humidity(min, max) {  
  return Math.floor(
    Math.random() * (max - min) + min
  )
}

// Example:
console.log(  
  humidity(0, 100)
)

const swaggerUi = require('swagger-ui-express'),
    swaggerDocument = require('./swagger.json');

// GET Call Humidity sensor
app.get("/health", (req, res) => {
  return res.status(200).send({
    success: "true",
    message: "Healty",
    humidity
  });
});

app.get("/humidity", function (req, res) {
  return res.status(200).send(humidity(0, 100));
});


app.use(
  '/api',
  swaggerUi.serve, 
  swaggerUi.setup(swaggerDocument)
);

app.listen(8000, () => {
  console.log("server listening on port 8000!");
});
