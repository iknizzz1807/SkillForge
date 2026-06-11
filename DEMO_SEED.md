# SkillForge Demo Seed

Seed rich local demo data for investor videos: businesses, students, projects, applications, invitations, chat, tasks, badges, XP, feedback, notifications, talent pool, and portfolios.

## Run With Docker MongoDB

Start MongoDB first:

```bash
docker compose up -d mongodb
```

If portfolio generation prints `permission denied`, fix the local storage ownership once:

```bash
sudo chown -R "$USER:$USER" backend/storage
```

Seed the local database from the host:

```bash
cd backend
JWT_SECRET=local-dev-secret-change-in-production \
MONGO_URI=mongodb://localhost:27017 \
go run ./cmd/seed_demo -reset
```

Then start the full app:

```bash
docker compose up --build
```

## Demo Accounts

All accounts use:

```text
Demo@123456
```

Business accounts:

```text
business@skillforge.demo  Aurora Labs
nova@skillforge.demo      Nova Commerce
```

Student accounts:

```text
an@skillforge.demo        Nguyen Minh An
linh@skillforge.demo      Tran Gia Linh
khoa@skillforge.demo      Pham Dang Khoa
mai@skillforge.demo       Le Hoang Mai
quang@skillforge.demo     Do Quang Huy
```

## Suggested Investor Demo Flow

1. Login as `business@skillforge.demo`.
2. Show business dashboard, project list, applications, talent pool, and invitations.
3. Open `AI Resume Screening Platform` and show task board + chat.
4. Logout, login as `an@skillforge.demo`.
5. Show marketplace, project matching, profile, badges, XP, and portfolio.
6. Show invitation/application states and realtime collaboration story.

## Notes

- `-reset` drops demo collections before seeding.
- Portfolio HTML files are written under `backend/storage/portfolios/`.
- The seed is safe to rerun; documents use stable demo IDs and are upserted.
