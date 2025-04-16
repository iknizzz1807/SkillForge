import type { PageServerLoad } from "./$types";

export const load = (async (event) => {
  // Make a get request to the backend to fetch all the applications base on the role
  // This is different when it is student or it is business
  return {
    name: event.locals.user?.name,
    email: event.locals.user?.email,
    role: event.locals.user?.role,
  };
}) satisfies PageServerLoad;
