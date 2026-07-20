# Misha Go Backend Curriculum Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build a self-contained, Russian-language, 12-week Go backend curriculum with runnable labs, a Codex-native tutor/examiner, evidence-based progress tracking, Dockerized dependencies, and a final `ogen` microservice.

**Architecture:** The repository separates teaching content (`modules/`), runnable experiments (`labs/`), the cumulative service (`project/backend-lab/`), Codex behavior (`AGENTS.md` and `.agents/skills/`), and learner state (`progress/`). A repository validator enforces lesson structure and links; language- or service-specific checks compile and exercise runnable artifacts independently.

**Tech Stack:** Markdown, Bash, Go, PostgreSQL, Redis, Kafka-compatible broker, Docker Compose, Nginx, Kubernetes manifests, OpenAPI 3, `ogen`.

## Global Constraints

- All learner-facing prose and Codex prompts are in Russian, with English technical names preserved where conventional.
- Start date is `2026-07-21`; target date is `2026-10-21`; baseline pace is 10–12 hours per week.
- Go runs on the host; PostgreSQL, Redis, and Kafka run through Docker Compose.
- Every required lab provides an exact run command, expected observation, deliberate breakage, repair task, retelling prompt, and acceptance check.
- A topic is a logical bundle: focused theory lives in `modules/`, runnable practice lives in `labs/` with `LAB.md`, `CHECK.md`, and starter code; `DEEP_DIVE.md` exists only for topics that genuinely need an advanced layer.
- Solutions are physically separated from starter tasks and Codex must not reveal them before a real attempt.
- Progress is updated only from evidence: command output, passing tests, working code, or an oral explanation accepted by the examiner.
- Kubernetes stays at interview-ready middle overview level; no local cluster is required.
- No script may delete learner data or install system software without explicit confirmation.
- Use current stable tool versions discovered from primary documentation during implementation; pin versions in `tool-versions.md` and Docker files.

---

## File Map

- `README.md`: entry point, quick start, commands, 12-week overview.
- `AGENTS.md`: global Codex teaching contract and learner profile.
- `.agents/skills/*/SKILL.md`: onboarding, lesson, exam, review, progress, mock interview, and next-task workflows.
- `00-onboarding/*`: environment diagnosis and Go/Docker installation learning path.
- `roadmap/*`: ordered syllabus, calendar, readiness rubric, and interview checklist.
- `modules/<module>/<topic>.md`: focused theory units with links to labs and exams.
- `labs/*`: runnable isolated experiments and deliberately broken variants.
- `exams/*`: question banks, grading rubrics, and cumulative interview scenarios.
- `progress/*`: editable learner state, exam log, weak-topic queue, and evidence ledger.
- `scripts/*`: environment, dependency, seed, verification, and progress validation commands.
- `deploy/*`: Compose, Nginx, and Kubernetes learning artifacts.
- `project/backend-lab/*`: cumulative production-shaped Go service.
- `internal/coursecheck/*`: Go-based structural validator and tests.

---

### Task 1: Repository Contract, Course Validator, and Progress Schema

**Files:**
- Create: `README.md`
- Create: `AGENTS.md`
- Create: `tool-versions.md`
- Create: `go.mod`
- Create: `internal/coursecheck/check.go`
- Create: `internal/coursecheck/check_test.go`
- Create: `cmd/coursecheck/main.go`
- Create: `progress/PROFILE.md`
- Create: `progress/STATUS.md`
- Create: `progress/EVIDENCE.md`
- Create: `progress/EXAMS.md`
- Create: `progress/REVIEW_QUEUE.md`
- Create: `progress/README.md`
- Create: `scripts/check-course.sh`
- Create: `scripts/test-go-labs.sh`

**Interfaces:**
- Consumes: the directory contract in the design specification.
- Produces: `coursecheck.Check(root string) []error`; stable progress headings consumed by all Codex skills.

- [ ] **Step 1: Write structural tests**

Create table-driven Go tests that build temporary valid and invalid courses. Assert that `Check` detects: missing root files, missing required progress headings, broken relative Markdown links, a required lab without `LAB.md`, and starter code placed under a solutions directory.

- [ ] **Step 2: Run the focused test and observe failure**

Run: `go test ./internal/coursecheck -run TestCheck -v`

Expected: compilation failure because `coursecheck.Check` does not exist.

- [ ] **Step 3: Implement the minimal validator**

