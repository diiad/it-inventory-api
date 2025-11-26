package app

//                                 Table "public.model"
// Column |         Type          | Collation | Nullable |           Default
//--------+-----------------------+-----------+----------+------------------------------
// id     | integer               |           | not null | generated always as identity
// name   | character varying(50) |           | not null |
//--------+-----------------------+-----------+----------+------------------------------
//Indexes:
//    "model_pkey" PRIMARY KEY, btree (id)
//Referenced by:
//    TABLE "machine" CONSTRAINT "machine_id_model_fkey" FOREIGN KEY (id_model) REFERENCES model(id)
