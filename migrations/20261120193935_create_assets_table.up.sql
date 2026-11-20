CREATE TABLE assets (
    id TEXT PRIMARY KEY,
    symbol TEXT NOT NULL,
    network TEXT NOT NULL,
    decimals SMALLINT NOT NULL CHECK (decimals BETWEEN 0 AND 36),
    contract TEXT,
    enabled BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX assets_network_contract_uniq 
    ON assets (network, contract)
    WHERE contract IS NOT NULL;

CREATE UNIQUE INDEX assets_network_native_uniq
    ON assets (network)
    WHERE contract IS NULL;