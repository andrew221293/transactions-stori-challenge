# Transactions Stori Challenge
Stori Software Engineer Technical Challenge

For this challenge, clean architecture is being worked on, which has the following Layers:
1. entity (It is the layer that contains the struct)
2. store (handles the connection to the database)
3. transport (handles the data and validates the information)
4. usecase (handles business logic)

Steps to run the code in localhost
======
1. [Clone the repository.](https://github.com/andrew221293/transactions-stori-challenge)
2. Run `go get` to install dependencies.
3. Set the environment variables, which were sent by email

Steps to run the code in Docker (assuming you have docker installed)
======
1. [Clone the repository.](https://github.com/andrew221293/transactions-stori-challenge)
2. Add the dockerfile that will be sent by email
3. Run `docker build -t transactions-stori-challenge .` to build a image.
4. Run `docker run -it -d -p 8080:8080 transactions-stori-challenge` to run system.
5. The system should be up on localhost

List of required environment variables
======
1. MONGO_DATABASE
2. BASIC_AUTH_PASSWORD
3. BASIC_AUTH_USER
4. MONGO_USER
5. MONGO_PASSWORD
6. MONGO_HOST
7. BACKEND_HOST
8. SENDGRID_API_KEY

