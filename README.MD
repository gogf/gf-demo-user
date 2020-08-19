# GoFrame Demos

This repo implements some demos using GoFrame.
1. A simple websocket chat service.
1. Basic API example for user SignUp/SignIn.
1. Universal CURD service.

## Installation

### 1. You need a go development environment setup before everything starts taking off.
### 2. Use `git clone` cloing the repo to your local folder. 
```
git clone https://github.com/gogf/gf-demos
```

### 3. Create configuration file from `config.example.toml`. 
```
cp config/config.example.toml config/config.toml
```
Update `config.toml` according to your local configurations if necessary.

### 4. Run command `go run main.go`, and you'll see something as follows if success:
```
  SERVER  | DOMAIN  | ADDRESS | METHOD |        ROUTE        |                              HANDLER                              |           MIDDLEWARE             
|---------|---------|---------|--------|---------------------|-------------------------------------------------------------------|---------------------------------|
  default | default | :8199   | ALL    | /chat               | github.com/gogf/gf-demos/app/api/chat.(*Controller).Index         | middleware.CORS                  
|---------|---------|---------|--------|---------------------|-------------------------------------------------------------------|---------------------------------|
  default | default | :8199   | ALL    | /chat/index         | github.com/gogf/gf-demos/app/api/chat.(*Controller).Index         | middleware.CORS                  
|---------|---------|---------|--------|---------------------|-------------------------------------------------------------------|---------------------------------|
  default | default | :8199   | ALL    | /chat/setname       | github.com/gogf/gf-demos/app/api/chat.(*Controller).SetName       | middleware.CORS                  
|---------|---------|---------|--------|---------------------|-------------------------------------------------------------------|---------------------------------|
  default | default | :8199   | ALL    | /chat/websocket     | github.com/gogf/gf-demos/app/api/chat.(*Controller).WebSocket     | middleware.CORS                  
|---------|---------|---------|--------|---------------------|-------------------------------------------------------------------|---------------------------------|
  default | default | :8199   | ALL    | /curd/:table/all    | github.com/gogf/gf-demos/app/api/curd.(*Controller).All           | middleware.CORS                  
|---------|---------|---------|--------|---------------------|-------------------------------------------------------------------|---------------------------------|
  default | default | :8199   | ALL    | /curd/:table/delete | github.com/gogf/gf-demos/app/api/curd.(*Controller).Delete        | middleware.CORS                  
|---------|---------|---------|--------|---------------------|-------------------------------------------------------------------|---------------------------------|
  default | default | :8199   | ALL    | /curd/:table/one    | github.com/gogf/gf-demos/app/api/curd.(*Controller).One           | middleware.CORS                  
|---------|---------|---------|--------|---------------------|-------------------------------------------------------------------|---------------------------------|
  default | default | :8199   | ALL    | /curd/:table/save   | github.com/gogf/gf-demos/app/api/curd.(*Controller).Save          | middleware.CORS                  
|---------|---------|---------|--------|---------------------|-------------------------------------------------------------------|---------------------------------|
  default | default | :8199   | ALL    | /curd/:table/update | github.com/gogf/gf-demos/app/api/curd.(*Controller).Update        | middleware.CORS                  
|---------|---------|---------|--------|---------------------|-------------------------------------------------------------------|---------------------------------|
  default | default | :8199   | ALL    | /user/checknickname | github.com/gogf/gf-demos/app/api/user.(*Controller).CheckNickName | middleware.CORS                  
|---------|---------|---------|--------|---------------------|-------------------------------------------------------------------|---------------------------------|
  default | default | :8199   | ALL    | /user/checkpassport | github.com/gogf/gf-demos/app/api/user.(*Controller).CheckPassport | middleware.CORS                  
|---------|---------|---------|--------|---------------------|-------------------------------------------------------------------|---------------------------------|
  default | default | :8199   | ALL    | /user/issignedin    | github.com/gogf/gf-demos/app/api/user.(*Controller).IsSignedIn    | middleware.CORS                  
|---------|---------|---------|--------|---------------------|-------------------------------------------------------------------|---------------------------------|
  default | default | :8199   | ALL    | /user/profile       | github.com/gogf/gf-demos/app/api/user.(*Controller).Profile       | middleware.CORS                  
|---------|---------|---------|--------|---------------------|-------------------------------------------------------------------|---------------------------------|
  default | default | :8199   | ALL    | /user/profile       | github.com/gogf/gf-demos/app/api/user.(*Controller).Profile       | middleware.CORS,middleware.Auth  
|---------|---------|---------|--------|---------------------|-------------------------------------------------------------------|---------------------------------|
  default | default | :8199   | ALL    | /user/signin        | github.com/gogf/gf-demos/app/api/user.(*Controller).SignIn        | middleware.CORS                  
|---------|---------|---------|--------|---------------------|-------------------------------------------------------------------|---------------------------------|
  default | default | :8199   | ALL    | /user/signout       | github.com/gogf/gf-demos/app/api/user.(*Controller).SignOut       | middleware.CORS                  
|---------|---------|---------|--------|---------------------|-------------------------------------------------------------------|---------------------------------|
  default | default | :8199   | ALL    | /user/signup        | github.com/gogf/gf-demos/app/api/user.(*Controller).SignUp        | middleware.CORS                  
|---------|---------|---------|--------|---------------------|-------------------------------------------------------------------|---------------------------------|
```

# GoFrame Sites
### GoFrame Repository
* [https://github.com/gogf/gf](https://github.com/gogf/gf)
* [https://gitee.com/johng/gf](https://gitee.com/johng/gf)

### GoFrame Home
* [https://goframe.org](https://goframe.org)
