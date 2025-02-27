txsort
======

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)

Package txsort provides the transaction sorting according to [BIP 69](https://github.com/bitcoin/bips/blob/master/bip-0069.mediawiki).

BIP 69 defines a standard lexicographical sort order of transaction inputs and
outputs.  This is useful to standardize transactions for faster multi-party
agreement as well as preventing information leaks in a single-party use case.

The BIP goes into more detail, but for a quick and simplistic overview, the
order for inputs is defined as first sorting on the previous output hash and
then on the index as a tie breaker.  The order for outputs is defined as first
sorting on the amount and then on the raw public key script bytes as a tie
breaker.

A comprehensive suite of tests is provided to ensure proper functionality.