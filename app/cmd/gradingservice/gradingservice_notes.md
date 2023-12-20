# Random notes while building and testing the grading service

## Testing with curl

Something should show up in the grading service logs if sending a message like this
```
curl -X POST http://gradingservice:6000/students
curl -X POST http://gradingservice:6000/students/1
or localy in docker
curl -X POST http://localhost:6000/students
```