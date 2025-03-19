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
<!-- Form tạo dự án với layout đã tối ưu -->
<main class="flex-1 pr-4 pl-4 ml-56">
  <div class="flex justify-between items-center mb-4">
    <h2 class="text-xl font-semibold">Projects</h2>
    {#if data.role === "business"}
      <button class="btn">
        <div class="flex items-center">
          <svg
            class="w-4 h-4 mr-1"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M12 6v6m0 0v6m0-6h6m-6 0H6"
            ></path>
          </svg>
          Create Project
        </div>
      </button>
    {/if}
  </div>

  <div class="flex space-x-4">
    <!-- Phần danh sách dự án active -->
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
      {:else}
        <div class="card p-4 text-center">
          <p class="text-gray-500">No active projects found.</p>
          <button class="btn mt-2">Create Your First Project</button>
        </div>
      {/if}
    </div>

    {#if data.role === "business"}
      <!-- Right sidebar with analytics and talent pool -->
      <div class="w-2/5 space-y-4">
        <!-- Business Analytics -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Project Analytics</h3>
          <div class="space-y-4">
            <div>
              <div class="flex justify-between mb-1">
                <span class="text-sm">Active Projects</span>
                <span class="text-sm font-medium">4</span>
              </div>
              <div class="h-2 bg-gray-200 rounded-full overflow-hidden">
                <div class="h-full bg-[#6b48ff]" style="width: 80%"></div>
              </div>
            </div>
            <div>
              <div class="flex justify-between mb-1">
                <span class="text-sm">Talent Allocation</span>
                <span class="text-sm font-medium">12/15</span>
              </div>
              <div class="h-2 bg-gray-200 rounded-full overflow-hidden">
                <div class="h-full bg-[#ff6f61]" style="width: 75%"></div>
              </div>
            </div>
            <div>
              <div class="flex justify-between mb-1">
                <span class="text-sm">On-time Completion</span>
                <span class="text-sm font-medium">92%</span>
              </div>
              <div class="h-2 bg-gray-200 rounded-full overflow-hidden">
                <div class="h-full bg-green-500" style="width: 92%"></div>
              </div>
            </div>
            <!-- <div class="flex justify-center mt-2">
            <button class="text-sm text-[#6b48ff] hover:underline">
              View Detailed Analytics
            </button>
          </div> -->
          </div>
        </div>

        <!-- Active Talent Pool -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Talent Pool</h3>
          <div class="space-y-2">
            <div
              class="flex items-center justify-between p-2 bg-gray-100 rounded"
            >
              <div class="flex items-center">
                <img
                  class="w-8 h-8 rounded-full mr-2"
                  src="https://avatars.githubusercontent.com/u/234567?v=4"
                  alt="Developer"
                />
                <div>
                  <p class="text-sm font-medium">Sarah Johnson</p>
                  <p class="text-xs text-gray-500">React, Node.js</p>
                </div>
              </div>
              <button class="btn-secondary text-xs">Chat</button>
            </div>
            <div
              class="flex items-center justify-between p-2 bg-gray-100 rounded"
            >
              <div class="flex items-center">
                <img
                  class="w-8 h-8 rounded-full mr-2"
                  src="https://avatars.githubusercontent.com/u/345678?v=4"
                  alt="Developer"
                />
                <div>
                  <p class="text-sm font-medium">Michael Chen</p>
                  <p class="text-xs text-gray-500">Python, MongoDB</p>
                </div>
              </div>
              <button class="btn-secondary text-xs">Chat</button>
            </div>
            <div
              class="flex items-center justify-between p-2 bg-gray-100 rounded"
            >
              <div class="flex items-center">
                <img
                  class="w-8 h-8 rounded-full mr-2"
                  src="https://avatars.githubusercontent.com/u/456789?v=4"
                  alt="Developer"
                />
                <div>
                  <p class="text-sm font-medium">Alex Rivera</p>
                  <p class="text-xs text-gray-500">UI/UX, Vue.js</p>
                </div>
              </div>
              <button class="btn-secondary text-xs">Chat</button>
            </div>
            <div class="flex justify-center mt-2">
              <button class="text-sm text-[#6b48ff] hover:underline">
                View All Talent
              </button>
            </div>
          </div>
        </div>
      </div>
    {/if}
  </div>
</main>

<style>
  [contenteditable]:empty:before {
    content: attr(placeholder);
    color: #9ca3af;
    display: block;
  }
</style>
