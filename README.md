# National In-service Training API

This is the TEST API for the National In-service Training program. It provides endpoints for managing training sessions, personnel, courses, and related data.

## Setup

### Prerequisites

- Go (version 1.25.0 or newer)
- PostgreSQL
- `migrate` CLI tool

### Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/Syha-01/national-inservice-training.git
    cd national-inservice-training
    ```

2.  **Database Setup:**
    - Make sure your PostgreSQL server is running.
    - Set the `TRAINING_DB_DSN` environment variable with your database connection string. You can add it to a `.envrc`
      ```bash
      export TRAINING_DB_DSN="postgres://user:password@localhost/training_db?sslmode=disable"
      ```
    - Run the database migrations:
      ```bash
      make db/migrations/up
      ```

3.  **Run the application:**
    - For development (allows all CORS origins):
      ```bash
      make run/api
      ```
    - For production (restricts CORS origins):
      ```bash
      make run/api/prod
      ```
    The API will be running at `http://localhost:4000`.

## Available Endpoints

### Healthcheck

-   **GET `/v1/healthcheck`**
    -   **Description:** Checks the status of the API.
    -   **Response:**
        ```json
        {
            "status": "available",
            "system_info": {
                "environment": "development",
                "version": "1.0.0"
            }
        }
        ```

### National Inservice Training (NIT)

-   **POST `/v1/nits`**
    -   **Description:** Creates a new training session.
    -   **Request Body:**
        ```json
        {
            "course_id": 1,
            "start_date": "2025-12-01T00:00:00Z",
            "end_date": "2025-12-05T00:00:00Z",
            "location": "Training Academy"
        }
        ```

### Officers

-   **GET `/v1/officers`**
    -   **Description:** Retrieves a list of all officers.

-   **POST `/v1/officers`**
    -   **Description:** Creates a new officer.
    -   **Request Body:**
        ```json
        {
            "regulation_number": "12345",
            "first_name": "John",
            "last_name": "Doe",
            "sex": "Male",
            "rank_id": 1,
            "formation_id": 1,
            "posting_id": 1,
            "is_active": true
        }
        ```

-   **GET `/v1/officers/:id`**
    -   **Description:** Retrieves a specific officer by their ID.

-   **PATCH `/v1/officers/:id`**
    -   **Description:** Updates a specific officer's details.

-   **DELETE `/v1/officers/:id`**
    -   **Description:** Deletes a specific officer.

### Courses

-   **GET `/v1/courses`**
    -   **Description:** Retrieves a list of all courses.

-   **POST `/v1/courses`**
    -   **Description:** Creates a new course.
    -   **Request Body:**
        ```json
        {
            "title": "Advanced Investigation Techniques",
            "description": "A course on modern investigation methods.",
            "category": "Mandatory",
            "credit_hours": 40.5
        }
        ```

-   **GET `/v1/courses/:id`**
    -   **Description:** Retrieves a specific course by its ID.

-   **PATCH `/v1/courses/:id`**
    -   **Description:** Updates a specific course's details.

-   **DELETE `/v1/courses/:id`**
    -   **Description:** Deletes a specific course.

## Coming Soon Endpoints

The following endpoints are planned for future:

-   **Training Session Management:**
    -   `GET /v1/trainings`: List all training sessions.
    -   `POST /v1/trainings`: Create a new training session.
-   **User Authentication:**
    -   Endpoints for user registration, login, and token management.
-   **Enrollment:**
    -   Endpoints for enrolling officers in training sessions.
-   **Facilitators:**
    -   Endpoints for managing course facilitators.
-   **Ratings:**
    -   Endpoints for rating courses and facilitators.
