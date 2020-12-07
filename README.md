## goCalculator-ms
### ilegra JT - Yan Haeffner - Data Engineering

--------------------------------------------------------

### 1) Requirements (based on Linux OS):
	a) Go: https://golang.org/doc/tutorial/getting-started
	b) Available port 8080
	
### 2) Running:
	a) In a terminal enter: go run calculator-ms.go
	b) In your browser, go to: localhost:8080/calc to access the calculator microservice
	
### 3) Endpoints:
	a) localhost:8080/calc/sum/$a/$b = returns the sum between "a" (valid float number) and "b" (valid float number)
	b) localhost:8080/calc/sub/$a/$b = returns the subtraction between "a" (valid float number) and "b" (valid float number)
	c) localhost:8080/calc/mul/$a/$b = returns the multiplication between "a" (valid float number) and "b" (valid float number)
	d) localhost:8080/calc/div/$a/$b = returns the division between "a" (valid float number) and "b" (valid float number and not zero)
	e) localhost:8080/calc/history = returns the operation history of the calculator service
	f) localhost:8080/calc = returns the available endpoints
