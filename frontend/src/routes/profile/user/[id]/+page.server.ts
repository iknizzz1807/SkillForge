import type { PageServerLoad } from "./$types";
import { redirect } from "@sveltejs/kit";

export const load = (async ({ params, locals, fetch }) => {
  const current_user_id = locals.user?.id;
  const target_user_id = params.id;
  const token = locals.token;

  if (current_user_id === target_user_id) {
    throw redirect(302, "/profile");
  }

  let userName = "";
  let userTitle = "";
  let userAvatar = "";

  try {
    const response = await fetch(
      `http://backend:8080/api/users/${target_user_id}/profile`,
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );
    if (response.ok) {
      const data = await response.json();
      const u = data.user || {};
      userName = u.name || "";
      userTitle = u.title || "";
      userAvatar = u.avatar_name || "";
    }
  } catch (error) {
    console.error("Failed to fetch user profile:", error);
  }

  return {
    id: target_user_id,
    name: userName,
    title: userTitle,
    avatarUrl: userAvatar,
  };
}) satisfies PageServerLoad;
