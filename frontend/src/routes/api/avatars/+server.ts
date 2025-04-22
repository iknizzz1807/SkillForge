import type { RequestHandler } from "./$types";

export const GET: RequestHandler = async ({ fetch, locals }) => {
  const token = locals.token; // Get auth token from locals

  try {
    // Add Authorization header with token if available
    const headers: Record<string, string> = {};
    if (token) {
      headers["Authorization"] = `Bearer ${token}`;
    }

    const response = await fetch(`http://backend:8080/avatars`, {
      headers,
    });

    if (!response.ok) {
      // Instead of throwing an error, serve a default avatar
      // You can either redirect to a static asset or generate one
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
    // Serve default avatar on error too
    return new Response(null, {
      status: 302,
      headers: {
        Location: "/images/default-avatar.png",
      },
    });
  }
};
