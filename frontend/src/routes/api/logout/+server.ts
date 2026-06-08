import { redirect } from "@sveltejs/kit";
import { dev } from "$app/environment";

export const POST = async ({ cookies, url }) => {
  cookies.delete("auth_token", {
    path: "/",
    secure: !dev,
    sameSite: "strict",
  });

  const redirectTo = url.searchParams.get("redirectTo") || "/login";

  throw redirect(303, redirectTo);
};
