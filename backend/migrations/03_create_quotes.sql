CREATE TABLE quotes (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    vehicle_id BIGINT NOT NULL
        REFERENCES vehicles(id),
    driver_age SMALLINT NOT NULL,
    years_licensed SMALLINT NOT NULL,
    penalty_points SMALLINT NOT NULL,
    duration_minutes INTEGER NOT NULL,
    price_pence INTEGER NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);