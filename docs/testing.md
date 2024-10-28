# Testing

* Unit testing the repository layer that uses an ORM is considerable effort (requires mocking  GORM's functionality), will lead to brittle tests that are prone to breakage and will give false confidence.  TODO - research the best way of integration testing against a real database e.g. in-memory SQlite database.