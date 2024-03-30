select 'hello world'

select DISTINCT original_language_id  from film

SELECT film_id as id, title as film_name from film

INSERT INTO film (title, description, release_year, language_id, rental_duration, rental_rate, length, replacement_cost, rating, special_features) 
VALUES ('The Dark Knight', 'Batman fights the Joker', 2008, 1, 3, 4.99, 152, 19.99, 'PG-13', 'Trailers'),
       ('The Dark Knight Rises', 'Batman fights Bane', 2012, 1, 3, 4.99, 165, 19.99, 'PG-13', 'Trailers'),
       ('The Dark Knight Returns', 'Batman fights Superman', 2016, 1, 3, 4.99, 152, 19.99, 'PG-13', 'Trailers');
       
INSERT INTO film
VALUES (default, 'The Dark Knight', 'Batman fights the Joker', 2008, 1, NULL, 3, 4.99, 152, 19.99, 'PG-13', 'Trailers', default);      

SELECT DISTINCT rating, release_year  FROM film WHERE rating = 'PG-13';

SELECT title, (length/60) as hour, (length/(24*60)) as day FROM film;

SELECT title, ROUND(length/60) FROM film;

SELECT * FROM film WHERE rating = 'PG-13' or release_year = 2006;

SELECT * FROM film WHERE NOT rating = 'PG-13';
SELECT * FROM film WHERE rating != 'PG-13';
SELECT * FROM film WHERE rating <> 'PG-13';

SELECT * FROM film WHERE rating NOT IN ('PG-13', 'R');