Implement `Check(root string) []error` with deterministic, sorted diagnostics. It must validate root files, Markdown links, progress schema, lab structure, and solution separation without modifying files.

- [ ] **Step 4: Add the CLI and shell wrapper**

`cmd/coursecheck` prints every diagnostic and exits 1 when any error exists. `scripts/check-course.sh` uses `set -euo pipefail`, resolves the repository root, and runs `go run ./cmd/coursecheck "$root"`.

`scripts/test-go-labs.sh` finds repository-owned `go.mod` files in deterministic order, runs `go test ./...` from each module directory, reports the module being tested, and exits on the first failure. This is the canonical repository-wide Go test command because labs and the cumulative project intentionally use isolated modules.

- [ ] **Step 5: Create the initial learner state and root documentation**

Populate the profile with React experience, backend starting point, dates, pace, preferred energetic tone, and evidence-only grading. `STATUS.md` starts at onboarding with zero fabricated completions. `README.md` lists natural-language Codex commands and the dependency model.

- [ ] **Step 6: Verify**

Run: `go test ./internal/coursecheck -v && bash scripts/check-course.sh && bash scripts/test-go-labs.sh`

Expected: all tests pass and the course checker exits 0.

- [ ] **Step 7: Commit**

```bash
git add README.md AGENTS.md tool-versions.md go.mod internal cmd progress scripts/check-course.sh
git commit -m "feat: establish curriculum contract and progress tracking"
```

### Task 2: Codex-Native Tutor Skills and Onboarding

**Files:**
- Create: `.agents/skills/misha-onboarding/SKILL.md`
- Create: `.agents/skills/misha-lesson/SKILL.md`
- Create: `.agents/skills/misha-exam/SKILL.md`
- Create: `.agents/skills/misha-review/SKILL.md`
- Create: `.agents/skills/misha-progress/SKILL.md`
- Create: `.agents/skills/misha-mock-interview/SKILL.md`
- Create: `.agents/skills/misha-next/SKILL.md`
- Create: `00-onboarding/README.md`
- Create: `00-onboarding/01-diagnose.md`
- Create: `00-onboarding/02-install-go.md`
- Create: `00-onboarding/03-install-docker.md`
- Create: `00-onboarding/04-first-module.md`
- Create: `scripts/doctor.sh`
- Create: `scripts/check-skills.sh`
- Test: `internal/coursecheck/check_test.go`

**Interfaces:**
- Consumes: progress headings and evidence rules from Task 1.
- Produces: seven discoverable repository-local Codex skills and a read-only `scripts/doctor.sh` report.

- [ ] **Step 1: Extend validator tests for skill manifests**

Assert each expected skill exists, has YAML frontmatter with `name` and `description`, uses one-question-at-a-time interaction, and references the progress files by exact path.

- [ ] **Step 2: Run tests and observe failure**

Run: `go test ./internal/coursecheck -run Skill -v`

Expected: failure listing seven missing skills.

- [ ] **Step 3: Implement the skills**

Write full workflows with trigger phrases, entry checks, dialogue rules, evidence gates, exact files they may update, refusal to reveal solutions early, concrete praise, deadline-aware status, and a final handoff format. The examiner must ask exactly one question per turn and record wrong-model/correct-model/follow-up/practical-proof.

- [ ] **Step 4: Implement safe environment diagnosis**

`scripts/doctor.sh` reports OS, architecture, and availability/version of `go`, `git`, `docker`, `docker compose`, `make`, `psql`, and `curl`. Missing tools produce installation guidance, not automatic installation. The script exits non-zero only for an internal script error, not for missing optional tools.

- [ ] **Step 5: Write onboarding lessons**

Cover choosing OS-specific installation instructions, verifying checksums or package-manager source, `GOROOT`, `GOPATH`, `GOMODCACHE`, creating a module, running tests, Docker Desktop/Engine, Compose, and requesting confirmation before system changes.

- [ ] **Step 6: Verify**

Run: `bash scripts/doctor.sh; bash scripts/check-skills.sh; go test ./internal/coursecheck -v`

Expected: a readable environment report, seven valid skills, and passing tests.

- [ ] **Step 7: Commit**

```bash
git add .agents 00-onboarding scripts internal/coursecheck
git commit -m "feat: add Codex tutor and onboarding workflows"
```

### Task 3: Roadmap, Lesson Template, and Exam System

