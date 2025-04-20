import type { PageServerLoad, Actions } from "./$types";
import { fail, redirect } from "@sveltejs/kit";

export const load = (async ({ url, cookies }) => {
  // Check if user is already logged in
  const token = cookies.get("auth_token");
  if (token) {
    throw redirect(303, "/dashboard");
  }

  return {};
}) satisfies PageServerLoad;

export const actions = {
  default: async ({ request }) => {
    const data = await request.formData();
    const email = data.get("email");
    const password = data.get("password");
    const confirmPassword = data.get("confirmPassword");
    const name = data.get("name");
    const website = data.get("website");
    const role = data.get("role");

    // Validate form data
    if (!email || !password || !confirmPassword || !name) {
      return fail(400, {
        error: true,
        message: "All fields are required",
        formData: {
          email,
          name,
          website,
        },
      });
    }

    if (password !== confirmPassword) {
      return fail(400, {
        error: true,
        message: "Passwords do not match",
        formData: {
          email,
          name,
          website,
        },
      });
    }

    try {
      //   Call the registration endpoint
      const response = await fetch("http://backend:8080/auth/register", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email, password, name, role }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        return fail(response.status, {
          error: true,
          message: errorData.error || "Registration failed",
          formData: {
            email,
            name,
            website,
          },
        });
      }

      //   // Registration successful, redirect to login page
      //   //   throw redirect(303, "/login?registered=true");
      return {
        success: true,
      };
    } catch (err) {
      if (err instanceof Response) throw err; // Rethrow redirect

      console.error("Registration error:", err);
      return fail(500, {
        error: true,
        message: "An error occurred during registration",
        formData: {
          email,
          name,
          website,
        },
      });
    }
  },
} satisfies Actions;
