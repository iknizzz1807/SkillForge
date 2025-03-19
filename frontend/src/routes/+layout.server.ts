import type { Url } from "url";
import type { LayoutServerLoad } from "./$types";

export const load = (async (event) => {
  return {
    userName: event.locals.user?.name,
  };
}) satisfies LayoutServerLoad;
