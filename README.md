# contract-poller

## Overview
> Service that sources and stores contract data

## Repository Structure
Our main service is the poller which serves as how we find and load contracts and their corresponding ABIs into the database. The poller is a go service that runs on a cron job and is responsible for finding new contracts and loading them into the database. 

### Env variables
There are a couple of environment variables that need to be set in order to run the poller locally. Create a .env file in the root of the project and add the following variables

```  
ETHERSCAN_API_KEY={INSERT YOUR ETHERSCAN API KEY HERE}
POLYGONSCAN_API_KEY={INSERT YOUR POLYGONSCAN API KEY HERE}
```

### Setting up the DB locally

There are a few commands you need to run to set up the DB.

First run the infra locally

```
make infra-up
```

Then make sure you have psql installed

```
brew install postgresql
```

Next connect to the postgres server running locally and create the database

```
psql -h localhost -p 5432 -U postgres 
enter password: postgres 
CREATE DATABASE db;
\c db
```

Now the database will be created on your local machine.