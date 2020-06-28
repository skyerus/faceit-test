# faceit-test

## Instructions

## Assumptions
1. The application is public facing.
2. Username and country are the only filterable fields (with wildcards either side). 
3. Event listeners are hardcoded and separated by event type e.g. post/put/delete
4. The user microservice receives a lot of traffic, creating the need for caching

## Room for improvement
1. Pagination for GET /users
2. Microservice for routing events
3. Outsource caching to a separate microservice

## Food for thought
Is it worth containerising the microservice with its database attached? Is this not costly if the application has a lot of microservices?
