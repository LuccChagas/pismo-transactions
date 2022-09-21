CREATE TABLE IF NOT EXISTS public.transactions (
	transactions_id serial NOT NULL,
	account_id int4  NULL,
	operation_type_id int4  NULL,
	amount decimal(5,2) NULL,
	event_date varchar NULL,
	CONSTRAINT transactions_pk PRIMARY KEY (transactions_id),
	CONSTRAINT transactions_account_fk FOREIGN KEY (account_id) REFERENCES public.accounts(account_id),
	CONSTRAINT transactions_operation_type_fk FOREIGN KEY (operation_type_id) REFERENCES public.operations_types(operation_type_id)
);
