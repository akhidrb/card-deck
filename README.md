# Toggl Cards

![alt text](Cards.png)

#### How To Run:
- The docker-compose file in this directory contains the images for the API service and PostgreSQL.
- Apply the following command to run docker containers: `make start`

#### Run BDD Tests
- All the feature files for the Integration Tests are located [here](tests/features)
- To test all scenarios run the following command `go test -v ./tests/`
- To run sopme unit tests run the following command `go test -v ./pkg/services/`

#### API Documentation
- To test with a [Postman Collection](Toggl.postman_collection.json)
