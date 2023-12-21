# BOOKSTORE MANAGEMENT APIS

1. Database - Mysql
2. GORM
3. Json marshall, unmarshall
4. Project Structure
5. Gorilla Mux

## Project Structure
### CMD
MAIN.GO
### PKG
config -> APP.GO
controllers -> BOOK_CONTROLLER
models -> BOOK.GO
routes -> BOOKSTORE_ROUTES
utils -> UTILS.GO

## ROUTES
controller funcs
GETBOOKS <- /book/ <- GET
CREATE BOOKS <- /book/ <- POST
CET BOOK BY ID <- /book/{bookID} <- GET
UPDATE BOOKS <- /book/{bookID} <- PUT
DELETE BOOKS <- /book/{bookID} <- DELETE