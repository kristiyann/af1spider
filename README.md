# <p align="center"> AF1 SPIDER <p align="center">

Service to periodically check whether a certain Nike (at the moment) product is available in the desired size. <br>

## Use:
1. Make sure you have Go & Git installed <br>
https://go.dev/ https://git-scm.com/downloads <br>
2. Clone the repository locally <br>
3. Install dependencies
```bash
go mod tidy
```
4. Create an .env file in the root of the project and populate the following variables <br>
```bash
FROM_EMAIL=
```
*Email from which the notification will be sent <br>
```bash
FROM_PASS=
```
*Special password for that goes along with your FROM_EMAIL value. This is not your usual Gmail password. If you are using Gmail, you generate it from the google accounts page: 'myaccount.google.com/u/4/apppasswords' <br>
```bash
TO_EMAIL=
```
*Email the notification will be sent to <br>
```bash
PRODUCT_LINK=
```
*Actual page url for a nike product from the official nike site (www.nike.com) <br>
```bash
SIZE_US=
```
*Size you're looking for in US terms <br>

5. From there on you can just run 
```bash
go run .
```
to start the service or 
```bash
go build
```
to create an executable to run in one click <br>
*Make sure that you have the .env file in the same directory as the executable otherwise the service will crash
