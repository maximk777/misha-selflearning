# Real-world Use Cases and Mini-projects Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Дополнить все 80 тем курса реальными backend-сценариями, глубоким разбором и самостоятельными мини-проектами, а tutor skills научить автоконспектам, review-циклу и работе со слабыми местами.

**Architecture:** `roadmap/SYLLABUS.md` задаёт порядок доступных знаний. Тематические Markdown-файлы содержат единый содержательный контракт, Go checker проверяет его структуру, а repo-local skills управляют уроком, конспектом, аудитом курса и доказательным прогрессом.

**Tech Stack:** Markdown, repo-local Codex skills, Go 1.24, shell verification scripts.

## Global Constraints

- Весь разговор с учеником ведётся по-русски; привычные English technical terms сохраняются.
- Каждый тематический файл получает не менее трёх конкретных backend-use cases, глубокий слой и самостоятельный мини-проект.
- Обязательный мини-проект использует только текущую и более ранние темы из `roadmap/SYLLABUS.md`.
- Решение, starter-код и архитектура не выдаются до самостоятельной попытки; подсказка даётся только по прямой просьбе и по одной за сообщение.
- Прогресс обновляется только по доказательству; файлы `progress/` и `solutions/` в этой реализации не изменяются.
- После `git pull` весь материал уже присутствует; локальная генерация учеником не требуется.
- Пользователь не просил создавать commits; все изменения остаются в рабочем дереве.

---

### Task 1: Structural checker contract

**Files:**
- Modify: `internal/coursecheck/check_test.go`
- Modify: `internal/coursecheck/check.go`
- Modify: `scripts/check-skills.sh`
- Test: `internal/coursecheck/check_test.go`

**Interfaces:**
- Consumes: topic files at `modules/<module>/<topic>.md`; repo-local skills at `.agents/skills/<name>/SKILL.md`.
- Produces: stable diagnostics for missing `## Где это применяется в реальном backend`, `## Глубокое погружение`, `## Мини-проект`, `### Результат`, `### Разрешённые знания`, and `### Критерии приёмки`.

- [ ] Add table-driven failing tests that create a minimal topic and remove each required heading one at a time; expect a diagnostic naming the exact topic and heading.
- [ ] Run `GOCACHE=/tmp/misha-coursecheck-cache go test ./internal/coursecheck` and verify the new cases fail because topic validation is absent.
- [ ] Add topic discovery limited to two-level Markdown files under `modules/`, excluding `README.md`, and validate the six exact headings.
- [ ] Run `GOCACHE=/tmp/misha-coursecheck-cache go test ./internal/coursecheck` and verify all cases pass.
- [ ] Extend `scripts/check-skills.sh` to include `misha-notes` and `misha-course-audit` while preserving the evidence-file contract for learner-facing skills.

### Task 2: Tutor, notes, audit, and interview skills

**Files:**
- Modify: `.agents/skills/misha-lesson/SKILL.md`
- Modify: `.agents/skills/misha-review/SKILL.md`
- Modify: `.agents/skills/misha-exam/SKILL.md`
- Modify: `.agents/skills/misha-mock-interview/SKILL.md`
- Modify: `.agents/skills/misha-progress/SKILL.md`
- Create: `.agents/skills/misha-notes/SKILL.md`
- Create: `.agents/skills/misha-course-audit/SKILL.md`
- Modify: `templates/TOPIC.md`
- Modify: `templates/LAB.md`
- Test: `scripts/check-skills.sh`

**Interfaces:**
- Consumes: syllabus order, topic blocks, `progress/PROFILE.md`, `STATUS.md`, `EVIDENCE.md`, `EXAMS.md`, `REVIEW_QUEUE.md`.
- Produces: one-question tutoring flow, hint-on-request review loop, evidence-filtered notes, weak-spot interview selection, and safe course enrichment.

- [ ] Record baseline agent outputs for lesson, notes, and course audit before editing skills.
- [ ] Rewrite `misha-lesson` as a compact positive contract: real use cases → prediction → independent project attempt → requested hints only → review loop → complication → tests and retelling.
- [ ] Make `misha-review` return one concrete defect plus a local principle example, then request another attempt until acceptance criteria pass.
- [ ] Make `misha-exam` and `misha-mock-interview` prioritize recorded wrong models and require transfer to a new example before clearing weakness.
- [ ] Make `misha-progress` synchronize evidence, current status, exam errors, and review queue without treating a completed lab as a completed topic.
- [ ] Create `misha-notes`: ask exactly one scope question when missing, include only evidenced material, internal mechanics, real uses, prior mistakes, project evidence, self-check questions, and honest gaps.
- [ ] Create `misha-course-audit`: run the checker, enrich only reported topics using earlier syllabus knowledge, preserve good text, and never edit `progress/`, `solutions/`, or learner attempts.
- [ ] Update `templates/TOPIC.md` and `templates/LAB.md` with the same project and review contract.
- [ ] Run `bash scripts/check-skills.sh` and forward-test the same three baseline scenarios with the revised skills.

### Task 3: Go Start and Go Core content

**Files:**
- Modify: `modules/01-go-start/01-syntax.md` through `modules/01-go-start/09-defer-panic-recover.md`
- Modify: `modules/02-go-core/01-interfaces.md` through `modules/02-go-core/09-pprof.md`

**Interfaces:**
- Consumes: exact order and prerequisites in `roadmap/SYLLABUS.md`.
- Produces: 18 complete topic contracts progressing through a cumulative CLI order/report tool and a testable domain library.

