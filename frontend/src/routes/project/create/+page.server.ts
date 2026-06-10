import { redirect } from "@sveltejs/kit";
import type { Actions, PageServerLoad } from "./$types";
// import { type Cookies, fail } from "@sveltejs/kit";

export const load = (async ({ parent }) => {
  const parentData = await parent();
  if (parentData.role !== "business") throw redirect(302, "/project");
  return {};
}) satisfies PageServerLoad;
