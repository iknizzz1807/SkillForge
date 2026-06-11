import type { PageServerLoad } from "./$types";
import { redirect } from "@sveltejs/kit";
import { BACKEND_URL } from '$lib/backend';

export const load = (async ({ params, locals, fetch }) => {
  const current_user_id = locals.user?.id;
  const target_user_id = params.id;
  const token = locals.token;

  if (current_user_id === target_user_id) {
    throw redirect(302, "/profile");
  }

  let profile: Record<string, any> = {};

  try {
    const response = await fetch(
      `${BACKEND_URL}/api/users/${target_user_id}/profile`,
      {
        headers: { Authorization: `Bearer ${token}` },
      }
    );
    if (response.ok) {
      const json = await response.json();
      profile = json;
    }
  } catch (error) {
    console.error("Failed to fetch user profile:", error);
  }

  const u = profile.user || {};
  const origin = locals.token ? "" : "";

  return {
    id: target_user_id,
    name: u.name || "",
    title: u.title || "",
    email: u.email || "",
    avatarUrl: u.id ? `/api/avatars/${u.id}` : null,
    skills: u.skills || [],
    badges: (profile.badges || []).map((b: any) => ({
      name: b.badge?.name ?? b.name ?? "Badge",
      description: b.badge?.description ?? b.description ?? "",
      icon: b.badge?.icon ?? b.icon ?? "🏅",
      date: b.awarded_at ?? "",
    })),
    projects: (profile.projects || []).map((p: any) => ({
      title: p.title ?? p.name ?? "Untitled",
      status: p.status ?? "",
      end_time: p.end_time ?? "",
      description: p.description ?? "",
    })),
    feedbacks: profile.feedbacks || [],
  };
}) satisfies PageServerLoad;
