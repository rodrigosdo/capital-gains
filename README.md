### README

# Capital Gains Calculator

This project is a Go application that calculates capital gains taxes based on a series of buy and sell operations.

## Usage

This requires Go to be installed on your machine. In case you don't have or want to install Go, you can use Docker to run the application.

### Build the Application

To build the application, run:

```sh
make build
```

### Run the Application

To run the application, use:

```sh
./bin/capital-gains < input.json
```

### Run Tests

To run the tests, execute:

```sh
make test
```

### Clean the Build

To clean the build artifacts, use:

```sh
make clean
```

### Docker

#### Build Docker Image

To build the Docker image, run:

```sh
make docker-build
```

#### Run Docker Container

To run the application inside a Docker container, use:

```sh
make docker-run
```

### Help

To display the available Makefile commands, run:

```sh
make help
```
