CREATE DATABASE todolist_rpl;

CREATE TABLE to_do_list(
  ID SERIAL PRIMARY KEY,
  Nama VARCHAR(200) NOT NULL,
  Aktifitas VARCHAR(200) NOT NULL,
  Mulai TIMESTAMP NOT NULL, 
  Selesai TIMESTAMP NOT NULL,  
  IsDone boolean NOT NULL
);

DROP TABLE to_do_list;
DROP DATABASE todolist_rpl;