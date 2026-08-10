CREATE TABLE policies (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    quote_id BIGINT UNIQUE NOT NULL
        REFERENCES quotes(id),

    vehicle_id BIGINT NOT NULL
        REFERENCES vehicles(id),

    starts_at TIMESTAMPTZ NOT NULL,
    ends_at TIMESTAMPTZ NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);