import { redirect, type Handle } from "@sveltejs/kit";

// List of routes that don't require authentication
const publicRoutes = ["/login", "/register", "/api/public"];

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

// Function to decode JWT without verification (verification happens on the backend)
function decodeJWT(token: string): UserData | null {
  try {
    // Split the token into parts
    const parts = token.split(".");
    if (parts.length !== 3) return null;

    // Decode the payload (middle part)
    const payload = JSON.parse(Buffer.from(parts[1], "base64").toString());

    return {
      id: payload.user_id || "",
      name: payload.name || "",
      email: payload.email || "",
      role: payload.role || "",
      isAuthenticated: true,
    };
  } catch (error) {
    console.error("Error decoding JWT:", error);
    return null;
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

  // Make token available in locals for server routes if it exists and is valid
  if (token) {
    const userData = decodeJWT(token);
    if (userData) {
      event.locals = {
        token: token,
        user: userData,
      };
    } else {
      // Token is invalid, clear it
      cookies.delete("auth_token", { path: "/" });
    }
  }

  // CASE 1: Authenticated users trying to access auth routes (login/register)
  // Redirect them to dashboard instead
  if (event.locals.user && authRoutes.some((route) => path.startsWith(route))) {
    // Don't redirect if they're already being redirected somewhere else
    const redirectParam = url.searchParams.get("redirectTo");
    if (redirectParam) {
      throw redirect(303, decodeURIComponent(redirectParam));
    }
    throw redirect(303, "/dashboard");
  }

  // CASE 2: Unauthenticated users trying to access protected routes
  // Redirect them to login
  if (
    !event.locals.user &&
    !publicRoutes.some((route) => path.startsWith(route))
  ) {
    const redirectTo = encodeURIComponent(path);
    throw redirect(303, `/login?redirectTo=${redirectTo}`);
  }

  try {
    // Process the request
    const response = await resolve(event);

    // Handle auth failure from API responses
    if (response.status === 401 || response.status === 403) {
      // cookies.delete("auth_token", { path: "/" });

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
