import type { Cookies } from "@sveltejs/kit";
import type { RequestHandler } from "../$types";
import type { Url } from "url";
import { redirect } from "@sveltejs/kit";

// logout handler
export const POST: RequestHandler = async ({
  cookies,
  url,
}: {
  cookies: Cookies;
  url: Url;
}) => {
  // Clear the authentication token
  cookies.delete("auth_token", { path: "/" });
  // Get the redirectTo parameter if it exists (optional)
  // const redirectTo = url.searchParams.get("redirectTo") || "/login";
  const redirectTo = "/login";

  // Redirect to login page
  throw redirect(303, redirectTo);
  // return new Response();
};
