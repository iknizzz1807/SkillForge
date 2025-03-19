<script lang="ts">
  import type { PageData } from "./$types";

  let { data }: { data: PageData } = $props();

  type ProjectDisplay = {
    id: string;
    title: string;
    description: string;
    skills: string[];
    timeline: string;
    created_by: string;
    status: string;
    repo_url: string;
    created_at: string;
  };

  let projectsDisplay: ProjectDisplay[] = $state(data.projects);
  let errorLoadingProjects: string | null = $state(data.error);
</script>

<main class="flex-1 pr-4 pl-4 ml-56">
  <!-- View Toggle
  <div class="flex justify-end mb-4">
    <div class="bg-gray-200 rounded-lg p-1 inline-flex">
      <button class="px-3 py-1 rounded-md text-sm bg-white shadow">
        Freelancer View
      </button>
      <button class="px-3 py-1 rounded-md text-sm text-gray-600">
        Business View
      </button>
    </div>
  </div> -->

  <!-- <div class="space-y-3 max-w-4xl"> -->
  <div class="flex space-x-4">
    <div class="w-3/5 space-y-3">
      <div class="flex justify-between items-center mb-4">
        <h3 class="text-lg font-semibold">Your Active Projects</h3>
        <div class="flex space-x-2">
          <button class="text-sm text-gray-500 hover:text-[#6b48ff]">
            All
          </button>
          <button class="text-sm text-[#6b48ff] border-b-2 border-[#6b48ff]">
            Active
          </button>
          <button class="text-sm text-gray-500 hover:text-[#6b48ff]">
            Completed
          </button>
        </div>
      </div>
      {#if projectsDisplay.length !== 0 && !errorLoadingProjects}
        {#each projectsDisplay as ProjectDisplay}
          <!-- Project Card demo -->
          <a href="/project/project-management" class="card p-3 block">
            <div class="flex justify-between items-center">
              <div>
                <h4 class="text-base font-medium">{ProjectDisplay.title}</h4>
                <p class="text-sm text-gray-500">
                  Skills: {ProjectDisplay.skills}
                </p>
                <p class="text-xs text-gray-400 mt-1">
                  Progress: <span class="accent-color">60%</span> | Due: {ProjectDisplay.timeline}
                  weeks
                </p>
              </div>
              <div class="flex space-x-2">
                <span
                  class="text-xs bg-green-100 text-green-800 px-2 py-1 rounded"
                  >Active</span
                >
                <svg
                  class="w-5 h-5 text-gray-400"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M9 5l7 7-7 7"
                  ></path>
                </svg>
              </div>
            </div>
          </a>
        {/each}
      {/if}
    </div>
  </div>
</main>
