import type { PageServerLoad } from "./$types";

export const load = (async (event) => {
  const user = event.locals.user;
  const origin = event.url.origin;

  return {
    id: user?.id,
    name: user?.name,
    email: user?.email,
    role: user?.role,
    // Add the avatar URL using the same pattern as in layout
    avatarUrl: user?.id ? `${origin}/api/avatars/${user.id}` : null,
  };
}) satisfies PageServerLoad;
