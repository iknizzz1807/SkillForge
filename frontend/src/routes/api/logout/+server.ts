import type { RequestHandler } from "./$types";

export const POST: RequestHandler = async ({ cookies, request }) => {
  // Set the auth cookie to empty with explicit domain
  cookies.set("auth_token", "", {
    path: "/",
    // domain: "143.244.147.141",
    httpOnly: true,
    secure: false, // Correct since you're using HTTP
    sameSite: "lax", // Using lax for better compatibility
    maxAge: 0, // Immediately expire the cookie
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
