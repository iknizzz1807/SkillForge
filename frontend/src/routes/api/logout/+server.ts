import { redirect, json } from "@sveltejs/kit";
import { dev } from "$app/environment";

export const POST = async ({ cookies, url, locals }) => {
  const token = locals.token;

  if (token) {
    try {
      await fetch("http://backend:8080/auth/logout", {
        method: "POST",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
    } catch (e) {
      console.error("Failed to invalidate session on backend:", e);
    }
  }

  cookies.delete("auth_token", {
    path: "/",
    secure: !dev,
    sameSite: "strict",
  });

  const redirectTo = url.searchParams.get("redirectTo") || "/login";

  if (url.searchParams.get("redirectTo")) {
    throw redirect(303, redirectTo);
  }

  return json({ success: true });
};
