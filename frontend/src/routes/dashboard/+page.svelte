<script lang="ts">
  import { browser } from "$app/environment";
  import { onMount } from "svelte";
  import Chart from "chart.js/auto";

  let progressChartCanvas: HTMLCanvasElement | null = $state(null);

  onMount(() => {
    if (browser && progressChartCanvas) {
      const ctx = progressChartCanvas.getContext("2d");
      if (ctx) {
        new Chart(ctx, {
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
            scales: {
              y: { beginAtZero: true, max: 100 },
            },
          },
        });
      }
    }
  });
</script>

<main class="flex-1 ml-56 pr-4 pl-4">
  <div class="space-y-4">
    <!-- Quick Stats -->
    <div class="flex space-x-3">
      <div class="card p-3 w-1/3">
        <h3 class="text-sm font-semibold">Active Projects</h3>
        <p class="text-xl mt-1 accent-color">3</p>
      </div>
      <div class="card p-3 w-1/3">
        <h3 class="text-sm font-semibold">Pending Tasks</h3>
        <p class="text-xl mt-1 accent-color">5</p>
      </div>
      <!-- <div class="card p-3 w-1/4">
        <h3 class="text-sm font-semibold">Skill Match</h3>
        <p class="text-xl mt-1 accent-color">87%</p>
      </div> -->
      <div class="card p-3 w-1/3">
        <h3 class="text-sm font-semibold">Completed Tasks</h3>
        <p class="text-xl mt-1 accent-color">12</p>
      </div>
    </div>

    <!-- Progress Chart -->
    <div class="card p-3 w-full overflow-x-auto">
      <h3 class="text-base font-semibold mb-2">Your Progress</h3>
      <canvas id="progressChart" height="60" bind:this={progressChartCanvas}
      ></canvas>
    </div>

    <!-- Suggested Projects & Project Status -->
    <div class="flex space-x-3">
      <!-- Suggested Projects -->
      <div class="card p-3 w-2/5">
        <h3 class="text-base font-semibold mb-2">Suggested Projects</h3>
        <div class="space-y-2">
          <div
            class="flex justify-between items-center p-2 bg-gray-100 rounded"
          >
            <div>
              <p class="text-sm font-medium">Build a Chat App</p>
              <p class="text-xs text-gray-500">
                Skills: React, Node.js | 90% Match
              </p>
            </div>
            <button class="btn text-xs">Apply</button>
          </div>
          <div
            class="flex justify-between items-center p-2 bg-gray-100 rounded"
          >
            <div>
              <p class="text-sm font-medium">E-commerce Backend</p>
              <p class="text-xs text-gray-500">
                Skills: Python, MongoDB | 85% Match
              </p>
            </div>
            <button class="btn text-xs">Apply</button>
          </div>
        </div>
      </div>

      <!-- Project Status -->
      <div class="card p-3 w-3/5">
        <h3 class="text-base font-semibold mb-2">Project Status</h3>
        <div class="space-y-2">
          <div class="p-2 bg-gray-100 rounded">
            <p class="text-sm font-medium">Portfolio Site</p>
            <p class="text-xs text-gray-500">75% Complete | Due: 22/03</p>
          </div>
          <div class="p-2 bg-gray-100 rounded">
            <p class="text-sm font-medium">API Integration</p>
            <p class="text-xs text-gray-500">40% Complete | Due: 25/03</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</main>
