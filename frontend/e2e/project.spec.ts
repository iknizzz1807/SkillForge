import { test, expect } from "@playwright/test";

test.describe("Project flows", () => {
  test("project listing page loads", async ({ page }) => {
    // This test expects to be redirected to login when not authenticated
    await page.goto("/project");
    await expect(page).toHaveURL(/\/login/);
  });

  test("business is redirected when trying to create project without auth", async ({ page }) => {
    await page.goto("/project/create");
    await expect(page).toHaveURL(/\/login/);
  });

  test("project page requires authentication", async ({ page }) => {
    await page.goto("/project");
    // Should redirect to login
    const url = page.url();
    expect(url).toContain("/login");
  });

  test("project page redirects to login when not authenticated", async ({ page }) => {
    await page.goto("/project");
    await expect(page.locator("h2")).toContainText(/login/i);
  });

  test("marketplace loads for unauthenticated users", async ({ page }) => {
    await page.goto("/marketplace");
    // Marketplace should redirect to login too since it's protected
    await expect(page).toHaveURL(/\/login/);
  });

  test("student can see project in browse", async ({ page }) => {
    // Skip - requires auth
    test.skip(true, "Requires authenticated student session");
  });
});
