# Load balancing

Round-robin — default; weight задаёт долю; least_conn полезен при разной длительности запросов; sticky привязывает session, но ухудшает перераспределение. Instance должен быть stateless, session — внешней.
