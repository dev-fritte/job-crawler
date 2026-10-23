<div align=center>

## Job Crawler

**Job crawler platform that matches job offers to your search criteria**

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/dev-fritte/job-crawler?filename=backend%2Fgo.mod)
![Angular](https://img.shields.io/badge/Angular-not%20yet-DD0031?logo=angular&logoColor=white)

</div>

---

> The main reason for this project is to learn [Go](https://go.dev/) and [Angular](https://angular.dev/).

Job Crawler collects job offers from job platforms and from the career pages of selected companies
in the DACH region (Germany, Austria, Switzerland). Every offer is normalized and classified once and
stored in a database. Users define search profiles (location and radius, job title, technologies,
benefits, salary, company size and industry) and get notified about matching offers by e-mail or
WhatsApp.

## Features

- **Crawler adapters** per source type: official APIs, applicant tracking systems (Greenhouse, Personio, …)
  and generic career pages that publish `schema.org/JobPosting` data
- **Normalization**: salary converted to a comparable yearly amount in EUR, locations geocoded for radius search
- **Deduplication**: the same offer from several sources becomes one job
- **Classification**: tags for technologies, skills and benefits, using a synonym dictionary and an LLM
- **Search profiles** with must-have, nice-to-have and exclude criteria
- **Notifications** by e-mail or WhatsApp, instantly or as a daily or weekly digest

## Architecture

```
┌──────────────────────────────────────────────────────────────┐
│ Sources: Bundesagentur API · Greenhouse · Personio · JSON-LD │
└──────────────────────────────┬───────────────────────────────┘
                               │ HTTP
┌──────────────────────────────▼───────────────────────────────┐
│ worker                                                       │
│ Scheduler ► Adapter ► Pipeline ► Matching ► Notifier         │
└──────────────┬──────────────────────────────────┬────────────┘
               │ stores                           │ sends
┌──────────────▼───────────────┐      ┌───────────▼────────────┐
│ PostgreSQL + PostGIS         │      │ E-Mail · WhatsApp      │
│ Jobs, profiles, matches,     │      └────────────────────────┘
│ job queue (River)            │
└──────────────▲───────────────┘
               │ reads / writes
┌──────────────┴───────────────┐
│ api                          │
│ Auth · Job search · Profiles │
└──────────────▲───────────────┘
               │ REST / JSON
┌──────────────┴───────────────┐
│ Angular app                  │
│ Login · Profiles · Matches   │
└──────────────────────────────┘
```

`worker` and `api` are the same Go binary started with different subcommands
(`job-crawler worker`, `job-crawler api`). They share the domain model, configuration and database code
but run as separate processes.

### How a job offer travels through the system

1. **Schedule**: the scheduler puts due sources into the job queue.
2. **Fetch**: one adapter per source loads the offers (API, ATS feed or JSON-LD).
3. **Store raw data**: the response and its hash are saved, unchanged content is skipped.
4. **Normalize**: convert the salary, geocode the location, detect remote mode and seniority.
5. **Deduplicate**: merge offers from several sources into one canonical job.
6. **Classify**: tags from the synonym dictionary, the rest is extracted by an LLM.
7. **Match**: the new job is checked against all active search profiles.
8. **Notify**: send the match instantly or collect it into a digest.

## Tech stack

| Area | Technology |
|---|---|
| Backend / API | Go, `net/http` with chi, OpenAPI via Huma |
| Crawler | Go, `net/http`, goquery |
| Background jobs | River (job queue on PostgreSQL) |
| Database | PostgreSQL with PostGIS, goose migrations, pgx and sqlc |
| Classification | Rules (synonym dictionary) and Claude through `anthropic-sdk-go` |
| Frontend | Angular with Angular Material |
| Local development | Docker Compose |
