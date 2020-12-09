# MyCart

A simle CLI app using go and cobra

---

## Steps to run the project

- Clone the repository to you Go Workspace

  ```
  git clone https://github.com/nasir-rabbani/go-cli-app.git
  ```

- Navigate inside `go-cli-app` directory
  ```cmd
  cd go-cli-app
  ```
- Run the below command to install the application
  ```go
  go install
  ```

## Using the Application

After installing the application will be accessible using `mycart` command.

```cmd
E:\Go Workspace\src\go-cli-app>mycart
CLI (Command Line Interface) based E-commerce app.

Usage:
  mycart [command]

Available Commands:
  categories  Displays all Categories
  help        Help about any command
  products    Displays all Products
  register    To register a User or Admin

Flags:
      --config string   config file (default is $HOME/.mycart.yaml)
  -h, --help            help for mycart
  -t, --toggle          Help message for toggle

Use "mycart [command] --help" for more information about a command.

E:\Go Workspace\src\go-cli-app>
```

## Examples of operations

1. To Register Admin (run `mycart register -h` for help menu)

   ```c
   >mycart register -a --name Admin --password Admin

   User Registered with userID :: 5
   ```

1. To Register normal User (run `mycart register -h` for help menu)

   ```c
   >mycart register --name User --password User

   User Registered with userID :: 5
   ```

1. To add a category

   ```c
   >mycart categories addCategory --name Interior
   Category Added with categoryID :: 6
   ```

1. To add a product

   ```c
   >mycart products addProduct --name Nokia-1200 --price 2500 --categoryID 1
   ```

1. To view categories

   ```
   >mycart categories
   List of available categories
   -----------------------------
   Category ID : 1
   Category name : Electronics
   -----------------------------
   Category ID : 2
   Category name : Gadgets
   -----------------------------
   Category ID : 3
   Category name : apparels
   -----------------------------
   Category ID : 4
   Category name : apparels
   -----------------------------
   Category ID : 5
   Category name : furniture
   -----------------------------
   Category ID : 6
   Category name : Interior
   -----------------------------
   ```

1. To view products

   ```
   >mycart products
   List of products
   -----------------------------
   Product ID : 1
   Product name : Nokia-1100
   Product Price : 1500.00
   Category ID : 1
   -----------------------------
   Product ID : 2
   Product name : Nokia-1200
   Product Price : 2500.00
   Category ID : 1
   -----------------------------
   Product ID : 3
   Product name : Nokia-1200
   Product Price : 2500.00
   Category ID : 1
   -----------------------------

   ```
