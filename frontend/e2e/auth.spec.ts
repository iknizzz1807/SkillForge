import { test, expect } from "@playwright/test";

test.describe("Authentication flows", () => {
  test("user can see login page", async ({ page }) => {
    await page.goto("/login");
    await expect(page.getByText("Login to SKILLFORGE")).toBeVisible();
    await expect(page.getByLabel("Email")).toBeVisible();
    await expect(page.getByLabel("Password")).toBeVisible();
  });

  test("user can see register page", async ({ page }) => {
    await page.goto("/register");
    await expect(page.getByText("Register to SKILLFORGE")).toBeVisible();
    await expect(page.getByLabel(/Full Name/)).toBeVisible();
    await expect(page.getByLabel("Email")).toBeVisible();
  });

  test("user can navigate from login to register", async ({ page }) => {
    await page.goto("/login");
    await page.getByText("Sign up").click();
    await expect(page).toHaveURL("/register");
  });

  test("user can navigate from register to login", async ({ page }) => {
    await page.goto("/register");
    await page.getByText("Login").click();
    await expect(page).toHaveURL("/login");
  });

  test("login form prevents empty submission", async ({ page }) => {
    await page.goto("/login");
    const submitBtn = page.getByRole("button", { name: /login/i });
    // HTML5 validation should prevent submission
    const emailInput = page.getByLabel("Email");
    await expect(emailInput).toHaveAttribute("required", "");
  });

  test("user cannot login with wrong password", async ({ page }) => {
    await page.goto("/login");
    await page.getByLabel("Email").fill("test@example.com");
    await page.getByLabel("Password").fill("wrongpassword");
    await page.getByRole("button", { name: /login/i }).click();
    // Should show error or stay on login page
    await expect(page).toHaveURL(/\/login/);
  });

  test("user can register as student", async ({ page }) => {
    await page.goto("/register");
    await page.getByLabel(/Full Name/).fill("Test Student");
    await page.getByLabel("Email").fill(`student_${Date.now()}@test.com`);
    await page.getByLabel("Student").click();
    await page.getByLabel("Password").fill("TestPass123!");
    await page.getByLabel("Confirm Password").fill("TestPass123!");
    // Select title after role is chosen
    await page.getByLabel(/Professional Title/).selectOption("Software Engineer");
    await page.getByRole("button", { name: /create account/i }).click();
  });

  test("user can register as business", async ({ page }) => {
    await page.goto("/register");
    await page.getByLabel(/Full Name/).fill("Test Business");
    await page.getByLabel("Email").fill(`biz_${Date.now()}@test.com`);
    await page.getByLabel("Business").click();
    await page.getByLabel("Password").fill("TestPass123!");
    await page.getByLabel("Confirm Password").fill("TestPass123!");
    await page.getByLabel(/Professional Title/).selectOption("CTO");
    await page.getByRole("button", { name: /create account/i }).click();
  });

  test("user can logout", async ({ page }) => {
    // Skip - requires being logged in first
    test.skip(true, "Requires authenticated session");
  });
});