**Files:**
- Create: `roadmap/12-WEEKS.md`
- Create: `roadmap/SYLLABUS.md`
- Create: `roadmap/READINESS.md`
- Create: `roadmap/INTERVIEW-CHECKLIST.md`
- Create: `templates/TOPIC.md`
- Create: `templates/LAB.md`
- Create: `templates/CHECK.md`
- Create: `templates/DEEP_DIVE.md`
- Create: `exams/README.md`
- Create: `exams/RUBRIC.md`
- Create: `exams/go-start.md`
- Create: `exams/go-core.md`
- Create: `exams/concurrency.md`
- Create: `exams/http.md`
- Create: `exams/postgresql.md`
- Create: `exams/redis-kafka.md`
- Create: `exams/architecture-infra.md`
- Create: `exams/full-middle.md`

**Interfaces:**
- Consumes: dates, five-level mastery rubric, and interaction contract.
- Produces: canonical topic template and question banks referenced from every module.

- [ ] **Step 1: Add validator fixtures for topic completeness**

Assert every syllabus topic has a theory file, lab/check link where practice applies, difficulty marker, prerequisites, expected duration, retelling prompt, and exam references.

- [ ] **Step 2: Write the 12-week calendar and syllabus**

Assign every design-spec topic to a week and mark it `start`, `core`, `interview`, `advanced`, or `production`. Include minimum weekly path and optional deep dives so advanced theory cannot block progress.

- [ ] **Step 3: Write grading and readiness documents**

Define `не изучено`, `узнаёт`, `объясняет`, `применяет`, `защищает решение`; specify evidence for each. Define interview readiness gates for Go, concurrency, HTTP, PostgreSQL, messaging, architecture, and infrastructure.

- [ ] **Step 4: Write question banks**

Each bank includes conceptual questions, code/SQL reading, deliberate misconception probes, practical tasks, follow-ups, and answer keys visible only to the examiner workflow. Include spaced-repetition tags and four cumulative mock-interview sets.

- [ ] **Step 5: Verify**

Run: `bash scripts/check-course.sh`

Expected: all roadmap links, templates, and exam references validate.

- [ ] **Step 6: Commit**

```bash
git add roadmap templates exams internal/coursecheck
git commit -m "feat: define roadmap and examination system"
```

### Task 4: Go Start and Go Core Curriculum

**Files:**
- Create: `modules/01-go-start/README.md`
- Create: `modules/01-go-start/01-syntax.md`
- Create: `modules/01-go-start/02-packages-modules.md`
- Create: `modules/01-go-start/03-functions.md`
- Create: `modules/01-go-start/04-arrays-strings.md`
- Create: `modules/01-go-start/05-slices.md`
- Create: `modules/01-go-start/06-maps.md`
- Create: `modules/01-go-start/07-structs-methods.md`
- Create: `modules/01-go-start/08-values-pointers.md`
- Create: `modules/01-go-start/09-defer-panic-recover.md`
- Create: `modules/02-go-core/README.md`
- Create: `modules/02-go-core/01-interfaces.md`
- Create: `modules/02-go-core/02-errors.md`
- Create: `modules/02-go-core/03-generics.md`
- Create: `modules/02-go-core/04-table-tests.md`
- Create: `modules/02-go-core/05-bench-fuzz.md`
- Create: `modules/02-go-core/06-slice-map-internals.md`
- Create: `modules/02-go-core/07-memory-escape.md`
- Create: `modules/02-go-core/08-gc.md`
- Create: `modules/02-go-core/09-pprof.md`
- Create: `labs/go/01-syntax/{LAB.md,CHECK.md,starter/go.mod,starter/main.go}`
- Create: `labs/go/02-values-pointers/{LAB.md,CHECK.md,starter/go.mod,starter/main.go,starter/main_test.go}`
- Create: `labs/go/03-arrays-strings-slices/{LAB.md,CHECK.md,starter/go.mod,starter/main.go,starter/main_test.go}`
- Create: `labs/go/04-maps/{LAB.md,CHECK.md,starter/go.mod,starter/main.go,starter/main_test.go}`
- Create: `labs/go/05-structs-methods/{LAB.md,CHECK.md,starter/go.mod,starter/main.go,starter/main_test.go}`
- Create: `labs/go/06-errors-panic/{LAB.md,CHECK.md,starter/go.mod,starter/main.go,starter/main_test.go}`
- Create: `labs/go/07-interfaces/{LAB.md,CHECK.md,starter/go.mod,starter/main.go,starter/main_test.go}`
- Create: `labs/go/08-tests-bench-fuzz/{LAB.md,CHECK.md,starter/go.mod,starter/calc.go,starter/calc_test.go}`
- Create: `labs/go/09-memory-gc/{LAB.md,CHECK.md,DEEP_DIVE.md,starter/go.mod,starter/main.go}`
- Create: `labs/go/10-pprof/{LAB.md,CHECK.md,DEEP_DIVE.md,starter/go.mod,starter/main.go}`
- Create: `labs/go/solutions/README.md`
- Create: `labs/go/solutions/01-syntax.md`
- Create: `labs/go/solutions/02-values-pointers.md`
- Create: `labs/go/solutions/03-arrays-strings-slices.md`
- Create: `labs/go/solutions/04-maps.md`
- Create: `labs/go/solutions/05-structs-methods.md`
- Create: `labs/go/solutions/06-errors-panic.md`
- Create: `labs/go/solutions/07-interfaces.md`
- Create: `labs/go/solutions/08-tests-bench-fuzz.md`
- Create: `labs/go/solutions/09-memory-gc.md`
- Create: `labs/go/solutions/10-pprof.md`
- Test: every `labs/go/**/go.mod` package

