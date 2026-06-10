import type { RequestHandler } from "./$types";

export const GET: RequestHandler = async ({ fetch, locals, url }) => {
  const token = locals.token;
  const userId = url.searchParams.get("user_id");

  if (!userId) {
    return new Response(null, {
      status: 302,
      headers: {
        Location: "/images/default-avatar.png",
      },
    });
  }

  const backendUrl = `http://backend:8080/avatars/${encodeURIComponent(userId)}`;

  try {
    const headers: Record<string, string> = {};
    if (token) {
      headers["Authorization"] = `Bearer ${token}`;
    }

    const response = await fetch(backendUrl, { headers });

    if (!response.ok) {
      return new Response(null, {
        status: 302,
        headers: {
          Location: "/images/default-avatar.png",
        },
      });
    }

    const imageData = await response.arrayBuffer();
    const contentType = response.headers.get("content-type") || "image/jpeg";

    return new Response(imageData, {
      headers: {
        "Content-Type": contentType,
        "Cache-Control": "no-cache",
      },
    });
  } catch (err) {
    console.error("Error fetching avatar:", err);
    return new Response(null, {
      status: 302,
      headers: {
        Location: "/images/default-avatar.png",
      },
    });
  }
};
