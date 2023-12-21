# Random notes while building and testing the teacher portal service

## Testing with curl

Something should show up in the teacher portal service logs if sending a message like this. Add a `-v` for more output if not seeing anything.
```
curl -X POST http://teacherportalservice:5001/students
curl -X POST http://teacherportalservice:5001/students/1
or localy in docker
curl -X POST http://localhost:5001/students
```