# Software Architecture Fundamental in Golang <sup>(Go)</sup>
This repository is for demonstration of Layered Software Architecture which 
is a fundamental concept of clean architecture in software engineering context. 

This is a hands-on approach, and it's result of spending weekends, afternoon, 
evenings, and holidays studying books and watching countless video trainings 
to master software architecture. Courses and books only give you the concepts 
with a few code examples but having a proper example of creation from layer 
to layer in a small project is impossible to find.


## Overview
Create a RESTful API with a single endpoint to store students card for all 
university's new admissions. You will be give a `SSH` access to one of 
university's VPS assigned to Computer Science students. You will need to create 
an Architecture Design, Software design, and its implementation as a final 
product. All submissions will be assessed and chosen to be used by University 
this year for admission process.

### Requirements

 * Portrait photo for all issued students card should be stored in VPS
 * API request will be made through web application with following JSON format:
 ```json
{
    "student_id": "",
    "full_name": "",
    "portrait_photo": ""
}
```
 * Design a Data Structure and store the data in chose format in database
 * Design a correlation between student data in database and stored photo
 * Store photos with in `jpg` format which is an accepted format submitted by agent


### Demo Environment
You will e given access to one of University's VPS with following hardware, 
software, and operating system.
 * Ubuntu Linux 20.04
 * 4 vCPU
 * 6 GB RAM
 * 30 GB SSD


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
