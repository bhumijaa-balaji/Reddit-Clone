# Reddit Clone with Engine and Simulator

The goal of this project is to implement a Reddit-like backend engine along with a client simulator using Go, Proto Actor model and Protobuf for message serialization. It replicates key functionalities of Reddit, such as account registration, subreddit management, posting, commenting, voting, and messaging.

## Prerequisites

#### Go - https://go.dev/doc/install
#### Proto Actor - https://github.com/asynkron/protoactor-go/tree/dev
#### Protobuf - http://google.golang.org/protobuf

## Installation

* Make sure Install Go is installed and added to your Environment PATH variable
* Unzip the project directory into desired location on local drive
* Make sure Proto Actor and Protobuf are installed and are available in go.mod file


## Instruction section

### Running the program
- Run the following command to compile the .proto file in the root directory
   ```bash 
   protoc --go_out=. messages/*.proto
- Based on the directory path given for go_package in protos.proto file, protos.pb.go file will be generated automatically post the compilation step
- Change to server directory and run the reddit engine:
    ```bash
    cd server
    go run main.go
- Similarly change to the client directory and run the simulator
    ```bash
    cd client
    go run main.go
    
    ```
    Note: 
        In order to run simulations for variable number of users, alter the numUsers variable in the client (main.go file)

## What is working

* We have spawned Reddit engine actor which processes the message requests sent by the users and sends appropriate response messages
* We have spawned Reddit simulator actor with thounsands of user actors which perform variety of operations as mentioned in the features
* The system also tracks performance and simulates connection management by having users disconnect and reconnect periodically. The simulation culminates in a shutdown message sent to the Reddit engine once all operations are completed.
* Once the Reddit engine is shut down, the simulator also shuts down

### Features

#### Engine Functionality
1. **Account Management**
   - Register new accounts.
2. **Subreddit Management**
   - Create subreddits.
   - Join and leave subreddits.
3. **Posts**
   - Post simple text-based messages in subreddits.
   - Support re-posting of popular messages.
4. **Comments**
   - Add comments to posts.
   - Support hierarchical comment threads (nested comments).
5. **Voting & Karma**
   - Upvote and downvote posts and comments.
   - Compute and track user Karma scores.
6. **Feeds**
   - Retrieve a feed of posts based on subreddit memberships.
7. **Direct Messaging**
   - Send and receive direct messages.
   - Reply to direct messages.

#### Client Simulator
1. Simulates thousands of users interacting with the engine.
2. Models live connection and disconnection periods for users.
3. Implements a **Zipf distribution** for subreddit memberships:
   - A few highly active subreddits with many members.
   - Increased post activity in subreddits with more members.
   - Includes re-posting of popular messages.
4. Measures parallelism by computing user+sys time and real time for analysing system performance under simulated load.
5. Measures total and average time taken to complete the simulation


## Performance

### Measured Metrics
1. **Largest Number of Users Simulated** -- 50000
2. **Parallelism Observed**
    
    |     numUsers    |    realTime     |  CPUTime (user+sys)   |     Ratio     |
    |-----------------|-----------------|-----------------------|---------------|
    |       1000      |     27.07       |        30.74          |     1.13      |
    |       2000      |     26.92       |        32.84          |     1.21      |
    |       5000      |     31.44       |        39.92          |     1.26      |
    |       10000     |     35.15       |        46.39          |     1.32      |
    |       50000     |     82.36       |        128.48         |     1.56      |

3. **Time**
    
    |     numUsers    |    Total Time(s)|      Average Time(ms) |
    |-----------------|-----------------|-----------------------|
    |       1000      |     20.74       |        20.74          |
    |       2000      |     20.93       |        10.46          |
    |       5000      |     25.35       |        5.07           |
    |       10000     |     29.17       |        2.91           |
    |       50000     |     76.35       |        1.52           |
