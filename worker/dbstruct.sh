# sudo -u postgres psql -c "DROP DATABASE IF EXISTS octopoddb;"
sudo -u postgres psql -c "CREATE DATABASE octopoddb;"

sudo -u postgres psql -d octopoddb -c "
DROP TABLE IF EXISTS countries;
CREATE TABLE countries (
    id SERIAL PRIMARY KEY,
    name_official_fr VARCHAR(255),
    flag_url VARCHAR(255),
    population INT
);
ALTER TABLE countries
ADD CONSTRAINT unique_name UNIQUE (name_official_fr)
"
