import type { RequestHandler } from "./$types";

export const POST: RequestHandler = async ({ cookies, request }) => {
  // Get the cookie domain from the request
  const host = request.headers.get("host");
  const domain = host?.includes(":") ? host.split(":")[0] : host;

  // Clear the authentication token with explicit domain matching
  cookies.delete("auth_token", {
    path: "/",
    domain: domain || undefined,
    httpOnly: true,
    secure: false, // Set to true if you're using HTTPS
    sameSite: "strict",
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
