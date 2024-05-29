# Dockerize

## Description
Dockerize is the dockerization [the ASCII Art Web] which is a simple web application that allows users to generate ASCII art from text input using various banner styles.

## Usage: How to Run
To dockerize the  [the ASCII Art Web] application, follow these steps:
1. Make sure [Docker]() is set on your machine.
1. Clone the repository
2 . Build the Docker image: `go build -o main .`
4. Run the Docker container: `docker run -p 8080:8080 main`
5. Access the application in your web browser at: `http://localhost:8080`

## Implementation Details: Algorithm
The ASCII Art generation algorithm works as follows:
1. The user provides input text and selects a banner style.
2. The server processes the input and retrieves the corresponding banner pattern from a file.
3. Each character in the input text is mapped to its corresponding pattern in the banner.
4. The banner pattern is applied to each character to create the ASCII art representation.
