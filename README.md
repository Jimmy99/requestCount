# requestCount
Dockerized multi-replica web request counter

The requirements call for the dockerised web server to generate 3 instances using the docker-compose replicas directive.

In order to synchronise the atomic counter of each instance, the Redis key-value store is used to increment the counter for each web request received by each server instance.

At the time of writing to the redis key-value store a mutex is used to lock the write to the key-value store.

With the exception of the redis key-value store, no additional packages apart from the standard library have been used to minimise dependencies.


