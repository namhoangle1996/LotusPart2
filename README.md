# Template

Tài liệu mô tả cấu trúc source đang sử dụng tại Sobanhang

# Note
* Framework : go-gin (https://github.com/gin-gonic/gin)
* Document  : swagger: (https://github.com/swaggo/gin-swagger)
* Go version : above 1.16 

## Getting started
![alt text](https://raw.githubusercontent.com/bxcodec/go-clean-arch/master/clean-arch.png)

Project có 4 lớp chính  :
* Models Layer     (pkg/model)   - lớp thực thể
* Repository Layer (pkg/repo)    - lớp tương tác dữ liệu database
* Usecase Layer    (pkg/service) - lớp xử lý logic 
* Delivery Layer   (pkg/handler) - lớp delivery, hiện đang dùng restful

# Chi tiết các folder, files

### Folder **_conf_**
* Mô tả config cấu hình mà project sử dụng. Hiện đang sử dụng biến môi trường (env)

### Folder **_docs_**
* Folder này chứa file swagger được auto generate 
```bash
# Command
$ swag init
```

### Folder **_pkg_**
#### apis
* Folder này chứa các interface và implement cho function gọi đến những microservice khác
#### handlers
* Folder này chứa các Restful handler , tương tác với lớp logic (pkg/service)
#### middleware
* Khai báo middleware sử dụng 
#### model
* Mô tả những đối tượng 
#### pubsub
* Viết những function gửi/nhận đến Queue qua restful
#### repo
* Tương tác database
#### route
* Khai báo những endpoint API 
#### service
* lớp xử lý logic 
#### utils
* chứa các hàm utility  

#### Run the Applications with hot reload (Local testing env) 

```bash
# Command
$ bash run-local.sh

```

