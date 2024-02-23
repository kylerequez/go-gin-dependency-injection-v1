# go-gin-dependency-injection-v1
A sample Go application using Gin framework that demonstrates dependency injection in a RESTful API that provides the following:
1. CRUD Functionality
2. Login Functionality (Email only)
3. JWT Authentication
4. Dependency Injection

# How to run the application?
1. Install Go
2. Install Postgres and create a Database called "go-gin-dependency-injection-v1"
3. Install Postman, or use any application that allows you to make HTTP Requests
4. Download/Clone the repository
5. Create a .env file inside the folder and add the following details:
![image](https://github.com/kylerequez/go-gin-dependency-injection-v1/assets/82488140/b881917d-bf6c-43ed-94bc-d6d2126152ac)
6. Run the "go mod download" command through the terminal, this will install all the necessary package (such as Gin)
8. Run the "CompileDaemon -command="./go-gin-dependency-injection-v1"" on the terminal, this will start the application
9. Using Postman, go to the register URL (localhost:3001/api/v1/auth/register) and submit the following request:
![image](https://github.com/kylerequez/go-gin-dependency-injection-v1/assets/82488140/ed04313c-ab53-4319-9312-5ff6a7fb61fe)
10. After this, try accessing the login URL (localhost:3001/api/v1/auth/login) and submit the following request:
![image](https://github.com/kylerequez/go-gin-dependency-injection-v1/assets/82488140/23c36f8a-3a1f-493d-8f96-4ff8bf40f4b0)
![image](https://github.com/kylerequez/go-gin-dependency-injection-v1/assets/82488140/9309c53a-7e26-4ecf-97de-53e5031d76e9)
11. Now test the following URLs with the following requests:
![image](https://github.com/kylerequez/go-gin-dependency-injection-v1/assets/82488140/36133f66-1b96-4b66-bed9-84a4f9ab6280)
![image](https://github.com/kylerequez/go-gin-dependency-injection-v1/assets/82488140/f883f3dc-1c75-4e76-b5ac-ba3caa591bed)
![image](https://github.com/kylerequez/go-gin-dependency-injection-v1/assets/82488140/67eb8d54-0c59-4ce3-94e0-b438020f4d7b)
![image](https://github.com/kylerequez/go-gin-dependency-injection-v1/assets/82488140/b4850a21-7345-4ba8-b7f3-31fab0bd7516)
![image](https://github.com/kylerequez/go-gin-dependency-injection-v1/assets/82488140/077c1530-6893-407c-8dd5-b97898f835f3)
![image](https://github.com/kylerequez/go-gin-dependency-injection-v1/assets/82488140/58adb318-2da5-4a25-aad1-2ac18ea1000f)
12. Using Postgres Admin, change the Authority column of your account to anything that is not "NORMAL_USER" and try all the URLs from step 11.
