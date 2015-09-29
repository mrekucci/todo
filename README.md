todo
====

[![Build Status](https://travis-ci.org/mrekucci/todo.svg)](https://travis-ci.org/mrekucci/todo)
[![GitHub license](https://img.shields.io/github/license/mashape/apistatus.svg)](LICENSE.txt)

A simple REST API based in-memory storage TODO application back-end.

Usage
-----

Compile and run the todo application back-end. By default, it listens on 127.0.0.1:8080.

Run the following commands to perform CRUD operations from the command line.

### Create

`curl -i -X POST -H "Content-Type: application/json" -d '{"title":"new"}' http://127.0.0.1:8080/task/`

### Read

`curl -i -X GET -H "Accept: application/json" http://127.0.0.1:8080/task/0`

### Read all

`curl -i -X GET -H "Accept: application/json" http://127.0.0.1:8080/task/`

### Update

`curl -i -X PUT -H "Content-Type: application/json" -d '{"id":0,"title":"update"}' http://127.0.0.1:8080/task/0`

### Delete

`curl -i -X DELETE http://127.0.0.1:8080/task/0`