# Voucher Purchasing API
Simple voucher API in Golang. Main functions are:
- Get Customer Detail (display customer info and points)
- Get Voucher Group List (display all available voucher and their stock)
- Purchase Voucher (purchase voucher using points. Note: a customer cannot purchase the same voucher more than once)

# How to Run
On your command line, go to the module directory 
<br />`cd D:\source-code\jwt-auth`
<br />Build the Docker image
<br />`docker build -t jwt-auth .`
<br />Create staging network inside docker
<br />`docker network create staging-network`
<br />Run the docker-compose
<br />`docker-compose up -d`
<br />Execute sql script into mysql inside Docker container
<br />`cat "D:\source-code\jwt-auth\test_reko.sql" | docker exec -i mysql /usr/bin/mysql -u root -proot

# Using Swagger
Open browser and go to 
<br />`http://localhost:8080/swagger/index.html#/`
<br />Use login API and get the accessToken. After that klik `Authorize` and paste the accessToken inside the text box: `bearer <accessToken>`

# API List
- [POST] /login
<br />Customer must login first and get the accessToken.
<br />Sample body:
<br />`{
  "username": "user02",
  "password": "password123"
}`
- [GET] /custromer/detail
<br />Display customer information
- [GET] /voucher-group
<br />Display voucher list
- [POST] /voucher-purchase
<br />Purchase voucher using points
<br />Sample body:
`{
  "voucherGroupId": 2
}`
<br /> If purchase success, the response will display list of vouchers purchased by the customer
