CREATE DATABASE bank;

\c bank

CREATE OR REPLACE FUNCTION public.tf_utils_setar_data_atualizacao()
 RETURNS trigger
 LANGUAGE plpgsql
AS $function$
DECLARE 
  tabela_raw TEXT; -- Armacena o nome da Tabela que recuperamos com TG_TABLE_NAME sem o "t_" padrão.    
  tabela TEXT; -- Armacenara o nome da Tabela já pronto para o mensagem de Erro.
  temp_mensagem TEXT; -- Para formatar mensagem de erro
BEGIN
  /*
  * A Função atualiza automaticamento o atributo data_atualizacao com o valor de NOW()
  * Alteração só será permitida ao usar um gatilho BEFORE UPDATE / row-level, para validação.
  */
  IF (TG_OP = 'UPDATE') THEN
    -- Setamos a data de atualizacao com os valor de NOW()
    NEW.updated_at = now();
    -- Retornamos o NEW.
    RETURN NEW;
  ELSE 
  
    -- Recuperamos o nome da Tabela e removemos o "t_" padrão.
    tabela_raw := replace(TG_TABLE_NAME, 't_', '');
    -- Removemos os "_" e reemplazamos por espacio. 
    tabela := replace(tabela_raw, '_', ' ');
    -- Formatamos o mensagem
    temp_mensagem := 'Atributo data_atualizacao, atualização NãO autorizada para gatilhos diferentes de UPDATE para ('||tabela||') !!'; 
    -- Retorna error se tentar remover um registro.
    RAISE EXCEPTION feature_not_supported USING HINT = temp_mensagem;

  END IF;
END;

$function$
;

CREATE TABLE IF NOT EXISTS public.t_bank (
	id SERIAL NOT NULL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT 'NOW()',
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
	"name" VARCHAR(255) NOT NULL UNIQUE,
	code INTEGER NOT NULL UNIQUE
);
CREATE TRIGGER t_bank_set_updated_data BEFORE
UPDATE ON public.t_bank FOR EACH ROW EXECUTE FUNCTION tf_utils_setar_data_atualizacao();

CREATE TABLE IF NOT EXISTS public.t_agency (
	id SERIAL NOT NULL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT 'NOW()',
	updated_at TIMESTAMP,
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
CREATE TRIGGER t_agency_set_updated_data BEFORE
UPDATE ON public.t_agency FOR EACH ROW EXECUTE FUNCTION tf_utils_setar_data_atualizacao();

CREATE TABLE IF NOT EXISTS public.t_account_owner(
	id SERIAL NOT NULL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT 'NOW()',
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
	"name" VARCHAR(255) NOT NULL,
	"document" VARCHAR(11) NOT NULL,
	birth_date TIMESTAMP NOT NULL,
	father_name VARCHAR(255) NOT NULL,
	mother_name VARCHAR(255) NOT NULL
);
CREATE TRIGGER t_account_owner_set_updated_data BEFORE
UPDATE ON public.t_account_owner FOR EACH ROW EXECUTE FUNCTION tf_utils_setar_data_atualizacao();

CREATE TABLE IF NOT EXISTS public.t_account_owner_address(
	id SERIAL NOT NULL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT 'NOW()',
	updated_at TIMESTAMP,
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
CREATE TRIGGER t_account_owner_address_set_updated_data BEFORE
UPDATE ON public.t_account_owner_address FOR EACH ROW EXECUTE FUNCTION tf_utils_setar_data_atualizacao();

--DROP TABLE IF EXISTS t_account;
CREATE TABLE IF NOT EXISTS public.t_account(
	id SERIAL NOT NULL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT 'NOW()',
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
	"number" SERIAL NOT NULL,
	"verifying_digit" INTEGER NOT NULL DEFAULT 1,
	agency_id BIGINT NOT NULL,
	account_owner_id BIGINT NOT NULL,
	"password" VARCHAR(255) NOT NULL
);
ALTER TABLE public.t_account ADD CONSTRAINT fk_account_owner_id FOREIGN KEY (account_owner_id) REFERENCES public.t_account_owner(id);
ALTER TABLE public.t_account ADD CONSTRAINT fk_agency_id FOREIGN KEY (agency_id) REFERENCES public.t_agency(id);
CREATE TRIGGER t_account_set_updated_data BEFORE
UPDATE ON public.t_account FOR EACH ROW EXECUTE FUNCTION tf_utils_setar_data_atualizacao();

--DROP TABLE IF EXISTS t_account_detail;
CREATE TABLE IF NOT EXISTS public.t_account_detail(
	id SERIAL NOT NULL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT 'NOW()',
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
	blocked BOOLEAN NOT NULL DEFAULT FALSE,
	balance float NOT NULL DEFAULT 0,
	account_id BIGINT NOT NULL
);
ALTER TABLE public.t_account_detail ADD CONSTRAINT fk_account_id FOREIGN KEY (account_id) REFERENCES public.t_account(id);
CREATE TRIGGER t_account_detail_set_updated_data BEFORE
UPDATE ON public.t_account_detail FOR EACH ROW EXECUTE FUNCTION tf_utils_setar_data_atualizacao();

--DROP TABLE IF EXISTS t_account_detail;
CREATE TABLE IF NOT EXISTS public.t_payment_slip(
	id SERIAL NOT NULL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT 'NOW()',
	header VARCHAR(255) NOT NULL,
	assignor VARCHAR(255) NOT NULL,
	issuance_date TIMESTAMP NOT NULL,
	due_date TIMESTAMP NOT NULL,
	agency_code INTEGER NOT NULL,
	account_code INTEGER NOT NULL,
	verifying_digit INTEGER NOT NULL,
	gross_value FLOAT NOT NULL,
	deduction FLOAT NOT NULL DEFAULT 0,
	discount FLOAT NOT NULL DEFAULT 0,
	penalty FLOAT NOT NULL DEFAULT 0,
	fees FLOAT NOT NULL DEFAULT 0,
	amount_charged FLOAT NOT NULL,
	paying_document VARCHAR(14) NOT NULL,
	paying_name VARCHAR(255) NOT NULL,
	receiver_document VARCHAR(14) NOT NULL,
	receiver_name VARCHAR(255) NOT NULL,
	payment_local VARCHAR(255) NOT NULL,
	digitable_line VARCHAR NOT NULL DEFAULT 'Sem linha digitável',
	instructions VARCHAR,
	msg_1 VARCHAR,
	msg_2 VARCHAR,
	msg_3 VARCHAR
);

INSERT INTO public.t_bank (id,created_at,updated_at,deleted_at,"name",code) VALUES (1,NOW(),NULL,NULL,'inBolso',234);
INSERT INTO public.t_agency (id,created_at,updated_at,deleted_at,bank_id,main_agency,zip_code,public_place,"number",complement,district,city,state,country,code) VALUES
	 (1,'2021-11-05 22:57:19.044',NULL,NULL,1,true,'63020060','Rua Santa Isabel','1631','PRÉDIO COMERCIAL','FRANCISCANOS','JUAZEIRO DO NORTE','CE','BRASIL',1),
	 (2,'2021-11-05 22:57:19.044',NULL,NULL,1,false,'63041155','Rua Profa. Maria Nilde Couto Bem','220','SALA 231','TRIÂNGULO','JUAZEIRO DO NORTE','CE','BRASIL',2);







