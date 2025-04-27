import type { PageServerLoad } from "./$types";
import { redirect } from "@sveltejs/kit";

// Cái này là get thông tin của user khác để phục vụ cho trang profile của người dùng khác
// Tuy nhiên nếu current_user_id trùng với target_user_id thì redirect về /profile
export const load = (async ({ params, locals, fetch }) => {
  const current_user_id = locals.user?.id;
  const target_user_id = params.id;
  const token = locals.token;

  // Nếu người dùng đang cố xem profile của chính mình, redirect về /profile
  if (current_user_id === target_user_id) {
    throw redirect(302, "/profile");
  }

  // Gửi request đến backend với token xác thực để lấy thông tin user
  const userResponse = await fetch(
    `http://backend:8080/api/user/${target_user_id}`,
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );

  if (!userResponse.ok) {
    const errorData = await userResponse.json();
    throw new Error(errorData.error || "Failed to fetch user profile");
  }

  const userData = await userResponse.json();

  // Gửi thêm request để lấy thông tin business
  const businessResponse = await fetch(
    `http://backend:8080/api/business-info`,
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );

  let businessData = null;
  if (businessResponse.ok) {
    businessData = await businessResponse.json();
  }

  return {
    user: userData,
    businessInfo: businessData,
    isOwnProfile: false,
  };
}) satisfies PageServerLoad;
