-- Copyright 2024 Google LLC
--
-- Licensed under the Apache License, Version 2.0 (the "License");
-- you may not use this file except in compliance with the License.
-- You may obtain a copy of the License at
--
--     http://www.apache.org/licenses/LICENSE-2.0
--
-- Unless required by applicable law or agreed to in writing, software
-- distributed under the License is distributed on an "AS IS" BASIS,
-- WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
-- See the License for the specific language governing permissions and
-- limitations under the License.

-- MySQL version of the Trillian Tessera database schema.

-- "Checkpoint" table stores a single row that records the current state of the log. It is updated after every sequence and integration.
CREATE TABLE IF NOT EXISTS `Checkpoint` (
  -- id is expected to be always 0 to maintain a maximum of a single row.
  `id`    INT UNSIGNED NOT NULL,
  -- note is the text signed by one or more keys in the checkpoint format. See https://c2sp.org/tlog-checkpoint and https://c2sp.org/signed-note.
  `note`  MEDIUMBLOB NOT NULL,
  PRIMARY KEY(`id`)
);

-- "Subtree" table is an internal tile consisting of hashes. There is one row for each internal tile, and this is updated until it is completed, at which point it is immutable.
CREATE TABLE IF NOT EXISTS `Subtree` (
  -- level is the level of the tile.
  `level` INT UNSIGNED NOT NULL,
  -- index is the index of the tile.
  `index` BIGINT UNSIGNED NOT NULL,
  -- nodes stores the hashes of the leaves.
  `nodes` MEDIUMBLOB NOT NULL,
  PRIMARY KEY(`level`, `index`)
);

-- "TiledLeaves" table stores the data committed to by the leaves of the tree. Follows the same evolution as Subtree.
CREATE TABLE IF NOT EXISTS `TiledLeaves` (
  `tile_index` BIGINT UNSIGNED NOT NULL,
  `data`       LONGBLOB NOT NULL,
  PRIMARY KEY(`tile_index`)
);
