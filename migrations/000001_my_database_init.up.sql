CREATE TABLE tags(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    tag_name VARCHAR NOT NULL UNIQUE,
    created TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE tasks(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    task_name VARCHAR NOT NULL UNIQUE,
    task_description VARCHAR,
    created TIMESTAMP NOT NULL DEFAULT NOW(),
    tag_id INT,
    CONSTRAINT fk_tags
      FOREIGN KEY(tag_id) 
	  REFERENCES tags(id)
);
