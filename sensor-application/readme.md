# Sensor Application #

Provide 2 GET request

1. for Humidity read
2. For oxygen read

## 5 HTTP GET ##

'/healty' for livenss prob

'/oxygen' for Oxygen reading

'/listoxygenreads' To get all oxygen readings

'/humidity' To get Humidity Reading

'/listhumidityreads' To get all humidity readings

---
## Postman Collection and Enviourment ##

Under the Postman folder there is an Export of all the GET Call's

## Container File ##

Build a new image from ubi8 NodeJS16 latest.

###a ready image can be found @quay.io/thason/sensor-app:latest###

Contianer Expose Port 8081 with http application

