## Table of contents

- [General info](#general-info)
- [Technologies](#technologies)
- [Setup](#setup)

# General info

This is an arbitrage betting application that calculates arbitrage opportunities and renders them to the client

# Technologies

The client is made using NextJS with TypeScript. Material UI is the component library of choice.\
The server is made using Gin Gonic Framework. Gorm is the orm of choice for a PostGreSQL database

# Setup

Clone this repo

```
git clone https://github.com/robinmuhia/arbitrageBetting.git

```

## Client folder

Create an env file in the client folder:

```
#for your backend port
NEXT_PUBLIC_BACKEND_URL=http://localhost:8000

```

To run the client folder, install it locally using npm:

```
$ cd client
$ npm install
$ npm run dev
```

## Server folder

Get an odds api key from [Sport Odds API](https://the-odds-api.com/)\
Create an env file in the client folder and:

```
PORT=8000
DB_CONNECTION="host={db host} user={db user} password={user password} dbname={db name} port={db port 5432 is default} sslmode=disable"
JWT_SECRET={random long string with letters and numbers}
ODDS_API_KEY={odds api key}

```

To run the server folder, install it locally using go:

```
$ cd server
$ go mod tidy
$ compileDaemon -command="./arbitrageBackend"

```
