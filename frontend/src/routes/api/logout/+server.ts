import type { RequestHandler } from "./$types";

export const POST: RequestHandler = async ({ cookies, request }) => {
  // Get the cookie domain from the request
  const host = request.headers.get("host");
  // Clear the authentication token with explicit domain matching
  cookies.delete("auth_token", {
    path: "/",
    httpOnly: true,
    secure: false, // Correct since you're using HTTP
    sameSite: "strict", // Try "lax" instead of "strict" for better compatibility
  });

  // Return a proper redirect response
  return new Response(null, {
    status: 303,
    headers: {
      Location: "/login",
      // Clear-Site-Data is a powerful way to ensure logout
      "Clear-Site-Data": '"cookies"',
    },
  });
};
