# K-Tax โปรแกรมคำนวนภาษี

K-Tax เป็น Application คำนวนภาษี ที่ให้ผู้ใช้งานสามารถคำนวนภาษีบุคคลธรรมดา ตามขั้นบันใดภาษี พร้อมกับคำนวนค่าลดหย่อน และภาษีที่ต้องได้รับคืน

## Getting Started

1. First you have to clone this repository
2. After that, you must run `assignment-tax-postgres-1` in Docker
3. Also, you have to uncomment and run `intiDB.go` to create your own Database for use this program.
4. And then, you can use this program from run `main.go`
5. So, you can follow the `User stories` or use with your own story

### Software Feature

1. **swagger**
2. **adjustable tax level**
3. **deduct value with version**
4. **MVC Style**
5. **Some Testing and Integration**
6. **Transaction Update**
7. ***Graceful Shutdown***

---
#### Testing the Application

Need to run application on port 8080 for integration test

---
### How to Run the Program

#### Setup Instructions

1. **Start the Database**:
   Ensure Docker is installed on your machine. Begin by starting the database using Docker Compose:
   ```
   docker-compose up 
   ```

2. **Initialize the Database**:
   After the database service is running, execute the following Go script to create tables based on the SQL files: ***before runnig please uncomment and comment back after run***
   ```
   go run initDb.go
   ```
   This script uses SQL files `allowance.sql` to set up the necessary database tables.

3. **Build the Docker Image**:
   Build a Docker image for the application:
   ```
   docker build -t tipsukanya-tax .
   ```

#### Running the Application

You have two options to run the service: using Docker or directly on your PC:

- **Using Docker**:
  Use the following command to run the service in Docker, ensuring all environment variables are set correctly:
  ```
  docker run -e DATABASE_URL='host=host.docker.internal port=5432 user=postgres password=postgres dbname=ktaxes sslmode=disable' -e PORT=8080 -e ADMIN_USERNAME=adminTax -e ADMIN_PASSWORD=admin! tipsukanya-tax
  ```

- **Directly on PC**:
  To run directly on your PC, you can use Go to execute the main application file. This method requires setting
  environment variables either in your system or within your IDE:
  ```
  go run main.go
  ```

#### Ready to Use

After following these steps, your application should be up and running and ready to use. Ensure that you have set all
the required environment variables before starting the service, especially when running directly on your PC.

---

