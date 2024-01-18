CREATE TABLE test_timestamp_table (
    "id" INTEGER PRIMARY KEY,
    "name" TEXT NOT NULL,
    "comment" TEXT,
    "create_time" TIMESTAMP NOT NULL,
    "update_time" TIMESTAMP NOT NULL,
    UNIQUE(name)
) ;
