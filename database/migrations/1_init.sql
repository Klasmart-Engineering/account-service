BEGIN;

CREATE TABLE IF NOT EXISTS account (
    id uuid NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS android_group (
    "id" uuid NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    "account_id" uuid NOT NULL REFERENCES account(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS android (
    "id" uuid NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    "android_group_id" uuid NOT NULL REFERENCES android_group(id) ON DELETE CASCADE
);

COMMIT;
