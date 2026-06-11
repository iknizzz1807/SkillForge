<script lang="ts">
  import { browser } from "$app/environment";
  import { goto } from "$app/navigation";
  import { page } from "$app/state";

  let profile: any = $state(null);
  let loading: boolean = $state(true);
  let error: string | null = $state(null);

  // Invite modal state
  let showInviteModal: boolean = $state(false);
  let projects: any[] = $state([]);
  let selectedProjectId: string | null = $state(null);
  let inviting: boolean = $state(false);
  let inviteSuccess: boolean = $state(false);
  let inviteError: string | null = $state(null);

  async function loadProfile() {
    if (!browser) return;
    loading = true;
    error = null;
    try {
      const res = await fetch(`/api/users/${page.params.id}/profile`);
      if (!res.ok) throw new Error("Failed to load profile");
      profile = await res.json();
    } catch (e) {
      error = e instanceof Error ? e.message : "An error occurred";
    } finally {
      loading = false;
    }
  }

  function getAvatarUrl(user: any): string {
    return user.avatar_name ? `/api/avatars/${user.id}` : "/api/avatars";
  }

  async function openInviteModal() {
    selectedProjectId = null;
    inviteSuccess = false;
    inviteError = null;
    showInviteModal = true;
    try {
      const res = await fetch("/api/projects/business");
      if (!res.ok) throw new Error("Failed to load projects");
      projects = await res.json();
    } catch (e) {
      inviteError = e instanceof Error ? e.message : "Failed to load projects";
    }
  }

  function closeInviteModal() {
    showInviteModal = false;
    selectedProjectId = null;
    projects = [];
  }

  async function sendInvite() {
    if (!profile?.user?.id || !selectedProjectId) return;
    inviting = true;
    inviteError = null;
    try {
      const res = await fetch("/api/invitations", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ student_id: profile.user.id, project_id: selectedProjectId }),
      });
      if (!res.ok) {
        const data = await res.json();
        throw new Error(data.error || "Failed to send invitation");
      }
      inviteSuccess = true;
      setTimeout(() => closeInviteModal(), 1500);
    } catch (e) {
      inviteError = e instanceof Error ? e.message : "Failed to send invitation";
    } finally {
      inviting = false;
    }
  }

  $effect(() => {
    loadProfile();
  });
</script>

