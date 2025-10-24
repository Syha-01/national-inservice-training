> This is a small sample, please take a look at our official documentation here: https://documentation-v2-iota.vercel.app/docs

# National In-service Training API

This API provides endpoints for managing training sessions, personnel, courses, and related data for the National In-service Training program.

## Quick Start

### Prerequisites

- Go (version 1.25.0 or newer)
- PostgreSQL
- `migrate` CLI tool

### Installation & Setup

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/Syha-01/national-inservice-training.git
    cd national-inservice-training
    ```

2.  **Database Setup:**
    - Ensure your PostgreSQL server is running.
    - Set the `TRAINING_DB_DSN` environment variable. You can add it to a `.envrc` file:
      ```bash
      export TRAINING_DB_DSN="postgres://user:password@localhost/training_db?sslmode=disable"
      ```
    - Run database migrations:
      ```bash
      make db/migrations/up
      ```

3.  **Run the application:**
    - For development (allows all CORS origins):
      ```bash
      make run/api
      ```
    - The API will be running at `http://localhost:4000`.

## Sample Requests (cURL)

Here are a few examples of how to interact with the API using cURL.

### 1. Create and Activate a New User

First, create a new user. The response will include an activation token.

```bash
BODY='{"email":"testuser@example.com", "password":"password123"}'
curl -i -d "$BODY" localhost:4000/v1/users
```

Next, use the activation token from the previous response to activate the user.

```bash
# Replace <ACTIVATION_TOKEN> with the token from the previous step
curl -X PUT -d '{"token": "<ACTIVATION_TOKEN>"}' localhost:4000/v1/users/activated
```

### 2. Authenticate and Get a Token

Exchange the user's credentials for an authentication token.

```bash
BODY='{"email": "testuser@example.com", "password": "password123"}'
curl -i -d "$BODY" localhost:4000/v1/tokens/authentication
```

### 3. Access a Protected Route

Use the authentication token to access protected endpoints. This example fetches a list of NITs (National In-service Trainings).

```bash
# Replace <AUTH_TOKEN> with the authentication token from the previous step
curl -i -H "Authorization: Bearer <AUTH_TOKEN>" localhost:4000/v1/nits
```

### 4. Submit Feedback for a Facilitator

To submit feedback, you need to be enrolled in a session. A user **must be enrolled in a training session** before they can provide feedback. The system links feedback directly to an enrollment, not just to a user.

```bash
# Replace <AUTH_TOKEN>, <ENROLLMENT_ID>, and <FACILITATOR_ID> with actual values
curl -X POST \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <AUTH_TOKEN>" \
  -d '{
  "session_enrollment_id": <ENROLLMENT_ID>,
  "score": 5,
  "comment": "Excellent facilitator!"
}' "http://localhost:4000/v1/facilitators/<FACILITATOR_ID>/feedback"
```

## Authorization and Permissions

The API uses a permission-based authorization system. Users are assigned permissions that grant them access to specific endpoints. By default, new users are assigned the `nits:read` permission, which allows them to view training information.

To perform actions such as creating or updating data, users will need additional permissions (e.g., `nits:write`). These can be granted by an administrator. For example, to grant a user `nits:write` permission, an admin would make the following request:

```bash
# Authenticate as an admin to get an admin token
ADMIN_BODY='{"email": "admin@example.com", "password": "adminpassword"}'
ADMIN_TOKEN=$(curl -s -d "$ADMIN_BODY" localhost:4000/v1/tokens/authentication | jq -r .authentication_token.token)

# Add the 'nits:write' permission to the user
# Replace <USER_ID> with the ID of the user
USER_ID=<USER_ID>
PERMISSION_BODY='{"code":"nits:write"}'
curl -i -d "$PERMISSION_BODY" -H "Authorization: Bearer $ADMIN_TOKEN" localhost:4000/v1/users/$USER_ID/permissions
```

## Available Endpoints

A summary of the main endpoints. For a full list, please see the [official documentation](https://documentation-v2-iota.vercel.app/docs).

- **Healthcheck:** `GET /v1/healthcheck`
- **Users:** `POST /v1/users`, `PUT /v1/users/activated`
- **Tokens:** `POST /v1/tokens/authentication`
- **Permissions:** `POST /v1/users/:id/permissions`
- **NITs:** `GET /v1/nits`, `POST /v1/nits`
- **Officers:** `GET /v1/officers`, `POST /v1/officers`, `GET /v1/officers/:id`, `PATCH /v1/officers/:id`, `DELETE /v1/officers/:id`
- **Courses:** `GET /v1/courses`, `POST /v1/courses`, `GET /v1/courses/:id`, `PATCH /v1/courses/:id`, `DELETE /v1/courses/:id`
- **Feedback:** `POST /v1/facilitators/:id/feedback`
