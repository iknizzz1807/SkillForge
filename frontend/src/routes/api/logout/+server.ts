import { redirect } from "@sveltejs/kit";

export const POST = async ({ cookies, url }) => {
  const domainProduction = "skillforge.ikniz.site";
  const domainDevelopment = "localhost";

  cookies.delete("auth_token", {
    path: "/",
    domain: domainProduction || "",
    secure: false,
    sameSite: "lax",
  });

  cookies.delete("auth_token", {
    path: "/",
    secure: false,
    sameSite: "lax",
  });

  cookies.delete("auth_token", {
    path: "/",
    domain: ".skillforge.ikniz.site",
    secure: false,
    sameSite: "lax",
  });

  cookies.delete("auth_token", {
    path: "/",
    domain: domainDevelopment || "",
  });

  const redirectTo = url.searchParams.get("redirectTo") || "/login";

  throw redirect(303, redirectTo);
};
