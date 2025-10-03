# Endpoint Table

This is a table of all the api endpoints the backend will provide to the client

## The Table

| Action       | HTTP Method | Endpoint                   | Description                                                      |
| ------------ | ----------- | -------------------------- | ---------------------------------------------------------------- |
| Create user  | POST        | /new-user                  | Add a newly logged user to the database                          |
| View user    | GET         | /users/{id}                | View information about a certain user                            |
| Create event | POST        | /new-event                 | Create a new event with user provided informatino                |
| Join event   | POST        | /join/{user_id}/{event_id} | Let a user communicate their intent of attending a created event |
