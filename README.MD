# GoFrame Demos

This repo implements some simple demos using GoFrame.
1. A simple websocket chat service.
1. Basic API examples for user service.

## Installation

### 1. You need a go development environment setup before everything starts taking off.

### 2. Use `git clone` cloing the repo to your local folder. 
```
git clone https://github.com/gogf/gf-demos
```

### 3. Import `manifest/sql/create.sql` to your database. 

### 4. Update `manifest/config/config.yaml` according to your local configurations if necessary.

### 5. Run command `go run main.go`, and you'll see something as follows if success:
```
  ADDRESS | METHOD |        ROUTE         |                                  HANDLER                                  |            MIDDLEWARE              
----------|--------|----------------------|---------------------------------------------------------------------------|------------------------------------
  :8199   | ALL    | /*                   | github.com/gogf/gf/v2/net/ghttp.internalMiddlewareServerTracing           | GLOBAL MIDDLEWARE                  
----------|--------|----------------------|---------------------------------------------------------------------------|------------------------------------
  :8199   | ALL    | /api.json            | github.com/gogf/gf/v2/net/ghttp.(*Server).openapiSpec                     |                                    
----------|--------|----------------------|---------------------------------------------------------------------------|------------------------------------
  :8199   | GET    | /chat                | github.com/gogf/gf-demos/v2/internal/handler.(*handlerChat).Index         | service.(*serviceMiddleware).Ctx   
          |        |                      |                                                                           | service.(*serviceMiddleware).CORS  
          |        |                      |                                                                           | ghttp.MiddlewareHandlerResponse    
----------|--------|----------------------|---------------------------------------------------------------------------|------------------------------------
  :8199   | POST   | /chat/name           | github.com/gogf/gf-demos/v2/internal/handler.(*handlerChat).Name          | service.(*serviceMiddleware).Ctx   
          |        |                      |                                                                           | service.(*serviceMiddleware).CORS  
          |        |                      |                                                                           | ghttp.MiddlewareHandlerResponse    
----------|--------|----------------------|---------------------------------------------------------------------------|------------------------------------
  :8199   | GET    | /chat/websocket      | github.com/gogf/gf-demos/v2/internal/handler.(*handlerChat).Websocket     | service.(*serviceMiddleware).Ctx   
          |        |                      |                                                                           | service.(*serviceMiddleware).CORS  
          |        |                      |                                                                           | ghttp.MiddlewareHandlerResponse    
----------|--------|----------------------|---------------------------------------------------------------------------|------------------------------------
  :8199   | ALL    | /swagger/*           | github.com/gogf/gf/v2/net/ghttp.(*Server).swaggerUI                       | HOOK_BEFORE_SERVE                  
----------|--------|----------------------|---------------------------------------------------------------------------|------------------------------------
  :8199   | POST   | /user/check-passport | github.com/gogf/gf-demos/v2/internal/handler.(*handlerUser).CheckPassport | service.(*serviceMiddleware).Ctx   
          |        |                      |                                                                           | service.(*serviceMiddleware).CORS  
          |        |                      |                                                                           | ghttp.MiddlewareHandlerResponse    
----------|--------|----------------------|---------------------------------------------------------------------------|------------------------------------
  :8199   | POST   | /user/is-signed-in   | github.com/gogf/gf-demos/v2/internal/handler.(*handlerUser).IsSignedIn    | service.(*serviceMiddleware).Ctx   
          |        |                      |                                                                           | service.(*serviceMiddleware).CORS  
          |        |                      |                                                                           | ghttp.MiddlewareHandlerResponse    
----------|--------|----------------------|---------------------------------------------------------------------------|------------------------------------
  :8199   | GET    | /user/profile        | github.com/gogf/gf-demos/v2/internal/handler.(*handlerUser).Profile       | service.(*serviceMiddleware).Ctx   
          |        |                      |                                                                           | service.(*serviceMiddleware).CORS  
          |        |                      |                                                                           | ghttp.MiddlewareHandlerResponse    
          |        |                      |                                                                           | service.(*serviceMiddleware).Auth  
----------|--------|----------------------|---------------------------------------------------------------------------|------------------------------------
  :8199   | POST   | /user/sign-in        | github.com/gogf/gf-demos/v2/internal/handler.(*handlerUser).SignIn        | service.(*serviceMiddleware).Ctx   
          |        |                      |                                                                           | service.(*serviceMiddleware).CORS  
          |        |                      |                                                                           | ghttp.MiddlewareHandlerResponse    
----------|--------|----------------------|---------------------------------------------------------------------------|------------------------------------
  :8199   | POST   | /user/sign-out       | github.com/gogf/gf-demos/v2/internal/handler.(*handlerUser).SignOut       | service.(*serviceMiddleware).Ctx   
          |        |                      |                                                                           | service.(*serviceMiddleware).CORS  
          |        |                      |                                                                           | ghttp.MiddlewareHandlerResponse    
----------|--------|----------------------|---------------------------------------------------------------------------|------------------------------------
  :8199   | POST   | /user/sign-up        | github.com/gogf/gf-demos/v2/internal/handler.(*handlerUser).SignUp        | service.(*serviceMiddleware).Ctx   
          |        |                      |                                                                           | service.(*serviceMiddleware).CORS  
          |        |                      |                                                                           | ghttp.MiddlewareHandlerResponse    
----------|--------|----------------------|---------------------------------------------------------------------------|------------------------------------
```

# GoFrame Sites
### GoFrame Repo
* [https://github.com/gogf/gf](https://github.com/gogf/gf)

### GoFrame Home
* [https://goframe.org](https://goframe.org)
