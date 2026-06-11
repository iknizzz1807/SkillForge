// Cái endpoint này là để cập nhật profile cơ bản trong /profile
import { json } from "@sveltejs/kit";
import type { RequestHandler } from "./$types";
import { BACKEND_URL } from '$lib/backend';

export const PUT: RequestHandler = async ({ request, fetch, cookies, locals }) => {
  try {
    const token = locals.token;

    if (!token) {
      return json({ error: "Unauthorized" }, { status: 401 });
    }

    // Forward the formdata to your Go backend
    const response = await fetch(BACKEND_URL + "/api/user", {
      method: "PUT",
      headers: {
        Authorization: `Bearer ${token}`,
      },
      body: await request.formData(),
    });

    if (!response.ok) {
      const errorData = await response.json();
      return json(
        { error: errorData.error || "Server error" },
        { status: response.status }
      );
    }

    // Parse the response
    const responseData = await response.json();

    // Update auth token in cookies if a new token was returned
    if (responseData.token) {
      cookies.set("auth_token", responseData.token, {
        path: "/",
        httpOnly: true,
        secure: process.env.NODE_ENV === "production",
        maxAge: 60 * 60 * 24 * 7, // 7 days
        sameSite: "strict",
      });
    }

    // Return the updated user data
    return json(responseData.user || responseData);
  } catch (error) {
    console.error("Profile update error:", error);
    return json({ error: "Failed to update profile" }, { status: 500 });
  }
};
