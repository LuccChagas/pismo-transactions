CREATE TABLE IF NOT EXISTS public.accounts (
	account_id serial NOT NULL,
	document_number varchar NULL,
	CONSTRAINT accounts_pk PRIMARY KEY (account_id)
);
