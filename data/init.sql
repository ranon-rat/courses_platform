-- sqlite
-- el admin puede agregar a profesores y a nuevos admins
-- pero debe de agregarse manualmente 
-- con esto tambien se puede guardar los comentarios
CREATE TABLE publications(
    ID INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    mineature TEXT NOT NULL,--url de un servicio de hosting 
    content TEXT NOT NULL,--con esto solo guardare el contenido en texto
    author INTEGER NOT NULL,-- solo usare el id del usuario que publico esto para poder cargar algunas cosas,
    -- en si solo es importante ver lo que son las publicaciones asi que con esto solo se debe de checar 
    -- las publicaciones que tengan el id de un admin
    topic VARCHAR(64),--no croe usar mar de esta cantidad de caracteres , de hecho me parece demasiado pero puede que tenga la necesidad
    datePublication integer,--tal vez lo vaya a cambiar
    Introduction VARCHAR(255)
);


CREATE TABLE users(
    ID iNTEGER PRIMARY KEY,
    privileges INT NOT NULL,--1:admin,2:teacher,3:pupil
    username VARCHAR(64) NOT NULL UNIQUE,--es innecesario tener mas
    pass VARCHAR(64) NOT NULL,--sha256(password+token) 
    token INTEGER NOT NULL , --tiempo de registro+numero aleatorio // solo se puede usar para el momento de entrar al a cuenta, no es algo que sirva de mucho
    email VARCHAR(64) NOT NULL UNIQUE, --deberia de encriptar esto pero posiblemente en caso de que se les
    ssid VARCHAR(64) --sha256(password+email+id+token+unix-time)
);