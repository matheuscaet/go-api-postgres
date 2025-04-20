# go-api-postgres

## Overview

This is a RESTful API built with Go. It provides endpoints for managing.

```/health (should implements database health)
/item -> POST to create new database item
/item:itemId -> GET to get an item by id
/items -> GET to return all database rows

## Prerequisites
1. **Go**: Ensure you have Go installed. You can download it from [golang.org](https://golang.org/).
2. **Database**: Install the required database Postgres, or use your cloud database.
3. **Environment Variables**: Set up the necessary environment variables.

## Environment Variables
Create a `.env` file in the root directory of your project and add the following variables:

```env
DBCONNECTION=postgres://YourUserName:YourPassword@YourHostname:5432/YourDatabaseName
GIN_MODE=debug
APP_PORT=8080