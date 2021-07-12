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

## Will be like..

![](https://github.com/j1mmyson/ToDoCalendar/blob/main/src/image/prototype.PNG?raw=true)

## Requirements

- Golang >= 1.16.2
- mysql

## Usage

1. Clone this repository
   `$git clone https://github.com/j1mmyson/todo_calendar.git`

2. Enter cloned repository
   `$cd todo_calendar`

3. Create `.env`file at root folder(where `main.go` exists)

   ```bash
   DBNAME=<your db name>
   User=<User>
   Host=127.0.0.1
   Password=<your db password>
   ```

4. Build binary file
   `$go build -o excuteServer`

5. Run binary file
   `$./excuteServer`

6. Open browser and enter http://localhost:8080/

