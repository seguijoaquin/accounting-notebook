# Accounting Notebook

![GitHub go.mod Go version](https://img.shields.io/badge/Golang-v1.13-blue.svg)
![technology React](https://img.shields.io/badge/React-v16.13.1-blue.svg)

## Responsibilities

1. Create credit and debit transactions.

2. List all transactions.

3. Find a transaction by id.

4. Get balance status of the account.

Information about the API endpoints may be found [here](https://github.com/seguijoaquin/accounting-notebook/tree/master/server#commiting-a-transaction)


## Running the app

To run the app open a terminal and located in the project's root directory execute

```bash
$ npm install
$ npm install client/
$ npm run start
```

### Launching details

In case of needing to launch the server and the client separately, here can be found info about each one:
 - [Client launching details](https://github.com/seguijoaquin/accounting-notebook/tree/master/client#available-scripts)
 - [Server launching details](https://github.com/seguijoaquin/accounting-notebook/tree/master/server#launching-the-app)

The command shown executes the following sentences, which may be run separately if preferred, as detailed below.

_Install the project dependencies to launch client and server simultaneusly_
```bash
$ npm install
```

_Install the application frontend dependencies_
```bash
$ npm install client/
```

_Build the backend docker image and name it "api-img"_
```bash
$ docker build -t api-img ./server
```

_Launch a container based on the image "api-img" previusly created_
```bash
$ docker run -p 8080:8080 api-img
```

_Launch a container based on the image "api-img" previusly created, mapping it's internal port 8080 with the host port 8080_
```bash
$ docker run -p 8080:8080 api-img
```

_Move to the client directory and launch the React App_
```bash
$ cd client/ && npm run start
```