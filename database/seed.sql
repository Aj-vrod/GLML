CREATE TABLE IF NOT EXISTS movies (
    movie_id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL
);
INSERT INTO movies
VALUES (1, 'D.E.B.S'),
    (2, 'But I am a Cheerleader'),
    (
        3,
        'Imagine me and you'
    ),
    (
        4,
        'PROM'
    ),
    (
        5,
        'Girltrash'
    );
