import type { PageServerLoad, Actions } from "./$types";
import { fail, type Cookies, redirect } from "@sveltejs/kit";

export const load = (async ({ url, cookies }) => {
  // Check if user is already logged in
  const token = cookies.get("auth_token");
  if (token) {
    throw redirect(303, "/dashboard");
  }

  // Get the redirectTo parameter if it exists
  const redirectTo = url.searchParams.get("redirectTo") || "/dashboard";

  return {
    redirectTo,
  };
}) satisfies PageServerLoad;

export const actions = {
  default: async ({
    request,
    cookies,
  }: {
    request: Request;
    cookies: Cookies;
  }) => {
    const data = await request.formData();
    const email = data.get("email");
    const password = data.get("password");
    const redirectTo = data.get("redirectTo") || "/dashboard";

    if (!email || !password) {
      return fail(400, {
        error: true,
        message: "Email and password are required",
        email: email,
        redirectTo,
      });
    }

    try {
      // Call the login endpoint
      const response = await fetch("http://localhost:8080/auth/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email, password }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        return fail(response.status, {
          error: true,
          message: errorData.error || "Login failed",
          email: email,
          redirectTo,
        });
      }

      const { token } = await response.json();

      // Store token in an HTTP-only cookie
      cookies.set("auth_token", token, {
        path: "/",
        httpOnly: true,
        secure: process.env.NODE_ENV === "production",
        maxAge: 60 * 60 * 24, // 1 day
        sameSite: "strict",
      });

      // Redirect to the requested page or dashboard
      throw redirect(303, redirectTo.toString());
    } catch (err) {
      console.error("Login error:", err);
      return fail(500, {
        error: true,
        message: "An error occurred during login",
        email: email,
        redirectTo,
      });
    }
  },
} satisfies Actions;
