CREATE TABLE users (
id serial,
first_name text NOT NULL,
last_name text NOT NULL,
age integer,
country text NOT NULL,
CONSTRAINT users_pk PRIMARY KEY (id)
);

CREATE TABLE items (
id serial,
user_id integer NOT NULL,
descr text NOT NULL,
position integer,
year integer NOT NULL,
CONSTRAINT results_pk PRIMARY KEY (id),
CONSTRAINT fk_results_runner_id FOREIGN KEY (user_id)
REFERENCES users (id) MATCH SIMPLE
ON UPDATE NO ACTION
ON DELETE NO ACTION
);



