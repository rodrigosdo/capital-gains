# Capital Gains Calculator

This project is a Go application that calculates capital gains taxes based on a series of buy and sell transactions.

## Technical and Architectural Decisions

- The project was developed using Go mainly due to its simplicity and binary compilation, which makes it easy to distribute and run the application on different platforms.
- Docker was used to facilitate the distribution and execution of the application, as well as to ensure that it runs in the same environment regardless of the host machine.
- There's also a Makefile to facilitate the building, testing, and running of the application, as well as the creation and running of Docker containers.

### Project Structure

The project was structured to clearly separate responsibilities:

- **bin**: Where the executable binary is stored when the application is compiled.
- **cmd/cli/main.go**: Contains the application's entrypoint.
- **cmd/cli/internal/testdata**: Contains test data files.
- **internal/domain**: Contains the domain entities and the logic for calculating capital gains.

## Build and Execution

### Build

To build the application, run:

**Note**: Go is required to build the application locally.

```sh
make build
```

A binary will be generated in the `bin` directory.

### Run Locally

To run the application, use:

```sh
./bin/capital-gains < input.json
```

There are some test data files in the `cmd/cli/internal/testdata` directory that can be used as input, like:

```sh
./bin/capital-gains < ./cmd/cli/internal/testdata/case_1_input.golden
```

It's also possible to run the application with the following command:

```sh
./bin/capital-gains
```

And then type the transactions manually.

### Tests

To run the tests, run:

```sh
make test
```

To run the tests with coverage, use:

```sh
make test-coverage
```

### Docker

To build the Docker image, run:

```sh
make docker-build
```

To run the application inside a Docker container, use:

```sh
make docker-run
```

You can interact the same way as running the application locally.

### Help

To display the commands available in the Makefile, run:

```sh
make help
```

## Additional Notes

- **Custom JSON Marshaller**: I have implemented a custom JSON marshaller for the `Tax` structure to ensure that values ​​are rounded to two decimal places. This is a silly workaround, but it was the simplest way to achieve this.
- **Parallel Tests**: Tests run using `t.Parallel()` to ensure that they are executed concurrently, improving test efficiency, and to guarantee that they are isolated from each other.
