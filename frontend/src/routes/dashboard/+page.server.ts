import type { PageServerLoad } from "./$types";

export const load = (async ({ locals }) => {
  const user = locals.user;
  return {
    user,
  };
}) satisfies PageServerLoad;
