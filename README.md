# GatorCan

GatorCan is a full-stack e-learning platform for course management,
assignment submission/grading, and student-instructor collaboration, built
as a Canvas-style academic project ("Code Mavericks").

[Demo video](https://drive.google.com/file/d/1LywboJhJ_LxYwHm6a4FTGKUGW-ZBsTMz/view?usp=sharing)

## Architecture

The project is composed of three services:

- **`gatorcan-backend`** — Go REST API (Gin) providing auth, course,
  assignment, submission, and user-management endpoints.
- **`gatorcan-frontend`** — React (Vite) single-page app for students,
  instructors, and admins.
- **`AIservice`** — Python FastAPI microservice that provides course
  recommendations (TF-IDF/cosine similarity) and text summarization,
  called by the backend.

## Features

- JWT-based authentication and role-based access control (Student,
  Instructor, TA, Admin) enforced via Gin middleware
- Course enrollment, browsing, and per-course detail views, plus a
  personalized weekly schedule (`@schedule-x` calendar)
- Assignment creation/update (instructors), submission, and file upload
- Grading workflow and grade retrieval
- Per-course real-time chat and discussions (React components backed by
  AWS AppSync GraphQL subscriptions)
- AI-powered course recommendations and text summarization via the
  AIservice microservice
- AWS SNS-backed notifications (`adapters/notificationAdapters.go`)

## API Overview (backend)

| Method | Path                                              | Access              |
|--------|----------------------------------------------------|---------------------|
| POST   | `/login`                                            | Public              |
| POST   | `/admin/add_user`                                   | Admin                |
| DELETE | `/admin/:username`                                  | Admin                |
| PUT    | `/admin/update_role`                                | Admin                |
| GET    | `/user/:username`                                   | Student, Admin       |
| PUT    | `/user/update`                                      | Student, Admin       |
| GET    | `/courses/`                                         | Student, Admin, Instructor |
| GET    | `/courses/enrolled`                                 | Student, Admin, Instructor |
| POST   | `/courses/enroll`                                   | Student, Admin, Instructor |
| GET    | `/courses/:cid`                                     | Student, Admin, Instructor |
| GET    | `/courses/recommendations`                          | Student, Admin, Instructor |
| POST   | `/courses/summarize`                                | Student, Admin, Instructor |
| GET    | `/courses/:cid/grades`                              | Student, Admin, Instructor |
| POST   | `/instructor/courses/:cid/upsertassignment`         | Instructor            |
| POST   | `/instructor/courses/:cid/assignments/:aid/grade`   | Instructor            |
| GET    | `/courses/:cid/assignments/`                        | Student, Instructor, TA |
| GET    | `/courses/:cid/assignments/:aid`                    | Student, Instructor, TA |
| GET    | `/courses/:cid/assignments/:aid/submission`         | Student, Instructor, TA |
| GET    | `/courses/:cid/assignments/:aid/submissions`        | Student, Instructor, TA |
| POST   | `/courses/:cid/assignments/:aid/submit`             | Student, Instructor, TA |
| POST   | `/courses/:cid/assignments/:aid/upload`             | Student, Instructor, TA |

See [`BackendAPIDocumentation.md`](BackendAPIDocumentation.md) for full
request/response details.

## Tech Stack

- **Backend:** Go 1.23, Gin, GORM (SQLite), JWT (`golang-jwt`), AWS SDK v2
  (SNS)
- **Frontend:** React 18, Vite, Material UI, AWS Amplify (AppSync GraphQL),
  `@schedule-x` calendar, Axios
- **AI service:** Python, FastAPI, pandas, scikit-learn
- **Testing:** Go `testify`, Cypress, Jest/React Testing Library

## Setup

### Backend

```bash
cd gatorcan-backend
go mod download
go mod tidy
go build
go run main.go
```

Configuration is read from environment variables (see
`gatorcan-backend/config/config.go`), including `DB_HOST`, `DB_USER`,
`DB_PASSWORD`, `SERVER_PORT`, `CORS_ALLOWED_ORIGINS`,
`NOTIFICATION_WEBHOOK_URL`, and AWS SNS credentials. Set these via a local
`.env` file or your shell — do not commit real credentials.

### AI Service

```bash
cd AIservice
uvicorn main:app --reload --port 8000
```

### Frontend

```bash
cd gatorcan-frontend
npm install
npm run dev
```

## Project Documentation

Sprint notes and test logs are tracked in this repo:
[`Sprint1.md`](Sprint1.md), [`Sprint2.md`](Sprint2.md),
[`Sprint3.md`](Sprint3.md), [`Sprint4.md`](Sprint4.md),
[`UnitTestLog.md`](UnitTestLog.md).

## License

Released under the [MIT License](LICENSE).
