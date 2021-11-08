CREATE DATABASE bank;

\c bank

CREATE TABLE IF NOT EXISTS public.t_bank (
	id SERIAL NOT NULL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT 'NOW()',
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
	"name" VARCHAR(255) NOT NULL UNIQUE,
	code INTEGER NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS public.t_agency (
	id SERIAL NOT NULL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT 'NOW()',
	updated_at INTEGER,
	deleted_at TIMESTAMP,
	bank_id BIGINT NOT NULL,
	code INTEGER NOT NULL UNIQUE,
	main_agency BOOLEAN NOT NULL DEFAULT 'false',
	zip_code VARCHAR(10) NOT NULL,
	public_place VARCHAR(255) NOT NULL,
	number VARCHAR(45) NOT NULL,
	complement VARCHAR(255),
	district VARCHAR(45) NOT NULL,
	city VARCHAR(45) NOT NULL,
	state VARCHAR(2) NOT NULL,
	country VARCHAR(45) NOT NULL
);
ALTER TABLE public.t_agency ADD CONSTRAINT fk_bank_id FOREIGN KEY (bank_id) REFERENCES public.t_bank(id);

CREATE TABLE IF NOT EXISTS public.t_account_owner(
	id SERIAL NOT NULL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT 'NOW()',
	updated_at INTEGER,
	deleted_at TIMESTAMP,
	"name" VARCHAR(255) NOT NULL,
	"document" VARCHAR(11) NOT NULL,
	birth_date TIMESTAMP NOT NULL,
	father_name VARCHAR(255) NOT NULL,
	mother_name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS public.t_account_owner_address(
	id SERIAL NOT NULL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT 'NOW()',
	updated_at INTEGER,
	deleted_at TIMESTAMP,
	zip_code VARCHAR(10) NOT NULL,
	public_place VARCHAR(255) NOT NULL,
	number VARCHAR(45) NOT NULL,
	complement VARCHAR(255),
	district VARCHAR(45) NOT NULL,
	city VARCHAR(45) NOT NULL,
	state VARCHAR(2) NOT NULL,
	country VARCHAR(45) NOT NULL,
	account_owner_id BIGINT NOT NULL
);
ALTER TABLE public.t_account_owner_address ADD CONSTRAINT fk_account_owner FOREIGN KEY (account_owner_id) REFERENCES public.t_account_owner(id);

--DROP TABLE IF EXISTS t_account;
CREATE TABLE IF NOT EXISTS public.t_account(
	id SERIAL NOT NULL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT 'NOW()',
	updated_at INTEGER,
	deleted_at TIMESTAMP,
	"number" SERIAL NOT NULL,
	"verifying_digit" INTEGER NOT NULL DEFAULT 1,
	agency_id BIGINT NOT NULL,
	account_owner_id BIGINT NOT NULL,
	"password" VARCHAR(255) NOT NULL
);
ALTER TABLE public.t_account ADD CONSTRAINT fk_account_owner_id FOREIGN KEY (account_owner_id) REFERENCES public.t_account_owner(id);
ALTER TABLE public.t_account ADD CONSTRAINT fk_agency_id FOREIGN KEY (agency_id) REFERENCES public.t_agency(id);

--DROP TABLE IF EXISTS t_account_detail;
CREATE TABLE IF NOT EXISTS public.t_account_detail(
	id SERIAL NOT NULL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT 'NOW()',
	updated_at INTEGER,
	deleted_at TIMESTAMP,
	blocked BOOLEAN NOT NULL DEFAULT FALSE,
	balance float NOT NULL DEFAULT 0,
	account_id BIGINT NOT NULL
);
ALTER TABLE public.t_account_detail ADD CONSTRAINT fk_account_id FOREIGN KEY (account_id) REFERENCES public.t_account(id);

--DROP TABLE IF EXISTS t_account_detail;
CREATE TABLE IF NOT EXISTS public.t_account_transaction_log(
	id SERIAL NOT NULL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT 'NOW()',
	from_account_id BIGINT NOT NULL,
	from_account_old_balance FLOAT NOT NULL,
	from_account_new_balance FLOAT NOT NULL,
	to_account_id BIGINT NOT NULL,
	to_account_old_balance FLOAT NOT NULL,
	to_account_new_balance FLOAT NOT NULL,
	transaction_value FLOAT NOT NULL
);
ALTER TABLE public.t_account_transaction_log ADD CONSTRAINT fk_from_account_id FOREIGN KEY (from_account_id) REFERENCES public.t_account(id);
ALTER TABLE public.t_account_transaction_log ADD CONSTRAINT fk_to_account_id FOREIGN KEY (to_account_id) REFERENCES public.t_account(id);

INSERT INTO public.t_bank (id,created_at,updated_at,deleted_at,"name",code) VALUES (1,NOW(),NULL,NULL,'inBolso',234);
INSERT INTO public.t_agency (id,created_at,updated_at,deleted_at,bank_id,main_agency,zip_code,public_place,"number",complement,district,city,state,country,code) VALUES
	 (1,'2021-11-05 22:57:19.044',NULL,NULL,1,true,'63020060','Rua Santa Isabel','1631','PRÉDIO COMERCIAL','FRANCISCANOS','JUAZEIRO DO NORTE','CE','BRASIL',1),
	 (2,'2021-11-05 22:57:19.044',NULL,NULL,1,false,'63041155','Rua Profa. Maria Nilde Couto Bem','220','SALA 231','TRIÂNGULO','JUAZEIRO DO NORTE','CE','BRASIL',2);







