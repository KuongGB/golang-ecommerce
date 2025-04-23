##Ecommerce With Golang Project

# You can start the project with below commands
docker-compose up -d
go run main.go
SIGNUP API CALL (POST REQUEST)
http://localhost:8080/users/signup

{
  "first_name": "User",
  "last_name": "Master",
  "email": "user@gmail.com",
  "password": "111111",
  "phone": "+123456789"
}
Response :"Successfully Signed Up!!"

LOGIN API CALL (POST REQUEST)

http://localhost:8080/users/login

{
  "email": "user@gmail.com",
  "password": "12345"
}
response will be like this

{
  "_id": "***********************",
  "first_name": "User",
  "last_name": "Master",
  "password": "$2a$14$ZdGZvMB0AMXbPdoN3E.dPO9RZXa9rx8XwomlsGDbpC54ZjdAWUveu",
  "email": "user@gmail.com",
  "phone": "+123456789",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImN1b25nbmU3NDAyQGdtYWlsLmNvbSIsImZpcnN0X25hbWUiOiJDdW9uZyIsImxhc3RfbmFtZSI6IlRyaW5oIiwidWlkIjoiNjdmZGY4OGIzMzJmM2I0YTYxYmVkMGY2IiwiZXhwIjoxNzQ0Nzg2NTkyfQ.Vb2NsCDIPH1_taoQvbU6nj0Pz8kmT4DXtANP9lORt4M",
  "Refresh_Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IiIsImZpcnN0X25hbWUiOiIiLCJsYXN0X25hbWUiOiIiLCJ1aWQiOiIiLCJleHAiOjE3NDUzMDIyODN9.hsr1N0YWEgRIqFFx-8F6G9cbF7mPyi4AAkiKS7ipg6s",
  "created_at": "2022-04-09T08:14:11Z",
  "updtaed_at": "2022-04-09T08:14:11Z",
  "user_id": "61614f539f29be942bd9df8e",
  "usercart": [],
  "address": [],
  "orders": []
}
Admin add Product API (POST REQUEST)

http://localhost:8080/admin/addproduct

{
  "product_name": "lenovo P50",
  "price": 1500,
  "rating": 10,
  "image": "lenovo.jpg"
}
Response : "Successfully added our Product Admin!!"

View all the Products API (GET REQUEST)

http://localhost:8000/users/productview

Response

[
  {
    "Product_ID": "67fe015f65c0061877c384ef",
    "product_name": "lenovo P50",
    "price": 1500,
    "rating": 10,
    "image": "lenovo.jpg"
  },
  {
    "Product_ID": "67fe01ae65c0061877c384f0",
    "product_name": "lenovo R7000",
    "price": 2000,
    "rating": 9,
    "image": "lenovo1.jpg"
  },
  {
    "Product_ID": "616152ee9f29be942bd9df90",
    "product_name": "iphone 13",
    "price": 1700,
    "rating": 4,
    "image": "ipho.jpg"
  }
]
Search Product by regex API (GET REQUEST)
defines the word search sorting http://localhost:8080/users/search?name=en

response:

[
  {
    "Product_ID": "67fe015f65c0061877c384ef",
    "product_name": "lenovo P50",
    "price": 1500,
    "rating": 10,
    "image": "lenovo.jpg"
  },
  {
    "Product_ID": "67fe01ae65c0061877c384f0",
    "product_name": "lenovo R7000",
    "price": 2000,
    "rating": 9,
    "image": "lenovo1.jpg"
  }
]
Adding the Products to the Cart API (GET REQUEST)

http://localhost:8080/addtocart?id=xxxproduct_idxxx&userID=xxxxxxuser_idxxxxxx

Corresponding mongodb query

Removing Item From the Cart API (GET REQUEST)

http://localhost:8080/removecartitem?id=xxxxxxx&userID=xxxxxxxxxxxx

Listing the item in the users cart API (GET REQUEST) and total price

http://localhost:8080/listcart?id=xxxxxxuser_idxxxxxxxxxx

Addding the Address API (POST REQUEST)

http://localhost:8080/addadress?id=user_id**\*\***\***\*\***

The Address array is limited to two values home and work address more than two address is not acceptable

{
  "house_name": "white house",
  "street_name": "white street",
  "city_name": "washington",
  "pin_code": "332423432"
}
Editing the Home Address API (PUT REQUEST)

http://localhost:8080/edithomeaddress?id=xxxxxxxxxxuser_idxxxxxxxxxxxxxxx

Editing the Work Address API (PUT REQUEST)

http://localhost:8080/editworkaddress?id=xxxxxxxxxxuser_idxxxxxxxxxxxxxxx

Delete Addresses API (GET REQUEST)

http://localhost:8080/deleteaddresses?id=xxxxxxxxxuser_idxxxxxxxxxxxxx

delete both addresses

Cart Checkout Function and placing the order API (GET REQUEST)

After placing the order the items have to be deleted from cart functonality added

http://localhost:8000?id=xxuser_idxxx

Instantly Buying the Products API (GET REQUEST) http://localhost:8080/instantbuy?userid=xxuser_idxxx&pid=xxxxproduct_idxxxx
