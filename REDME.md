
Getting Started
--------------- 
- For start use go run cmd/app/main.go
- Port config (by default 8081)
- TodoList/configs/config.yml 

API
---------------

POST
---------------
Reqest http://127.0.0.1:8081/api/v1/createTask
{
	"name":"task",
	"description":"description",
	"status":0
}

status can be 0 or 1 
- 0 - new task
- 1 - finish task

Response code 200
{
	"taskId": "taskId",
	"message": "message"
}

Response code 400, 500
{
	"message": "message"
}

PUT
---------------
Reqest http://127.0.0.1:8081/api/v1/updateTask/id (id - "taskId")
{
	"name":"task",
	"description":"description",
	"status":0
}

status can be 0 or 1 
- 0 - new task
- 1 - finish task

Response code 200
{
	"taskId": "taskId",
	"message": "message"
}

Response code 400, 404, 500
{
	"message": "message"
}


DELETE
---------------
Reqest http://127.0.0.1:8081//api/v1/deleteTask/id (id - "taskId")

response code 200
{
	"message": "message"
}

response code 404
{
	"message": "message"
}

GET
---------------
Reqest - http://127.0.0.1:8081/api/v1/getTaskItem/id (id - "taskId")

Response code 200
{
	"taskId": "taskId",
	"name": "name",
	"description": "description",
	"status": 0,
	"timeStump": 0
}

Response code 404
{
	"message": "message"
}

GET
---------------
Reqest - http://127.0.0.1:8081/api/v1/tasksList 

Response - code 200
{
	"tasks": [
		{
			"taskId": "taskId",
			"name": "taskId",
			"description": "description",
			"status": 0,
			"timeStump": 0
		}
	]
}

or "tasks" can be null if arr is empty

{
	"tasks": null
}

Response code 200
{
	"tasks": null
}

Response code 404
{
	"message": "message"
}