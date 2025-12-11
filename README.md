# Genealogy Tree

## Overview

Genealogy Tree is a comprehensive digital platform designed to build, manage, and visualize interactive family trees.

The project provides a secure and scalable system for documenting ancestry, defining complex familial relationships, and archiving associated media (photos, documents).

## Features

* **Person Data Management:** Secure storage for core individual information (dates, locations, notes).
* **Relationship Mapping:** Flexible system for defining and tracking various familial bonds (Parent, Child, Spouse, Guardian).
* **Media Archiving:** Metadata storage and management for attached files (images, PDFs) linked to individuals.
* **Interactive Visualization:** (Frontend) A graphical interface for exploring the family tree structure.

## Architecture

The system utilizes a Monorepo structure with decoupled services. 

### Backend Service

* **Language:** Go (Golang)
* **Protocol:** gRPC (for high-performance communication)
* **Database:** PostgreSQL (for relational data storage)
* **Structure:** Clean Architecture (Model, Repository, Service layers).

### Frontend Application

* **Technology:** [To be determined]
* **Purpose:** User interface for data entry, editing, and visualization.

## Development Setup

### Prerequisites

* Docker and Docker Compose
* Go (1.20+)

### Quick Start

1.  **Launch Database:** Start the PostgreSQL container using Docker Compose.
    ```bash
    docker-compose up -d db
    ```
2.  **Run Backend:** Navigate to the `backend/` directory and run the main application.
    ```bash
    go run cmd/genealogy/main.go
    ```
