## <p align="center"> AF1 SPIDER <p align="center">

Service to periodically check whether a certain Nike (at the moment) product is available in the desired size. <br>

<h3> Use: </h3>
1) Make sure you have golang & git installed <br>
https://go.dev/ https://git-scm.com/downloads <br>
2) Clone the repository locally <br>
3) Run 'go mod tidy' in the root of the project to install dependencies <br>
4) Create an .env file in the root of the project and populate the following variables: <br>
FROM_EMAIL= <br>
*Email from which the notification will be sent <br>
FROM_PASS= <br>
*Special password for the FROM_EMAIL. This is not your usual gmail email password. If you are using Gmail, you generate it from the google accounts page: 'myaccount.google.com/u/4/apppasswords' <br>
TO_EMAIL= <br>
*Email the notification will be sent to <br>
PRODUCT_LINK= <br>
*Actual page url for a nike product from the official nike site (www.nike.com) <br>
SIZE_US= <br>
*Size you're looking for in US terms <br>
5) From there on you can just run 'go run .' to start the service or 'go build' to create an executable to run in one click <br>
*Make sure that you have the .env file in the same directory as the executable otherwise the service will crash
