

CREATE TABLE `clientes`
(
    id INT(11) PRIMARY KEY auto_increment,
    nombre varchar(50) NOT NULL,
    apellido_1 varchar(50) NOT NULL,
    apellido_2 varchar(50) NOT NULL,
    edad INT(11) NOT NULL
);

INSERT INTO 
    `clientes` 
    (
       `nombre`,
       `apellido_1`,
       `apellido_2`,
       `edad`
    )
    VALUES (
      'Fernando',
      'Fernán',
      'González',
      23
    );
