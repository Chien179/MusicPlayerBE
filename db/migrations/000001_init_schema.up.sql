CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,  
  "password" varchar NOT NULL,
  "image" varchar NOT NULL,
  "role" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "songs" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "singer" varchar NOT NULL,
  "image" varchar NOT NULL,
  "file_url" varchar NOT NULL,
  "duration" BIGSERIAL NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "genres" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "image" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "playlists" (
  "id" BIGSERIAL PRIMARY KEY,
  "users_id" bigserial NOT NULL,
  "name" varchar NOT NULL,
  "image" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "songs_genres" (
  "id" BIGSERIAL PRIMARY KEY,
  "songs_id" bigserial NOT NULL,
  "genres_id" bigserial NOT NULL
);

CREATE TABLE "playlists_songs" (
  "id" BIGSERIAL PRIMARY KEY,
  "songs_id" bigserial NOT NULL,
  "playlists_id" bigserial NOT NULL
);

CREATE INDEX ON "playlists" ("users_id");

CREATE INDEX ON "songs_genres" ("songs_id");

CREATE INDEX ON "songs_genres" ("genres_id");

CREATE INDEX ON "songs_genres" ("songs_id", "genres_id");

CREATE INDEX ON "playlists_songs" ("songs_id");

CREATE INDEX ON "playlists_songs" ("playlists_id");

CREATE INDEX ON "playlists_songs" ("songs_id", "playlists_id");

ALTER TABLE "playlists" ADD FOREIGN KEY ("users_id") REFERENCES "users" ("id");

ALTER TABLE "songs_genres" ADD FOREIGN KEY ("songs_id") REFERENCES "songs" ("id");

ALTER TABLE "songs_genres" ADD FOREIGN KEY ("genres_id") REFERENCES "genres" ("id");

ALTER TABLE "playlists_songs" ADD FOREIGN KEY ("playlists_id") REFERENCES "playlists" ("id") ON DELETE CASCADE;

ALTER TABLE "playlists_songs" ADD FOREIGN KEY ("songs_id") REFERENCES "songs" ("id") ON DELETE CASCADE;

ALTER TABLE "playlists_songs" ADD CONSTRAINT "song_playlist_key" UNIQUE ("playlists_id", "songs_id");

ALTER TABLE "songs_genres" ADD CONSTRAINT "song_genre_key" UNIQUE ("genres_id", "songs_id");

