CREATE TABLE movies (
id SERIAL PRIMARY KEY,            -- Identificador único de la película
name VARCHAR(255) NOT NULL,     -- Nombre de la película
description VARCHAR(255) NOt NULL,                 -- Descripción o sinopsis
release_date DATE NOT NULL,           -- Fecha de estreno
runtime INT NOT NULL,             -- Duración en minutos
genre VARCHAR(50) NOT NULL,               -- Género de la película
director VARCHAR(100) NOT NULL            -- Nombre del director
);
