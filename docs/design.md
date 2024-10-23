# Application Design

## Domain Driven Design (DDD)

The aim of DDD is to build a deep understanding of a core business domain and using this to shape the software's structure and functionality.

It emphasises close collaboration between technical and domain experts to create a shared model and a common understanding using ubiquitous language.

I have tried to adhere to DDD principles while designing Rollout and the domain is at the core of the application.  This is why you will see domain packages encapsulating core logic concerned with folders, pipelines, runs etc.

## Clean Architecture

Clean Architecture is a design philosophy that aims to create software that is highly flexible, maintainable and easy to test.

It focuses on organising code into layers, separating core business logic from external details like databases and frameworks.  The central idea is to structure the application in such a way that changes to one part of the system have minimal impact on others.

This is why Rollout code is organised as follows:

#### Domain
The innermost layer, sometimes referred to as "Entities", containing the core business logic and rules, independent of any specific application.

#### Use Cases
The next layer, containing all the application-specific business rules that define what the software should do.

#### Interface Adapters
The next layer, responsible for translating data between use cases and external systems like the web, databases or user interfaces.  In the case of Rollout, we have:

- **Controllers**: handle incoming HTTP requests, decode their JSON payloads, and hand off to Use Cases to handle application logic.
- **Repositories**: implementations of repository interfaces (defined in Use Cases) that handle data persistence e.g. database read/write operations.

#### Frameworks & Drivers
The outermost layer, which includes actual implementation details, like web servers and APIs.  In Rollout's design, we have the HTTP API layer that starts up the web server, defines routes and the controllers that handle them.

### Unidirectional Dependencies
Another key concept is unidirectional dependencies.  Outer layers can depend on inner layers, but never the other way round.  This avoids circular dependencies, keeps a clear separation of concerns, and makes testing or swapping out a specific implementation (such as database engine) much easier.

For example, a Repository can import a Domain Model, but a Domain Model must not have any awareness of the Repository Layer or database.  A Domain should not care how it is being used or persisted.

### Typical Flow

1. The HTTP API layer receives a request and routes it to the appropriate Controller.
2. The Controller decodes the request and hands off to a Use Case e.g. `CreateFolder`.
3. The Use Case handles any application-specific logic, calling Repository methods to handle database interactions.
4. The response flows back up through the Use Case, Controller, and HTTP API layers, returning a response the calling client.