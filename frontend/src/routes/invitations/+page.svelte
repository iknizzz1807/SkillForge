<script lang="ts">
  import { browser } from "$app/environment";

  let invitations: any[] = $state([]);
  let loading: boolean = $state(true);
  let error: string | null = $state(null);
  let respondingId: string | null = $state(null);

  async function loadInvitations() {
    if (!browser) return;
    loading = true;
    error = null;
    try {
      const res = await fetch("/api/invitations/me");
      if (!res.ok) throw new Error("Failed to load invitations");
      invitations = await res.json();
    } catch (e) {
      error = e instanceof Error ? e.message : "An error occurred";
    } finally {
      loading = false;
    }
  }

  async function respond(invitationId: string, status: string) {
    respondingId = invitationId;
    try {
      const res = await fetch(`/api/invitations/${invitationId}/respond`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ status }),
      });
      if (!res.ok) {
        const data = await res.json();
        throw new Error(data.error || "Failed to respond");
      }
      invitations = invitations.filter((inv: any) => inv.id !== invitationId);
    } catch (e) {
      alert(e instanceof Error ? e.message : "Failed to respond");
    } finally {
      respondingId = null;
    }
  }

  $effect(() => {
    loadInvitations();
  });
</script>

<main class="flex-1 ml-64 pr-4 pl-4 pt-4 pb-4">
  <div class="mb-4">
    <h1 class="text-2xl font-bold">My Invitations</h1>
    <p class="text-gray-600">Manage your project invitations</p>
  </div>

  {#if loading}
    <div class="flex justify-center py-12">
      <svg class="animate-spin h-8 w-8 text-[#6b48ff]" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" />
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z" />
      </svg>
    </div>
  {:else if error}
    <div class="text-red-600 p-4 bg-red-50 rounded">{error}</div>
  {:else if invitations.length === 0}
    <div class="text-center py-12 text-gray-500">
      <svg class="w-16 h-16 mx-auto mb-4 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
      </svg>
      <p class="text-lg">No pending invitations</p>
    </div>
  {:else}
    <div class="space-y-4">
      {#each invitations as invitation}
        <div class="card p-4 flex items-center justify-between">
          <div class="flex-1">
            <h3 class="font-semibold">{invitation.project_title || "Project"}</h3>
            <p class="text-sm text-gray-500">From: {invitation.business_name || "Business"}</p>
            <p class="text-xs text-gray-400">Received: {new Date(invitation.created_at).toLocaleDateString()}</p>
          </div>
          <div class="flex gap-2">
            <button
              onclick={() => respond(invitation.id, "accepted")}
              disabled={respondingId === invitation.id}
              class="px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600 text-sm disabled:opacity-50"
            >
              {#if respondingId === invitation.id}
                Processing...
              {:else}
                Accept
              {/if}
            </button>
            <button
              onclick={() => respond(invitation.id, "rejected")}
              disabled={respondingId === invitation.id}
              class="px-4 py-2 bg-red-500 text-white rounded hover:bg-red-600 text-sm disabled:opacity-50"
            >
              {#if respondingId === invitation.id}
                Processing...
              {:else}
                Reject
              {/if}
            </button>
          </div>
        </div>
      {/each}
    </div>
  {/if}
</main>
