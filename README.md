# faceit-test

## Instructions
To start the containers:
```
docker-compose up -d
```
To run the tests:
```
docker exec -it faceit-test.local /bin/bash
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
4. Unit testing for the caching mechanism

## Food for thought
Is it worth containerising the microservice with its database attached? Is this not costly if the application has a lot of microservices?

## Troubleshooting
If you get an error similar to
```
dial tcp 172.18.0.2:3306: connection: connection refused
```
This is likely because the mysql container isn't ready yet (despite our service being dependant on it). Please wait a few seconds and try again.