- [ ] For each topic, add three concrete real-life uses with role and boundary; future systems may be mentioned but not required.
- [ ] Add `## Глубокое погружение` covering runtime model, invariants, costs, edge cases, proof method, and interview reasoning appropriate to that topic.
- [ ] Add a mini-project with exact result, business scenario, permitted earlier knowledge, verification, acceptance criteria, and one post-MVP complication.
- [ ] In slices cover descriptor, backing array, aliasing, reallocation, retention, and element copy; in interfaces cover method sets, typed nil, implicit satisfaction, and consumer-side interfaces.
- [ ] Run the structural checker and manually inspect syntax, slices, interfaces, and pprof for forward-knowledge violations.

### Task 4: Concurrency and HTTP content

**Files:**
- Modify: `modules/03-concurrency/01-gmp.md` through `modules/03-concurrency/11-graceful-shutdown.md`
- Modify: `modules/04-http/01-net-http.md` through `modules/04-http/07-graceful-shutdown.md`

**Interfaces:**
- Consumes: completed Go Core topics and the sequential order within concurrency and HTTP.
- Produces: 18 complete topic contracts culminating in an HTTP dispatcher and Task API.

- [ ] Add concrete service scenarios and deep internals for scheduler, goroutines, channels, select, context, sync, atomic, failure modes, worker pools, semaphores, pipelines, and shutdown.
- [ ] In channels cover send/receive wait queues conceptually, buffer states, nil/closed behavior, close ownership, blocking, happens-before, cancellation, and leaks without claiming unstable runtime implementation details as API guarantees.
- [ ] Keep early goroutine projects free of channels until the channels lesson; add each new primitive only after it appears in syllabus.
- [ ] Make the cumulative dispatcher project require fixed workers, semaphore-limited HTTP calls, context, body closure, race-free result collection, leak-free shutdown, and tests only after HTTP client is available; the concurrency-only precursor uses a function dependency instead of HTTP.
- [ ] Add HTTP server/client/middleware/test/shutdown projects using only earlier Go and concurrency knowledge.
- [ ] Run the structural checker and inspect goroutines, channels, worker pool, semaphore, HTTP client, and HTTP shutdown.

### Task 5: PostgreSQL, Redis, and Kafka content

**Files:**
- Modify: `modules/05-postgresql/01-ddl-dml.md` through `modules/05-postgresql/10-logical-replication-debezium.md`
- Modify: `modules/06-redis/01-data-types-ttl.md` through `modules/06-redis/05-streams.md`
- Modify: `modules/07-kafka/01-log-brokers-topics.md` through `modules/07-kafka/07-order-idempotency.md`

**Interfaces:**
- Consumes: earlier HTTP, Go, concurrency, and database topics.
- Produces: 22 complete contracts for an order store, cache, and event processor.

- [ ] Add three concrete real uses and deep internals per topic, including failure modes and proof queries/metrics.
- [ ] Sequence PostgreSQL projects from schema and transactions to MVCC, locks, indexes, plans, pool/vacuum, WAL, and logical replication without requiring later Kafka.
- [ ] Sequence Redis projects from TTL/data types to cache-aside, stampede, Pub/Sub, and Streams while preserving PostgreSQL as source of truth where relevant.
- [ ] Sequence Kafka projects from log/partitions to producers, groups, delivery, retries/DLQ, and idempotency; require PostgreSQL only where it has already been covered.
- [ ] Run the structural checker and inspect transaction isolation, indexes, cache-aside, Streams, delivery semantics, and idempotency.

### Task 6: Architecture, infrastructure, and final microservice content

**Files:**
- Modify: `modules/08-architecture/01-cap.md` through `modules/08-architecture/07-observability-security.md`
- Modify: `modules/09-infrastructure/01-docker.md` through `modules/09-infrastructure/07-kubernetes-delivery.md`
- Modify: `modules/10-microservice/01-contract.md` through `modules/10-microservice/08-containerization.md`

**Interfaces:**
- Consumes: the entire earlier minimum path.
- Produces: 22 complete contracts that evolve the existing backend lab through design, deployment, and final defense.

- [ ] Add concrete architecture and operations cases with measurable consequences rather than term definitions.
- [ ] Make design mini-projects produce scenarios, decision records, failure experiments, or runnable configurations with explicit acceptance criteria.
- [ ] Sequence Docker → Compose → Nginx → balancing → TLS/timeouts/retries → Kubernetes core → delivery without requiring future infrastructure early.
- [ ] Make all eight microservice checkpoints integrate only mechanisms already taught and end with tests, Compose, operational checks, and trade-off defense.
- [ ] Run the structural checker and inspect CAP, outbox, Docker, retry storms, Kubernetes probes, and final containerization.

### Task 7: Whole-course verification and quality review

**Files:**
- Modify only files with verified Critical or Important review findings; never edit `progress/` or `solutions/`.

**Interfaces:**
- Consumes: outputs of Tasks 1–6.
- Produces: a pull-ready learner checkout with complete content and passing verification.

- [ ] Count exactly 80 non-README topic files and verify each has all six required headings.
- [ ] Search for placeholders, generic «используется в высоконагруженных системах» claims, solution code, and requirements that reference later syllabus topics.
- [ ] Run `GOCACHE=/tmp/misha-coursecheck-cache go test ./internal/coursecheck ./cmd/coursecheck`.
- [ ] Run `bash scripts/check-skills.sh`, `bash scripts/check-course.sh`, and `bash scripts/check-course.sh` against a temporary intentionally incomplete topic to prove the diagnostic.
- [ ] Run `GOCACHE=/tmp/misha-go-labs-cache bash scripts/test-go-labs.sh` and `bash scripts/verify-all.sh`; distinguish Docker/environment failures from course failures.
- [ ] Forward-test lesson, notes, mock interview, and course audit skills in fresh contexts.
- [ ] Review the full diff for unintended changes and confirm the three existing user-modified `progress/` files are untouched by this implementation.
