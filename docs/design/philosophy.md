

## Objective

This document explains the rationale behind some of the philosophy and design choices underpinning Trillian Tessera.


## Simplicity

Tessera is intended to be: simple to use, adopt, and maintain; and cheaper/easier to operate than Trillian v1.

There are many tensions and trade-offs here, and while there is no guarantee that a single "right answer"
exists, we are shooting for a MVP, and must hold ourselves accountable whenever we're adding cost, complexity,
or [speculative abstractions](https://100go.co/#interface-pollution-5) - _"is the driver for this something
we *really need now*?", or otherwise restricting our ability to make large internal changes in the future.


## Multi-implementation storage

Each storage implementation for Trillian Tessera is independently implemented, and takes the most "native"
approach for the infrastructure it targets.

Trillian v1 defined `LogStorage` and embedded `TreeStorage` interfaces which all storage implementations had
to implement. These interfaces were created early, reflected implementation details of a small sampling of largely
similar storage implementations, and consequently turned out not to be a clean abstraction of what was _actually_
desired by higher levels in the stack. In turn, this made it hard to: 

1. Support non-single-domain/non-transactional storage implementations, and 
2. Refactor storage internals to improve performance.

With Trillian Tessera, we are learning from these mistakes, and acknowledging that:

1. The different storage implementations we are building now, and those which will come in the future, have their
   own unique properties which stem from the infrastructure they're built upon - e.g. _some_ infrastructure offers
   rich transactional semantics over multiple entities, others offer only check-and-set semantics.
2. We don't _necessarily_ need to use the more expensive transactional storage to serve reads.
3. Prematurely binding different storage implementations together (e.g. through inappropriate code reuse,
   interfaces, structures, etc.) which _appear_ similar today can lead to headaches later if we find we need to
   make structural changes.

For at least the early versions of Tessera, it is an explicit non-goal to try to reuse code between storage
implementations. Attempting to do this so early in the project lifecycle opens us up to the same pitfalls described
above, and any perceived benefits from this approach are unlikely to be worth the risk; storage implementations are
expected to be relatively small in terms of LoC and complexity.


## Asynchronous integration in storage implementation

In Trillian v1, the only supported mechanism for adding entries to a log was via a fully decoupled queue: the
caller requesting the addition of the entry was given nothing more than a timestamp and a promise that the entry
would be integrated at some point (note that 24h is the CT _policy_, but there's no specific parameter or deadline
in Trillian itself - it's _"as soon as possible"_).

With Trillian Tessera, we're tightening the storage contract up so that calls to add entries to the log will
return with a durably assigned sequence number, or an error.

It's not a requirement that the Merkle tree has already been extended to cryptographically commit to the new leaf
by the time the call to add returns, although it _is_ expected that this process will take place within a short
window (e.g. seconds).

This API represents a reasonable set of tradeoffs:

1. Keeping sequencing and integration separate enables:
    1. Storage to be implemented in the way which works best given the particular constraints of that
       infrastructure
    2. Higher write-throughput
        * E.g. A Bucket-type storage typically has roundtrip read/write latency which is far higher than DBMS.
          From our experiments, typically, a transactional DBMS is used for coordination and sequencing, and the
          slower, cheaper, bucket storage is used for serving the tree.
 
          Sequencing, which has typically been done within the DBMS only, is fast. Integration, however, must
          update the buckets with new tree structure, leaves, checkpoint, etc., and is by far the slower
          operation of the two.

          Coupling &lt;sequence>-&lt;integrate> within a call to add entries to the log (even if batched via a
          local pool in a single frontend), requires blocking updates to the tree for the long-pole integration
          duration.

          Allowing &lt;sequence> operations to happen asynchronously to &lt;integration> enables sequencing to
          proceed from multiple frontend pools concurrently with any integration operation (at least until some
          back-pressure is applied), which, in turn, allows the &lt;integration> operation to potentially
          amortise the long-pole cost over a larger number of sequenced entries.
2. Limiting the intended window between &lt;sequence> and &lt;integrate> operations to a low
   single-digit-seconds target enables a synchronous add API to be constructed at the layer above (i.e within
   the "personality" that is built using Trillian Tessera).

   This approach enables:
     1. synchronous "personalities" to benefit from improved write-throughput (compared with a naive 
        &lt;sequence>-&lt;integrate> approach) at the cost of a small increase of latency
     2. Other "personalities" are not forced to pay the cost of the synchronous operation they do not require.

Back-pressure for `Writer` requests to the storage implementation will be important to maintain the very short
window between sequence numbers being assigned to entries, and those entries being committed to by the Merkle
tree.


## Resilience and availability

Storage implementations should be implemented in such a way that multiple instances of a Tessera-backed
personality can *safely* be run concurrently for any given log storage instance. The manner in which this is
achieved is left for each storage implementation to decide, allowing for the simplest/most
infrastructure-native mechanism to be used in each case.

Having this property offers two benefits:

*   Tessera-backed logs can offer comparable availability as a similar log being managed by Trillian v1 on the
    same infrastructure would have.
*   Safety guard rails are in place against "silly" mistakes, such as "&lt;up>&lt;up>&lt;enter>" and
    _copy-n-paste_ errors, resulting in accidentally launching multiple instances pointing at the same storage
    configuration.

