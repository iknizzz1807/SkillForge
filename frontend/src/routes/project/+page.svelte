<script lang="ts">
  import type { PageData } from "./$types";

  let { data }: { data: PageData } = $props();

  type ProjectDisplay = {
    id: string;
    title: string;
    description: string;
    skills: string[];
    start_time: Date;
    end_time: Date;
    created_by: string;
    status: string;
    max_member: number;
    current_member: number;
    created_at: string;
  };

  let projectsDisplay: ProjectDisplay[] = $state(data.projects);
  let errorLoadingProjects: string | null = $state(data.error);

  let filterState: string = $state("all");
  let projectsDisplayCopy = [...projectsDisplay];

  $effect(() => {
    if (filterState) {
      if (filterState === "open") {
        projectsDisplay = projectsDisplayCopy.filter(
          (project) => project.status === "open"
        );
      } else if (filterState === "close") {
        projectsDisplay = projectsDisplayCopy.filter(
          (project) => project.status === "close"
        );
      } else if (filterState === "all") {
        projectsDisplay = [...projectsDisplayCopy];
      }
    }
  });

  // Thêm các biến mới để xử lý markdown
  let markdownContent: string = $state("");
  let previewMode: boolean = $state(false);

  // // Hàm để xử lý khi submit form
  // async function handleSubmitMarkdown() {
  //   try {
  //     const response = await fetch("/api/save-markdown", {
  //       method: "POST",
  //       headers: {
  //         "Content-Type": "application/json",
  //       },
  //       body: JSON.stringify({
  //         content: markdownContent,
  //         projectId: "your-project-id", // Thay thế bằng ID dự án thích hợp
  //       }),
  //     });

  //     if (response.ok) {
  //       alert("Markdown đã được lưu thành công!");
  //     } else {
  //       alert("Có lỗi khi lưu markdown!");
  //     }
  //   } catch (error) {
  //     console.error("Error saving markdown:", error);
  //     alert("Có lỗi khi lưu markdown!");
  //   }
  // }

  // Format date function
  function formatDate(dateString: string): string {
    if (!dateString) return "N/A";

    const date = new Date(dateString);

    // Check if date is valid
    if (isNaN(date.getTime())) return "Invalid date";

    // Format: "Mar 24, 2025"
    return date.toLocaleDateString("en-US", {
      month: "short",
      day: "numeric",
      year: "numeric",
    });
  }
</script>

