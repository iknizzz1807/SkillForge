import type { LayoutServerLoad } from "./$types";
import { redirect } from "@sveltejs/kit";

export const load = (async (event) => {
  // Kiểm tra người dùng đã đăng nhập chưa
  const user = event.locals.user;

  // Lấy đường dẫn hiện tại
  const path = event.url.pathname;

  // Tạo base URL dựa trên origin của request
  const origin = event.url.origin;

  // Chuẩn bị dữ liệu phản hồi - luôn trả về dữ liệu này
  const userData = {
    role: user?.role,
    id: user?.id,
    userName: user?.name,
    // Now using the SvelteKit API endpoint instead of direct backend URL
    avatarUrl: user?.id ? `${origin}/api/avatars/${user.id}` : null,
  };

  // Chỉ chuyển hướng nếu đáp ứng điều kiện cụ thể
  if (user && path === "/") {
    // Không sử dụng throw redirect để không cắt luồng xử lý
    return {
      ...userData,
      // redirect: {
      //   status: 302,
      //   location: "/dashboard",
      // },
    };
  }

  // Trả về dữ liệu người dùng mà không chuyển hướng
  return userData;
}) satisfies LayoutServerLoad;
