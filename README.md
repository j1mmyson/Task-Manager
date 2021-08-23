# ToDoCalendar

`ToDoCalendar` is side project that helps you to manage your task on a daily basis in web browser.  
maybe you can improve your task efficiency and trace what to do next.  
I inspired by my internship experiments. It helps me a lot to note lists every morning in below pattern.

```
--Done--------------------
- what you done 

--In process--------------
- what you are doing now
- 2

--Todo--------------------
- what to do next

```

## Demo
#### try try!  
`public id`: testid  
`password`: qwe123  
Also you can sign up and use it!  
<http://ec2-3-17-39-222.us-east-2.compute.amazonaws.com/>

## Requirements

- Golang >= 1.16.2
- mysql

## Run in local

1. Clone this repository
   `$ git clone https://github.com/j1mmyson/todo_calendar.git`

2. Enter cloned repository
   `$ cd todo_calendar`

3. Create `.env`file at root folder(where `main.go` exists)

   ```bash
   DBNAME=<your db name>
   User=<User>
   Host=127.0.0.1
   Password=<your db password>
   ```

4. Build binary file
   `$ go build -o executeServer`

5. Run binary file
   `$ ./executeServer`

6. Open browser and enter http://localhost:8080/

## Future Feature
...