**Interfaces:**
- Consumes: canonical templates and Go version pin.
- Produces: individually runnable labs with `go test`, `go test -bench`, `go test -fuzz`, escape analysis, and pprof commands.

- [ ] **Step 1: Create module indexes and focused topic files**

Cover syntax, packages/modules, functions, arrays, strings, slices, maps, structs, methods, value/pointer semantics, interfaces, errors, defer/panic/recover, generics, tests, benchmarks, fuzzing, runtime memory, GC, escape analysis, and pprof. Contrast with JavaScript only where it shortens the initial mental-model jump.

- [ ] **Step 2: Create fast-result labs**

Every lab must run in under two minutes, begin with a prediction, include one deliberate break, and end with an oral retelling prompt. Include observable capacity growth, UTF-8 byte/rune differences, shared slice backing arrays, nil map panic, copied structs, wrapped errors, interface nil trap, benchmark allocation counts, escape output, and a pprof endpoint.

- [ ] **Step 3: Add separated solutions**

Provide corrected code and explanation under `labs/go/solutions/`, mirroring lab IDs. Each solution states which misconception it resolves; learner-facing files link only through the examiner.

- [ ] **Step 4: Verify all Go labs**

Run: `bash scripts/test-go-labs.sh`

Expected: all normal starter packages and validator tests pass; intentionally broken snippets are stored as `.txt` or build-tagged examples and do not break the repository build.

- [ ] **Step 5: Verify teaching structure**

Run: `bash scripts/check-course.sh`

Expected: every Go topic and lab satisfies the canonical structure.

- [ ] **Step 6: Commit**

```bash
git add modules/01-go-start modules/02-go-core labs/go
git commit -m "feat: add Go start and core curriculum"
```

### Task 5: Concurrency and HTTP Curriculum

**Files:**
- Create: `modules/03-concurrency/{README.md,01-gmp.md,02-goroutines.md,03-channels.md,04-select.md,05-context.md,06-sync.md,07-atomic.md,08-races-deadlocks-leaks.md,09-worker-pool.md,10-semaphore-pipelines.md,11-graceful-shutdown.md}`
- Create: `modules/04-http/{README.md,01-net-http.md,02-client.md,03-server.md,04-middleware-json.md,05-timeouts-cancellation.md,06-httptest.md,07-graceful-shutdown.md}`
- Create: `labs/concurrency/{README.md,go.mod}`
- Create: `labs/concurrency/channels/{LAB.md,CHECK.md,starter/channels.go,starter/channels_test.go}`
- Create: `labs/concurrency/context/{LAB.md,CHECK.md,starter/context.go,starter/context_test.go}`
- Create: `labs/concurrency/race/{LAB.md,CHECK.md,starter/race.go,starter/race_test.go}`
- Create: `labs/concurrency/deadlock/{LAB.md,CHECK.md,starter/main.go}`
- Create: `labs/concurrency/workerpool/{LAB.md,CHECK.md,starter/pool.go,starter/pool_test.go}`
- Create: `labs/concurrency/semaphore/{LAB.md,CHECK.md,starter/semaphore.go,starter/semaphore_test.go}`
- Create: `labs/concurrency/pipeline/{LAB.md,CHECK.md,starter/pipeline.go,starter/pipeline_test.go}`
- Create: `labs/concurrency/shutdown/{LAB.md,CHECK.md,starter/main.go,starter/main_test.go}`
- Create: `labs/concurrency/solutions/{README.md,channels.md,context.md,race.md,deadlock.md,workerpool.md,semaphore.md,pipeline.md,shutdown.md}`
- Create: `labs/http/README.md`
- Create: `labs/http/01-task-service/{LAB.md,CHECK.md,starter/go.mod,starter/cmd/server/main.go,starter/internal/api/handler.go,starter/internal/api/middleware.go,starter/internal/api/handler_test.go,starter/internal/client/client.go,starter/internal/client/client_test.go}`
- Create: `labs/http/solutions/README.md`

