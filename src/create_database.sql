DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id VARCHAR(128) NOT NULL,
    first_name VARCHAR(128) NOT NULL,
    last_name VARCHAR(128) NOT NULL,
    height INT unsigned NOT NULL,
    feetOrInches INT unsigned NOT NULL,
    userWeight INT UNSIGNED NOT NUll,
    lbOrKg INT unsigned NOT NULL,
    PRIMARY KEY (id)
);
INSERT INTO users
(id, first_name, last_name, height,
    feetOrInches, userWeight, lbOrKg
)
VALUES
('nkim256', 'Nathan', 'Kim', 72, 0, 190, 0),
('nickdraggy', 'Nicholas', 'Kim', 69, 0, 165, 0);

DROP TABLE if exists workouts;
create table workouts (
    id varchar(128) not null,
    user_id varchar(128) not null,
    workout_date varchar(128) not null,
    workout_name varchar(128)
    PRIMARY KEY (id)
)

insert into workouts
(id, user_id, workout_date, workout_name)
VALUES
("nkim256_1", "nkim256", "1-1-24", "chest")
("")