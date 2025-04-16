<script lang="ts">
  import { browser } from "$app/environment";
  import { onMount } from "svelte";
  import Chart from "chart.js/auto";

  // Toggle this value to switch between roles
  let role: "student" | "business" = $state("business");

  // Canvas references for charts
  let progressChartCanvas: HTMLCanvasElement | null = $state(null);
  let skillsChartCanvas: HTMLCanvasElement | null = $state(null);
  let applicantsChartCanvas: HTMLCanvasElement | null = $state(null);
  let projectsChartCanvas: HTMLCanvasElement | null = $state(null);

  // Chart instances to destroy when recreating
  let progressChart: Chart | null = $state(null);
  let skillsChart: Chart | null = $state(null);
  let applicantsChart: Chart | null = $state(null);
  let projectsChart: Chart | null = $state(null);

  // Toggle function to switch roles
  function toggleRole() {
    role = role === "student" ? "business" : "student";
    // Re-render charts on next tick after DOM updates
    setTimeout(() => {
      initCharts();
    }, 0);
  }

  // Function to initialize all charts
  function initCharts() {
    if (!browser) return;

    // Destroy existing charts before recreating
    if (progressChart) progressChart.destroy();
    if (skillsChart) skillsChart.destroy();
    if (applicantsChart) applicantsChart.destroy();
    if (projectsChart) projectsChart.destroy();

    // Student charts
    if (progressChartCanvas) {
      const ctx = progressChartCanvas.getContext("2d");
      if (ctx) {
        progressChart = new Chart(ctx, {
          type: "line",
          data: {
            labels: ["10/03", "12/03", "14/03", "16/03", "18/03"],
            datasets: [
              {
                label: "Completion (%)",
                data: [20, 35, 50, 70, 85],
                borderColor: "#6b48ff",
                backgroundColor: "rgba(107, 72, 255, 0.1)",
                fill: true,
                tension: 0.4,
              },
            ],
          },
          options: {
            responsive: true,
            scales: {
              y: { beginAtZero: true, max: 100 },
            },
          },
        });
      }
    }

    if (skillsChartCanvas) {
      const ctx = skillsChartCanvas.getContext("2d");
      if (ctx) {
        skillsChart = new Chart(ctx, {
          type: "radar",
          data: {
            labels: [
              "JavaScript",
              "React",
              "Node.js",
              "MongoDB",
              "UI/UX",
              "Testing",
            ],
            datasets: [
              {
                label: "Your Skills",
                data: [85, 70, 65, 60, 75, 50],
                backgroundColor: "rgba(107, 72, 255, 0.2)",
                borderColor: "#6b48ff",
                pointBackgroundColor: "#6b48ff",
              },
              {
                label: "Industry Benchmark",
                data: [75, 80, 70, 65, 60, 75],
                backgroundColor: "rgba(99, 194, 222, 0.2)",
                borderColor: "rgb(99, 194, 222)",
                pointBackgroundColor: "rgb(99, 194, 222)",
              },
            ],
          },
          options: {
            responsive: true,
            scales: {
              r: {
                min: 0,
                max: 100,
                ticks: {
                  stepSize: 20,
                },
              },
            },
          },
        });
      }
    }

    // Business charts
    if (applicantsChartCanvas) {
      const ctx = applicantsChartCanvas.getContext("2d");
      if (ctx) {
        applicantsChart = new Chart(ctx, {
          type: "bar",
          data: {
            labels: [
              "Web App",
              "Mobile App",
              "API Integration",
              "UI Design",
              "Database Migration",
            ],
            datasets: [
              {
                label: "Applications",
                data: [12, 8, 15, 5, 7],
                backgroundColor: "#6b48ff",
              },
              {
                label: "Accepted",
                data: [5, 3, 6, 2, 3],
                backgroundColor: "#48c78e",
              },
            ],
          },
          options: {
            responsive: true,
            scales: {
              y: {
                beginAtZero: true,
                title: {
                  display: true,
                  text: "Number of Students",
                },
              },
              x: {
                title: {
                  display: true,
                  text: "Project Type",
                },
              },
            },
          },
        });
      }
    }

    if (projectsChartCanvas) {
      const ctx = projectsChartCanvas.getContext("2d");
      if (ctx) {
        projectsChart = new Chart(ctx, {
          type: "doughnut",
          data: {
            labels: ["Completed", "In Progress", "Not Started", "Delayed"],
            datasets: [
              {
                data: [35, 40, 15, 10],
                backgroundColor: ["#48c78e", "#6b48ff", "#f0f0f0", "#ff6b6b"],
                borderWidth: 1,
              },
            ],
          },
          options: {
            responsive: true,
            plugins: {
              legend: {
                position: "right",
              },
            },
          },
        });
      }
    }
  }

  onMount(() => {
    // Initialize charts on component mount
    initCharts();
  });

  // Use $effect to monitor canvas references
  $effect(() => {
    // If canvas references change and role is student, reinitialize
    if (role === "student" && progressChartCanvas && skillsChartCanvas) {
      initCharts();
    }
  });
