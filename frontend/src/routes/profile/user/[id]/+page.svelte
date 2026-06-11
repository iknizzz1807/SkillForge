<script lang="ts">
  import type { PageData } from "./$types";

  let { data }: { data: PageData } = $props();
  let linkCopied = $state(false);

  const copyPortfolioLink = async () => {
    const url = window.location.origin + "/portfolios/" + data.id + ".html";
    try {
      await navigator.clipboard.writeText(url);
      linkCopied = true;
      setTimeout(() => (linkCopied = false), 2000);
    } catch {
      const input = document.createElement("input");
      input.value = url;
      document.body.appendChild(input);
      input.select();
      document.execCommand("copy");
      document.body.removeChild(input);
      linkCopied = true;
      setTimeout(() => (linkCopied = false), 2000);
    }
  };

  const formatDate = (d: string) => d ? new Date(d).toLocaleDateString() : "";
</script>

<main class="flex-1 p-6">
  <div class="max-w-3xl mx-auto space-y-6">
    <a href="/talent" class="inline-flex items-center text-sm text-gray-500 hover:text-[#6b48ff] mb-2">
      <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
      </svg>
      Back to Talent Pool
    </a>

    <!-- Profile Header -->
    <div class="card p-6">
      <div class="flex flex-col sm:flex-row items-center sm:items-start gap-4">
        <img
          class="w-20 h-20 rounded-full object-cover border-4 border-[#6b48ff] flex-shrink-0"
          src={`https://ui-avatars.com/api/?name=${encodeURIComponent(data.name || "U")}&background=6b48ff&color=fff`}
          alt={data.name}
        />
        <div class="text-center sm:text-left flex-1">
          <h1 class="text-2xl font-bold">{data.name}</h1>
          {#if data.title}
            <p class="text-gray-500">{data.title}</p>
          {/if}
          {#if data.email}
            <p class="text-sm text-gray-400">{data.email}</p>
          {/if}
          <div class="flex flex-wrap gap-2 mt-3 justify-center sm:justify-start">
            <a
              href={"/portfolios/" + data.id + ".html"}
              target="_blank"
              class="btn text-sm inline-flex items-center"
            >
              <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
              </svg>
              View Portfolio
            </a>
            <button class="btn-secondary text-sm inline-flex items-center" onclick={copyPortfolioLink}>
              <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 5H6a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2v-1M8 5a2 2 0 002 2h2a2 2 0 002-2M8 5a2 2 0 012-2h2a2 2 0 012 2m0 0h2a2 2 0 012 2v3m2 4H10m0 0l3-3m-3 3l3 3" />
              </svg>
              {linkCopied ? "Copied!" : "Copy Portfolio Link"}
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Skills -->
    {#if data.skills && data.skills.length > 0}
      <div class="card p-4">
        <h2 class="text-base font-semibold mb-3">Skills</h2>
        <div class="flex flex-wrap gap-2">
          {#each data.skills as skill}
            <span class="bg-purple-100 text-purple-700 px-3 py-1 rounded-full text-sm">{typeof skill === "string" ? skill : skill.name ?? skill.skill ?? skill}</span>
          {/each}
        </div>
      </div>
    {/if}

    <!-- Badges -->
    {#if data.badges && data.badges.length > 0}
      <div class="card p-4">
        <h2 class="text-base font-semibold mb-3">Badges & Achievements</h2>
        <div class="grid grid-cols-2 sm:grid-cols-3 gap-3">
          {#each data.badges as badge}
            <div class="text-center p-3 rounded border border-[#6b48ff] bg-purple-50">
              <div class="text-2xl mb-1">{badge.icon}</div>
              <p class="text-sm font-medium">{badge.name}</p>
              {#if badge.description}
                <p class="text-xs text-gray-500">{badge.description}</p>
              {/if}
              {#if badge.date}
                <p class="text-xs text-[#6b48ff] mt-1">{formatDate(badge.date)}</p>
              {/if}
            </div>
          {/each}
        </div>
      </div>
    {/if}

    <!-- Projects -->
    {#if data.projects && data.projects.length > 0}
      <div class="card p-4">
        <h2 class="text-base font-semibold mb-3">Projects</h2>
        <div class="space-y-3">
          {#each data.projects as project}
            <div class="p-3 bg-gray-50 rounded border border-gray-200">
              <div class="flex justify-between items-start">
                <p class="text-sm font-medium">{project.title}</p>
                {#if project.status}
                  <span class="text-xs px-2 py-0.5 rounded-full {project.status === "completed" || project.status === "close" ? "bg-green-100 text-green-700" : project.status === "in_progress" ? "bg-blue-100 text-blue-700" : "bg-yellow-100 text-yellow-700"}">{project.status}</span>
                {/if}
              </div>
              {#if project.description}
                <p class="text-xs text-gray-500 mt-1">{project.description}</p>
              {/if}
              {#if project.end_time}
                <p class="text-xs text-gray-400 mt-1">Due: {formatDate(project.end_time)}</p>
              {/if}
            </div>
          {/each}
        </div>
      </div>
    {/if}

    <!-- Feedbacks -->
    {#if data.feedbacks && data.feedbacks.length > 0}
      <div class="card p-4">
        <h2 class="text-base font-semibold mb-3">Feedback</h2>
        <div class="space-y-3">
          {#each data.feedbacks as fb}
            <div class="p-3 bg-gray-50 rounded border border-gray-200">
              <div class="flex justify-between items-start">
                <p class="text-sm font-medium">{fb.from_name ?? fb.from ?? "Anonymous"}</p>
                {#if fb.rating}
                  <span class="text-sm text-yellow-500">{Array(fb.rating).fill("★").join("")}</span>
                {/if}
              </div>
              {#if fb.content}
                <p class="text-xs text-gray-600 mt-1">{fb.content}</p>
              {/if}
            </div>
          {/each}
        </div>
      </div>
    {/if}
  </div>
</main>
