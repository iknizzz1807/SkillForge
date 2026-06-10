import { test, expect } from "@playwright/test";

test.describe("Navigation and routing", () => {
  test("root path redirects to login", async ({ page }) => {
    await page.goto("/");
    await expect(page).toHaveURL(/\/login/);
  });

  test("dashboard redirects to login", async ({ page }) => {
    await page.goto("/dashboard");
    await expect(page).toHaveURL(/\/login/);
  });

  test("settings redirects to login", async ({ page }) => {
    await page.goto("/settings");
    await expect(page).toHaveURL(/\/login/);
  });

  test("account redirects to login", async ({ page }) => {
    await page.goto("/account");
    await expect(page).toHaveURL(/\/login/);
  });

  test("login page is accessible without auth", async ({ page }) => {
    await page.goto("/login");
    await expect(page.getByText("Login to SKILLFORGE")).toBeVisible();
  });

  test("register page is accessible without auth", async ({ page }) => {
    await page.goto("/register");
    await expect(page.getByText("Register to SKILLFORGE")).toBeVisible();
  });

  test("login page has forgot password link", async ({ page }) => {
    await page.goto("/login");
    await expect(page.getByText("Forgot Password?")).toBeVisible();
  });

  test("register page has avatar upload option", async ({ page }) => {
    await page.goto("/register");
    const avatarLabel = page.getByLabel("Upload avatar");
    await expect(avatarLabel).toBeVisible();
  });

  test("login page shows email input placeholder", async ({ page }) => {
    await page.goto("/login");
    const emailInput = page.getByPlaceholder("you@example.com");
    await expect(emailInput).toBeVisible();
  });

  test("login page shows password input placeholder", async ({ page }) => {
    await page.goto("/login");
    const passwordInput = page.getByPlaceholder("••••••••••••••••");
    await expect(passwordInput).toBeVisible();
  });
});
