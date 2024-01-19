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

DROP TABLE if exists workouts;
create table workouts (
    id INT unsigned not null AUTO_INCREMENT,
    user_id VARCHAR(128) not null,
    workout_date varchar(128) not null,
    PRIMARY KEY (id)
);

--Randomly Assign id to a workout, wokrout will contain USER ID to show
-- who is doing workout, and date for tracking purposes
-- Further Sets will all include this workout ID to associate workouts
DROP TABLE if exists exercises;
create table exercises (
    id INT unsigned not null AUTO_INCREMENT,
    workout_id INT unsigned not null,
    workout_type varchar(128),
    PRIMARY KEY (id)
);

DROP TABLE if exists exercise_sets;
create table exercise_sets (
    id INT unsigned not null AUTO_INCREMENT,
    exercise_id INT unsigned not null,
    weight_amt INT unsigned NOT NULL,
    reps INT unsigned NOT NULL
    PRIMARY KEY (id)
);
INSERT INTO users
(id, first_name, last_name, height,
    feetOrInches, userWeight, lbOrKg
)
VALUES
('nkim256', 'Nathan', 'Kim', 72, 0, 190, 0),
('nickdraggy', 'Nicholas', 'Kim', 69, 0, 165, 0);


insert into workouts
(id, user_id, workout_date)
VALUES
(0, "nkim256", "1-1-24");


insert into exercises
(id, workout_id, workout_type)
VALUES
(0, "nkim256_1", "Bench Press");


INSERT INTO exercise_sets
(id, exercise_id, weight_amt, reps)
VALUES
(0, "123", 225, 5);
