# go-test-skeleton

A small Go service skeleton, used as the starting point for the Codurance Go
exercise.

**You have 30 minutes.** Two coding tasks and a discussion, time-boxed below.
The boxes are there so you can pace yourself and so we assess everyone on the
same footing — they are not a race.

Task 2 has the largest box because it is the one we are most interested in.
Getting that right and leaving the stretch work untouched is a good outcome,
not a shortfall.

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

## Task 1 — Add a v2

Serve `v2` alongside `v1`, both from the same process at the same time, with
`v1` unchanged.

`v2` needs one endpoint: a status endpoint like `v1`'s, except that its
response body must include the API version as well as the status. Keep it to
that one endpoint — we are not assessing API design, but we are interested in
how you handle two versions that are *nearly* the same.

## Task 2 — Shut down gracefully

This is the task we are most interested in, and it has the largest time box
for that reason. It is as much about goroutines, channels and `context` as it
is about HTTP, so take the time rather than rushing on to task 3.

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

## Task 3 — Discussion

Down tools and talk us through two things. Bullets in a scratch file are fine
if you would rather write than talk.

**Testing.** What would you test about the work you just did, what would you
deliberately *not* bother testing, and why? We are more interested in this
than in test code written against the clock.

**The skeleton itself.** You have now read most of this codebase. Name two or
three things you would change about it — design, structure, error handling,
testability, production-readiness, anything — and why. Be blunt. There is
plenty here worth criticising, and we would much rather hear it from you than
not.

**Stretch, only if the time is genuinely there:** write one of the tests you
described. `api/routes/v1_test.go` is a working example to adapt. Do not start
it at the expense of task 2 — an unfinished shutdown with a neat test beside
it is the worse outcome.

## Before you finish

```sh
make check     # gofmt + go vet + race tests
```

## What we are looking for

- Idiomatic, readable Go. Three clear things beat six clever ones.
- Errors that say what went wrong and where.
- A clear sense of what is worth testing, and what isn't.
- Judgement about scope. If you run out of time, say what you would do next
  and why — an honest `TODO` beats a rushed implementation, and we would much
  rather discuss it than watch you rush.

Thinking out loud is welcome throughout and counts for as much as the code.
Standard library, docs and tooling are all fair game.
