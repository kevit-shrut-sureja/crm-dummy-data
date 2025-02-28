# CRM Data Seeding Script

This script automates adding data to your CRM system. Follow these steps:

### 1. Setup the Database
- Create the required database for your CRM.

### 2. Start the Server & Run Migrations
- Ensure the server is running.
- Apply database migrations to set up tables.

### 3. Create a Super Admin User
- Insert a **Super Admin** user into the `users` table.

### 4. Configure the Script
- Retrieve the **Super Admin token** and **Base URL** from the system.
- Update the `main.go` file with these values.

### 5. Run the Script
- Choose how the leads should be distributed (**yes/no** prompt).
- The script will generate and insert leads based on your configuration.

### 6. Important Configuration Variables
Define the following variables before running the script:

```plaintext
TOTAL_RECORDS  
MAX_RECORDS_PER_WORKSPACE  
BATCH_SIZE  
MAX_WORKSPACES  
```

> **Note:** This script uses random number generation. If any variable is set too high or too low, it may cause inconsistencies in data distribution.

### 7. Happy Coding! 🚀💻😊