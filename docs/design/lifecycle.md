# Log Lifecycle

Log lifecycle is a useful concept for outlining expected modes the log can be in, and state transitions between these states.
The lifecycle states outlined below exist conceptually, but there are no explicit enums or state defined or maintained in code.
The requirements for each state are documented, and the log operator takes responsibility to ensure that they only migrate between these states in the supported directions.

[![](https://mermaid.ink/img/pako:eNqFUD1rAzEM_StGYzkvHV3odGMKodkSdzC2cjE9y4kiF0LIf6_P1_QoFKJJ70s27wo-BwQDWmtLEmVEo1Zxj_7iR7TU6LM4wT66gV3SX8-WVJ0QGb3ETGr1_mJpJndPH0rrV7XBU0HykYa__Joxc0DGMCWatlDN0bOL9Ju7o3-PLvhRsscRZXqzKTNoQv2XJeggIScXQ-3hOpksyAETWjB1DY4_LVi6VZ8rkjcX8mCEC3bAuQwHMHs3nisqx7AUdbdgiJL5bW65ld3B0dE25_QTvH0DWex-jw?type=png)](https://mermaid.live/edit#pako:eNqFUD1rAzEM_StGYzkvHV3odGMKodkSdzC2cjE9y4kiF0LIf6_P1_QoFKJJ70s27wo-BwQDWmtLEmVEo1Zxj_7iR7TU6LM4wT66gV3SX8-WVJ0QGb3ETGr1_mJpJndPH0rrV7XBU0HykYa__Joxc0DGMCWatlDN0bOL9Ju7o3-PLvhRsscRZXqzKTNoQv2XJeggIScXQ-3hOpksyAETWjB1DY4_LVi6VZ8rkjcX8mCEC3bAuQwHMHs3nisqx7AUdbdgiJL5bW65ld3B0dE25_QTvH0DWex-jw)

## Terminology

The definitions below will use terms that we'll define here:
 - sequenced: an entry has been submitted to the log and durably assigned a sequence number, however it may not be integrated
 - integrated: a sequenced entry that has been included in the Merkle tree, and a checkpoint committing to it has been created

## States

### `Sequencing`

This is the "normal" state of most active logs.
The purpose of this state is to allow entries to be sequenced by, and integrated into, the log.

This state can start from an empty tree, or from the `Draining` state, in which case there can be any number of entries already in the tree, but they must all be integrated.
This state is characterized by the writer personality only calling the `Add` method.
The only valid transition outwards is to [`Draining`](#Draining).

### `Preordered`

This state is used for logs where the precise index assigned to each entry in the log is critical.
The most common example of this is when a log is mirroring another log.

This state must start from an empty tree.
It is characterized by the personality that writes only via calls to the `Set` method.
The only valid transition outwards is to [`Draining`](#Draining).

### `Draining`

The purpose of this state is to prevent any new entries being added and to integrate any pending sequenced entries.
This state may be realized as a terminal state, in which case the log is frozen.

This state requires [`Preordered`](#Preordered) or [`Sequencing`](#Sequencing) first.
No calls to the write methods should be made while in this state.
The only valid transitions are to [`Sequencing`](#Sequencing), or [`Deleted`](#Deleted).

### `Deleted`

Each storage implementation will define instructions for deleting the contents of the log when no longer required.
It can only be reached from the [`Draining`](#Draining) state.
The concept of a soft-delete is not supported by Tessera, though the deployer may be able to realize this via their own infrastructure (e.g. by deleting URL mappings to the log handlers).

## Lifecycle in Trillian v1

This lifecycle proposal was inspired by Trillian v1, but simplified as much as possible.
For comparative purposes, the states possible in Trillian are documented below.

Using Trillian log [TreeState](https://github.com/google/trillian/blob/master/trillian.proto#L66) and [TreeType](https://github.com/google/trillian/blob/master/trillian.proto#L92) as inspiration, the largest conceivable lifecycle is:
 - No log / unknown
 - Empty log
 - Pre-ordered
 - Active
 - Draining (this is the only state that allows a back-transition. This can move back to Active)
 - Frozen
 - No log (deleted)

