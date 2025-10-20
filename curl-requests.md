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
