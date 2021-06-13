# Building a realtime message streaming system using REDIS Pub/Sub.

How do we build a message streaming system which is very highly scalable, where users can join a stream and able to recieve message in realtime. 

For simpilicty we are not considering having historical messages and persistance of message for later use. Which are very simple to add if required

## Requirements
1. Users can join a a stream and recieve messages.
2. Users can be distrbuted across different servers across different regions.
3. Users can send messages if they want.

## Implementation details

1. Our backend servers will connect to redis cluster and use the pub sub feature of redis. Redis cluster supports high performance and scalability of upto 1000 nodes. More on https://redis.io/topics/cluster-spec

2. Backend servers will keep a persistant HTTP connection with the clients and send new messages vis SSE. We are not going ahead with websockets right now for simplicity reasons.

3. Users who subscribe to a certain chanell will recieve realtime messages.