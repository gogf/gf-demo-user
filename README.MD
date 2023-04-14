# A simple user API service

## Installation

### 1. You need a go development environment setup before everything starts taking off.

### 2. Use `git clone` to clone the repo to your local folder.

```
git clone https://github.com/gogf/gf-demo-user
```

### 3. Import `manifest/sql/create.sql` to your database.

### 4. Update `manifest/config/config.yaml` according to your local configurations if necessary.

### 5. Run command `go run main.go`, and you'll see something as follows if success:

```
  ADDRESS | METHOD |         ROUTE         |                                       HANDLER                                        |        MIDDLEWARE
----------|--------|-----------------------|--------------------------------------------------------------------------------------|---------------------------
  :8000   | ALL    | /*                    | github.com/gogf/gf/v2/net/ghttp.internalMiddlewareServerTracing                      | GLOBAL MIDDLEWARE
----------|--------|-----------------------|--------------------------------------------------------------------------------------|---------------------------
  :8000   | ALL    | /*                    | github.com/gogf/gf/v2/net/ghttp.MiddlewareHandlerResponse                            | GLOBAL MIDDLEWARE
----------|--------|-----------------------|--------------------------------------------------------------------------------------|---------------------------
  :8000   | ALL    | /api.json             | github.com/gogf/gf/v2/net/ghttp.(*Server).openapiSpec                                |
----------|--------|-----------------------|--------------------------------------------------------------------------------------|---------------------------
  :8000   | ALL    | /swagger/*            | github.com/gogf/gf/v2/net/ghttp.(*Server).swaggerUI                                  | HOOK_BEFORE_SERVE
----------|--------|-----------------------|--------------------------------------------------------------------------------------|---------------------------
  :8000   | POST   | /user/check-nick-name | github.com/gogf/gf-demo-user/v2/internal/controller/user.(*Controller).CheckNickName | service.IMiddleware.Ctx   
          |        |                       |                                                                                      | ghttp.MiddlewareCORS      
----------|--------|-----------------------|--------------------------------------------------------------------------------------|---------------------------
  :8000   | POST   | /user/check-passport  | github.com/gogf/gf-demo-user/v2/internal/controller/user.(*Controller).CheckPassport | service.IMiddleware.Ctx   
          |        |                       |                                                                                      | ghttp.MiddlewareCORS      
----------|--------|-----------------------|--------------------------------------------------------------------------------------|---------------------------
  :8000   | POST   | /user/is-signed-in    | github.com/gogf/gf-demo-user/v2/internal/controller/user.(*Controller).IsSignedIn    | service.IMiddleware.Ctx   
          |        |                       |                                                                                      | ghttp.MiddlewareCORS      
----------|--------|-----------------------|--------------------------------------------------------------------------------------|---------------------------
  :8000   | GET    | /user/profile         | github.com/gogf/gf-demo-user/v2/internal/controller/user.(*Controller).Profile       | service.IMiddleware.Ctx
          |        |                       |                                                                                      | ghttp.MiddlewareCORS
----------|--------|-----------------------|--------------------------------------------------------------------------------------|---------------------------
  :8000   | GET    | /user/profile         | github.com/gogf/gf-demo-user/v2/internal/controller/user.(*Controller).Profile       | service.IMiddleware.Ctx
          |        |                       |                                                                                      | ghttp.MiddlewareCORS
          |        |                       |                                                                                      | service.IMiddleware.Auth
----------|--------|-----------------------|--------------------------------------------------------------------------------------|---------------------------
  :8000   | POST   | /user/sign-in         | github.com/gogf/gf-demo-user/v2/internal/controller/user.(*Controller).SignIn        | service.IMiddleware.Ctx
          |        |                       |                                                                                      | ghttp.MiddlewareCORS
----------|--------|-----------------------|--------------------------------------------------------------------------------------|---------------------------
  :8000   | POST   | /user/sign-out        | github.com/gogf/gf-demo-user/v2/internal/controller/user.(*Controller).SignOut       | service.IMiddleware.Ctx
          |        |                       |                                                                                      | ghttp.MiddlewareCORS
----------|--------|-----------------------|--------------------------------------------------------------------------------------|---------------------------
  :8000   | POST   | /user/sign-up         | github.com/gogf/gf-demo-user/v2/internal/controller/user.(*Controller).SignUp        | service.IMiddleware.Ctx
          |        |                       |                                                                                      | ghttp.MiddlewareCORS
----------|--------|-----------------------|--------------------------------------------------------------------------------------|---------------------------
```

# GoFrame Sites

### GoFrame Repo

* [https://github.com/gogf/gf](https://github.com/gogf/gf)

### GoFrame Home

* [https://goframe.org](https://goframe.org) (中文)
