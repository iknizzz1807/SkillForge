import { redirect } from "@sveltejs/kit";
import type { Actions, PageServerLoad } from "./$types";
// import { type Cookies, fail } from "@sveltejs/kit";

export const load = (async ({ parent, cookies }) => {
  const parentData = await parent();
  const token = cookies.get("auth_token");
  if (parentData.role !== "business") throw redirect(302, "/project");
  return {
    token: token,
  };
}) satisfies PageServerLoad;
