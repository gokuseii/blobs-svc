-- +migrate Up

CREATE TABLE blobs (
    id                  character(52)  NOT NULL,
    owner_address       character(64),
    type                integer        NOT NULL,
    value               jsonb          NOT NULL,
    creator_signer_role integer        NOT NULL
);

ALTER TABLE ONLY blobs
    ADD CONSTRAINT blobs_pkey PRIMARY KEY (id);

-- +migrate Down

DROP TABLE blobs cascade;