# wristband-nb-server
This project for wristband used.

Run this project follow the step below the description:

● Create database nb_schema <br>
● Import database nb_schema.sql <br>
● Set mysql/mysql.go
  1. Check database set like userName、password...etc.
  2. go build mysql.go
● Set tcpserver/tcpserver.go
  1. Check port.
  2. go build tcpserver.go
● Command below the step: <br>
  1. cd /webservice
  2. go run webservice.go
  3. allow the port
