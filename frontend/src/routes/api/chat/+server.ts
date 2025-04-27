import { json } from "@sveltejs/kit";

/**
 * Retrieves all chat groups, list of chat names and ids
 */
export async function GET({ fetch, locals }) {
  try {
    const token = locals.token;

    const response = await fetch("http://backend:8080/api/chats", {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      return json(
        { error: "Failed to fetch chat groups" },
        { status: response.status }
      );
    }

    const chatGroups = await response.json();
    return json(chatGroups);
  } catch (error) {
    console.error("Error fetching chat groups:", error);
    return json({ error: "Internal server error" }, { status: 500 });
  }
}
