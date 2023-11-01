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


INSERT INTO transacao (tipo)
VALUES ('deposito');

INSERT INTO transacao (tipo)
VALUES ('transferencia'); 
INSERT INTO transacao (tipo)
VALUES ('levantamento'); 

INSERT INTO transacao (tipo)
VALUES ('deposito');

INSERT INTO deposito (idconta, idtransacao, montante)
VALUES (1, 1, 300.0); 

UPDATE conta
              SET  saldo=saldo+300.0
              WHERE id=1;

INSERT INTO Transferencia (idcontaemissora, idcontareceptora, idtransacao, montante)
VALUES (1, 2, 2, 200.0); 
UPDATE conta
              SET  saldo=saldo+200.0
              WHERE id=2;

UPDATE conta
              SET  saldo=saldo-200.0
              WHERE id=1;

INSERT INTO deposito (idconta, idtransacao, montante)
VALUES (3, 4, 700.0); 

UPDATE conta
              SET  saldo=saldo+700.0
              WHERE id=3;

INSERT INTO levantamento (idconta, idtransacao, montante)
VALUES (3, 3, 50.0); 

UPDATE conta
              SET  saldo=saldo-50.0
              WHERE id=2;