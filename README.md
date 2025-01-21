# CRUD API

Project for testing go's net/http package in version 1.22+

Also test the following:

- Embed package
- Database repository pattern

### Endpoints

| Endpoint   | Method | Objective                                                              |
| ---------- | ------ | ---------------------------------------------------------------------- |
| /          | any    | General catch endpoint, returns a simple message                       |
| /coffee    | any    | Reference to Hyper Text Coffee Pot Control Protocol                    |
| /data/{id} | GET    | Search DB for entry by id, returning full object                       |
| /data/     | POST   | Insert new data on DB, the trailing slash is required due to a bug. :) |
| /data/{id} | DELETE | Delete a entry using it's id                                           |
| /data/{id} | PUT    | Update entry value using ids as search value                           |
