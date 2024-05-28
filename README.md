# Docker Container Efficiency Benchmark

This repository contains a benchmark comparing the efficiency of Docker containers using Go and Node.js to create simple HTTP servers. The results demonstrate the advantages of using Go for smaller and more efficient containers.

## Prerequisites

Before running the benchmarks, ensure you have the following installed:

- Docker
- Node.js
- Go
- `autocannon` or any other HTTP benchmarking tool

### Installing Docker

Follow the instructions on the [Docker website](https://docs.docker.com/get-docker/) to install Docker on your machine.

### Installing Autocannon

To install `autocannon`, run:

```bash
npm install -g autocannon
```

## Project Structure

- `golang`: Folder to golang code
- `nodejs`: Folder to node js code

## Building the Docker Containers

### Go

1. Navigate to the directory `golang`.
2. Build the Docker image:

    ```bash
    docker build . -t go-lang-http-basic
    ```

3. Run the Docker container:

    ```bash
    docker run -p 8080:8080 --cpus 1 --memory 256m go-lang-http-basic:latest
    ```

### Node.js

1. Navigate to the directory `node-js`.
2. Build the Docker image:

    ```bash
    docker build . -t node-http-basic
    ```

3. Run the Docker container:

    ```bash
    docker run -p 9090:9090 --cpus 1 --memory 256m node-http-basic:latest
    ```

## Running the Benchmark

### Go

Run the benchmark against the Go server:

```bash
autocannon "http://localhost:8080/test_go_122?start=1&end=100" -c 1000 -d 60
```

### Node.js

Run the benchmark against the Node.js server:

```bash
autocannon "http://localhost:9090/test_node_20?start=1&end=100" -c 1000 -d 60
```

## Results

After running the benchmarks, you can compare the results:

- **Go:**
  - Container size: 9 MB
  - Average requests/second: 19,360.74
  - Average latency: 51.23 ms

- **Node.js:**
  - Container size: 266 MB
  - Average requests/second: 6,000.52
  - Average latency: 58.76 ms

## Conclusion

The benchmarks demonstrate that Go produces smaller and more efficient Docker containers compared to Node.js. Smaller containers mean less time to start and less resource usage, which is crucial for serverless functions like AWS Lambda and Google Cloud Functions.

## Disclaimer

Go is not a magic solution. Its implementation should be analyzed by the team to assess whether it is worth implementing. Depending on the team's experience, Node.js solutions might meet most scenarios and be easier for developers to maintain in the long term.

---

Feel free to contribute and raise issues if you encounter any problems or have suggestions for improvements.

---

### License

This project is licensed under the MIT License.

---

### Contact

For any questions, feel free to reach out.

---