# Random notes while building and testing the log service

## Testing with curl
A good command to test the log service
```
curl -X POST -d 'This is a logging test message' http://logservice:4000/log
curl -X POST http://logservice:4000/services
```