**Interfaces:**
- Consumes: Go Core knowledge and course templates.
- Produces: runnable GMP demonstrations, race/deadlock/leak exercises, concurrency patterns, and a small `net/http` service with tests.

- [ ] **Step 1: Write concurrency theory**

Cover GMP, scheduling/preemption, goroutine lifecycle, channels and close ownership, `select`, context trees, sync primitives, typed atomics/CAS, race/deadlock/leak diagnosis, worker pools, semaphores, fan-in/out, pipelines, and graceful shutdown.

- [ ] **Step 2: Build isolated concurrency labs**

Provide labs that visibly block on unbuffered channels, fill buffered channels, panic on send-after-close, leak without cancellation, fail under `-race`, deadlock through lock ordering, enforce semaphore bounds, cancel a worker pool, and drain on shutdown. Expected failures must have dedicated commands and cleanup guidance.

- [ ] **Step 3: Write HTTP theory and lab**

Create a dependency-free `net/http` JSON task service. Include server/client timeouts, body close, middleware, request ID, validation, error envelope, cancellation, `httptest`, and signal-driven shutdown.

- [ ] **Step 4: Verify normal behavior**

Run: `bash scripts/test-go-labs.sh` and then run `go test -race ./...` from `labs/concurrency` and `labs/http`.

Expected: all corrected/default paths pass with no races.

- [ ] **Step 5: Verify documented failures**

Run each deliberate failure command from its `LAB.md`; confirm its observed race, deadlock timeout, panic, or cancellation matches the documented signature without hanging indefinitely.

- [ ] **Step 6: Commit**

```bash
git add modules/03-concurrency modules/04-http labs/concurrency labs/http
git commit -m "feat: add concurrency and HTTP laboratories"
```

### Task 6: PostgreSQL Curriculum and Reproducible Data Labs

**Files:**
- Create: `modules/05-postgresql/{README.md,01-ddl-dml.md,02-transactions-acid.md,03-isolation-mvcc.md,04-locks.md,05-deadlocks.md,06-indexes.md,07-explain.md,08-pool-migrations-vacuum.md,09-wal-replication.md,10-logical-replication-debezium.md}`
- Create: `deploy/compose.yaml`
- Create: `deploy/postgres/init/001_schema.sql`
- Create: `labs/postgres/README.md`
- Create: `labs/postgres/01-ddl-dml/{LAB.md,CHECK.md,starter/queries.sql}`
- Create: `labs/postgres/02-transactions/{LAB.md,CHECK.md,starter/transactions.sql}`
- Create: `labs/postgres/03-isolation/{LAB.md,CHECK.md,starter/isolation.sql}`
- Create: `labs/postgres/04-locks/{LAB.md,CHECK.md,starter/locks.sql,starter/locks.sh}`
- Create: `labs/postgres/05-skip-locked/{LAB.md,CHECK.md,starter/worker.sql}`
- Create: `labs/postgres/06-deadlock/{LAB.md,CHECK.md,starter/deadlock.sh}`
- Create: `labs/postgres/07-indexes/{LAB.md,CHECK.md,starter/indexes.sql}`
- Create: `labs/postgres/08-explain/{LAB.md,CHECK.md,starter/explain.sql}`
- Create: `labs/postgres/09-logical-replication/{LAB.md,CHECK.md,starter/logical-replication.sql}`
- Create: `scripts/seed-postgres.sql`
- Create: `scripts/deps-up.sh`
- Create: `scripts/deps-down.sh`
- Create: `scripts/wait-deps.sh`
- Create: `scripts/seed-postgres.sh`

**Interfaces:**
- Consumes: Docker Compose and PostgreSQL version pins.
- Produces: healthy PostgreSQL service, idempotent schema/seed scripts, query-plan and locking experiments.

- [ ] **Step 1: Define Compose PostgreSQL and health check**

