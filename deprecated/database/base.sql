CREATE TABLE search_queries (
    id serial primary key,
    query text unique,
	count int
);