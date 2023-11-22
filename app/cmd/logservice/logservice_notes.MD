# Random notes while building and testing the log service

## Testing with curl
A good command to test the log service
```
curl -X POST -d 'This is a logging test message' http://localhost:4000/log
```

Something should show up in the registry service if sending a message like this
```
curl -X POST -d 'This is a test message to the registry service' http://localhost:3000/services
```
