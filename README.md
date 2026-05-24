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
    