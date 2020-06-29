# faceit-test

## Instructions
To start the containers:
```
docker-compose up -d
```
Exec into the go container:
```
docker exec -it faceit-test.local /bin/bash
```
To launch the service:
```
run
```
To run the tests:
```
run-tests
```
Note: You'll see non-fatal errors in the terminal when making POST/PUT/DELETE requests, this is because an event is being fired into nothing.
To disable events firing you can change the environment variable NO_EVENT_BROADCASTS=true in docker-compose.yml

## Assumptions
1. The application is public facing
2. Username and country are the only filterable fields (with wildcards either side)
3. Event listeners are hardcoded and separated by event type e.g. POST/PUT/DELETE
4. The user microservice receives a lot of traffic, creating the need for caching

## Room for improvement
1. Pagination for GET /users
2. Microservice for routing events
3. Outsource caching to a separate microservice

## Food for thought
Is it worth containerising the microservice with its database attached? Is this not costly if the application has a lot of microservices?