Use a named volume, explicit credentials intended only for local learning, port override support, and `pg_isready`. `deps-down.sh` must preserve volumes unless the user explicitly supplies a documented destructive flag and confirms it.

- [ ] **Step 2: Write schema and deterministic seed**

Create customers, orders, jobs, documents, tags, and outbox tables with constraints. Generate enough skewed rows through SQL `generate_series` to demonstrate both Seq Scan and Index Scan without external data downloads.

- [ ] **Step 3: Write SQL labs**

Cover DDL/DML, transactions, isolation anomalies, MVCC, row/table/advisory locks, `FOR UPDATE`, `SKIP LOCKED`, deadlock, B-tree/hash/GIN, composite/partial indexes, selectivity, `EXPLAIN (ANALYZE, BUFFERS)`, pool pressure, vacuum, WAL, physical/logical replication concepts, publication/subscription, replication identity, and Debezium overview.

- [ ] **Step 4: Provide controlled multi-session scripts**

Scripts print which terminal/session owns which lock, use bounded timeouts, clean up transactions on exit, and finish without leaving blocked sessions.

- [ ] **Step 5: Verify**

Run: `docker compose -f deploy/compose.yaml up -d postgres`

Run: `bash scripts/wait-deps.sh postgres && bash scripts/seed-postgres.sh && bash scripts/check-course.sh`

Expected: PostgreSQL becomes healthy, seed is repeatable, representative EXPLAIN output demonstrates the documented plan changes, and lock scripts terminate.

- [ ] **Step 6: Commit**

```bash
git add modules/05-postgresql deploy labs/postgres scripts
git commit -m "feat: add PostgreSQL data and locking curriculum"
```

### Task 7: Redis, Kafka, Architecture, and Infrastructure Curriculum

**Files:**
- Create: `modules/06-redis/{README.md,01-data-types-ttl.md,02-cache-aside.md,03-stampede-locks.md,04-pubsub.md,05-streams.md}`
- Create: `modules/07-kafka/{README.md,01-log-brokers-topics.md,02-partitions-offsets.md,03-producers-consumers.md,04-groups-rebalance.md,05-delivery-semantics.md,06-retries-dlq.md,07-order-idempotency.md}`
- Create: `modules/08-architecture/{README.md,01-cap.md,02-sharding-consistent-hashing.md,03-outbox.md,04-saga.md,05-cdc-debezium.md,06-monolith-microservices.md,07-observability-security.md}`
- Create: `modules/09-infrastructure/{README.md,01-docker.md,02-compose-networking.md,03-nginx-reverse-proxy.md,04-load-balancing.md,05-tls-timeouts-retries.md,06-kubernetes-core.md,07-kubernetes-delivery.md}`
- Modify: `deploy/compose.yaml`
- Create: `deploy/nginx/nginx.conf`
- Create: `deploy/kubernetes/namespace.yaml`
- Create: `deploy/kubernetes/configmap.yaml`
- Create: `deploy/kubernetes/secret.example.yaml`
- Create: `deploy/kubernetes/deployment.yaml`
- Create: `deploy/kubernetes/service.yaml`
- Create: `deploy/kubernetes/ingress.yaml`
- Create: `deploy/kubernetes/hpa.yaml`
- Create: `labs/redis/README.md`
- Create: `labs/redis/01-cache/{LAB.md,CHECK.md,starter/go.mod,starter/cache.go,starter/cache_test.go}`
- Create: `labs/redis/02-pubsub/{LAB.md,CHECK.md,starter/go.mod,starter/pubsub.go}`
- Create: `labs/redis/03-streams/{LAB.md,CHECK.md,starter/go.mod,starter/streams.go}`
- Create: `labs/kafka/README.md`
- Create: `labs/kafka/01-producer-consumer/{LAB.md,CHECK.md,starter/go.mod,starter/producer.go,starter/consumer.go}`
- Create: `labs/kafka/02-idempotency/{LAB.md,CHECK.md,starter/go.mod,starter/idempotency.go,starter/idempotency_test.go}`
- Create: `labs/architecture/README.md`
- Create: `labs/architecture/01-cap/{LAB.md,CHECK.md,starter/scenarios.md}`
- Create: `labs/architecture/02-consistent-hashing/{LAB.md,CHECK.md,starter/go.mod,starter/hashring.go,starter/hashring_test.go}`
- Create: `labs/architecture/03-outbox-saga/{LAB.md,CHECK.md,starter/scenarios.md}`
- Create: `labs/nginx/README.md`
- Create: `labs/nginx/01-balancing/{LAB.md,CHECK.md,starter/go.mod,starter/cmd/replica/main.go}`

