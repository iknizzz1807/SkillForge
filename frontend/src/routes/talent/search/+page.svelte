<script lang="ts">
  import { browser } from "$app/environment";
  import { goto } from "$app/navigation";

  let students: any[] = $state([]);
  let loading: boolean = $state(true);
  let error: string | null = $state(null);
  let page: number = $state(1);
  let totalPages: number = $state(1);
  let skillFilter: string = $state("");
  let searchTimeout: ReturnType<typeof setTimeout> | null = null;

  // Invite modal state
  let showInviteModal: boolean = $state(false);
  let selectedStudentId: string | null = $state(null);
  let projects: any[] = $state([]);
  let selectedProjectId: string | null = $state(null);
  let inviting: boolean = $state(false);
  let inviteSuccess: boolean = $state(false);
  let inviteError: string | null = $state(null);

  async function loadStudents() {
    if (!browser) return;
    loading = true;
    error = null;
    try {
      const params = new URLSearchParams({ page: String(page), limit: "12" });
      if (skillFilter) params.set("skill", skillFilter);
      const res = await fetch(`/api/users/students?${params}`);
      if (!res.ok) throw new Error("Failed to load students");
      const data = await res.json();
      students = data.students;
      totalPages = Math.ceil((data.total || 1) / 12);
    } catch (e) {
      error = e instanceof Error ? e.message : "An error occurred";
    } finally {
      loading = false;
    }
  }

  function onSkillInput() {
    if (searchTimeout) clearTimeout(searchTimeout);
    searchTimeout = setTimeout(() => {
      page = 1;
      loadStudents();
    }, 400);
  }

  function prevPage() {
    if (page > 1) { page--; loadStudents(); }
  }

  function nextPage() {
    page++; loadStudents();
  }

  async function openInviteModal(studentId: string) {
    selectedStudentId = studentId;
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
    selectedStudentId = null;
    selectedProjectId = null;
    projects = [];
  }

  async function sendInvite() {
    if (!selectedStudentId || !selectedProjectId) return;
    inviting = true;
    inviteError = null;
    try {
      const res = await fetch("/api/invitations", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ student_id: selectedStudentId, project_id: selectedProjectId }),
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

  function viewProfile(studentId: string) {
    goto(`/talent/search/${studentId}`);
  }

  $effect(() => {
    loadStudents();
  });
</script>

<main class="flex-1 ml-64 pr-4 pl-4 pt-4 pb-4">
  <div class="mb-4">
    <h1 class="text-2xl font-bold">Browse Students</h1>
    <p class="text-gray-600">Find talented students for your projects</p>
  </div>

  <div class="mb-4 flex gap-2">
    <input
      type="text"
      placeholder="Search by skill..."
      bind:value={skillFilter}
      oninput={onSkillInput}
      class="w-full max-w-md p-2 border border-gray-300 rounded focus:ring-2 focus:ring-[#6b48ff] focus:border-transparent outline-none"
    />
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
  {:else if students.length === 0}
    <div class="text-center py-12 text-gray-500">No students found.</div>
  {:else}
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      {#each students as student}
        <div class="card p-4 flex flex-col">
          <div class="flex items-center space-x-3 mb-3">
            <img
              class="w-12 h-12 rounded-full object-cover border-2 border-[#6b48ff]"
              src={student.avatar_name ? `/api/avatars/${student.id}` : "/api/avatars"}
              alt={student.name}
            />
            <div class="flex-1 min-w-0">
              <h3 class="font-semibold truncate">{student.name}</h3>
              <p class="text-sm text-gray-500 truncate">{student.title || "Student"}</p>
            </div>
          </div>
          <div class="flex flex-wrap gap-1 mb-3">
            {#if student.skills && student.skills.length > 0}
              {#each student.skills.slice(0, 4) as skill}
                <span class="text-xs px-2 py-0.5 bg-purple-100 text-purple-800 rounded-full">{skill}</span>
              {/each}
              {#if student.skills.length > 4}
                <span class="text-xs px-2 py-0.5 bg-gray-100 text-gray-600 rounded-full">+{student.skills.length - 4}</span>
              {/if}
            {:else}
              <span class="text-xs text-gray-400">No skills listed</span>
            {/if}
          </div>
          <div class="mt-auto flex gap-2">
            <button onclick={() => viewProfile(student.id)} class="btn-secondary text-xs py-1.5 px-3 flex-1">
              View Profile
            </button>
            <button onclick={() => openInviteModal(student.id)} class="btn text-xs py-1.5 px-3 flex-1">
              Invite to Project
            </button>
          </div>
        </div>
      {/each}
    </div>

    <div class="flex justify-center items-center gap-4 mt-6">
      <button onclick={prevPage} disabled={page <= 1} class="btn-secondary text-sm px-3 py-1" class:opacity-50={page <= 1}>
        Previous
      </button>
      <span class="text-sm text-gray-600">Page {page}</span>
      <button onclick={nextPage} disabled={page >= totalPages} class="btn-secondary text-sm px-3 py-1" class:opacity-50={page >= totalPages}>
        Next
      </button>
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
            <select id="invite-project" bind:value={selectedProjectId} class="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-[#6b48ff] focus:border-transparent outline-none">
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
