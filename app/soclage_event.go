package app

//                                   Table "public.soclage_event"
//    Column     |            Type             | Collation | Nullable |           Default
//---------------+-----------------------------+-----------+----------+------------------------------
// id            | integer                     |           | not null | generated always as identity
// comment       | character varying(255)      |           |          | NULL::character varying
// date_soclage  | timestamp without time zone |           |          |
// num_inventory | character varying(30)       |           |          | NULL::character varying
// mail          | character varying(255)      |           |          | NULL::character varying
// floor         | integer                     |           |          |
// desk          | character varying(10)       |           |          | NULL::character varying
// version       | character varying(10)       |           |          | NULL::character varying
// ip_address    | character varying(45)       |           |          | NULL::character varying
// project       | character varying(50)       |           |          | NULL::character varying
// id_collab     | integer                     |           | not null |
// id_machine    | integer                     |           |          |
// id_site       | integer                     |           |          |
//Indexes:
//    "soclage_event_pkey" PRIMARY KEY, btree (id)
//    "idx_soclage_event_collab" btree (id_collab)
//    "idx_soclage_event_machine" btree (id_machine)
//    "idx_soclage_event_site" btree (id_site)
//Foreign-key constraints:
//    "soclage_event_id_collab_fkey" FOREIGN KEY (id_collab) REFERENCES collab(id)
//    "soclage_event_id_machine_fkey" FOREIGN KEY (id_machine) REFERENCES machine(id)
//    "soclage_event_id_site_fkey" FOREIGN KEY (id_site) REFERENCES site(id)
