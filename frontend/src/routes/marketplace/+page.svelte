<script lang="ts">
  import type { PageData } from "./$types";

  let { data }: { data: PageData } = $props();

  const projectsDisplay = data.projects;
  const projectsSuggest = data.projects; // Tạm thời để nó như vậy, khi nào hoàn thiện hệ thống AI

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

<div class="flex space-x-4 ml-64 pl-4 pr-4 pt-4">
  <!-- Left Column - Contains both Filter and Suggested Projects -->
  <div class="w-2/5 space-y-4 flex-shrink-0">
    <!-- Filter Panel -->
    <div class="card p-4">
      <h3 class="text-base font-semibold mb-3">Filters</h3>
      <form class="space-y-3">
        <!-- Skills và Difficulty đặt cùng hàng -->
        <div class="flex space-x-3">
          <div class="w-1/2">
            <label class="block text-sm font-medium mb-1">Fields</label>
            <select
              class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
            >
              <option>All Fields</option>
              <option>Web Development</option>
              <option>Mobile App</option>
              <option>Data Science</option>
              <option>AI engineer</option>
            </select>
          </div>
          <div class="w-1/2">
            <label class="block text-sm font-medium mb-1">Difficulty</label>
            <select
              class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
            >
              <option>All Levels</option>
              <option>Beginner</option>
              <option>Intermediate</option>
              <option>Advanced</option>
              <option>Expert</option>
            </select>
          </div>
        </div>

        <!-- Field với checkbox chia làm 2 cột -->
        <div>
          <label class="block text-sm font-medium mb-1">Skills</label>
          <div class="grid grid-cols-2 gap-1">
            <label class="flex items-center text-sm">
              <input type="checkbox" class="mr-2" /> ReactJS
            </label>
            <label class="flex items-center text-sm">
              <input type="checkbox" class="mr-2" /> Python
            </label>
            <label class="flex items-center text-sm">
              <input type="checkbox" class="mr-2" /> JavaScript
            </label>
            <label class="flex items-center text-sm">
              <input type="checkbox" class="mr-2" /> Docker
            </label>
            <label class="flex items-center text-sm">
              <input type="checkbox" class="mr-2" /> Machine Learning
            </label>
            <label class="flex items-center text-sm">
              <input type="checkbox" class="mr-2" /> Golang
            </label>
            <label class="flex items-center text-sm">
              <input type="checkbox" class="mr-2" /> C++
            </label>
            <label class="flex items-center text-sm">
              <input type="checkbox" class="mr-2" /> Rust
            </label>
            <label class="flex items-center text-sm">
              <input type="checkbox" class="mr-2" /> Java
            </label>
            <label class="flex items-center text-sm">
              <input type="checkbox" class="mr-2" /> TypeScript
            </label>
          </div>
        </div>
        <button type="submit" class="btn w-full">Apply Filters</button>
      </form>
    </div>

    <!-- Suggested Projects -->
    <div class="card p-3">
      <h3 class="text-base font-semibold mb-2">Suggested Projects</h3>
      <div class="space-y-2">
        {#each projectsSuggest as project}
          <div
            class="flex justify-between items-center p-2 bg-gray-100 rounded"
          >
            <div>
              <p class="text-sm font-medium">{project.title}</p>
              <p class="text-xs text-gray-500">
                Skills: {project.skills} | 90% Match
              </p>
              <p class="text-xs text-gray-500">
                Timeline: {formatDate(project.start_time.toString())} - {formatDate(
                  project.end_time.toString()
                )} | By
                <a
                  href={"/user/profile/" + project.created_by_id}
                  class="text-xs text-[#6b48ff]"
                >
                  {project.created_by_name}
                </a>
                <!-- Change the "created_by" to the name of business and href link to
                business profile -->
              </p>
            </div>
            <a href={"/marketplace/" + project.id}>
              <button class="btn text-xs">View</button>
            </a>
          </div>
        {/each}
      </div>
    </div>
  </div>

  <!-- Project List -->
  <div class="w-3/5 space-y-3">
    <!-- Search Bar -->
    <div class="flex items-center space-x-2">
      <input
        type="text"
        class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
        placeholder="Search projects..."
      />
      <button class="btn">
        <svg
          class="w-6.5 h-6.5"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
          ></path>
        </svg>
      </button>
    </div>

    <!-- Project Cards (unchanged) -->
    {#each projectsDisplay as project}
      <div class="card p-3">
        <a href={"/marketplace/" + project.id}>
          <div class="flex justify-between items-center">
            <div>
              <h4 class="text-base font-medium">{project.title}</h4>
              <p class="text-sm text-gray-500">
                {project.skills} | 90% Match
              </p>
              <p class="text-xs text-gray-400 mt-1">
                Timeline: {formatDate(project.start_time.toString())} - {formatDate(
                  project.end_time.toString()
                )} | Posted: {formatDate(project.created_at.toString())}
              </p>
              <p class="text-xs text-gray-400 mt-1">
                By {project.created_by_name}
              </p>
            </div>

            <div class="flex space-x-2 items-center">
              <!-- Hiển thị số lượng thành viên -->
              <div class="flex items-center bg-gray-100 px-2 py-1 rounded-lg">
                <div class="flex -space-x-2 mr-2">
                  <div
                    class="w-6 h-6 rounded-full bg-purple-500 border-2 border-white flex items-center justify-center"
                  >
                    <span class="text-[10px] font-bold text-white"
                      >{project.current_member > 0
                        ? project.current_member
                        : ""}</span
                    >
                  </div>
                  {#if project.current_member >= 2}
                    <div
                      class="w-6 h-6 rounded-full bg-indigo-500 border-2 border-white flex items-center justify-center"
                    ></div>
                  {/if}
                  {#if project.current_member >= 3}
                    <div
                      class="w-6 h-6 rounded-full bg-blue-500 border-2 border-white flex items-center justify-center"
                    ></div>
                  {/if}
                </div>

                <div class="flex items-center">
                  <span class="text-xs font-medium">
                    {project.current_member}/{project.max_member}
                  </span>
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 24 24"
                    fill="currentColor"
                    class="w-4 h-4 ml-1 text-gray-500"
                  >
                    <path
                      d="M4.5 6.375a4.125 4.125 0 118.25 0 4.125 4.125 0 01-8.25 0zM14.25 8.625a3.375 3.375 0 116.75 0 3.375 3.375 0 01-6.75 0zM1.5 19.125a7.125 7.125 0 0114.25 0v.003l-.001.119a.75.75 0 01-.363.63 13.067 13.067 0 01-6.761 1.873c-2.472 0-4.786-.684-6.76-1.873a.75.75 0 01-.364-.63l-.001-.122zM17.25 19.128l-.001.144a2.25 2.25 0 01-.233.96 10.088 10.088 0 005.06-1.01.75.75 0 00.42-.643 4.875 4.875 0 00-6.957-4.611 8.586 8.586 0 011.71 5.157v.003z"
                    />
                  </svg>
                </div>
              </div>

              <!-- Thêm hiển thị status với màu sắc tương ứng -->
              <span
                class={`text-xs px-2 py-1 rounded font-bold ${
                  project.status?.toLowerCase() === "open"
                    ? "bg-green-100 text-green-800"
                    : "bg-red-100 text-red-800"
                }`}
              >
                {project.status ? project.status.toUpperCase() : "OPEN"}
              </span>
            </div>
          </div>
        </a>
      </div>
    {/each}

    <!-- Other project cards continue here -->
  </div>
</div>
