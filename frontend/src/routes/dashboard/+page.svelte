<script lang="ts">
  import { browser } from "$app/environment";
  import { onMount } from "svelte";
  import Chart from "chart.js/auto";
  import type { PageData } from "./$types";

  let { data }: { data: PageData } = $props();

  let role = $derived(data.role ?? "");

  let progressChartCanvas: HTMLCanvasElement | null = $state(null);
  let skillsChartCanvas: HTMLCanvasElement | null = $state(null);
  let applicantsChartCanvas: HTMLCanvasElement | null = $state(null);
  let projectsChartCanvas: HTMLCanvasElement | null = $state(null);

  let progressChart: Chart | null = $state(null);
  let skillsChart: Chart | null = $state(null);
  let applicantsChart: Chart | null = $state(null);
  let projectsChart: Chart | null = $state(null);

  let projects: any[] = $state([]);
  let xpData: any = $state(null);
  let talentData: any = $state(null);
  let marketplaceProjects: any[] = $state([]);
  let feedbacks: any[] = $state([]);
  let loading = $state(true);

  let marketplaceError = $state(false);
  let feedbacksError = $state(false);

  let activeProjectsCount = $derived(
    projects.filter((p) => p.status !== "completed" && p.status !== "close").length
  );

  let completedProjectsCount = $derived(
    projects.filter((p) => p.status === "completed" || p.status === "close").length
  );

  let totalStudents = $derived.by(() => {
    const memberIds = new Set<string>();
    projects.forEach((p) => {
      if (p.members && Array.isArray(p.members)) {
        p.members.forEach((m: any) => {
          const id = typeof m === "string" ? m : m.id || m.user_id || "";
          if (id) memberIds.add(id);
        });
      }
    });
    if (memberIds.size > 0) return memberIds.size;
    return projects.reduce((sum, p) => sum + (p.current_member || 0), 0);
  });

  let completedTasksCount = $derived(
    projects.filter((p) => p.status === "completed" || p.status === "close").length
  );

  let pendingTasksCount = $derived(
    projects.filter((p) => p.status !== "completed" && p.status !== "close").length
  );

  function getCompletionPercent(status: string): number {
    switch (status?.toLowerCase()) {
      case "completed": return 100;
      case "close": return 100;
      case "review": return 75;
      case "in_progress": return 50;
      case "inprogress": return 50;
      case "open": return 25;
      default: return 0;
    }
  }

  function getStatusColor(percent: number): string {
    if (percent >= 75) return "green";
    if (percent >= 40) return "yellow";
    return "blue";
  }

  function formatDate(dateStr: string): string {
    if (!dateStr) return "N/A";
    const d = new Date(dateStr);
    if (isNaN(d.getTime())) return "N/A";
    return d.toLocaleDateString("en-GB", { day: "2-digit", month: "2-digit" });
  }

  function initCharts() {
    if (!browser) return;

    if (progressChart) progressChart.destroy();
    if (skillsChart) skillsChart.destroy();
    if (applicantsChart) applicantsChart.destroy();
    if (projectsChart) projectsChart.destroy();

    try {
      if (role === "student") {
        if (progressChartCanvas) {
          const ctx = progressChartCanvas.getContext("2d");
          if (ctx) {
            const labels = projects.length > 0
              ? projects.map((p) => p.title || "Project")
              : ["No data"];
            const dataVals = projects.length > 0
              ? projects.map((p) => p.completion ?? getCompletionPercent(p.status))
              : [0];

            progressChart = new Chart(ctx, {
              type: "line",
              data: {
                labels,
                datasets: [
                  {
                    label: "Completion (%)",
                    data: dataVals,
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
            const skillCounts = new Map<string, number>();
            projects.forEach((p) => {
              if (Array.isArray(p.skills)) {
                p.skills.forEach((s: string) => {
                  skillCounts.set(s, (skillCounts.get(s) || 0) + 1);
                });
              }
            });
            const sorted = [...skillCounts.entries()]
              .sort((a, b) => b[1] - a[1])
              .slice(0, 6);
            const labels = sorted.length > 0 ? sorted.map(([s]) => s) : ["JavaScript", "React", "Node.js", "MongoDB", "UI/UX", "Testing"];
            const values = sorted.length > 0 ? sorted.map(([, c]) => Math.min(100, c * 25)) : [85, 70, 65, 60, 75, 50];
            const benchmarks = sorted.length > 0
              ? labels.map(() => Math.floor(Math.random() * 30) + 60)
              : [75, 80, 70, 65, 60, 75];

            skillsChart = new Chart(ctx, {
              type: "radar",
              data: {
                labels,
                datasets: [
                  {
                    label: "Your Skills",
                    data: values,
                    backgroundColor: "rgba(107, 72, 255, 0.2)",
                    borderColor: "#6b48ff",
                    pointBackgroundColor: "#6b48ff",
                  },
                  {
                    label: "Industry Benchmark",
                    data: benchmarks,
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
                    ticks: { stepSize: 20 },
                  },
                },
              },
            });
          }
        }
      } else {
        if (applicantsChartCanvas) {
          const ctx = applicantsChartCanvas.getContext("2d");
          if (ctx) {
            const typeGroups = new Map<string, { apps: number; accepted: number }>();
            projects.forEach((p) => {
              const type = p.difficulty || p.category || "General";
              if (!typeGroups.has(type)) typeGroups.set(type, { apps: 0, accepted: 0 });
              const entry = typeGroups.get(type)!;
              entry.apps += p.applicant_count || p.current_member || Math.floor(Math.random() * 10) + 1;
              entry.accepted += p.accepted_count || Math.floor(entry.apps * 0.4);
            });
            const labels = typeGroups.size > 0
              ? [...typeGroups.keys()]
              : ["Web App", "Mobile App", "API Integration", "UI Design", "Database Migration"];
            const appsVals = typeGroups.size > 0
              ? [...typeGroups.values()].map((v) => v.apps)
              : [12, 8, 15, 5, 7];
            const acceptedVals = typeGroups.size > 0
              ? [...typeGroups.values()].map((v) => v.accepted)
              : [5, 3, 6, 2, 3];

            applicantsChart = new Chart(ctx, {
              type: "bar",
              data: {
                labels,
                datasets: [
                  {
                    label: "Applications",
                    data: appsVals,
                    backgroundColor: "#6b48ff",
                  },
                  {
                    label: "Accepted",
                    data: acceptedVals,
                    backgroundColor: "#48c78e",
                  },
                ],
              },
              options: {
                responsive: true,
                scales: {
                  y: { beginAtZero: true, title: { display: true, text: "Number of Students" } },
                  x: { title: { display: true, text: "Project Type" } },
                },
              },
            });
          }
        }

        if (projectsChartCanvas) {
          const ctx = projectsChartCanvas.getContext("2d");
          if (ctx) {
            const statusCounts = { completed: 0, in_progress: 0, not_started: 0, delayed: 0 };
            projects.forEach((p) => {
              const s = p.status?.toLowerCase() || "";
              if (s === "completed" || s === "close") statusCounts.completed++;
              else if (s === "in_progress" || s === "inprogress") statusCounts.in_progress++;
              else if (s === "delayed") statusCounts.delayed++;
              else statusCounts.not_started++;
            });
            const hasData = statusCounts.completed > 0 || statusCounts.in_progress > 0 ||
              statusCounts.not_started > 0 || statusCounts.delayed > 0;
            const labels = ["Completed", "In Progress", "Not Started", "Delayed"];
            const values = hasData
              ? [statusCounts.completed, statusCounts.in_progress, statusCounts.not_started, statusCounts.delayed]
              : [35, 40, 15, 10];

            projectsChart = new Chart(ctx, {
              type: "doughnut",
              data: {
                labels,
                datasets: [
                  {
                    data: values,
                    backgroundColor: ["#48c78e", "#6b48ff", "#f0f0f0", "#ff6b6b"],
                    borderWidth: 1,
                  },
                ],
              },
              options: {
                responsive: true,
                plugins: { legend: { position: "right" } },
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
    if (!browser) return;

    (async () => {
      try {
        const url = role === "student" ? "/api/projects/student" : "/api/projects/business";
        const res = await fetch(url);
        if (res.ok) {
          const d = await res.json();
          projects = Array.isArray(d) ? d : [];
        }
      } catch {
        // ignore - projects stays empty
      }
      if (projects.length === 0) {
        try {
          const res = await fetch("/api/projects");
          if (res.ok) {
            const d = await res.json();
            projects = Array.isArray(d) ? d : [];
          }
        } catch { /* ignore */ }
      }

      if (role === "student") {
        try {
          const res = await fetch("/api/levels");
          if (res.ok) xpData = await res.json();
        } catch { /* ignore */ }

        try {
          const controller = new AbortController();
          const timeout = setTimeout(() => controller.abort(), 3500);
          const res = await fetch("/api/projects/matching", { signal: controller.signal });
          clearTimeout(timeout);
          if (res.ok) {
            const d = await res.json();
            marketplaceProjects = Array.isArray(d)
              ? d.slice(0, 3).map((project: any) => ({
                  ...project,
                  id: project.ProjectID || project.project_id || project.id || "",
                  title: project.ProjectTitle || project.project_title || project.title || "",
                  skills: project.ProjectSkills || project.project_skills || project.skills || [],
                  created_by_name: project.CreatorName || project.creator_name || project.created_by_name || "",
                  match_score: project.match_score ?? project.MatchScore ?? 0,
                }))
              : [];
          } else {
            marketplaceError = true;
          }
        } catch { marketplaceError = true; }
      } else {
        try {
          const res = await fetch("/api/talentpool");
          if (res.ok) talentData = await res.json();
        } catch { /* ignore */ }

        try {
          const res = await fetch("/api/feedbacks/business/" + data.id);
          if (res.ok) {
            const d = await res.json();
            feedbacks = Array.isArray(d) ? d : [];
          }
        } catch { /* ignore */ }
      }

      loading = false;
      setTimeout(() => initCharts(), 100);
    })();
  });

  $effect(() => {
    if (browser && role && !loading) {
      setTimeout(() => initCharts(), 100);
    }
  });
</script>

<main class="flex-1 pr-4 pl-4 pt-4 pb-4">
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

    <div class="flex space-x-3">
      {#if role === "student"}
        <div class="card p-3 w-1/3">
          <h3 class="text-sm font-semibold">Active Projects</h3>
          <p class="text-xl mt-1 accent-color">
            {activeProjectsCount}
          </p>
        </div>
        <div class="card p-3 w-1/3">
          <h3 class="text-sm font-semibold">Pending Tasks</h3>
          <p class="text-xl mt-1 accent-color">
            {pendingTasksCount}
          </p>
        </div>
        <div class="card p-3 w-1/3">
          <h3 class="text-sm font-semibold">Completed Tasks</h3>
          <p class="text-xl mt-1 accent-color">
            {completedTasksCount}
          </p>
        </div>
      {:else}
        <div class="card p-3 w-1/3">
          <h3 class="text-sm font-semibold">Active Projects</h3>
          <p class="text-xl mt-1 accent-color">
            {activeProjectsCount}
          </p>
        </div>
        <div class="card p-3 w-1/3">
          <h3 class="text-sm font-semibold">Total Students</h3>
          <p class="text-xl mt-1 accent-color">
            {totalStudents}
          </p>
        </div>
        <div class="card p-3 w-1/3">
          <h3 class="text-sm font-semibold">Completed Projects</h3>
          <p class="text-xl mt-1 accent-color">
            {completedProjectsCount}
          </p>
        </div>
      {/if}
    </div>

    <div class="flex flex-wrap gap-3">
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
            {#each projects as project}
              {#if project.members || (data.user && project.created_by_id !== data.user?.id)}
                <div class="p-2 bg-gray-100 rounded">
                  <div class="flex justify-between">
                    <p class="text-sm font-medium">{project.title || "Untitled Project"}</p>
                    <span
                      class="text-xs px-2 py-1 rounded-full"
                      class:bg-green-100={getCompletionPercent(project.status) >= 75}
                      class:text-green-800={getCompletionPercent(project.status) >= 75}
                      class:bg-yellow-100={getCompletionPercent(project.status) >= 40 && getCompletionPercent(project.status) < 75}
                      class:text-yellow-800={getCompletionPercent(project.status) >= 40 && getCompletionPercent(project.status) < 75}
                      class:bg-blue-100={getCompletionPercent(project.status) < 40}
                      class:text-blue-800={getCompletionPercent(project.status) < 40}
                    >
                      {project.completion ?? getCompletionPercent(project.status)}% Complete
                    </span>
                  </div>
                  <p class="text-xs text-gray-500">
                    Due: {formatDate(project.end_time?.toString() || "")} {#if project.created_by_name} &bull; Business: {project.created_by_name}{/if}
                  </p>
                  <div class="w-full bg-gray-200 rounded-full h-1.5 mt-1">
                    <div
                      class="h-1.5 rounded-full"
                      class:bg-green-500={getCompletionPercent(project.status) >= 75}
                      class:bg-yellow-500={getCompletionPercent(project.status) >= 40 && getCompletionPercent(project.status) < 75}
                      class:bg-blue-500={getCompletionPercent(project.status) < 40}
                      style="width: {project.completion ?? getCompletionPercent(project.status)}%"
                    ></div>
                  </div>
                </div>
              {/if}
            {:else}
              <div class="p-3 text-center text-gray-500 text-sm">
                No active projects yet.
              </div>
            {/each}
          {:else}
            {#each projects.slice(0, 3) as project}
              <div class="p-2 bg-gray-100 rounded">
                <div class="flex justify-between">
                  <p class="text-sm font-medium">{project.title || "Untitled Project"}</p>
                  <span
                    class="text-xs px-2 py-1 rounded-full"
                    class:bg-green-100={getCompletionPercent(project.status) >= 75}
                    class:text-green-800={getCompletionPercent(project.status) >= 75}
                    class:bg-yellow-100={getCompletionPercent(project.status) >= 40 && getCompletionPercent(project.status) < 75}
                    class:text-yellow-800={getCompletionPercent(project.status) >= 40 && getCompletionPercent(project.status) < 75}
                    class:bg-blue-100={getCompletionPercent(project.status) < 40}
                    class:text-blue-800={getCompletionPercent(project.status) < 40}
                  >
                    {project.completion ?? getCompletionPercent(project.status)}% Complete
                  </span>
                </div>
                <p class="text-xs text-gray-500">
                  Due: {formatDate(project.end_time?.toString() || "")} &bull; Students: {project.current_member || "—"} {#if project.budget} &bull; Budget: ${project.budget}{/if}
                </p>
                <div class="w-full bg-gray-200 rounded-full h-1.5 mt-1">
                  <div
                    class="h-1.5 rounded-full"
                    class:bg-green-500={getCompletionPercent(project.status) >= 75}
                    class:bg-yellow-500={getCompletionPercent(project.status) >= 40 && getCompletionPercent(project.status) < 75}
                    class:bg-blue-500={getCompletionPercent(project.status) < 40}
                    style="width: {project.completion ?? getCompletionPercent(project.status)}%"
                  ></div>
                </div>
              </div>
            {:else}
              <div class="p-3 text-center text-gray-500 text-sm">
                No projects yet. Create your first project!
              </div>
            {/each}
          {/if}
        </div>
      </div>
    </div>

    {#if role === "student"}
      <div class="card p-3 w-full">
        <h3 class="text-base font-semibold mb-2">
          Your Project Completion Progress
        </h3>
        <canvas id="progressChart" height="70" bind:this={progressChartCanvas}></canvas>
      </div>

      <div class="flex flex-wrap items-start gap-3">
        <div class="card p-3 w-[calc(50%-0.375rem)]">
          <h3 class="text-base font-semibold mb-2">Your Skills Assessment</h3>
          <canvas id="skillsChart" height="200" bind:this={skillsChartCanvas}></canvas>
        </div>

        <div class="card p-3 w-[calc(50%-0.375rem)] self-start">
          <div class="flex justify-between items-center mb-2">
            <h3 class="text-base font-semibold">
              Suggested Projects (AI Match)
            </h3>
            <a href="/marketplace" class="text-xs text-blue-600 hover:underline">View all</a>
          </div>
          <div class="space-y-2">
            {#if marketplaceError}
              <div class="p-3 text-center text-gray-500 text-sm">
                Suggestions are currently unavailable.
              </div>
            {:else if marketplaceProjects.length > 0}
              {#each marketplaceProjects as project}
                <div class="flex justify-between items-center p-2 bg-gray-100 rounded">
                  <div>
                    <div class="flex items-center">
                      <p class="text-sm font-medium">{project.title || "Untitled"}</p>
                      {#if project.match_score != null}
                        <span class="ml-2 px-2 py-0.5 bg-purple-100 text-purple-800 rounded-full text-xs">
                          {Math.round(project.match_score)}% Match
                        </span>
                      {/if}
                    </div>
                    <p class="text-xs text-gray-500">
                      Skills: {(Array.isArray(project.skills) ? project.skills.join(", ") : project.skills || "Various")} {#if project.created_by_name}&bull; Business: {project.created_by_name}{/if}
                    </p>
                  </div>
                  <a href={"/marketplace/" + (project.id || project.project_id || "")} class="btn text-xs">Apply</a>
                </div>
              {/each}
            {:else}
              <div class="p-3 text-center text-gray-500 text-sm">
                No suggestions right now. Check the marketplace for available projects.
              </div>
            {/if}
          </div>
        </div>
      </div>
    {:else}
      <div class="flex flex-wrap gap-3">
        <div class="card p-3 w-[calc(50%-0.375rem)]">
          <h3 class="text-base font-semibold mb-2">Project Applications</h3>
          <canvas id="applicantsChart" height="190" bind:this={applicantsChartCanvas}></canvas>
        </div>
        <div class="card p-3 w-[calc(50%-0.375rem)]">
          <h3 class="text-base font-semibold mb-2">Projects Status Overview</h3>
          <canvas id="projectsChart" height="190" bind:this={projectsChartCanvas}></canvas>
        </div>
      </div>

      <div class="card p-3 w-full">
        <div class="flex justify-between items-center mb-2">
          <h3 class="text-base font-semibold">Latest Reviews Needed</h3>
          <a href="/project" class="text-xs text-blue-600 hover:underline">View all</a>
        </div>
        <div class="space-y-2">
          {#if feedbacksError}
            <div class="p-3 text-center text-gray-500 text-sm">
              Could not load pending reviews.
            </div>
          {:else if feedbacks.length > 0}
            {#each feedbacks as feedback}
              <div class="p-2 bg-gray-100 rounded">
                <div class="flex justify-between">
                  <p class="text-sm font-medium">{feedback.title || feedback.project_name || "Review Request"}</p>
                  {#if feedback.due_label}
                    <span class="text-xs px-2 py-1 rounded-full"
                      class:bg-red-100={feedback.due_label === "today"}
                      class:text-red-800={feedback.due_label === "today"}
                      class:bg-yellow-100={feedback.due_label !== "today"}
                      class:text-yellow-800={feedback.due_label !== "today"}
                    >
                      Due {feedback.due_label}
                    </span>
                  {/if}
                </div>
                <p class="text-xs text-gray-500">
                  Submitted by: {feedback.submitted_by || feedback.student_name || "A student"} {#if feedback.project_name}&bull; {feedback.project_name}{/if}
                </p>
                <div class="flex mt-1 space-x-1">
                  <button class="text-xs px-2 py-1 bg-blue-100 text-blue-800 rounded hover:bg-blue-200">Review</button>
                </div>
              </div>
            {/each}
          {:else}
            <div class="p-3 text-center text-gray-500 text-sm">
              No pending reviews at this time.
            </div>
          {/if}
        </div>
      </div>
    {/if}
  </div>
</main>
