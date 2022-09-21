CREATE TABLE IF NOT EXISTS public.operations_types (
	operation_type_id serial NOT NULL,
	description varchar NULL,
	CONSTRAINT operations_types_pk PRIMARY KEY (operation_type_id)
);
