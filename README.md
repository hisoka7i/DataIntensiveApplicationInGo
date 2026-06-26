# DataIntensiveApplicationInGo
I am trying to learn computer science concepts by implementing them. This Repo which contain different protocols and different concepts of computer science, This will also contain different algorithm and there implementation.

# Thrift Binary protocol
- Problem statement.    
    - Why do we need it? The current issue is with the serializers. Different programming languages have there own serializers. Java objects are serialized using java serializers their structure is different. Python does this in its own way. There serialization formats are different. So communication over the network is difficult in that case.
    To solve the problem thirft was introduced. A binary protocol allowing data to be serialised in a standard way which can be effectively used by different languages.
    Same kind of work is done by protocol buffer as well. 

- Approach 
    - Thrift protocol is made of 2 protocols: Binary + compact protocol
    Binary protocol sends id and data type separately while compact protocol compresses them together.
    - Forward compatibility: Do not change the tag of data, since tag and data type is mapped to the data. Now all the new additional fields need to be not required, as it can break backwards compatibility.
    - Backward compatibility: Old code will simply ignore the new data that is added to the file. But new data that is being added should be not required otherwise it will break the old code in validation.
    remove a field is same as add a new field. Required field can not be deleted and tag can not be reused.


# Apache Avro
    - Why?
        - problem with thift and buffer protocol is their tag. It is hard to maintain where in avro there is no need for tag matching and all those things as whenever there is change in schema it we can simply generate a new shcema read file. 
    - How? 
        - It works using a reader schema and a writer schema. Reader schema should match the writer schema. 
        - There is a schema reader and a schema writer. Who's work is to solve descripencies between the schema. 
        - Backwards compatibility is maintained using the older read schema and newwer write schema. 


# Modes of data flow 
    - Some of the modes for dataflow are Dataflow through database, rest and message passing

# Chapter 5: Distributed Data
    - Issues with the distributed data
    
# Chapter 6: Replication
    - Popular algorithm for replicating changes between nodes: single-leader, multi-leader, and leaderless replication.
     -single leader: Write comes to a single point when that write is successful it is replicated to other databases i.e., followers. Issue with this approach is every this relies on a single leader if that leader is down then the whole system is down. There can be a outage until next leader is elected. 
        - Another reason is that sync replication, if any of the node is down then the whole replication is not working. 
        - Other issue is with async approch that there is no way to recover data
        - Another issue is with data that client might not be getting the latest data. Then in this case client need to maintain a timestamp and compare with the db timestamp. There is one issue with time as well because different time zones will have different time data. Metadata is also needed to syncronised for it.
    -multi leader selection: Does not make sense in a single data center. 
        Setting up new followers: 
            a. Take snapshots of the leader at regular interval. 
            b. Follower takes the latest snapshot
            c. Follower after adding the snapshot data ask the leader for the data that was added after that snapshot. 
            d. When all that remaining data is added to the follower then we say that the follower is caught up.
        (note: This snapshot is associated with exact leader position. This position has various different names like log sequence number(postgresSQL) and binlof coordinates(MySql))
        Handling nodes outages:
            We are trying to reboot the node without and downtime. 
            Follower failure: This is quite easy in the sense that it can recover from the log. Take last snapshot and then ask the additional data changes from the leader.
            Leader failure: This is tricky: 
                a. How to identify if the leader is down? Most system use timeout, but timeout should not be less. 
                b. How to select the next leader? This is done through election or perviously set controller.
                c. How to followers identify its new leader? 
                d. What happens when the prev leader comes back online? Some times both leaders starts to accept the write request which leads to data curruption. Some times prev leader has its counter which is slightly less than the newly selected leader and this leads to data inaccuracy.
                    d.1: We need to should down one of the leader to solve this issue, but some times both the leaders are shut down.


# TODO: Write sstable and LSM trees. Write ahead log.