**Interfaces:**
- Consumes: Compose lifecycle scripts and the small HTTP service.
- Produces: healthy Redis and Kafka-compatible dependencies, messaging/cache experiments, Nginx multi-replica lab, and readable Kubernetes equivalents.

- [ ] **Step 1: Add Redis and Kafka-compatible services**

Pin images, configure single-node local development, health checks, named volumes, and documented ports. Keep the broker API Kafka-compatible so Go exercises use standard Kafka concepts.

- [ ] **Step 2: Write Redis labs**

Cover strings/hashes/sets, TTL, cache-aside, stale cache, stampede mitigation, Pub/Sub message loss without a subscriber, Streams persistence and consumer groups, plus limitations of naive distributed locks.

- [ ] **Step 3: Write Kafka labs**

Cover topic/partition/offset inspection, key-based ordering, consumer groups, rebalance, manual commits, at-most/at-least-once, retry/backoff, DLQ, poison messages, and idempotent handling. Every lab includes CLI observation plus a minimal Go producer or consumer where useful.

- [ ] **Step 4: Write architecture labs**

Use scenario cards and small simulations for CAP, consistent hashing, transactional outbox, saga orchestration/choreography, CDC/Debezium, and the monolith-versus-microservices decision.

- [ ] **Step 5: Build Nginx and Kubernetes learning artifacts**

Run multiple labeled HTTP replicas behind Nginx and observe balancing. Provide alternate configurations for round-robin, weight, least connections, and sticky behavior. Map the lab to commented Pod/Deployment/Service/Ingress/ConfigMap/Secret/probe/resource/rolling-update/HPA manifests without requiring a cluster.

- [ ] **Step 6: Verify**

Run: `docker compose -f deploy/compose.yaml config`

Run: `bash scripts/deps-up.sh && bash scripts/wait-deps.sh && bash scripts/test-go-labs.sh && bash scripts/check-course.sh`

Expected: Compose config is valid, dependencies become healthy, Go clients pass their integration checks when enabled, and every curriculum link validates.

- [ ] **Step 7: Commit**

```bash
git add modules/06-redis modules/07-kafka modules/08-architecture modules/09-infrastructure deploy labs scripts
git commit -m "feat: add messaging architecture and infrastructure curriculum"
```

### Task 8: Cumulative `ogen` Backend Lab

**Files:**
- Create: `modules/10-microservice/{README.md,01-contract.md,02-domain-http.md,03-postgres.md,04-redis.md,05-outbox-kafka.md,06-workers-shutdown.md,07-testing.md,08-containerization.md}`
- Create: `project/backend-lab/README.md`
- Create: `project/backend-lab/go.mod`
- Create: `project/backend-lab/Makefile`
- Create: `project/backend-lab/api/openapi.yaml`
- Create: `project/backend-lab/cmd/api/main.go`
- Create: `project/backend-lab/internal/config/config.go`
- Create: `project/backend-lab/internal/task/model.go`
- Create: `project/backend-lab/internal/task/service.go`
- Create: `project/backend-lab/internal/task/service_test.go`
- Create: `project/backend-lab/internal/task/repository.go`
- Create: `project/backend-lab/internal/platform/postgres/repository.go`
- Create: `project/backend-lab/internal/platform/postgres/outbox.go`
- Create: `project/backend-lab/internal/platform/redis/cache.go`
- Create: `project/backend-lab/internal/platform/kafka/publisher.go`
- Create: `project/backend-lab/internal/platform/kafka/consumer.go`
- Create: `project/backend-lab/internal/platform/httpserver/server.go`
- Create: `project/backend-lab/internal/platform/httpserver/handler.go`
- Create: `project/backend-lab/migrations/001_init.up.sql`
- Create: `project/backend-lab/migrations/001_init.down.sql`
- Create: `project/backend-lab/tests/integration_test.go`
- Create: `project/backend-lab/Dockerfile`

**Interfaces:**
- Consumes: Compose PostgreSQL/Redis/Kafka endpoints and generated `ogen` interfaces.
- Produces: task API with PostgreSQL source of truth, Redis cache-aside reads, transactional outbox, Kafka publisher/consumer, worker pool, cancellation, probes, and integration tests.

- [ ] **Step 1: Define contract-first acceptance tests**

