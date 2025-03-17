import { redirect, type Handle } from "@sveltejs/kit";

// List of routes that don't require authentication
const publicRoutes = ["/login", "/register"];

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

  // Make token available in locals for server routes
  if (token) {
    // Decode the token to extract user information
    const userData = decodeJWT(token);
    if (userData) {
      event.locals = {
        token: token,
        user: userData,
      };
    } else {
      // If token cannot be decoded, consider it invalid
      event.locals = {
        token: null,
        user: null,
      };
      cookies.delete("auth_token", { path: "/" });
    }
  } else {
    event.locals = {
      token: null,
      user: null,
    };
  }

  // console.log(event.locals);

  // Check path
  const path = url.pathname;

  // CASE 1: Authenticated users trying to access auth routes (login/register)
  // Redirect them to dashboard instead
  if (token && authRoutes.some((route) => path.startsWith(route))) {
    throw redirect(303, "/dashboard");
  }

  // CASE 2: Unauthenticated users trying to access protected routes
  // Redirect them to login
  if (!token && !publicRoutes.some((route) => path.startsWith(route))) {
    const redirectTo = encodeURIComponent(url.pathname);
    throw redirect(303, `/login?redirectTo=${redirectTo}`);
  }

  // For API calls that need authentication, add the token to the headers
  try {
    // Handle the request and catch any errors that might occur
    const response = await resolve(event, {
      transformPageChunk: ({ html }) => html,
      filterSerializedResponseHeaders: (name) => name === "content-type",
    });

    // Check if the response is an API response that indicates auth failure
    if (response.status === 401 || response.status === 403) {
      // If this is an API call and the backend rejected our auth,
      // clear the invalid token and redirect to login
      cookies.delete("auth_token", { path: "/" });

      // Only redirect to login for HTML responses, not for API calls
      const isHtmlResponse = response.headers
        .get("content-type")
        ?.includes("text/html");
      if (isHtmlResponse) {
        throw redirect(303, "/login");
      }
    }

    return response;
  } catch (error) {
    // Handle error
    if (error instanceof Response) {
      return error;
    }

    console.error("Hooks error:", error);
    // If the error is from a redirect, allow it to propagate
    if (error instanceof Error && "status" in error && "location" in error) {
      throw error;
    }

    // For other errors, return a generic error response
    return new Response("Internal Server Error", { status: 500 });
  }
};

// Helper for API fetch calls to add the auth token automatically
export async function fetchWithAuth(
  fetch: typeof globalThis.fetch,
  url: string,
  options: RequestInit = {}
) {
  // Get the token from cookies on the client side
  const token = document.cookie
    .split(";")
    .find((c) => c.trim().startsWith("auth_token="))
    ?.split("=")[1];

  // Add the authorization header if token exists
  const headers = {
    ...options.headers,
    "Content-Type": "application/json",
    ...(token ? { Authorization: `Bearer ${token}` } : {}),
  };

  // Make the API call
  const response = await fetch(url, {
    ...options,
    headers,
  });

  // Handle auth errors
  if (response.status === 401 || response.status === 403) {
    // Clear the auth cookie if it's invalid
    document.cookie =
      "auth_token=; Path=/; Expires=Thu, 01 Jan 1970 00:00:00 GMT";

    // Redirect to login
    window.location.href = "/login";
    return new Response(null, { status: 401 });
  }

  return response;
}
