# Parcel-Service

[![Actions Status](https://github.com/teachmind/Parcel-Service/workflows/build/badge.svg)](https://github.com/teachmind/Parcel-Service/actions)
[![codecov](https://codecov.io/gh/teachmind/Parcel-Service/branch/master/graph/badge.svg?token=HivKkjhfjl)](https://codecov.io/gh/teachmind/Parcel-Service)
[![Go Report Card](https://goreportcard.com/badge/github.com/teachmind/Parcel-Service)](https://goreportcard.com/report/github.com/teachmind/Parcel-Service)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/934b654ea9eb4f72b98138b21b5aea94)](https://www.codacy.com/gh/teachmind/Parcel-Service/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=teachmind/Parcel-Service&amp;utm_campaign=Badge_Grade)
[![](https://godoc.org/github.com/teachmind/Parcel-Service?status.svg)](https://godoc.org/github.com/teachmind/Parcel-Service)

## Features 
-   Server Preparation for Running the project on localhost
-   Database Migration
-   Parcel Create
-   Parcel Details
-   Parcel Update
-   Carrier Request
-   Carrier Selection
-   Available Parcel List

## Feature Details
### Database Migration
-   Database schema will be created when migrating
-   Rollback query is added for making the database empty
### Parcel Create
-   User can create a parcel to send a location
-   Validation for required fields
### Parcel Details
-   Parcel details endpoint for user and carrier
-   Request parcel details by Parcel ID
### Parcel Update
-   Update parcel status based on the user or carrier action

## Project Structure
    .
    |-- cmd                 # Contains the commands for the project
    |-- images              # Contains all image file
    |-- internal            # Configuration files and Constants
    |-- migration           # Contains migration files
    |-- .env.example        # example/structure of .env file
    |-- Dockerfile          # Used to build docker image.
    |-- go.mode             # Define's the module's import path used for root directory
    |-- go.sum              # Contains the expected cryptographic checksums of the content of specific module versions
    |-- Makefile            # Makefile to run commands after docker up
    |-- readme.md           # Explains project installation and other informations

## Tools and Technology
-   Golang
-   PostgreSQL

## Installation
-   **Step-1:** Copy/rename `.env.example` file as `.env`. Change the `APP_PORT`, `DB_PORT`, `DB_NAME`,`DB_HOST`, `DB_USER`, `DB_PASSWORD` value as per your DB and Project setup.
-   **Step-2:** Run migration command `make migrate` for Database migration
-   **Step-3:** To start server run `make server`
