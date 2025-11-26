package app

//  Column  |         Type          | Collation | Nullable |           Default
//----------+-----------------------+-----------+----------+------------------------------
// id       | integer               |           | not null | generated always as identity
// city     | character varying(50) |           |          | NULL::character varying
// name     | character varying(50) |           |          | NULL::character varying
// building | character varying(50) |           |          | NULL::character varying
//----------+-----------------------+-----------+----------+------------------------------
//Indexes:
//    "site_pkey" PRIMARY KEY, btree (id)
//Referenced by:
//    TABLE "soclage_event" CONSTRAINT "soclage_event_id_site_fkey" FOREIGN KEY (id_site) REFERENCES site(id)
