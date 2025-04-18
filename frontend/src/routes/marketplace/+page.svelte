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
                <!-- 90% match is for AI calculate this, consider hardcoding instead -->
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

            <div class="flex flex-row gap-2">
              <p class="text-sm text-gray-500">
                {project.current_member}/{project.max_member}
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
          </div>
        </a>
      </div>
    {/each}

    <!-- Other project cards continue here -->
  </div>
</div>
