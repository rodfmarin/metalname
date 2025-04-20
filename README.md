# MetalName Generator

MetalName Generator is a Go-based application that generates random staff names by combining first and last names from a predefined list. It's simple, fast, and perfect for creating unique names for your projects or characters.

## Features

- **Random Name Generation**: Combines random first and last names to create unique staff names.
- **Customizable**: Easily extend the name lists to suit your needs.
- **Thread-Safe Randomization**: Uses a local random generator for better concurrency.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/riot-rmarin/metalname.git
   cd metalname

2. Build the application:
   ```bash
    go build -o metalname
    ```
3. Run the application:
   ```bash
    ./metalname
   > Staff Name: randomized snake
    ```