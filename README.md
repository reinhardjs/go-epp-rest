# Domain EPP Rest in Go

![image](https://github.com/user-attachments/assets/d064b798-2294-4fbe-ae2b-e020ded32260)

The Extensible Provisioning Protocol (EPP) is a flexible protocol designed for allocating objects within registries over the Internet. The motivation for the creation of EPP was to create a robust and flexible protocol that could provide communication between domain name registries and domain name registrars.

The `go-epp-rest` is a RESTful API written in Go for managing domain registrations and interactions using the EPP (Extensible Provisioning Protocol) standard. This project focuses on leveraging Go's performance and concurrency features to provide a robust solution for domain registrars and resellers.

## Service Architecture

![image](https://github.com/user-attachments/assets/1461dad9-3360-4c91-89d3-c4ffa180f09e)


## Features

- **EPP Protocol Support**: Implements the EPP protocol for domain registration, transfer, renewal, and management.
- **RESTful API**: Provides a modern REST API interface for interacting with domain operations.
- **Performance**: Utilizes Go's concurrency model and efficiency for handling high loads.
- **Scalability**: Designed to scale horizontally to accommodate increasing demands.

## Installation

1. **Prerequisites**: Ensure you have Go installed on your system. You can download it from [golang.org](https://golang.org/dl/).

2. **Clone the Repository**:
   ```bash
   git clone https://github.com/reinhardjs/go-epp-rest.git
   cd go-epp-rest
   
3. **Build and Run**:
   ```bash
   go build -o domain-epp-rest
   ./domain-epp-rest
   
4. **Configuration**: Use the .env.example file to set up database connections, EPP server settings, and any other configuration specific to your environment.

## Usage
### API Endpoints
- **Check Domain Availability**: GET `/domain/check`
- **Register Domain**: POST `/domain/create`
- **Renew Domain**: POST `/domain/renew`
- **Get Contacts Info**: GET `/contact/info`

For detailed API endpoints list, go to `internal/delivery/http/router/router.go`.

## Contributing
Contributions are welcome! If you'd like to contribute to this project, please follow these steps:

Fork the repository and create your branch from main.
Make your changes and test thoroughly.
Ensure your code follows the Go coding style guidelines.
Submit a pull request describing your changes.

## License
This project is licensed under the MIT License - see the [LICENSE](https://github.com/reinhardjs/go-epp-rest/blob/main/LICENSE) file for details.

## Contact
For questions or support, please contact me at reinhardjsilalahi@gmail.com
