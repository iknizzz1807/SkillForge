import { test, expect } from "@playwright/test";

test.describe("Profile flows", () => {
  test("profile page redirects to login when not authenticated", async ({ page }) => {
    await page.goto("/profile");
    await expect(page).toHaveURL(/\/login/);
  });

  test("profile page shows login form", async ({ page }) => {
    await page.goto("/profile");
    await expect(page.getByText("Login to SKILLFORGE")).toBeVisible();
  });

  test("profile redirects with correct redirectTo parameter", async ({ page }) => {
    await page.goto("/profile");
    const url = page.url();
    expect(url).toContain("redirectTo");
    expect(url).toContain("profile");
  });

  test("user can view profile when authenticated", async ({ page }) => {
    // Skip - requires authenticated session
    test.skip(true, "Requires authenticated session");
  });

  test("user can update profile", async ({ page }) => {
    // Skip - requires authenticated session
    test.skip(true, "Requires authenticated session");
  });

  test("profile changes persist after refresh", async ({ page }) => {
    // Skip - requires authenticated session
    test.skip(true, "Requires authenticated session");
  });
});
