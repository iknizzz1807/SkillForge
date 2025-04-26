import type { PageServerLoad, Actions } from "./$types";
import { fail, redirect } from "@sveltejs/kit";

export const actions = {
  default: async ({ request }) => {
    const formData = await request.formData();
    const email = formData.get("email") as string;
    const password = formData.get("password") as string;
    const confirmPassword = formData.get("confirmPassword") as string;
    const name = formData.get("name") as string;
    const role = formData.get("role") as string;
    const website = (formData.get("website") as string) || "";
    const title = formData.get("title") as string;
    const avatar = formData.get("avatar") as File | null;

    // Validate form data
    if (!email || !password || !confirmPassword || !name || !role || !title) {
      return fail(400, {
        error: true,
        message: "All fields are required",
        formData: { email, name, website, role, title },
      });
    }

    if (password !== confirmPassword) {
      return fail(400, {
        error: true,
        message: "Passwords do not match",
        formData: { email, name, website, role, title },
      });
    }

    try {
      // Create a new FormData to send to the backend
      const apiFormData = new FormData();

      // Add all the text fields
      apiFormData.append("email", email);
      apiFormData.append("password", password);
      apiFormData.append("name", name);
      apiFormData.append("role", role);
      apiFormData.append("title", title);

      // Add website field if it's a business account
      if (role === "business" && website) {
        apiFormData.append("website", website);
      }

      // Add avatar if it exists
      if (avatar && avatar.size > 0) {
        apiFormData.append("avatar", avatar);
      }

      // Call the registration endpoint with FormData
      const response = await fetch("http://backend:8080/auth/register", {
        method: "POST",
        body: apiFormData, // No content-type header - browser sets it automatically with boundary
      });

      if (!response.ok) {
        const errorData = await response.json();
        return fail(response.status, {
          error: true,
          message: errorData.error || "Registration failed",
          formData: { email, name, website, role, title },
        });
      }

      return { success: true };
    } catch (err) {
      if (err instanceof Response) throw err; // Rethrow redirect

      console.error("Registration error:", err);
      return fail(500, {
        error: true,
        message: "An error occurred during registration",
        formData: { email, name, website, role, title },
      });
    }
  },
} satisfies Actions;
