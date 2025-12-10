
### How to run

#### 1. Initialze mongo cluster

First, we need to run a mongo instance with the docker compose file and then connect to a any node (the file runs 3 shards)

start the cluster:

```bash
docker compose up -d
```

then, connect to the container
```bash
docker exec -it mongo1 mongosh
```

or just run 

```bash
docker exec -it mongo1 mongo
```

once inside the container we need to initiate

```bash
rs.initiate({
  _id: "rs0",
  members: [
    { _id: 0, host: "mongo1:27017" },
    { _id: 1, host: "mongo2:27017" },
    { _id: 2, host: "mongo3:27017" }
  ]
})
```

check status:

```bash
rs.status()
```