{#if data.role === "business"}
  <header class="flex justify-between items-center mb-4 ml-64 pr-4 pl-4 pt-4">
    <button class="btn">
      <a href="/project/create">
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
      </a>
    </button>
  </header>
{/if}

<main class="flex-1 pr-4 pl-4 ml-64 pt-4">
  <div class="flex space-x-4">
    <!-- Phần danh sách dự án active -->
    <div class="w-3/5 space-y-3">
      <div class="flex justify-between items-center mb-4">
        <h3 class="text-lg font-semibold">Your Active Projects</h3>
        <div class="flex space-x-2">
          <button
            class={filterState === "all"
              ? "text-sm text-[#6b48ff] border-b-2 hover:text-[#6b48ff]"
              : "text-sm text-gray-500 hover:text-[#6b48ff]"}
            onclick={() => (filterState = "all")}
          >
            All
          </button>
          <button
            class={filterState === "open"
              ? "text-sm text-[#6b48ff] border-b-2 hover:text-[#6b48ff]"
              : "text-sm text-gray-500 hover:text-[#6b48ff]"}
            onclick={() => (filterState = "open")}
          >
            Open
          </button>
          <button
            class={filterState === "close"
              ? "text-sm text-[#6b48ff] border-b-2 hover:text-[#6b48ff]"
              : "text-sm text-gray-500 hover:text-[#6b48ff]"}
            onclick={() => (filterState = "close")}
          >
            Close
          </button>
        </div>
      </div>
      {#if projectsDisplay.length !== 0 && !errorLoadingProjects}
        {#each projectsDisplay as ProjectDisplay}
          <!-- Project Card demo -->
          <a href={"/project/" + ProjectDisplay.id} class="card p-3 block">
            <div class="flex justify-between items-center">
              <div>
                <h4 class="text-base font-medium">{ProjectDisplay.title}</h4>
                <p class="text-sm text-gray-500">
                  Skills: {ProjectDisplay.skills}
                </p>
                <p class="text-xs text-gray-400 mt-1">
                  <!-- Thay số 60% bằng số tasks đã hoàn thành / số task tổng -->
                  Progress: <span class="accent-color">60%</span> | Start: {formatDate(
                    ProjectDisplay.start_time.toString()
                  )}
                  | End: {formatDate(ProjectDisplay.end_time.toString())}
                </p>
              </div>

              <div class="flex space-x-2">
                <div class="flex flex-row gap-1">
                  <p class="text-sm text-gray-500">
                    {ProjectDisplay.current_member}/{ProjectDisplay.max_member}
                  </p>
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="gray"
                    height="20px"
                    width="px"
                    version="1.1"
                    id="Layer_1"
                    viewBox="0 0 315.539 315.539"
                    xml:space="preserve"
                  >
                    <g>
                      <g>
                        <g>
                          <path
                            d="M38.877,38.884C17.44,38.884,0,56.325,0,77.761c0,21.436,17.44,38.877,38.877,38.877s38.877-17.44,38.877-38.877     S60.314,38.884,38.877,38.884z M38.877,101.638C25.711,101.638,15,90.927,15,77.761s10.711-23.877,23.877-23.877     s23.877,10.711,23.877,23.877S52.043,101.638,38.877,101.638z"
                          />
                          <path
                            d="M38.877,131.65c-20.073,0-36.404,16.331-36.404,36.404v101.101c0,4.143,3.358,7.5,7.5,7.5h57.808     c4.142,0,7.5-3.357,7.5-7.5V168.054C75.281,147.981,58.95,131.65,38.877,131.65z M60.281,261.655H17.473v-93.601     c0-11.802,9.602-21.404,21.404-21.404c11.803,0,21.404,9.601,21.404,21.404V261.655z"
                          />
                          <path
                            d="M157.77,38.884c-21.437,0-38.877,17.44-38.877,38.877s17.44,38.877,38.877,38.877s38.877-17.44,38.877-38.877     S179.207,38.884,157.77,38.884z M157.77,101.638c-13.166,0-23.877-10.711-23.877-23.877s10.711-23.877,23.877-23.877     s23.877,10.711,23.877,23.877S170.936,101.638,157.77,101.638z"
                          />
                          <path
                            d="M157.769,131.65c-20.073,0-36.404,16.331-36.404,36.404v101.101c0,4.143,3.358,7.5,7.5,7.5h57.808     c4.142,0,7.5-3.357,7.5-7.5V168.054C194.173,147.981,177.842,131.65,157.769,131.65z M179.174,261.655h-42.808v-93.601     c0-11.802,9.602-21.404,21.404-21.404s21.404,9.601,21.404,21.404V261.655z"
                          />
                          <path
                            d="M276.662,116.638c21.437,0,38.877-17.44,38.877-38.877s-17.44-38.877-38.877-38.877s-38.877,17.44-38.877,38.877     S255.225,116.638,276.662,116.638z M276.662,53.884c13.166,0,23.877,10.711,23.877,23.877s-10.711,23.877-23.877,23.877     c-13.166,0-23.877-10.711-23.877-23.877S263.496,53.884,276.662,53.884z"
                          />
                          <path
                            d="M276.662,131.65c-20.073,0-36.404,16.331-36.404,36.404v101.101c0,4.143,3.358,7.5,7.5,7.5h57.808     c4.142,0,7.5-3.357,7.5-7.5V168.054C313.066,147.981,296.735,131.65,276.662,131.65z M298.066,261.655h-42.808v-93.601     c0-11.802,9.602-21.404,21.404-21.404s21.404,9.601,21.404,21.404V261.655z"
                          />
                        </g>
                      </g>
                    </g>
                  </svg>
                </div>
                <span
                  class="text-xs bg-green-100 text-green-800 px-2 py-1 rounded font-bold"
                  >{ProjectDisplay.status.toUpperCase()}</span
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
          <!-- <button class="btn mt-2">Create Your First Project</button> -->
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
              <button class="btn-secondary text-xs">View</button>
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
              <button class="btn-secondary text-xs">View</button>
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
              <button class="btn-secondary text-xs">View</button>
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
