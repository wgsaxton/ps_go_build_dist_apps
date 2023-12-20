# Random notes while building and testing the registry service

## Testing with curl

Something should show up in the registry service if sending a message like this.
Will cause an error since not adding a service to the registry
```
curl -X POST -d 'This is a test message to the registry service' http://registryservice:3000/services
```

Need a better curl command to make something show up in the service's logs
