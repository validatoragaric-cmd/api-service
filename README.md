# api-service
================

## Description
------------

The api-service is a RESTful API built using Node.js and Express.js. It provides a simple interface for managing users, products, and orders. The service is designed to be scalable, secure, and easy to maintain.

## Features
------------

### User Management

*   User creation and authentication
*   User retrieval by ID, username, or email
*   User update and deletion

### Product Management

*   Product creation and retrieval
*   Product update and deletion
*   Product filtering by category and price range

### Order Management

*   Order creation and retrieval
*   Order update and deletion
*   Order status updates (e.g., "pending", "shipped", "delivered")

### Other Features

*   API key authentication and authorization
*   Error handling and logging
*   API documentation using Swagger

## Technologies Used
-------------------

### Front-end

*   Node.js 14.x
*   Express.js 4.x
*   TypeScript
*   Swagger

### Back-end

*   MySQL 8.x
*   Sequelize ORM
*   JSON Web Tokens (JWT) for authentication

### Development

*   Docker
*   Docker Compose
*   ESLint
*   Prettier

## Installation
------------

### Prerequisites

*   Node.js 14.x installed
*   MySQL 8.x installed and running
*   Docker installed

### Project Setup

1.  Clone the repository: `git clone https://github.com/your-username/api-service.git`
2.  Install dependencies: `npm install`
3.  Create a `.env` file with your MySQL database credentials
4.  Run the database migrations: `npm run migration`
5.  Start the application: `npm start`
6.  Start the Docker containers: `docker-compose up`

### Running the Tests

1.  Run the tests: `npm test`

## API Documentation
-------------------

API documentation is available at `http://localhost:3000/api/docs`.

## Contributing
------------

Contributions are welcome! Please submit a pull request or create an issue to report any bugs or feature requests. Make sure to include a clear description of your changes and any relevant tests.