import { test, expect, type Page } from "@playwright/test";

const business = {
  email: "business@skillforge.demo",
  password: "Demo@123456",
  name: "Aurora Labs",
};

const student = {
  email: "an@skillforge.demo",
  password: "Demo@123456",
  name: "Nguyen Minh An",
};

const projectTitle = "AI Resume Screening Platform";

test.describe.configure({ mode: "serial" });

test("investor demo walkthrough", async ({ page }) => {
  test.setTimeout(10 * 60 * 1000);
  await page.setViewportSize({ width: 1440, height: 900 });

  await page.context().clearCookies();

  await scene(page, "SkillForge", "A project marketplace where businesses discover student talent through verified work.");
  await login(page, business.email, business.password);

  await scene(page, "Business Dashboard", "Aurora Labs sees active projects, applications, and team momentum at a glance.");
  await page.goto("/dashboard");
  await waitForApp(page);
  await pause(3500);

  await scene(page, "Project Portfolio", "Seeded projects show different stages: active, open, and completed.");
  await page.goto("/project");
  await waitForApp(page);
  await expect(page.getByText(projectTitle).first()).toBeVisible({ timeout: 15000 });
  await pause(4000);

  await scene(page, "Collaborative Workspace", "The business and students manage tasks together in a live project workspace.");
  await page.goto("/project/demo-project-ai-resume");
  await waitForApp(page);
  await expect(page.getByText(projectTitle).first()).toBeVisible({ timeout: 15000 });

  // Drag a task from "To-do" to "In Progress" to show kanban interaction
  const taskCard = page.locator('[data-title="Write integration test suite"]').first();
  const inProgressCol = page.locator('.kanban-column[data-status="in_progress"]').first();
  await taskCard.dragTo(inProgressCol, { force: true, timeout: 5000 }).catch(() => undefined);
  await pause(3500);

  await scene(page, "Realtime Chat", "Team messages show how collaboration and delivery evidence are captured inside the platform.");
  await page.goto("/chat");
  await waitForApp(page);
  await page.getByText(projectTitle).first().click({ timeout: 15000 }).catch(() => undefined);
  await pause(5500);

  await scene(page, "Applications Review", "Businesses can review applicants and convert strong candidates into project members.");
  await page.goto("/project/application");
  await waitForApp(page);
  await pause(5000);

  await scene(page, "Talent Search", "SkillForge gives businesses a searchable talent pool with skills, project history, and invitations.");
  await page.goto("/talent/search");
  await waitForApp(page);
  await page.getByText("Failed to load students").first().waitFor({ timeout: 1500 }).then(async () => {
    await page.reload();
    await waitForApp(page);
  }).catch(() => undefined);
  await pause(3500);

  await scene(page, "Switch To Student View", "Now we switch to a student account to show the other side of the marketplace.");
  await logout(page);
  await login(page, student.email, student.password);

  await scene(page, "Student Dashboard", "Students see suggested projects, progress, skills, and opportunities in one place.");
  await page.goto("/dashboard");
  await waitForApp(page);
  await pause(4500);

  await scene(page, "AI-Matched Marketplace", "The marketplace highlights projects matching the student's skills and goals.");
  await page.goto("/marketplace");
  await waitForApp(page);
  await expect(page.getByText(projectTitle).first()).toBeVisible({ timeout: 15000 });
  await expect(page.getByText("Suggested Projects").first()).toBeVisible({ timeout: 15000 });
  await pause(4500);

  await scene(page, "Project Detail", "Students can inspect scope, required skills, team members, and application state before applying.");
  await page.goto("/marketplace/demo-project-ai-resume");
  await waitForApp(page);
  await expect(page.getByText(projectTitle).first()).toBeVisible({ timeout: 15000 });
  await pause(5500);

  await scene(page, "Student Profile", "Profiles combine skills, badges, XP, feedback, and portfolio generation into verified proof of ability.");
  await page.goto("/profile");
  await waitForApp(page);
  await expect(page.getByText(/Student Profile|Nguyen Minh An/).first()).toBeVisible({ timeout: 15000 });
  await pause(6500);

  await scene(page, "SkillForge Demo Complete", "Businesses get execution-ready talent. Students build proof through real work.");
  await pause(5000);
});

async function login(page: Page, email: string, password: string) {
  await scene(page, "Login", `Signing in as ${email}`);
  await page.goto("/login");
  await waitForApp(page);
  await page.getByLabel(/Email/i).fill(email);
  await page.getByLabel(/Password/i).fill(password);
  await page.getByRole("button", { name: /^login$/i }).click();
  await page.waitForURL(/dashboard|project|marketplace|profile|chat|talent|invitations/, { timeout: 20000 }).catch(() => undefined);
  await waitForApp(page);
  await pause(2000);
}

async function logout(page: Page) {
  await page.getByRole("button", { name: /logout/i }).click({ timeout: 10000 }).catch(async () => {
    await page.goto("/api/logout?redirectTo=/login");
  });
  await page.waitForURL(/login/, { timeout: 15000 }).catch(() => undefined);
  await pause(1500);
}

async function waitForApp(page: Page) {
  await page.waitForLoadState("domcontentloaded");
  await page.waitForLoadState("networkidle", { timeout: 10000 }).catch(() => undefined);
}

async function scene(page: Page, title: string, subtitle: string) {
  await page.evaluate(
    ({ title, subtitle }) => {
      const existing = document.getElementById("skillforge-demo-caption");
      if (existing) existing.remove();

      const el = document.createElement("div");
      el.id = "skillforge-demo-caption";
      el.innerHTML = `<div style="font-weight:800;font-size:30px;margin-bottom:6px">${title}</div><div style="font-size:17px;line-height:1.45;opacity:.95">${subtitle}</div>`;
      el.style.cssText = [
        "position:fixed",
        "left:28px",
        "bottom:28px",
        "z-index:2147483647",
        "max-width:720px",
        "padding:18px 22px",
        "border-radius:18px",
        "background:linear-gradient(135deg,rgba(17,24,39,.94),rgba(76,29,149,.92))",
        "color:white",
        "box-shadow:0 22px 70px rgba(0,0,0,.35)",
        "font-family:Inter,ui-sans-serif,system-ui,-apple-system,BlinkMacSystemFont,'Segoe UI',sans-serif",
        "pointer-events:none",
      ].join(";");
      document.body.appendChild(el);
    },
    { title, subtitle }
  );
  await pause(2200);
}

function pause(ms: number) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}
