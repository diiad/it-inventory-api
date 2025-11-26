-- ============================================================================
--   ALLER DANS db/db.go et changer la config pour faire fonctionner la bdd en
--   local !
--
--   DROP DES OBJETS
-- ============================================================================

DROP INDEX IF EXISTS idx_soclage_event_site;
DROP INDEX IF EXISTS idx_soclage_event_machine;
DROP INDEX IF EXISTS idx_soclage_event_collab;

DROP TABLE IF EXISTS soclage_event;
DROP TABLE IF EXISTS machine;
DROP TABLE IF EXISTS site;
DROP TABLE IF EXISTS model;
DROP TABLE IF EXISTS collab;

DROP TYPE IF EXISTS category_type;


-- ============================================================================
--   CREATION DU TYPE ENUM
-- ============================================================================

CREATE TYPE category_type AS ENUM ('PC Fixe', 'PC Portable');


-- ============================================================================
--   TABLE collab
-- ============================================================================

CREATE TABLE collab
(
    id         INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    last_name  VARCHAR(50) NOT NULL,
    first_name VARCHAR(50) NOT NULL
);


-- ============================================================================
--   TABLE model
-- ============================================================================

CREATE TABLE model
(
    id       INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name     VARCHAR(50) NOT NULL,
    brand    VARCHAR(100),
    category category_type
);


-- ============================================================================
--   TABLE machine
-- ============================================================================

CREATE TABLE machine
(
    id            INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    company       VARCHAR(50) NOT NULL,
    mac_address   VARCHAR(30) UNIQUE,
    serial_number VARCHAR(30) UNIQUE NOT NULL,
    id_model      INTEGER NOT NULL REFERENCES model(id)
);


-- ============================================================================
--   TABLE site
-- ============================================================================

CREATE TABLE site
(
    id       INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    city     VARCHAR(50),
    name     VARCHAR(50),
    building VARCHAR(50)
);


-- ============================================================================
--   TABLE soclage_event
-- ============================================================================

CREATE TABLE soclage_event
(
    id            INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    comment       VARCHAR(255),
    date_soclage  TIMESTAMP,
    num_inventory VARCHAR(30),
    mail          VARCHAR(255),
    floor         INTEGER,
    desk          VARCHAR(10),
    version       VARCHAR(10),
    ip_address    VARCHAR(45),
    project       VARCHAR(50),
    id_collab     INTEGER NOT NULL REFERENCES collab(id),
    id_machine    INTEGER REFERENCES machine(id),
    id_site       INTEGER REFERENCES site(id)
);


-- ============================================================================
--   INDEXES
-- ============================================================================

CREATE INDEX idx_soclage_event_collab ON soclage_event (id_collab);
CREATE INDEX idx_soclage_event_machine ON soclage_event (id_machine);
CREATE INDEX idx_soclage_event_site ON soclage_event (id_site);