# TransactionService


## Directory and File Structure

The application follows the hexagonal architecture (ports and adapters pattern), with a clear separation of responsibilities. Here is the proposed directory and file structure:

```
/myapp
├── cmd
│   └── server
│       └── main.go        # Application entry point, sets up and launches the HTTP server
├── pkg
│   ├── config             # Application configurations (e.g., environment variables)
│   ├── core               # Central business logic
│   │   └── service.go     # Business logic implementation
│   ├── handlers           # HTTP handlers, convert HTTP requests/responses
│   ├── infrastructure     # Everything related to external infrastructure (e.g., transaction)
│   ├── ports              # Interfaces (ports) that define how to communicate with the outside
│   │   ├── http           # HTTP port
│   │   │   └── router.go  # Router configuration and endpoints
│   │   └── transaction    # Port for interacting with transaction
│   └── repository         # Data access (e.g., persistent storage)
└── internal               # Project-specific code that should not be exposed

```

### Component Description
- **cmd/server/main.go**: Application entry point. Configures and launches the HTTP server.
- **pkg/config**: Contains application configurations.
- **pkg/core**: Central business logic.
- **pkg/handlers**: Handlers for HTTP requests.
- **pkg/infrastructure**: Management of external infrastructures, such as transaction.
- **pkg/ports**: Interfaces for external interaction.
- **pkg/repository**: Data access and management.
- **internal**: Project-specific code, not accessible externally.

## Getting Started

This guide assumes that you have PostgreSQL and Go already installed and configured on your system.

### Environment Variables

Before starting, ensure you have set the following environment variables or be ready to use the defaults:

- `DB_USER` - Database username (default: `root`)
- `DB_PASSWORD` - Database user password (default: `rootPassword`)
- `DB_NAME` - Database name for the service (default: `stori_service`)
- `DB_HOST` - Database host (default: `localhost`)
- `DB_PORT` - Database port (default: `5432`)

### Steps

1. **Create the Database**:

   Run the following command to create a new database using the credentials specified by the environment variables or the defaults set in the Makefile:
```
make createdb
```

2. **Run Migrations**:

Execute the database migrations with the following command:

```
make migrate
```

This will apply all the migration scripts located in the `database/` directory.


3. **Tidy Dependencies**:

Clean up the module by removing unused dependencies and adding indirect ones needed by other dependencies:
```
make tidy
```

4. **Run the Service**:

Finally, to start the service, use the following command:
```
make run
```

This will start the `user-service` application using the `main.go` file at the root of the project.

After following these steps, your `user-service` should be up and running, connected to your PostgreSQL database, and ready to handle requests.


