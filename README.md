# Software Architecture Fundamental in Golang <sup>(Go)</sup>
IT's NOT FINISHED YET


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
    "full_name": ""
}
```
 * Design a Data Structure and store the data in chose format in database
 * Design a correlation between student data in database and stored photo
 * Store photos with in `jpeg` with format and `.jpg` extension which is an 
accepted format submitted by agent


### Demo Environment
You will be given access to one of University's VPS with the following hardware, 
software, and operating system spec.
 * Ubuntu Linux 20.04
 * 4 vCPU
 * 6 GB RAM
 * 30 GB SSD


## Architecture Design
<img src="./docs/architecture-design-diagram.png" style="max-width: 100%">


## Software Design
Going through each layer with an example and address of relevant file.

### Infrastructure Layer
Discussing postgres and http, and docker volume

### Service Layer
discuss abstract layers


### Business Layer
patterns for creating reusable services 


### Core Layer
All business related and requirement logic is here


## Deploy and Orchestrate with Docker
There is a `make` script at the root of this repository. You can run and stop 
container with `make up` and `make down`.


## RESTful Server

### Health Endpoint
When server is running successfully, you should be able to get a http status 
code `200` with following response from `/health` endpoint:

```sh
curl --location 'http://127.0.0.1/health'
```

```json
{
  "status": "healthy"
}
```

### Student Card Creation Endpoint(s)
There are currently two endpoint to facilitate Student Card registration.

#### Create Student Card Record
```sh
curl --location --request POST 'http://127.0.0.1/v1/student/card' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9=' \
--data '{
    "full_name": "John Doe",
    "student_id": "276163278231"
}'
```

#### Upload Photo Identification
```sh
curl --location --request PUT 'http://127.0.0.1/v1/student/card/photo' \
--header 'Content-Type: image/jpeg' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9=' \
--data '@/Users/hadi/Downloads/bart.jpeg'
```


```sh
curl \ 
-H 'Content-Type: application/json' \ 
-d "data=@path/to/my-file.txt" \ 
-d '{ "student_id":"JDOE1988232", "full_name":"John Doe" }' \ 
-X POST \ 
'http://127.0.0.1/v1/student/card'
```

`enctype="multipart/form-data"`
```sh
curl -X POST \ 
-F 'data=@path/to/local/file' http://127.0.0.1/v1/student/card
```


### Todo(s)
 * [ ] Correct Pattern use in Diagram Adapter is for database and etc
 * [ ] Make all docker containers production ready
 * [ ] Finalise submission and data structure with correlation to photo id
 * [ ] Finalise documentation
 * [ ] Investigate Volume Write permission issue for non-root users 


#### Credit
Author: [Hadi Tajallaei](mailto:hadi@syniol.com)

Copyright &copy; 2024 Syniol Limited. All rights reserved.
