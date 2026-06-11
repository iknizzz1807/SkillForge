import { redirect, json, type Handle } from "@sveltejs/kit";
import { BACKEND_URL } from '$lib/backend';

// List of routes that don't require authentication
const publicRoutes = ["/login", "/register", "/api/public", "/api/logout"];

// List of auth routes that authenticated users shouldn't access
const authRoutes = ["/login", "/register"];

// Interface for the user data extracted from JWT
interface UserData {
  id: string;
  name: string;
  email: string;
  role: string;
  isAuthenticated: boolean;
}

function safeRedirectPath(value: string | null, fallback = "/dashboard") {
  if (!value) return fallback;
  try {
    const decoded = decodeURIComponent(value);
    if (!decoded.startsWith("/") || decoded.startsWith("//")) return fallback;
    return decoded;
  } catch {
    return fallback;
  }
}

export const handle: Handle = async ({ event, resolve }) => {
  const { url, cookies } = event;
  const token = cookies.get("auth_token");
  const path = url.pathname;

  // Initialize locals with a default state
  event.locals = {
    token: null,
    user: null,
  };

  // Make token available in locals only after the backend verifies it.
  if (token) {
    const response = await fetch(BACKEND_URL + "/api/user", {
      headers: { Authorization: `Bearer ${token}` },
    }).catch(() => null);

    if (response?.ok) {
      const user = await response.json();
      event.locals = {
        token: token,
        user: {
          id: user.id || "",
          name: user.name || "",
          email: user.email || "",
          role: user.role || "",
          isAuthenticated: true,
        } satisfies UserData,
      };
    } else {
      cookies.delete("auth_token", { path: "/" });
    }
  }

  // CASE 1: Authenticated users trying to access auth routes (login/register)
  // Redirect them to dashboard instead
  if (event.locals.user && authRoutes.some((route) => path.startsWith(route))) {
    // Don't redirect if they're already being redirected somewhere else
    throw redirect(303, safeRedirectPath(url.searchParams.get("redirectTo")));
  }

  // CASE 2: Unauthenticated users trying to access protected routes
  // Redirect them to login
  if (
    !event.locals.user &&
    !publicRoutes.some((route) => path.startsWith(route))
  ) {
    if (path.startsWith("/api/")) {
      return json({ error: "Unauthorized" }, { status: 401 });
    }
    const redirectTo = encodeURIComponent(path);
    throw redirect(303, `/login?redirectTo=${redirectTo}`);
  }

  try {
    // Process the request
    const response = await resolve(event);

    // Handle auth failure from API responses
    if (response.status === 401 || response.status === 403) {
      cookies.delete("auth_token", { path: "/" });

      const contentType = response.headers.get("content-type") || "";
      if (contentType.includes("text/html")) {
        throw redirect(303, "/login");
      }
    }

    return response;
  } catch (error) {
    // Key part: Properly handle redirects that happen during the flow
    if (
      error &&
      typeof error === "object" &&
      "status" in error &&
      "location" in error
    ) {
      throw error; // Make sure redirects from page actions are properly propagated
    }

    if (error instanceof Response) {
      return error;
    }

    console.error("Hooks error:", error);
    return new Response("Internal Server Error", { status: 500 });
  }
};
