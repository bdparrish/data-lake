SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

DROP SCHEMA IF EXISTS data_lake CASCADE;

CREATE SCHEMA data_lake;

ALTER SCHEMA data_lake OWNER to webserver;
