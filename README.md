# DGABloom

This product uses bloom filters to help users to identify if a domain is used as
a malware domain.

## Why using bloom filter ?

Bloom filters are probabilistic space efficient data algorithm, it is
generally used for checking if a given data is present in a dataset.

## UseCase(s)

  * Check a domain using the cli
  * Use the database in a proxy (not developed)

## Pro / Cons

  * Pro :
    * Space efficient (database takes only few MBs)
    * Quick lookup for checking if the domain is not is the set

  * Con :
    * Could generate false negative
    * Could not list all members of a dataset
