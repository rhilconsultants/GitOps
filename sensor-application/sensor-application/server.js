var express = require('express');
var app = express();
var fs = require("fs");

// Get All Humidity Reads
app.get('/listhumidityreads', function (req, res) {
   fs.readFile("humidity.json", 'utf8', function (err, data) {
      console.log( data );
      res.end( data );
   });
})

 // Get a random Read
 app.get('/humidity', function (req, res) {
    // First read existing users.
    fs.readFile("humidity.json", 'utf8', function (err, data) {
       var users = JSON.parse( data );
       function humidity(min, max){
        return Math.floor(
            Math.random() * (max - min + 1 ) + min 
        )
       }
       var user = users["read" + humidity(1, 10)] 
       console.log( user );
       res.end( JSON.stringify(user));
    });
 })


var server = app.listen(8081, function () {
   var host = server.address().address
   var port = server.address().port
   console.log("Example app listening at http://%s:%s", host, port)
})