<main class="flex-1 pr-4 pl-4 pt-4 pb-4">
  <button onclick={() => goto("/talent/search")} class="text-sm text-[#6b48ff] hover:underline mb-4 flex items-center">
    <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
    </svg>
    Back to Students
  </button>

  {#if loading}
    <div class="flex justify-center py-12">
      <svg class="animate-spin h-8 w-8 text-[#6b48ff]" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" />
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z" />
      </svg>
    </div>
  {:else if error}
    <div class="text-red-600 p-4 bg-red-50 rounded">{error}</div>
  {:else if profile}
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-4">
      <!-- Left Column - Profile Info -->
      <div class="space-y-4">
        <div class="card p-4">
          <div class="flex items-center space-x-4 mb-4">
            <img class="w-16 h-16 rounded-full object-cover border-2 border-[#6b48ff]" src={getAvatarUrl(profile.user)} alt={profile.user.name} />
            <div>
              <h2 class="text-xl font-bold">{profile.user.name}</h2>
              <p class="text-sm text-gray-500">{profile.user.title || "Student"}</p>
              <p class="text-sm text-gray-400">{profile.user.email}</p>
            </div>
          </div>
          {#if profile.user.website}
            <a href={profile.user.website.startsWith("http") ? profile.user.website : "https://" + profile.user.website} target="_blank" class="text-sm text-[#6b48ff] hover:underline flex items-center">
              <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1" />
              </svg>
              {profile.user.website}
            </a>
          {/if}
          <button onclick={openInviteModal} class="btn w-full mt-4 text-sm">Invite to Project</button>
        </div>

        <!-- Skills -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Skills</h3>
          {#if profile.user.skills && profile.user.skills.length > 0}
            <div class="flex flex-wrap gap-2">
              {#each profile.user.skills as skill}
                <span class="px-3 py-1 bg-purple-100 text-purple-800 rounded-full text-sm">{skill}</span>
              {/each}
            </div>
          {:else}
            <p class="text-sm text-gray-500">No skills listed</p>
          {/if}
        </div>
      </div>

      <!-- Right Column - Badges, Feedback, Projects -->
      <div class="lg:col-span-2 space-y-4">
        <!-- Badges -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Badges</h3>
          {#if profile.badges && profile.badges.length > 0}
            <div class="grid grid-cols-2 sm:grid-cols-3 gap-3">
              {#each profile.badges as ub}
                <div class="text-center p-3 rounded border border-[#6b48ff] bg-purple-50">
                  <p class="text-lg mb-1">{ub.badge?.name || "Badge"}</p>
                  <p class="text-xs text-gray-500">{ub.badge?.description || ""}</p>
                  <p class="text-xs text-[#6b48ff] mt-1">{new Date(ub.awarded_at).toLocaleDateString()}</p>
                </div>
              {/each}
            </div>
          {:else}
            <p class="text-sm text-gray-500">No badges earned yet</p>
          {/if}
        </div>

        <!-- Feedback -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Feedback & Reviews</h3>
          {#if profile.feedbacks && profile.feedbacks.length > 0}
            <div class="space-y-3">
              {#each profile.feedbacks as fb}
                <div class="p-3 bg-gray-50 rounded border border-gray-200">
                  <div class="flex items-center gap-1 mb-1">
                    {#each Array(fb.rating) as _}
                      <svg class="w-4 h-4 text-yellow-500" fill="currentColor" viewBox="0 0 20 20">
                        <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                      </svg>
                    {/each}
                  </div>
                  <p class="text-sm text-gray-700">{fb.content}</p>
                  <p class="text-xs text-gray-400 mt-1">{new Date(fb.created_at).toLocaleDateString()}</p>
                </div>
              {/each}
            </div>
          {:else}
            <p class="text-sm text-gray-500">No feedback yet</p>
          {/if}
        </div>

        <!-- Completed Projects -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Projects</h3>
          {#if profile.projects && profile.projects.length > 0}
            <div class="space-y-3">
              {#each profile.projects as project}
                <div class="p-3 bg-gray-50 rounded border border-gray-200">
                  <div class="flex justify-between items-start">
                    <div>
                      <p class="font-medium">{project.title}</p>
                      <p class="text-xs text-gray-500">{project.description?.slice(0, 100)}</p>
                    </div>
                    <span class="text-xs px-2 py-1 rounded-full capitalize" class:bg-green-100:text-green-800={project.status === "active" || project.status === "open"} class:bg-gray-100:text-gray-600={project.status !== "active" && project.status !== "open"}>
                      {project.status}
                    </span>
                  </div>
                </div>
              {/each}
            </div>
          {:else}
            <p class="text-sm text-gray-500">No projects yet</p>
          {/if}
        </div>
      </div>
    </div>
  {/if}
</main>

{#if showInviteModal}
  <div class="fixed inset-0 bg-black bg-opacity-50 z-50 flex items-center justify-center">
    <div class="bg-white rounded-lg w-full max-w-md mx-4 p-6">
      <h3 class="text-lg font-semibold mb-4">Invite to Project</h3>
      {#if inviteSuccess}
        <div class="text-green-600 flex items-center mb-4">
          <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
          </svg>
          Invitation sent successfully!
        </div>
      {:else}
        <div class="mb-4">
          <label class="block text-sm font-medium mb-1" for="invite-project">Select Project</label>
          {#if projects.length === 0}
            <p class="text-sm text-gray-500">No projects available. Create a project first.</p>
          {:else}
            <select id="invite-project" bind:value={selectedProjectId} class="w-full p-2 border border-gray-200 rounded focus:ring-2 focus:ring-[#6b48ff] focus:border-transparent outline-none">
              <option value={null} disabled>Choose a project...</option>
              {#each projects as project}
                <option value={project.id}>{project.title}</option>
              {/each}
            </select>
          {/if}
        </div>
        {#if inviteError}
          <div class="text-red-600 text-sm mb-3">{inviteError}</div>
        {/if}
        <div class="flex justify-end gap-2">
          <button onclick={closeInviteModal} class="btn-secondary px-4 py-2 text-sm">Cancel</button>
          <button onclick={sendInvite} disabled={!selectedProjectId || inviting} class="btn px-4 py-2 text-sm" class:opacity-50={!selectedProjectId || inviting}>
            {#if inviting}
              Sending...
            {:else}
              Send Invitation
            {/if}
          </button>
        </div>
      {/if}
    </div>
  </div>
{/if}
