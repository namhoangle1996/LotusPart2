
# Note
* Framework : go-gin (https://github.com/gin-gonic/gin)
* Go version : above 1.16 

# Step
- Run docker-compose.yaml ( running Postgresql)
- Create database user_auth 
- Run command : go run main.go

## Register
```bash
curl --location 'localhost:8088/api/v1/user/register' \
--header 'Content-Type: application/json' \
--data '{
    "user_name" : "namhoangle1996", 
    "password" : "123456"
}'
```

## Login 
```bash
curl --location 'localhost:8088/api/v1/user/login' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhdWRpZW5jZSIsImV4cCI6MTY5ODA4MjM3NywiaXNzIjoiaXNzdWVyIiwic3ViIjoic3ViamVjdCIsImF1dGhJZCI6OCwidXNlcklkIjoxfQ.b06W30R1VkPpXDv-mN0D-Kz7zbrGbS6aYn9EzN6wtfc' \
--header 'Content-Type: application/json' \
--data '{
    "user_name" : "namhoangle1996", 
    "password" : "123456"
}'
```

## Log out 
```bash
curl --location --request POST 'localhost:8088/api/v1/user/logout' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhdWRpZW5jZSIsImV4cCI6MTY5ODA4MjM3NywiaXNzIjoiaXNzdWVyIiwic3ViIjoic3ViamVjdCIsImF1dGhJZCI6OCwidXNlcklkIjoxfQ.b06W30R1VkPpXDv-mN0D-Kz7zbrGbS6aYn9EzN6wtfc'
```


## Upload 
```bash
curl --location 'localhost:8088/api/v1/file/upload' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhdWRpZW5jZSIsImV4cCI6MTY5ODEyMDI0NiwiaXNzIjoiaXNzdWVyIiwic3ViIjoic3ViamVjdCIsImF1dGhJZCI6MTIsInVzZXJJZCI6MX0.hPtgWuoofRHV004x05_CZV7EsAVY8jOClBkE6HabsIA' \
--form 'data=@"/Users/macprom116256/Downloads/11111.jpg"'
```


## Register
![img_1.png](img_1.png)

## Login 
![img.png](img.png)

## Upload file
![img_2.png](img_2.png)


