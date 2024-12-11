# Software Architecture Fundamental in Golang <sup>(Go)</sup>
This repository is for demonstration of Layered Software Architecture which 
is a fundamental concept of clean architecture in software engineering context. 

This is a hands-on approach, and it's result of spending weekends, afternoon, 
evenings, and holidays studying books and watching countless video trainings 
to master software architecture. Courses and books only give you the concepts 
with a few code examples but having a proper example of creation from layer 
to layer in a small project is impossible to find.


## Overview



## Architecture Design
<img src="./docs/architecture-design-diagram.png" style="max-width: 100%">


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
