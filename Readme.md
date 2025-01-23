![logo](./public/logo.png)


**RouterFlow** is an open-source Command-Line Interface (CLI) tool designed to streamline the configuration, monitoring, and management of Mikrotik routers. Tailored for Internet Service Providers (ISPs), this tool emphasizes scalability, security, and ease of use while encouraging community collaboration and contributions.

## Getting Started

### Prerequisites

Ensure you have the following installed on your system:

- [Go](https://golang.org/doc/install) (version go 1.23.3 or higher)
- Git
- A Mikrotik router with API access enabled

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/hermangoncalves/RouterFlow.git
   cd RouterFlow
   ```

2. Build the CLI tool:
   ```bash
   go build -o routerflow .
   ```

3. Add the executable to your PATH for easy access (optional):
   ```bash
   export PATH=$PATH:$(pwd)
   ```

### Usage

Run the tool to see available commands:
```bash
routerflow --help
```

Example of connecting to a router:
```bash
routerflow connect --address=192.168.88.1 --user=admin --password=yourpassword
```


## Contributing

Contributions are welcome! To get started:

1. Fork the repository.
2. Create a new branch for your feature or fix.
3. Commit your changes and open a pull request.

Make sure to follow the [contributing guidelines](docs/contributing.md) (to be added).

## License
This project is licensed under the Apache 2.0 License. See the [LICENSE](https://github.com/hermangoncalves/RouterFlow/blob/main/LICENSE) file for details.

---

For questions or support, open an issue on GitHub.

