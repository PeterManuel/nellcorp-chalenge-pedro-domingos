-- create_table.sql
CREATE TABLE conta (
    id serial PRIMARY KEY,
    nome text NOT NULL,
    saldo double precision NOT NULL default 0
);

CREATE TABLE transacao (
    id serial PRIMARY KEY,
    tipo text NOT NULL,
    tempo timestamp
);

CREATE TABLE deposito (
    id serial PRIMARY KEY,
    idconta integer,
    idtransacao integer,
    FOREIGN KEY (idconta) REFERENCES conta(id),
    FOREIGN KEY (idtransacao) REFERENCES transacao(id)
);

-- insert_data.sql
INSERT INTO conta (nome)
VALUES ('Pedro');

INSERT INTO conta (nome)
VALUES ('Juventina');

INSERT INTO transacao (tipo)
VALUES ('deposito');

INSERT INTO deposito (idconta,idtransacao)
VALUES (1,1);

