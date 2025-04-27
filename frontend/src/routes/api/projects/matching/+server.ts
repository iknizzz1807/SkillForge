import { json } from "@sveltejs/kit";
import type { RequestHandler } from "./$types";

export const GET: RequestHandler = async ({ fetch, locals }) => {
  // Lấy token từ locals
  const token = locals.token;

  try {
    // Gọi trực tiếp đến API backend
    const response = await fetch("http://backend:8080/api/matches", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`, // Thêm token vào header
      },
    });

    if (!response.ok) {
      console.error(`Backend responded with ${response.status}`);
      return json([], { status: response.status });
    }

    // Lấy dữ liệu từ response và trả về trực tiếp
    const matchedProjects = await response.json();
    return json(matchedProjects);
  } catch (error) {
    console.error("Error fetching project matches:", error);
    return json([], { status: 500 });
  }
};
