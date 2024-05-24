# Even Todos Fetcher

This command-line tool fetches even-numbered (even IDs) todos from the API at [https://jsonplaceholder.typicode.com/todos/{id}](https://jsonplaceholder.typicode.com/todos/{id}). Currently, it fetches the first 20 even-numbered todos without accepting any command-line arguments. Future releases may include support for passing arguments.

## Motivation

To ensure optimal performance, the tool fires all API requests simultaneously using separate goroutines. Since these requests are independent, they are initiated concurrently. The resulting todo objects are collected in a shared slice, with a mutex lock used to prevent race conditions. A WaitGroup ensures all goroutines complete before the results are printed to the console.

### Performance

- Without goroutines: ~5 seconds
- With goroutines: ~500 milliseconds

## Tests

Unit tests have been written for the `FetchEvenTodos` function, which spawns goroutines. The API call function itself is not currently tested.

## Dockerfile

The project includes a multi-stage Dockerfile to create a very lean image. The final image size is approximately 12MB.

## Running the Tool

1. Clone the GitHub repository.
2. Navigate into the repository directory on your local machine.
3. Build the Docker image:
   `docker build -t demyst .`
   
4. Run the Docker container:
   `docker run demyst`

## Running Tests

To run the tests, navigate to the code directory and execute:
`go test ./...`





