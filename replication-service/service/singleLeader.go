package service

type Leader struct {
	followers []string
	data      map[string]string
}

func (l *Leader) SetData(key, value string) {
	l.data[key] = value

	//TODO: this should be a bulk write. either time based or log based
	for _, follower := range l.followers {
		replicate(follower, key, value)
	}
}

func replicate(follower, key, value string) {
	panic("unimplemented")
}

//TODO: Write a function to read from the replica and that should be load balanced
//TODO: Write a load balancer
//TODO: Health check to see if the leader/replicas are healthy
//TODO: A fall back machanism when the leader
//TODO: Make a statement based replication system and a trigger based replication
//TODO: Write a system for eventual replication
//TODO: Write a system to find the lag in the system
