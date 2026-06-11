<script lang="ts">
  import type { PageData } from "./$types";

  let { data }: { data: PageData } = $props();
  let students: any[] = $state([]);
  let loading = $state(true);
  let searchQuery = $state("");

  $effect(() => {
    const fetchStudents = async () => {
      try {
        const res = await fetch("/api/users/students");
        if (res.ok) {
          const json = await res.json();
          students = Array.isArray(json) ? json : json.students ?? json.users ?? [];
        }
      } catch (e) {
        console.error("Failed to fetch students:", e);
      } finally {
        loading = false;
      }
    };
    fetchStudents();
  });

  let filtered = $derived(
    students.filter(
      (s) =>
        !searchQuery ||
        (s.name ?? "").toLowerCase().includes(searchQuery.toLowerCase()) ||
        (s.title ?? s.skills ?? "").toLowerCase().includes(searchQuery.toLowerCase())
    )
  );
</script>

<main class="flex-1 p-6">
  <div class="max-w-6xl mx-auto space-y-6">
    <div class="flex justify-between items-center">
      <div>
        <h1 class="text-2xl font-bold">Talent Pool</h1>
        <p class="text-gray-500 text-sm">Browse and search for student talent</p>
      </div>
    </div>

    <div class="relative">
      <input
        type="text"
        placeholder="Search by name or skill..."
        bind:value={searchQuery}
        class="w-full p-3 pl-10 border border-gray-300 rounded-lg focus:ring-2 focus:ring-[#6b48ff] focus:border-transparent outline-none"
      />
      <svg class="absolute left-3 top-3.5 h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
      </svg>
    </div>

    {#if loading}
      <div class="text-center py-12">
        <svg class="animate-spin h-8 w-8 text-[#6b48ff] mx-auto" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" />
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z" />
        </svg>
        <p class="text-gray-500 mt-2">Loading students...</p>
      </div>
    {:else if filtered.length === 0}
      <div class="text-center py-12 text-gray-500">
        {#if searchQuery}
          No students match your search.
        {:else}
          No students found in the talent pool.
        {/if}
      </div>
    {:else}
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {#each filtered as student}
          <div class="card p-4 hover:shadow-md transition-shadow">
            <div class="flex items-center space-x-3 mb-3">
              <img
                class="w-12 h-12 rounded-full object-cover border-2 border-gray-200"
                src={`https://ui-avatars.com/api/?name=${encodeURIComponent(student.name ?? "U")}&background=6b48ff&color=fff`}
                alt=""
              />
              <div class="min-w-0">
                <p class="font-semibold text-sm truncate">{student.name ?? "Unknown"}</p>
                <p class="text-xs text-gray-500 truncate">{student.title ?? student.email ?? ""}</p>
              </div>
            </div>
            {#if student.skills && student.skills.length > 0}
              <div class="flex flex-wrap gap-1">
                {#each student.skills.slice(0, 4) as skill}
                  <span class="text-xs bg-purple-100 text-purple-700 px-2 py-0.5 rounded">{typeof skill === "string" ? skill : skill.name ?? skill.skill ?? ""}</span>
                {/each}
                {#if student.skills.length > 4}
                  <span class="text-xs text-gray-400">+{student.skills.length - 4}</span>
                {/if}
              </div>
            {/if}
            <div class="mt-3 flex justify-between items-center">
              <span class="text-xs text-gray-400">{student.email ?? ""}</span>
              <a href="/profile/user/{student.id}" class="text-xs text-[#6b48ff] hover:underline">View Profile</a>
            </div>
          </div>
        {/each}
      </div>
    {/if}
  </div>
</main>
