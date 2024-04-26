DROP TABLE IF EXISTS tax.allowance;

create table tax.allowance
(
    allowanceid integer not null
        primary key,
    amount      double precision,
    type        varchar(50)
);

alter table allowance
    owner to postgres;

INSERT INTO tax.allowance (allowanceid, amount, type) VALUES (1, 0, 'donation');
INSERT INTO tax.allowance (allowanceid, amount, type) VALUES (3, 0, 'k-receipt');
INSERT INTO tax.allowance (allowanceid, amount, type) VALUES (2, 0, 'Personal');
