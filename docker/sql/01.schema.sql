
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE public.ab_atena (
    id uuid DEFAULT public.gen_random_uuid() NOT NULL PRIMARY KEY,
    names text[] NOT NULL,
    postcode text,
    address text,
    group_name text,
    kind integer,
    is_sender boolean DEFAULT false NOT NULL,
    deleted boolean DEFAULT false NOT NULL
);
