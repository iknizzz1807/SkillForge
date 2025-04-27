import { json } from "@sveltejs/kit";

// Get one chat group, return messages and list of users in the group
export async function GET({ fetch, locals, params }) {
  try {
    const token = locals.token;
    const id = params.id;

    const response = await fetch("http://backend:8080/api/chats/" + id, {
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
