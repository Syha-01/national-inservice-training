# Authorization Testing

## 1. Create a new user

This command will create a new user with the specified email and password. By default, new users are assigned the `nits:read` permission.

```bash
BODY='{"email":"testuser@example.com", "password":"password123"}'
curl -i -d "$BODY" localhost:4000/v1/users
```

From the response, you will need to copy the `token` value from the `user` object. This is the activation token.

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

## 6. (Optional) Give a user `nits:write` permission

To test the `nits:write` permission, you will need to manually add the permission to a user in the database. You can do this with the following SQL commands:

```sql
-- First, find the user's ID
SELECT id FROM users WHERE email = 'testuser@example.com';

-- Then, find the 'nits:write' permission's ID
SELECT id FROM permissions WHERE code = 'nits:write';

-- Finally, insert a new record into the users_permissions table
INSERT INTO users_permissions (user_id, permission_id) VALUES (<USER_ID>, <PERMISSION_ID>);
```

After running these commands, you can re-run the POST request from step 5, and it should succeed.
