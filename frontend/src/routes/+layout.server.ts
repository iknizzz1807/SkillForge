import type { LayoutServerLoad } from "./$types";

export const load = (async (event) => {
  return {
    role: event.locals.user?.role,
    userName: event.locals.user?.name,
  };
}) satisfies LayoutServerLoad;
