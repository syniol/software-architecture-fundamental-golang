# Software Architecture Fundamental in Golang <sup>(Go)</sup>
This repository is for demonstration of Layered Software Architecture which 
is a fundamental concept of clean architecture in software engineering context.


## Up and Running with Docker
There is a `make` script at the root of this repository. You can run and stop 
container with `make up` and `make down`.


## RESTful Server
When server is running successfully, you should be able to get http status 
code `200` with following response from `/health` endpoint:

```sh
curl --location 'http://127.0.0.1/health'
```

```json
{
  "status": "healthy"
}
```


#### Credit
Author: [Hadi Tajallaei](mailto:hadi@syniol.com)

Copyright &copy; 2024 Syniol Limited. All rights reserved.
