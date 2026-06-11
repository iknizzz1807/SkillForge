import { error, json } from "@sveltejs/kit";
import type { RequestHandler } from "./$types";
import { BACKEND_URL } from '$lib/backend';

export const GET: RequestHandler = async ({ params, fetch, locals }) => {
  try {
    // Lấy project_id từ params URL
    const projectId = params.project_id;
    const token = locals.token;

    if (!projectId) {
      throw error(400, "Project ID is required");
    }

    if (!token) {
      throw error(401, "Unauthorized - Authentication required");
    }

    // Gọi đến API backend
    const response = await fetch(
      `${BACKEND_URL}/api/projects/students/${projectId}`,
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      }
    );

    // Kiểm tra response
    if (!response.ok) {
      const errorData = await response.json();
      throw error(
        response.status,
        errorData.message || "Failed to fetch students for this project"
      );
    }

    // Chuyển đổi response thành JSON và trả về
    const students = await response.json();
    return json(students);
  } catch (err) {
    console.error("Error fetching project students:", err);

    // Xử lý lỗi
    if (err instanceof Error) {
      throw error(500, err.message);
    }

    throw error(500, "Failed to fetch students for this project");
  }
};
