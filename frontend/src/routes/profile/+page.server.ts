import type { PageServerLoad } from "./$types";

export const load = (async (event) => {
  return {
    name: event.locals.user?.name,
    email: event.locals.user?.email,
    role: event.locals.user?.role,
  };
}) satisfies PageServerLoad;
