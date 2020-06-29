# faceit-test

## Instructions
To start the containers:
```
docker-compose up -d
docker ps
```
Check that both containers are listed in `docker ps`. If our service isn't listed, please follow [troubleshooting](#troubleshooting)

To run the tests:
```
docker exec -it faceit-test.local /bin/bash
run-tests
```

Note: You'll see non-fatal errors in the logs when making POST/PUT/DELETE requests, this is because an event is being fired into nothing.
To disable events firing you can change the environment variable `NO_EVENT_BROADCASTS=true` in docker-compose.yml

## Assumptions
1. The application is public facing
2. Username and country are the only filterable fields (with wildcards either side)
3. Event listeners are hardcoded and separated by event type e.g. POST/PUT/DELETE
4. The user microservice receives a lot of traffic, creating the need for caching

## Room for improvement
1. Pagination for GET /users
2. For production: Package the application in an alpine container containing just the executable
3. Microservice for routing events
4. Outsource caching to a separate microservice
5. Unit testing for the caching mechanism

## Food for thought
Is it worth containerising the microservice with its database attached? Is this not costly if the application has a lot of microservices?

## Troubleshooting
```
docker logs faceit-test.local
```
If you get an error similar to
```
dial tcp 172.18.0.2:3306: connection: connection refused
```
This is likely because the mysql container isn't healthy yet (despite our service being dependant on it). Please wait for the mysql container to become healthy (`docker ps`) and `docker-compose up -d` again.
