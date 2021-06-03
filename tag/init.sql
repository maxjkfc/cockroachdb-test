


CREATE DATABASE IF NOT EXISTS tag;

USE tag;

CREATE TABLE records (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_info JSONB,
  user_name STRING AS (user_info->>'name') STORED,
  user_id STRING AS (user_info->>'id') STORED,
  tags STRING[],
  createdAt TIMESTAMPTZ DEFAULT now(),
  INDEX user_info_2(user_id , user_name , createdAt),
  -- INVERTED INDEX user_info(user),
  INVERTED INDEX tags(tags)
);