</script>

<main class="flex-1 ml-64 pr-4 pl-4 pt-4 pb-4">
  <!-- Header with role toggle -->
  <div class="mb-4 flex justify-between items-center">
    <div>
      <h1 class="text-2xl font-bold">
        {#if role === "student"}
          Student Dashboard
        {:else}
          Business Dashboard
        {/if}
      </h1>
      <p class="text-gray-600">
        {#if role === "student"}
          Welcome back! Here's your learning progress.
        {:else}
          Welcome back! Here's your talent pipeline overview.
        {/if}
      </p>
    </div>
    <button
      on:click={toggleRole}
      class="px-3 py-2 bg-purple-500 text-white rounded-lg hover:bg-purple-600 transition-colors"
    >
      Switch to {role === "student" ? "Business" : "Student"} View
    </button>
  </div>

  <div class="space-y-4">
    <!-- Quick Stats -->
    <div class="flex space-x-3">
      {#if role === "student"}
        <div class="card p-3 w-1/4">
          <h3 class="text-sm font-semibold">Active Projects</h3>
          <p class="text-xl mt-1 accent-color">3</p>
        </div>
        <div class="card p-3 w-1/4">
          <h3 class="text-sm font-semibold">Pending Tasks</h3>
          <p class="text-xl mt-1 accent-color">5</p>
        </div>
        <div class="card p-3 w-1/4">
          <h3 class="text-sm font-semibold">Completed Tasks</h3>
          <p class="text-xl mt-1 accent-color">12</p>
        </div>
        <div class="card p-3 w-1/4">
          <h3 class="text-sm font-semibold">Earned Badges</h3>
          <p class="text-xl mt-1 accent-color">4</p>
        </div>
      {:else}
        <div class="card p-3 w-1/4">
          <h3 class="text-sm font-semibold">Active Projects</h3>
          <p class="text-xl mt-1 accent-color">5</p>
        </div>
        <div class="card p-3 w-1/4">
          <h3 class="text-sm font-semibold">Total Students</h3>
          <p class="text-xl mt-1 accent-color">17</p>
        </div>
        <div class="card p-3 w-1/4">
          <h3 class="text-sm font-semibold">Pending Reviews</h3>
          <p class="text-xl mt-1 accent-color">8</p>
        </div>
        <div class="card p-3 w-1/4">
          <h3 class="text-sm font-semibold">Project Completion</h3>
          <p class="text-xl mt-1 accent-color">72%</p>
        </div>
      {/if}
    </div>

    <!-- Projects section -->
    <div class="flex space-x-3">
      <!-- Current/Active Projects -->
      <div class="card p-3 w-1/2">
        <div class="flex justify-between items-center mb-2">
          <h3 class="text-base font-semibold">
            {#if role === "student"}
              Your Active Projects
            {:else}
              Your Active Projects
            {/if}
          </h3>
          <a
            href={role === "student" ? "/projects" : "/projects/manage"}
            class="text-xs text-blue-600 hover:underline"
          >
            {role === "student" ? "View all" : "Manage all"}
          </a>
        </div>
        <div class="space-y-2">
          {#if role === "student"}
            <div class="p-2 bg-gray-100 rounded">
              <div class="flex justify-between">
                <p class="text-sm font-medium">Portfolio Site</p>
                <span
                  class="text-xs px-2 py-1 bg-green-100 text-green-800 rounded-full"
                  >75% Complete</span
                >
              </div>
              <p class="text-xs text-gray-500">
                Due: 22/03 • Business: DesignPro
              </p>
              <div class="w-full bg-gray-200 rounded-full h-1.5 mt-1">
                <div
                  class="bg-green-500 h-1.5 rounded-full"
                  style="width: 75%"
                ></div>
              </div>
            </div>
            <div class="p-2 bg-gray-100 rounded">
              <div class="flex justify-between">
                <p class="text-sm font-medium">API Integration</p>
                <span
                  class="text-xs px-2 py-1 bg-yellow-100 text-yellow-800 rounded-full"
                  >40% Complete</span
                >
              </div>
              <p class="text-xs text-gray-500">
                Due: 25/03 • Business: TechSolutions
              </p>
              <div class="w-full bg-gray-200 rounded-full h-1.5 mt-1">
                <div
                  class="bg-yellow-500 h-1.5 rounded-full"
                  style="width: 40%"
                ></div>
              </div>
            </div>
            <div class="p-2 bg-gray-100 rounded">
              <div class="flex justify-between">
                <p class="text-sm font-medium">Mobile App UI</p>
                <span
                  class="text-xs px-2 py-1 bg-blue-100 text-blue-800 rounded-full"
                  >10% Complete</span
                >
              </div>
              <p class="text-xs text-gray-500">
                Due: 30/03 • Business: AppWorks
              </p>
              <div class="w-full bg-gray-200 rounded-full h-1.5 mt-1">
                <div
                  class="bg-blue-500 h-1.5 rounded-full"
                  style="width: 10%"
                ></div>
              </div>
            </div>
          {:else}
            <div class="p-2 bg-gray-100 rounded">
              <div class="flex justify-between">
                <p class="text-sm font-medium">E-commerce Web App</p>
                <span
                  class="text-xs px-2 py-1 bg-green-100 text-green-800 rounded-full"
                  >85% Complete</span
                >
              </div>
              <p class="text-xs text-gray-500">
                Due: 22/03 • Students: 4 • Budget: $600
              </p>
              <div class="w-full bg-gray-200 rounded-full h-1.5 mt-1">
                <div
                  class="bg-green-500 h-1.5 rounded-full"
                  style="width: 85%"
                ></div>
              </div>
            </div>
            <div class="p-2 bg-gray-100 rounded">
              <div class="flex justify-between">
                <p class="text-sm font-medium">CRM Integration</p>
                <span
                  class="text-xs px-2 py-1 bg-yellow-100 text-yellow-800 rounded-full"
                  >45% Complete</span
                >
              </div>
              <p class="text-xs text-gray-500">
                Due: 01/04 • Students: 3 • Budget: $450
              </p>
              <div class="w-full bg-gray-200 rounded-full h-1.5 mt-1">
                <div
                  class="bg-yellow-500 h-1.5 rounded-full"
                  style="width: 45%"
                ></div>
              </div>
            </div>
            <div class="p-2 bg-gray-100 rounded">
              <div class="flex justify-between">
                <p class="text-sm font-medium">Mobile App Redesign</p>
                <span
                  class="text-xs px-2 py-1 bg-blue-100 text-blue-800 rounded-full"
                  >20% Complete</span
                >
              </div>
              <p class="text-xs text-gray-500">
                Due: 15/04 • Students: 5 • Budget: $750
              </p>
              <div class="w-full bg-gray-200 rounded-full h-1.5 mt-1">
                <div
                  class="bg-blue-500 h-1.5 rounded-full"
                  style="width: 20%"
                ></div>
              </div>
            </div>
          {/if}
        </div>
      </div>

      <!-- Second Column: Tasks for Student / Top Talent for Business -->
      <div class="card p-3 w-1/2">
        <div class="flex justify-between items-center mb-2">
          <h3 class="text-base font-semibold">
            {#if role === "student"}
              Upcoming Tasks
            {:else if role === "business"}
              Top Talent
            {/if}
          </h3>
          <a
            href={role === "student" ? "/tasks" : "/talent/all"}
            class="text-xs text-blue-600 hover:underline"
          >
            View all
          </a>
        </div>
        <div class="space-y-2">
          {#if role === "student"}
            <div class="p-2 bg-gray-100 rounded flex items-center">
              <div class="w-2 h-2 rounded-full bg-red-500 mr-2"></div>
              <div class="flex-1">
                <p class="text-sm font-medium">Complete homepage design</p>
                <p class="text-xs text-gray-500">
                  Due tomorrow • Portfolio Site
                </p>
              </div>
              <div
                class="text-xs px-2 py-1 bg-red-100 text-red-800 rounded-full"
              >
                High
              </div>
            </div>
            <div class="p-2 bg-gray-100 rounded flex items-center">
              <div class="w-2 h-2 rounded-full bg-yellow-500 mr-2"></div>
              <div class="flex-1">
                <p class="text-sm font-medium">Implement authentication</p>
                <p class="text-xs text-gray-500">
                  Due in 3 days • API Integration
                </p>
              </div>
              <div
                class="text-xs px-2 py-1 bg-yellow-100 text-yellow-800 rounded-full"
              >
                Medium
              </div>
            </div>
            <div class="p-2 bg-gray-100 rounded flex items-center">
              <div class="w-2 h-2 rounded-full bg-blue-500 mr-2"></div>
              <div class="flex-1">
                <p class="text-sm font-medium">Initial wireframes</p>
                <p class="text-xs text-gray-500">
                  Due in 5 days • Mobile App UI
                </p>
              </div>
              <div
                class="text-xs px-2 py-1 bg-blue-100 text-blue-800 rounded-full"
              >
                Standard
              </div>
            </div>
          {:else}
            <div class="p-2 bg-gray-100 rounded flex items-center">
              <div
                class="w-8 h-8 rounded-full bg-gray-300 mr-2 flex items-center justify-center text-sm font-bold"
              >
                JS
              </div>
              <div class="flex-1">
                <p class="text-sm font-medium">Jane Smith</p>
                <p class="text-xs text-gray-500">
                  Full Stack Developer • 5 Projects • 4.9/5 ⭐
                </p>
              </div>
              <div class="flex space-x-1">
                <button
                  class="text-xs px-2 py-1 bg-blue-100 text-blue-800 rounded hover:bg-blue-200"
                  >Profile</button
                >
                <button
                  class="text-xs px-2 py-1 bg-green-100 text-green-800 rounded hover:bg-green-200"
                  >Invite</button
                >
              </div>
            </div>
            <div class="p-2 bg-gray-100 rounded flex items-center">
              <div
                class="w-8 h-8 rounded-full bg-gray-300 mr-2 flex items-center justify-center text-sm font-bold"
              >
                MJ
              </div>
              <div class="flex-1">
                <p class="text-sm font-medium">Michael Johnson</p>
                <p class="text-xs text-gray-500">
                  UI/UX Designer • 3 Projects • 4.8/5 ⭐
                </p>
              </div>
              <div class="flex space-x-1">
                <button
                  class="text-xs px-2 py-1 bg-blue-100 text-blue-800 rounded hover:bg-blue-200"
                  >Profile</button
                >
                <button
                  class="text-xs px-2 py-1 bg-green-100 text-green-800 rounded hover:bg-green-200"
                  >Invite</button
                >
              </div>
            </div>
            <div class="p-2 bg-gray-100 rounded flex items-center">
              <div
                class="w-8 h-8 rounded-full bg-gray-300 mr-2 flex items-center justify-center text-sm font-bold"
              >
                AL
              </div>
              <div class="flex-1">
                <p class="text-sm font-medium">Amanda Lee</p>
                <p class="text-xs text-gray-500">
                  Backend Developer • 4 Projects • 4.7/5 ⭐
                </p>
              </div>
              <div class="flex space-x-1">
                <button
                  class="text-xs px-2 py-1 bg-blue-100 text-blue-800 rounded hover:bg-blue-200"
                  >Profile</button
                >
                <button
                  class="text-xs px-2 py-1 bg-green-100 text-green-800 rounded hover:bg-green-200"
                  >Invite</button
                >
              </div>
            </div>
          {/if}
        </div>
      </div>
    </div>

    <!-- Charts Section -->
    {#if role === "student"}
      <!-- Progress Chart for Student -->
      <div class="card p-3 w-full">
        <h3 class="text-base font-semibold mb-2">
          Your Project Completion Progress
        </h3>
        <canvas id="progressChart" height="70" bind:this={progressChartCanvas}
        ></canvas>
      </div>
    {:else}
      <!-- Charts for Business -->
      <div class="flex space-x-3">
        <div class="card p-3 w-1/2">
          <h3 class="text-base font-semibold mb-2">Project Applications</h3>
          <canvas
            id="applicantsChart"
            height="190"
            bind:this={applicantsChartCanvas}
          ></canvas>
        </div>
        <div class="card p-3 w-1/2">
          <h3 class="text-base font-semibold mb-2">Projects Status Overview</h3>
          <canvas
            id="projectsChart"
            height="190"
            bind:this={projectsChartCanvas}
          ></canvas>
        </div>
      </div>
    {/if}

    <!-- Skills/Suggestions for Student or Reviews/Applications for Business -->
    <div class="flex space-x-3">
      {#if role === "student"}
        <!-- Skills Radar -->
        <div class="card p-3 w-1/2">
          <h3 class="text-base font-semibold mb-2">Your Skills Assessment</h3>
          <canvas id="skillsChart" height="200" bind:this={skillsChartCanvas}
          ></canvas>
        </div>

        <!-- Suggested Projects -->
        <div class="card p-3 w-1/2">
          <div class="flex justify-between items-center mb-2">
            <h3 class="text-base font-semibold">
              Suggested Projects (AI Match)
            </h3>
            <a
              href="/projects/suggestions"
              class="text-xs text-blue-600 hover:underline">View all</a
            >
          </div>
          <div class="space-y-2">
            <div
              class="flex justify-between items-center p-2 bg-gray-100 rounded"
            >
              <div>
                <div class="flex items-center">
                  <p class="text-sm font-medium">Build a Chat App</p>
                  <span
                    class="ml-2 px-2 py-0.5 bg-purple-100 text-purple-800 rounded-full text-xs"
                    >90% Match</span
                  >
                </div>
                <p class="text-xs text-gray-500">
                  Skills: React, Node.js • Business: ChatTech
                </p>
              </div>
              <button class="btn text-xs">Apply</button>
            </div>
            <div
              class="flex justify-between items-center p-2 bg-gray-100 rounded"
            >
              <div>
                <div class="flex items-center">
                  <p class="text-sm font-medium">E-commerce Backend</p>
                  <span
                    class="ml-2 px-2 py-0.5 bg-purple-100 text-purple-800 rounded-full text-xs"
                    >85% Match</span
                  >
                </div>
                <p class="text-xs text-gray-500">
                  Skills: Python, MongoDB • Business: ShopSmart
                </p>
              </div>
              <button class="btn text-xs">Apply</button>
            </div>
            <div
              class="flex justify-between items-center p-2 bg-gray-100 rounded"
            >
              <div>
                <div class="flex items-center">
                  <p class="text-sm font-medium">Analytics Dashboard</p>
                  <span
                    class="ml-2 px-2 py-0.5 bg-purple-100 text-purple-800 rounded-full text-xs"
                    >82% Match</span
                  >
                </div>
                <p class="text-xs text-gray-500">
                  Skills: JavaScript, D3.js • Business: DataViz
                </p>
              </div>
              <button class="btn text-xs">Apply</button>
            </div>
          </div>
        </div>
      {:else}
        <!-- Latest Reviews -->
        <div class="card p-3 w-1/2">
          <div class="flex justify-between items-center mb-2">
            <h3 class="text-base font-semibold">Latest Reviews Needed</h3>
            <a href="/reviews" class="text-xs text-blue-600 hover:underline"
              >View all</a
            >
          </div>
          <div class="space-y-2">
            <div class="p-2 bg-gray-100 rounded">
              <div class="flex justify-between">
                <p class="text-sm font-medium">Homepage Redesign</p>
                <span
                  class="text-xs px-2 py-1 bg-red-100 text-red-800 rounded-full"
                  >Due today</span
                >
              </div>
              <p class="text-xs text-gray-500">
                Submitted by: Michael Johnson • E-commerce Web App
              </p>
              <div class="flex mt-1 space-x-1">
                <button
                  class="text-xs px-2 py-1 bg-blue-100 text-blue-800 rounded hover:bg-blue-200"
                  >Review</button
                >
                <button
                  class="text-xs px-2 py-1 bg-gray-100 text-gray-800 rounded hover:bg-gray-200"
                  >Later</button
                >
              </div>
            </div>
            <div class="p-2 bg-gray-100 rounded">
              <div class="flex justify-between">
                <p class="text-sm font-medium">User Authentication API</p>
                <span
                  class="text-xs px-2 py-1 bg-yellow-100 text-yellow-800 rounded-full"
                  >Due tomorrow</span
                >
              </div>
              <p class="text-xs text-gray-500">
                Submitted by: Amanda Lee • CRM Integration
              </p>
              <div class="flex mt-1 space-x-1">
                <button
                  class="text-xs px-2 py-1 bg-blue-100 text-blue-800 rounded hover:bg-blue-200"
                  >Review</button
                >
                <button
                  class="text-xs px-2 py-1 bg-gray-100 text-gray-800 rounded hover:bg-gray-200"
                  >Later</button
                >
              </div>
            </div>
          </div>
        </div>

        <!-- Project Applications -->
        <div class="card p-3 w-1/2">
          <div class="flex justify-between items-center mb-2">
            <h3 class="text-base font-semibold">New Project Applications</h3>
            <a
              href="/applications"
              class="text-xs text-blue-600 hover:underline">View all</a
            >
          </div>
          <div class="space-y-2">
            <div class="p-2 bg-gray-100 rounded">
              <div class="flex justify-between">
                <p class="text-sm font-medium">Sarah Williams</p>
                <span
                  class="text-xs px-2 py-1 bg-purple-100 text-purple-800 rounded-full"
                  >95% Match</span
                >
              </div>
              <p class="text-xs text-gray-500">
                Applied for: Database Migration • Skills: MongoDB, PostgreSQL
              </p>
              <div class="flex mt-1 space-x-1">
                <button
                  class="text-xs px-2 py-1 bg-green-100 text-green-800 rounded hover:bg-green-200"
                  >Accept</button
                >
                <button
                  class="text-xs px-2 py-1 bg-red-100 text-red-800 rounded hover:bg-red-200"
                  >Decline</button
                >
                <button
                  class="text-xs px-2 py-1 bg-blue-100 text-blue-800 rounded hover:bg-blue-200"
                  >Profile</button
                >
              </div>
            </div>
            <div class="p-2 bg-gray-100 rounded">
              <div class="flex justify-between">
                <p class="text-sm font-medium">David Chen</p>
                <span
                  class="text-xs px-2 py-1 bg-purple-100 text-purple-800 rounded-full"
                  >87% Match</span
                >
              </div>
              <p class="text-xs text-gray-500">
                Applied for: Mobile App Redesign • Skills: React Native, UI/UX
              </p>
              <div class="flex mt-1 space-x-1">
                <button
                  class="text-xs px-2 py-1 bg-green-100 text-green-800 rounded hover:bg-green-200"
                  >Accept</button
                >
                <button
                  class="text-xs px-2 py-1 bg-red-100 text-red-800 rounded hover:bg-red-200"
                  >Decline</button
                >
                <button
                  class="text-xs px-2 py-1 bg-blue-100 text-blue-800 rounded hover:bg-blue-200"
                  >Profile</button
                >
              </div>
            </div>
          </div>
        </div>
      {/if}
    </div>

    <!-- Final Section: Recent Activity for Student or Quick Actions for Business -->
    {#if role === "student"}
      <!-- Recent Activity -->
      <div class="card p-3 w-full">
        <h3 class="text-base font-semibold mb-2">Recent Activity</h3>
        <div class="space-y-2">
          <div class="p-2 bg-gray-100 rounded flex items-center">
            <div
              class="w-8 h-8 rounded-full bg-green-100 text-green-600 flex items-center justify-center mr-2"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-4 w-4"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                />
              </svg>
            </div>
            <div>
              <p class="text-sm">
                You completed the task <span class="font-medium"
                  >CSS Styling for Homepage</span
                >
              </p>
              <p class="text-xs text-gray-500">Today, 14:30</p>
            </div>
          </div>
          <div class="p-2 bg-gray-100 rounded flex items-center">
            <div
              class="w-8 h-8 rounded-full bg-blue-100 text-blue-600 flex items-center justify-center mr-2"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-4 w-4"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path d="M10 12a2 2 0 100-4 2 2 0 000 4z" />
                <path
                  fill-rule="evenodd"
                  d="M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z"
                  clip-rule="evenodd"
                />
              </svg>
            </div>
            <div>
              <p class="text-sm">
                <span class="font-medium">DesignPro</span> reviewed your submission
              </p>
              <p class="text-xs text-gray-500">Yesterday, 16:45</p>
            </div>
          </div>
          <div class="p-2 bg-gray-100 rounded flex items-center">
            <div
              class="w-8 h-8 rounded-full bg-purple-100 text-purple-600 flex items-center justify-center mr-2"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-4 w-4"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  fill-rule="evenodd"
                  d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-11a1 1 0 10-2 0v2H7a1 1 0 100 2h2v2a1 1 0 102 0v-2h2a1 1 0 100-2h-2V7z"
                  clip-rule="evenodd"
                />
              </svg>
            </div>
            <div>
              <p class="text-sm">
                You've been assigned a new task in <span class="font-medium"
                  >Mobile App UI</span
                >
              </p>
              <p class="text-xs text-gray-500">Yesterday, 09:15</p>
            </div>
          </div>
        </div>
      </div>
    {:else}
      <!-- Quick Actions for Business -->
      <div class="card p-3 w-full">
        <h3 class="text-base font-semibold mb-3">Quick Actions</h3>
        <div class="flex space-x-3">
          <a
            href="/projects/create"
            class="flex flex-col items-center w-1/4 p-3 bg-gray-100 rounded hover:bg-gray-200"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-6 w-6 text-blue-600 mb-2"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 6v6m0 0v6m0-6h6m-6 0H6"
              />
            </svg>
            <span class="text-sm">Create Project</span>
          </a>
          <a
            href="/talent/search"
            class="flex flex-col items-center w-1/4 p-3 bg-gray-100 rounded hover:bg-gray-200"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-6 w-6 text-green-600 mb-2"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
              />
            </svg>
            <span class="text-sm">Find Talent</span>
          </a>
          <a
            href="/reviews/pending"
            class="flex flex-col items-center w-1/4 p-3 bg-gray-100 rounded hover:bg-gray-200"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-6 w-6 text-yellow-600 mb-2"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"
              />
            </svg>
            <span class="text-sm">Review Work</span>
          </a>
          <a
            href="/messages"
            class="flex flex-col items-center w-1/4 p-3 bg-gray-100 rounded hover:bg-gray-200"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-6 w-6 text-purple-600 mb-2"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z"
              />
            </svg>
            <span class="text-sm">Messages</span>
          </a>
        </div>
      </div>
    {/if}
  </div>
</main>
