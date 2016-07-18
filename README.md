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
    * Could generate false positives
    * Could not list all members of a dataset

## DGA Algorithms

  *  [Locky v3](https://github.com/baderj/domain_generation_algorithms/blob/master/locky/dgav3.py) (20 seeds)

## Todo

  - [ ] Redis storage
  - [ ] gRPC microservice
  - [ ] HTTP Proxy

## Usage

### Generate the filter

```
#> dgabloom build -b 365 -a 10
```
Generates bloom filter since 365 in the past and 10 days in the future.

### Check if a domain is known as a DGA domain

```
#> dgabloom check zenithar.org
zenithar.org,false
```

### Multiple checks

```
#> dgabloom check ruoectomayrj.pw hoflddlod.org zenithar.org
ruoectomayrj.pw,true
hoflddlod.org,true
zenithar.org,false
```
