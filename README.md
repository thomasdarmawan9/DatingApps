# Dating App Backend REST API with Golang

Welcome to the repository for a **Dating App Backend REST API** built using **Golang**. This project serves as the backend service for a dating application, providing all the necessary API endpoints for user management, profile creation, photo uploads, matching, and more.

## Features

- **User Authentication & Authorization:**
  - JWT-based authentication for secure login and session management.
  - Middleware to protect endpoints and restrict access based on user roles or profile completeness.

- **Profile Management:**
  - CRUD operations for user profiles, including creating, updating, and retrieving profile details.
  - Integration with photo uploads for profile pictures.
  - Swipe functionality with limits based on user subscription (e.g., free or premium users).

- **Matching System:**
  - Swipe left/right feature allowing users to express interest in other profiles.
  - Real-time matching algorithm that identifies mutual interests and triggers a match notification.
  - Match history and management, allowing users to view and interact with their matches.

- **Subscription Management:**
  - Support for different user subscription tiers (e.g., free and premium).
  - Enforcement of feature restrictions based on the user's subscription level (e.g., swipe limits for free users).

- **Security:**
  - Comprehensive error handling and input validation to ensure API security and robustness.
  - Secure password storage using hashing and salting techniques.

## Tech Stack

- **Golang**: High-performance and efficient server-side language.
- **Gin Framework**: Lightweight and fast HTTP web framework for building RESTful services in Go.
- **GORM**: Powerful ORM library for Golang, used for database interactions.
- **PostgreSQL**: Relational database system used to store user data, profiles, matches, and more.
- **JWT (JSON Web Tokens)**: For secure user authentication and authorization.

## Project Structure

- `controllers/`: Contains the logic for handling API requests and responses.
- `models/`: Defines the data models and structures used in the application.
- `middlewares/`: Middleware functions for tasks such as authentication, authorization, and request validation.
- `helpers/`: Utility functions and helpers for tasks like token generation and password hashing.
- `routes/`: Defines the routing for all API endpoints.
- `config/`: Configuration files for database connections and other settings.

## Getting Started

### Prerequisites

- Go 1.16 or later installed on your machine.
- PostgreSQL database setup.
- A tool like `Postman` for testing API endpoints.

### Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/yourusername/dating-app-backend.git
   cd dating-app-backend
   ```

2. **Install dependencies:**

   ```sh
   go mod tidy
   ```

3. **Set up environment variables:**

   Create a `.env` file and configure your database connection string, JWT secret, and other necessary environment variables.

4. **Run database migrations:**

   Apply migrations to set up the database schema.

   ```sh
   go run main.go migrate
   ```

5. **Start the server:**

   ```sh
   go run main.go
   ```

6. **Access the API:**

   The API will be accessible at `http://localhost:8080`. You can use tools like `Postman` to test the endpoints.

## API Documentation

Swagger documentation is available for all endpoints. To access the documentation, visit:

```
http://localhost:8080/swagger/index.html
```

## Contributing

We welcome contributions from the community! If you'd like to contribute, please fork the repository and submit a pull request with your changes. Be sure to include tests for your new features or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

---

Feel free to customize this description to better fit your project's specifics.
