import type { PageServerLoad } from "./$types";
import { redirect } from "@sveltejs/kit";

// Cái này là get thông tin của user khác để phục vụ cho trang profile của người dùng khác
// Tuy nhiên nếu current_user_id trùng với target_user_id thì redirect về /profile
export const load = (async ({ params, locals, fetch }) => {
  const current_user_id = locals.user?.id;
  const target_user_id = params.id;

  if (current_user_id === target_user_id) {
    throw redirect(302, "/profile");
  } else throw redirect(302, "/dashboard");
}) satisfies PageServerLoad;
