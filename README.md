# wristband-nb-server
This project for wristband used.

Run this project follow the step below the description:

----
# Set database 
- Create database name: nb_schema 
- Import database nb_schema.sql 
# Set mysql/mysql.go
- Check database set like userName、password...etc.
- go build mysql.go <br>
# Set tcpserver/tcpserver.go 
- Check port
- go build tcpserver.go 
# Run command below the step: 
- cd /webservice
- go run webservice.go
- allow the port
