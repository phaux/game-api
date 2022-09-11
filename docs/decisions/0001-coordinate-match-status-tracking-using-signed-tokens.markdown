---
status: proposed
date: 2022-09-11
deciders: @szabba @phaux
consulted: none
informed: none
---

# Coordinate match status tracking using signed tokens

## Context and Problem Statement

When players are matched to play in a ranked match, they must then play a match on some server (which needs to be selected) and the match result needs to be reported so that their skill estimates and/or rankings can be updated appropriately.

## Decision Drivers

* Modularity is preferable.
  We want different people develop different parts of the system.

* Low number of cross-service interactions.
  Obviously some are necessary, but each additional interface and/or data format we need to specify, implement and integrat requires the effort of at least two people.

* Low number of database interactions.
  Whether we pay for capacity or per-operation - they increase costs.

## Considered Options

* Coordinate across several services with synchronous interactions.

* Coordinate with a monolith and a single database.

* Coordinate with signed tokens certifying past events.

## Decision Outcome

We elected to use signed tokens to certify that particular events in the ranked match flow have already occurred.

We identified at least the following services to be implemented:

* A `matchmaker` receives requests for matches and sends back the same match ID for requests that are matched.
  The match ID is a signed token with:
  * the `matchmaker` as the issuer,
  * the `serverlist` as the audience,
  * the participant it was returned to as the subject,
  * a globally unique match ID (`match_id`),
  * details of the match configuration (`config`),
  * expected participants (`participants`),
  * a short-lived expiration (on the order of minutes).
* A `serverlist` maintains a list of available servers.

  Given a match ID it can produce a match allocation.
  The allocation is a signed token with:
  * the `serverlist` as the issuer,
  * the particular game server as the audience,
  * the participant it was returned to as the subject,
  * a globally unique match ID (`match_id`),
  * details of the match configuration (`config`),
  * expected participants (`participants`),
  * information on how to connect to the selected server (`conn`),
  * a short-lived expiration (on the order of minutes),
    * another match can be assigned to the server if it does not report starting the allocated match before this,
    * the `serverlist` needs to maintain this information in stable storage.

  Additionally, the allocated server receives a match report token.
  This is a signed token with:
  * the `serverlist` as the issuer,
  * the `matchmaker` as the audience,
  * a globally unique match ID (`match_id`),
  * details of the match configuration (`config`),
  * expected participants (`participants`),
  * an expiration time sufficient for a match to complete.

* Once the match starts, the server confirms that to the `serverlist`.

* After a completed match, a game server uses the match report token to authenticates a request to the matchmaker that reports the results of the match.

From the POV of a client the happy path flow is:

* request a match from the `matchmaker`,
* using a match ID, request an allocation from the `serverlist`,
* using the match allocation, connect to the server using the,
* play through the match.

From the POV of the server:

* it registers itself with the `serverlist`,
* it awaits for a match report token - receiving one confirms to it that it should run a match,
* it awaits for a sufficient number of participants to connect,
* it reports to the `serverlist` that the match has started,
* once the match is complete it reports the results to the `matchmaker`.

## More Information

### Matchmaking separate from ranking/skill-eval

It might be possible for matchmaking and ranking/skill-evaluation to be separated into different services.

### Many games

We start out with the tokens being specific to a particular game.

In the future it will be possible to add optional field identifying what game the token is meant for.
(The default given a missing/empty field would be to assume the original game the system was first built for.)

This could either be a top-level field or part of the match configuration.

### Match configuration opaqueness

It is easier to support many games if the system can mostly treat them as opaque.
On the other hand some features might require to use information from the config:

* For some games, matchmaking requests might specify criteria that an acceptable configuration would match.
  This would make it necessary for the `matchmaker` to understand the match configurations.
* Not all available servers might be able to run a match with a given configuration.
  This would make it necessary for the `serverlist` to understand the match configurations.
* For some games the results of matches with different configurations might have different effets on skill evaluation.
  This would make it necessary for the `matchmaker` (or a separate future service) to understand the match configurations.

### Token claim JSON schemas

TBD
