# Authorization Testing

## 1. Create a new user

This command will create a new user with the specified email and password. By default, new users are assigned the `nits:read` permission.

```bash
BODY='{"email":"testuser@example.com", "password":"password123"}'
curl -i -d "$BODY" localhost:4000/v1/users
```

From the response, you will need to copy the `id` of the user and the `token` value from the `user` object. This is the activation token.

## 2. Activate the new user

Replace `<ACTIVATION_TOKEN>` with the token you copied from the previous step.

```bash
curl -X PUT -d '{"token": "<ACTIVATION_TOKEN>"}' localhost:4000/v1/users/activated
```

## 3. Authenticate as the new user

This command will exchange the user's email and password for an authentication token.

```bash
BODY='{"email": "testuser@example.com", "password": "password123"}'
curl -i -d "$BODY" localhost:4000/v1/tokens/authentication
```

From the response, you will need to copy the `token` value from the `authentication_token` object.

## 4. Test `nits:read` permission (should succeed)

Replace `<AUTH_TOKEN>` with the authentication token you copied from the previous step. This request should succeed because the user has the `nits:read` permission.

```bash
curl -i -H "Authorization: Bearer <AUTH_TOKEN>" localhost:4000/v1/nits
```

## 5. Test `nits:write` permission (should fail)

This request should fail with a `403 Forbidden` error because the user does not have the `nits:write` permission.

```bash
BODY='{"title":"New NIT", "content":"This should not work."}'
curl -i -d "$BODY" -H "Authorization: Bearer <AUTH_TOKEN>" localhost:4000/v1/nits
```

## 6. Give a user `nits:write` permission

To give a user the `nits:write` permission, you will need to be authenticated as a user with the `admin:all` permission.

First, authenticate as an admin user (you may need to create one and assign the `admin:all` permission manually in the database if you don't have one).

```bash
# Authenticate as admin
ADMIN_BODY='{"email": "admin@example.com", "password": "adminpassword"}'
ADMIN_TOKEN=$(curl -s -d "$ADMIN_BODY" localhost:4000/v1/tokens/authentication | jq -r .authentication_token.token)
```

Now, you can add the `nits:write` permission to the user you created in step 1. Replace `<USER_ID>` with the ID of the user.

```bash
USER_ID=<USER_ID>
PERMISSION_BODY='{"code":"nits:write"}'
curl -i -d "$PERMISSION_BODY" -H "Authorization: Bearer $ADMIN_TOKEN" localhost:4000/v1/users/$USER_ID/permissions
```

After running this command, you can re-run the POST request from step 5 (with the non-admin user's token), and it should succeed.

---

# Facilitator and Course Feedback

## Submitting Feedback for a Facilitator

This command sends feedback for a specific facilitator.

```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <AUTH_TOKEN>" \
  -d '{
  "session_enrollment_id": <ENROLLMENT_ID>,
  "score": 5,
  "comment": "Excellent facilitator!"
}' "http://localhost:4000/v1/facilitators/<FACILITATOR_ID>/feedback"
```

### Command Breakdown:

*   `curl`: The command-line tool used to make HTTP requests.
*   `-X POST`: Specifies that this is an HTTP `POST` request, which is used to create a new resource (in this case, a new feedback entry).
*   `-H "Content-Type: application/json"`: This is a header that tells the server the body of our request is in JSON format.
*   `-H "Authorization: Bearer <AUTH_TOKEN>"`: This is the authorization header. You must include a valid authentication token (obtained in step 3) to prove who you are.
*   `-d '{...}'`: This is the data, or body, of the request. It contains the actual feedback information.
    *   `"session_enrollment_id"`: **CRITICAL**. This must be the ID of a user's enrollment in a specific session.
    *   `"score"`: The numerical rating.
    *   `"comment"`: The text feedback.
*   `"http://localhost:4000/v1/facilitators/<FACILITATOR_ID>/feedback"`: The API endpoint. You must replace `<FACILITATOR_ID>` with the ID of the facilitator you are rating.

### **Important Note on Enrollment**

A user **must be enrolled in a training session** before they can provide feedback. The system links feedback directly to an enrollment, not just to a user.

1.  **Enroll a user** in a session. This will generate a unique `session_enrollment_id`.
2.  **Use that `session_enrollment_id`** in the body of your feedback request.

If you try to use a `session_enrollment_id` that doesn't belong to the authenticated user or one that has already been used to submit feedback for the same facilitator, the request will fail. This is the reason for the "duplicate key" error.
