var fs = require("fs");

console.log("Going to write into existing file");

var user = {
    "user4" : {
       "name" : "mohit",
       "password" : "password4",
       "profession" : "teacher",
       "id": 4
    } 
 }


fs.appendFile('input.txt', JSON.stringify(user)+ '\n' , err => {
    if (err) {
     console.error(err);
   }
   // done!
});
