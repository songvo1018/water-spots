CREATE TABLE public."water-spots"
(
    id SERIAL PRIMARY KEY,
	"identifier" numeric NOT NULL DEFAULT 0,
    "Model" text,
    "LastService" text,
    "Owner" text,
    "IsUnsecured" boolean,
    "OzonPressure" double precision,
    "WaterPressure" double precision
);

ALTER TABLE IF EXISTS public."water-spots"
    OWNER to pguser;