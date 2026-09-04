# go-test-skeleton

A small Go service skeleton, used as the starting point for the Codurance Go
exercise.

**You have 30 minutes.** Four tasks, time-boxed below. The boxes are there so
you can pace yourself and so we assess everyone on the same footing — they are
not a race. Finishing three tasks well beats four in a hurry.

## Before the clock starts

Requires Go 1.26+. This part is not part of your 30 minutes.

```sh
make test                     # should pass
make run                      # starts the server on :3000
curl -i localhost:3000/v1     # => 200 {"status":"ok"}
```

## Orientation (2 min)

```
main.go               entrypoint
cmd/serve.go          wires the router into the server
api/                  server lifecycle (New, Serve, Stop)
api/routes/           API versions: routes.go mounts them, v1.go implements one
```

Start with `api/routes/routes.go` and `api/routes/v1.go` — together they are
about 70 lines. The existing tests in `api/` and `api/routes/` show the house
style for testing this codebase.

## Task 1 — Add a v2 (5 min)

Serve `v2` alongside `v1`, both from the same process at the same time, with
`v1` unchanged.

`v2` needs one endpoint: a status endpoint like `v1`'s, except that its
response body must include the API version as well as the status. Keep it to
that one endpoint — we are not assessing API design, but we are interested in
how you handle two versions that are *nearly* the same.

## Task 2 — Shut down gracefully (12 min)

Ctrl-C currently kills the process instantly and drops in-flight requests.
Make it shut down cleanly:

- `SIGINT` and `SIGTERM` begin the shutdown.
- New connections stop being accepted; in-flight requests finish, under a
  bounded timeout.
- A clean shutdown exits zero; a real failure exits non-zero.

`TODO(candidate)` in `api/rest.go` marks the spot. `api.RestService.Stop` is
already written and should not need changing.

**Out of scope** — don't spend time here: making the timeout configurable (a
constant is fine), draining anything other than HTTP requests, and
connection-level tracking.

## Task 3 — Test it (5 min)

Add a test covering the v2 endpoint from task 1.

Then tell us, in a sentence or a comment, **what you deliberately did not
test, and why.** We are more interested in that answer than in the number of
tests you got written.

## Task 4 — Tell us what you'd change (4 min, no code)

You have now read most of this codebase. Name **two or three things you would
change** about the skeleton itself — design, structure, error handling,
testability, production-readiness, anything — and why. Talking them through is
fine; a few bullets in a scratch file is fine too.

Be blunt. There is plenty here worth criticising, and we would rather hear it
from you than not.

## Before you finish (2 min)

```sh
make check     # gofmt + go vet + race tests
```

## What we are looking for

- Idiomatic, readable Go. Three clear things beat six clever ones.
- Errors that say what went wrong and where.
- A test that would catch a real regression.
- Judgement about scope. If you run out of time, say what you would do next
  and why — an honest `TODO` beats a rushed implementation, and we would much
  rather discuss it than watch you rush.

Thinking out loud is welcome throughout and counts for as much as the code.
Standard library, docs and tooling are all fair game.
