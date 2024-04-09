CREATE TABLE "logs" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "data" JSONB NOT NULL
);
