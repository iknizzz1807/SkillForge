import { test, expect } from "@playwright/test";

test.describe("Chat functionality", () => {
  test("chat page redirects to login when not authenticated", async ({ page }) => {
    await page.goto("/chat");
    await expect(page).toHaveURL(/\/login/);
  });

  test("chat page shows redirect param", async ({ page }) => {
    await page.goto("/chat");
    const url = page.url();
    expect(url).toContain("redirectTo");
  });

  test("chat page requires authentication", async ({ page }) => {
    await page.goto("/chat");
    await expect(page.locator("h2")).toContainText(/login/i);
  });

  test("user can send message", async ({ page }) => {
    // Skip - requires authenticated session + project membership
    test.skip(true, "Requires authenticated session with chat access");
  });

  test("messages appear in chat", async ({ page }) => {
    // Skip - requires authenticated session with existing messages
    test.skip(true, "Requires authenticated session with messages");
  });
});
