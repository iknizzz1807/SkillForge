import type { LayoutServerLoad } from "./$types";
import { redirect } from "@sveltejs/kit";

export const load = (async (event) => {
  // Kiểm tra người dùng đã đăng nhập chưa
  const user = event.locals.user;

  // Lấy đường dẫn hiện tại
  const path = event.url.pathname;

  // Chuẩn bị dữ liệu phản hồi - luôn trả về dữ liệu này
  const userData = {
    role: user?.role,
    userName: user?.name,
  };

  // Chỉ chuyển hướng nếu đáp ứng điều kiện cụ thể
  if (user && path === "/") {
    // Không sử dụng throw redirect để không cắt luồng xử lý
    return {
      ...userData,
      redirect: {
        status: 302,
        location: "/dashboard",
      },
    };
  }

  // Hoặc kiểm tra role để điều hướng đến dashboard tương ứng
  // if (user?.role === 'admin' && path === '/') {
  //   return {
  //     ...userData,
  //     redirect: {
  //       status: 302,
  //       location: "/admin/dashboard"
  //     }
  //   };
  // } else if (user?.role === 'user' && path === '/') {
  //   return {
  //     ...userData,
  //     redirect: {
  //       status: 302,
  //       location: "/user/dashboard"
  //     }
  //   };
  // }

  // Trả về dữ liệu người dùng mà không chuyển hướng
  return userData;
}) satisfies LayoutServerLoad;
