-- create_table.sql
CREATE TABLE conta (
    id serial PRIMARY KEY,
    nome text NOT NULL,
    saldo double precision NOT NULL default 0
);

CREATE TABLE transacao (
    id serial PRIMARY KEY,
    tipo text NOT NULL,
    estado text default 'activo'
);

CREATE TABLE deposito (
    id serial PRIMARY KEY,
    idconta integer,
    idtransacao integer,
    montante double precision NOT NULL,
    FOREIGN KEY (idconta) REFERENCES conta(id),
    FOREIGN KEY (idtransacao) REFERENCES transacao(id)
);

CREATE TABLE levantamento (
    id serial PRIMARY KEY,
    idconta integer,
    idtransacao integer,
    montante double precision NOT NULL,
    FOREIGN KEY (idconta) REFERENCES conta(id),
    FOREIGN KEY (idtransacao) REFERENCES transacao(id)
);


CREATE TABLE Transferencia (
    id serial PRIMARY KEY,
    idcontaemissora integer,
    idcontareceptora integer,
    idtransacao integer,
    montante double precision NOT NULL,
    FOREIGN KEY (idcontaemissora) REFERENCES conta(id),
    FOREIGN KEY (idcontareceptora) REFERENCES conta(id),
    FOREIGN KEY (idtransacao) REFERENCES transacao(id)
);


-- insert_data.sql

INSERT INTO conta (nome)
VALUES ('restaurante');

INSERT INTO conta (nome)
VALUES ('empregado');

INSERT INTO conta (nome)
VALUES ('cliente');
