import { error, json } from "@sveltejs/kit";
import type { RequestHandler } from "./$types";

export const DELETE: RequestHandler = async ({ params, fetch, locals }) => {
  try {
    // Lấy student_id và project_id từ params URL
    const { project_id, student_id } = params;
    const token = locals.token;

    if (!project_id || !student_id) {
      throw error(400, "Project ID and Student ID are required");
    }

    if (!token) {
      throw error(401, "Unauthorized - Authentication required");
    }

    // Gọi đến API backend
    const response = await fetch(
      `http://backend:8080/api/projects/students/${student_id}/${project_id}`,
      {
        method: "DELETE",
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
        errorData.message || "Failed to remove student from project"
      );
    }

    return json({ success: true, message: "Student removed successfully" });
  } catch (err) {
    console.error("Error removing student from project:", err);

    // Xử lý lỗi
    if (err instanceof Error) {
      throw error(500, err.message);
    }

    throw error(500, "Failed to remove student from project");
  }
};
