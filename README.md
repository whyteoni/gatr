# Gatr

A log aggregator inspired by Will Farrel's Gator in the hit movie _The Other Guys_.

## Prerequisites

### Go(lang)

If you're a real peacock, you need to install the Go toolchain so you can fly free! While there are several ways to install and manage the Go toolchain, the official guidance is [here](https://go.dev/doc/install).

### PostgreSQL

This application uses PostgreSQL to store information. Again, many paths, but [here](https://www.postgresql.org/download/)'s the official documentation.

## Installation

1. Clone the repo: 

    ```sh
    git clone https://github.com/whyteoni/gatr
    ```

1. Use the Go toolchain to build and install from the source code:

    ```sh
    cd gatr
    go install .
    ```

1. Next we need to setup our database. For this example we are using the user `postgres` and the password `postgres`. This will invoke `db_setup.sql` which will create the `gatr` database and all tables, columns, and constraints needed for the application.

    ```sh
    psql "postgres://postgres:postgres@localhost:5432" -f db_setup.sql
    ``` 

    
1. Add the database name and optional `?sslmode=disable` to the URI and we have our _Connection String_. 
    
    ```sh
    "postgres://postgres:postgres@localhost:5432/gatr?sslmode=disable"
    ```

1. Create a local config file for Gatr.

    ```sh
    echo "{ \"db_url\": \"<YOUR CONNECTION STRING>\" }" > ~/.gatrconfig.json
    ```

    The result should look like (using the example from the previous step):

    ```json
    { "db_url": "postgres://postgres:postgres@localhost:5432/gatr?sslmode=disable" }
    ```

1. You should now be able to run `gatr help` and get back a help message explaining the different commands available. To do anything we'll first need to register a user and add some RSS feeds to track.

    1. Register a user: `gatr register <user name>`
    2. Track a feed: `gatr addfeed "<feed name>" "<feed URL>"`
    3. List users: `gatr users` You will see `(current)` next to your user
    4. List tracked feeds: `gatr feeds`

## Collecting posts

As an aggregator, Gatr needs to be checking in the background for new posts. While you can, and I recommend, setting up a service for this, that is beyond the scope of this README. The easy way to get started would be to run a backgrounded service.

On Linux: 

```sh
nohup gatr agg <frequency> >> gatr.log &
```

> The **frequency** is how often to check a feed. If you have a frequency of 1 hour and you have 6 feeds, each feed will only be checked 4 times a day. It is important to not overly tax RSS feeds though, so you should not set the frequency to something very short (like `5s`).

Once the aggregation service is running in the background you can use `Gatr` as normal to register new users, add new feeds, and browse recent posts.

## Basic usage

To get a full list of commands simply run `gatr help`. Here we'll just cover the most essential commands.

* `gatr login <user>`: Login as a user. If your user does not exist you will recieve an error, but you can...
* `gatr register <user>`: Register a new user. Once registered, they automatically become the logged in user.
* `gatr feeds`: List currently tracked feeds.
* `gatr following`: List feeds the current user is following.
* `gatr addfeed "<feed name>" "<feed url>"`: Add a new feed to follow. The current user automatically starts following.
* `gatr browse [<limit>]`: Look at the most recent posts from feeds you are following. Defaults to the 2 most recent posts, but you can optionally supply a number of posts to look at.

## The monster at the end of this README

This is part of a project course through [Boot.Dev](https://boot.dev), a wonderful online learning platform that teaches students the skills they need to become backend developers. This repo is from their `Go Developer` track. For anyone interested in learning more I cannot recommend them enough.

Now, this also means this was all part of a school assignment. If you somehow stumbled onto this repo and thought to yourself, "This guy gets it! Gatr is everything I ever wanted in an RSS aggregator, if not everything I ever wanted out of life in general," then I am sorry to dissappoint you but there will be no updates, bug patches, feature enhancements, or AI integrations in the future. I hit submit on this assignment, and then I clear my cache to avoid a memory leak.
