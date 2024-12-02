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

-- "Checkpoint" table stores a single row that records the latest _published_ checkpoint for the log.
-- This is stored separately from the TreeState in order to enable publishing of commitments to updated tree states to happen
-- on an indepentent timeframe to the internal updating of state.
CREATE TABLE IF NOT EXISTS `Checkpoint` (
  -- id is expected to be always 0 to maintain a maximum of a single row.
  `id`    INT UNSIGNED NOT NULL,
  -- note is the text signed by one or more keys in the checkpoint format. See https://c2sp.org/tlog-checkpoint and https://c2sp.org/signed-note.
  `note`  MEDIUMBLOB NOT NULL,
  -- published_at is the millisecond UNIX timestamp of when this row was written.
  `published_at` BIGINT NOT NULL,
  PRIMARY KEY(`id`)
);

-- "TreeState" table stores the current state of the integrated tree.
-- This is not the same thing as a Checkpoint, which is a signed commitment to such a state.
CREATE TABLE IF NOT EXISTS `TreeState` (
  -- id is expected to be always 0 to maintain a maximum of a single row.
  `id`    INT UNSIGNED NOT NULL,
  -- size is the extent of the currently integrated tree.
  `size`  BIGINT UNSIGNED NOT NULL,
  -- root is the root hash of the tree at the size stored in `size`.
  `root`  TINYBLOB NOT NULL,
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
