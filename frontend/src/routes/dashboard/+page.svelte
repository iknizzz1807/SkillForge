<script lang="ts">
  import { browser } from "$app/environment";
  import { onMount } from "svelte";
  import Chart from "chart.js/auto";
  import type { PageData } from "./$types";

  let { data }: { data: PageData } = $props();

  // Toggle this value to switch between roles
  let role: string | null = $state(data.role);

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

  // Function to initialize all charts
  function initCharts() {
    if (!browser) return;

    // Destroy existing charts before recreating
    if (progressChart) progressChart.destroy();
    if (skillsChart) skillsChart.destroy();
    if (applicantsChart) applicantsChart.destroy();
    if (projectsChart) projectsChart.destroy();

    try {
      // Student charts
      if (role === "student") {
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
      } else {
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
                    backgroundColor: [
                      "#48c78e",
                      "#6b48ff",
                      "#f0f0f0",
                      "#ff6b6b",
                    ],
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
    } catch (error) {
      console.error("Error initializing charts:", error);
    }
  }

  onMount(() => {
    // Initialize charts on component mount
    if (browser) {
      setTimeout(() => {
        initCharts();
      }, 100);
    }
  });

  // Use $effect to monitor canvas references and role changes
  $effect(() => {
    if (browser && role) {
      // Add a small delay to ensure DOM elements are fully rendered
      setTimeout(() => {
        initCharts();
      }, 100);
    }
  });
</script>

<main class="flex-1 ml-64 pr-4 pl-4 pt-4 pb-4">
  <!-- Header with role display -->
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
  </div>

  <div class="space-y-4">
    <!-- Quick Actions for Business (moved to top) -->
    {#if role === "business"}
      <div class="card p-3 w-full">
        <h3 class="text-base font-semibold mb-3">Quick Actions</h3>
        <div class="flex justify-between w-full">
          <a
            href="/project/create"
            class="flex flex-col items-center w-[32%] p-3 bg-gray-100 rounded hover:bg-gray-200"
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
            class="flex flex-col items-center w-[32%] p-3 bg-gray-100 rounded hover:bg-gray-200"
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
            href="/chat"
            class="flex flex-col items-center w-[32%] p-3 bg-gray-100 rounded hover:bg-gray-200"
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

    <!-- Quick Stats -->
    <div class="flex space-x-3">
      {#if role === "student"}
        <div class="card p-3 w-1/3">
          <h3 class="text-sm font-semibold">Active Projects</h3>
          <p class="text-xl mt-1 accent-color">3</p>
        </div>
        <div class="card p-3 w-1/3">
          <h3 class="text-sm font-semibold">Pending Tasks</h3>
          <p class="text-xl mt-1 accent-color">5</p>
        </div>
        <div class="card p-3 w-1/3">
          <h3 class="text-sm font-semibold">Completed Tasks</h3>
          <p class="text-xl mt-1 accent-color">12</p>
        </div>
      {:else}
        <div class="card p-3 w-1/3">
          <h3 class="text-sm font-semibold">Active Projects</h3>
          <p class="text-xl mt-1 accent-color">5</p>
        </div>
        <div class="card p-3 w-1/3">
          <h3 class="text-sm font-semibold">Total Students</h3>
          <p class="text-xl mt-1 accent-color">17</p>
        </div>
        <div class="card p-3 w-1/3">
          <h3 class="text-sm font-semibold">Completed Projects</h3>
          <p class="text-xl mt-1 accent-color">8</p>
        </div>
      {/if}
    </div>

    <!-- Projects section -->
    <div class="flex flex-wrap gap-3">
      <!-- Current/Active Projects -->
      <div class="card p-3 w-full">
        <div class="flex justify-between items-center mb-2">
          <h3 class="text-base font-semibold">
            {#if role === "student"}
              Your Active Projects
            {:else}
              Your Active Projects
            {/if}
          </h3>
          <a href="/project" class="text-xs text-blue-600 hover:underline">
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

      <!-- Skills and Suggested Projects for Student -->
      <div class="flex flex-wrap gap-3">
        <!-- Skills Radar -->
        <div class="card p-3 w-[calc(50%-0.375rem)]">
          <h3 class="text-base font-semibold mb-2">Your Skills Assessment</h3>
          <canvas id="skillsChart" height="200" bind:this={skillsChartCanvas}
          ></canvas>
        </div>

        <!-- Suggested Projects -->
        <div class="card p-3 w-[calc(50%-0.375rem)]">
          <div class="flex justify-between items-center mb-2">
            <h3 class="text-base font-semibold">
              Suggested Projects (AI Match)
            </h3>
            <a href="/marketplace" class="text-xs text-blue-600 hover:underline"
              >View all</a
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
      </div>
    {:else}
      <!-- Charts for Business -->
      <div class="flex flex-wrap gap-3">
        <div class="card p-3 w-[calc(50%-0.375rem)]">
          <h3 class="text-base font-semibold mb-2">Project Applications</h3>
          <canvas
            id="applicantsChart"
            height="190"
            bind:this={applicantsChartCanvas}
          ></canvas>
        </div>
        <div class="card p-3 w-[calc(50%-0.375rem)]">
          <h3 class="text-base font-semibold mb-2">Projects Status Overview</h3>
          <canvas
            id="projectsChart"
            height="190"
            bind:this={projectsChartCanvas}
          ></canvas>
        </div>
      </div>

      <!-- Latest Reviews for Business -->
      <div class="card p-3 w-full">
        <div class="flex justify-between items-center mb-2">
          <h3 class="text-base font-semibold">Latest Reviews Needed</h3>
          <a href="/project" class="text-xs text-blue-600 hover:underline"
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
            </div>
          </div>
        </div>
      </div>
    {/if}
  </div>
</main>