Write the OpenAPI contract and tests for create/get/complete task, validation errors, request IDs, readiness, idempotent creation, cache miss/hit behavior, and outbox delivery state. Generate code using a pinned `ogen` version rather than hand-writing generated output.

- [ ] **Step 2: Implement the minimal HTTP/domain slice**

Implement config, task model, service boundaries, generated handler adapter, structured errors, timeouts, request IDs, and graceful shutdown. Keep generated code isolated from handwritten code.

- [ ] **Step 3: Add PostgreSQL and migrations**

Implement repository operations, transaction boundaries, optimistic conflict handling, and task-plus-outbox atomic commit. Integration tests use the Compose database and clean their own rows.

- [ ] **Step 4: Add Redis cache-aside**

Cache task reads with TTL, invalidate/update after writes, handle Redis outage as degraded behavior, and test stale/miss/error paths.

- [ ] **Step 5: Add Kafka outbox worker and consumer**

Use `FOR UPDATE SKIP LOCKED`, a bounded worker pool, context cancellation, retries with backoff, delivery marking, stable event IDs, and an idempotent consumer example.

- [ ] **Step 6: Containerize and document exercises**

Create a non-root multi-stage image and step-by-step tutorial checkpoints. Each checkpoint asks the learner to predict, modify, run, observe, explain, and pass an exam before advancing.

- [ ] **Step 7: Verify**

Run: `bash scripts/test-go-labs.sh`

Run inside project: `go generate ./... && go test ./...`

Run with dependencies: integration tests for API, PostgreSQL, Redis, outbox, Kafka, cancellation, and shutdown.

Expected: all unit tests pass; generated code is reproducible; enabled integration tests pass against healthy Compose services.

- [ ] **Step 8: Commit**

```bash
git add modules/10-microservice project/backend-lab
git commit -m "feat: add cumulative ogen backend laboratory"
```

### Task 9: Full Verification, Learner Dry Run, and Handoff

**Files:**
- Modify: `README.md`
- Modify: `progress/STATUS.md`
- Create: `docs/verification/2026-07-21-initial-course-check.md`
- Create: `scripts/verify-all.sh`

**Interfaces:**
- Consumes: every deliverable from Tasks 1–8.
- Produces: one bounded verification command and a clean first-session handoff.

- [ ] **Step 1: Create the verification orchestrator**

`scripts/verify-all.sh` runs shell syntax checks, course structure checks, Go formatting checks, unit tests, race tests for safe packages, Compose config validation, YAML parsing where available, and optional integration tests behind `RUN_INTEGRATION=1`. It prints skipped checks explicitly.

- [ ] **Step 2: Perform a clean learner dry run**

Follow only `README.md` and onboarding files: run doctor, initialize/inspect progress, start the first Go lesson, execute its lab, invoke a mini-exam, and verify that only evidence-backed progress is recorded.

- [ ] **Step 3: Perform a fresh Codex dry run**

Start from the repository root instructions and verify that the agent can answer: current status, next task, deadline pace, weak topics, and interview readiness without relying on conversation history.

- [ ] **Step 4: Run full verification**

Run: `bash scripts/verify-all.sh`

Run: `RUN_INTEGRATION=1 bash scripts/verify-all.sh`

Expected: mandatory checks pass; optional unavailable tooling is reported as skipped; integration checks pass when Docker is available.

- [ ] **Step 5: Record evidence**

Write exact commands, versions, pass/fail counts, intentionally skipped checks, and any environmental limitations to `docs/verification/2026-07-21-initial-course-check.md`. Do not mark learner topics completed during repository verification.

- [ ] **Step 6: Final commit**

```bash
git add README.md progress/STATUS.md docs/verification scripts/verify-all.sh
git commit -m "docs: verify and hand off Misha backend curriculum"
```

---

## Plan Self-Review

- Spec coverage: all thirteen design sections map to Tasks 1–9.
- Isolation: course tooling, agent behavior, topic content, dependency labs, cumulative service, and final verification have separate review gates.
- Type consistency: the shared validator interface is `coursecheck.Check(root string) []error`; all skills use the exact `progress/` files created in Task 1.
- Safety: installation is advisory, teardown preserves volumes by default, deliberate failures are bounded, and integration checks are opt-in where they mutate local dependency data.
- Scope control: Kubernetes, replication, and Debezium remain interview-level or optional; the main runnable path concentrates on Go, concurrency, HTTP, PostgreSQL, Redis, Kafka, and the cumulative service.
