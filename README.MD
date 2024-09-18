# CrowdFunding-Golang-NuxtJS
## Summary:
This project implements Domain-Driven Design (DDD) to organize the code structure, focusing on business logic and scalability.
- Key Concepts:

    - Domain Layer: Contains the core business logic, including entities like User and services that manage user-related operations.
    - Application Layer: Handles the interaction between the user and the domain layer, processing requests and orchestrating business rules.
    - Infrastructure Layer: Manages external services like databases using GORM and PostgreSQL, implementing repositories for data persistence.
    - Interface Layer: Manages HTTP requests through Gin, providing endpoints and formatting responses for the API.

- Why DDD?

    - Focus on Business Logic: The architecture emphasizes the core domain, making it easier to maintain and evolve as the business grows.
    - Separation of Concerns: By separating the domain, application, infrastructure, and interface layers, the codebase is more modular and testable.
    - Scalable Architecture: DDD supports clear boundaries between components, making the system scalable and adaptable to change.

This repository has purposes:
- Learning how to build website
- Learning the algorithm mechanism of the website
- Learning how to code review and secure our code from malicious attackers.

### Technologies
- Database: postgresql
- Golang Framework: Gin
- ORM: Gorm
- Token Auth: JWT

## Documentation
### Entity Relationship Diagram
<a href="https://drive.google.com/file/d/17UWkw1TEQ6WGlj367Ysj8Pdl2TFtt30D/view?usp=sharing">click here!</a>

<img src="./repo_images/ERD.png">

### Database Diagram

<a href ="https://dbdiagram.io/d/CrowdFunding-66e85cb56dde7f4149453df2"> Click here!</a>

<img src="./repo_images/Database Diagram.png">


