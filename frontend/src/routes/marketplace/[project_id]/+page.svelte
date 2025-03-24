<script lang="ts">
  import type { PageData } from "./$types";

  let { data }: { data: PageData } = $props();

  const project = data.project;
</script>

<svelte:head>
  <title>Project Detail</title>
</svelte:head>

<main class="flex-1 ml-64 pl-4 pr-4 pt-4">
  <div class="flex space-x-4">
    <!-- Project Info Card -->
    <div class="card p-4 w-2/3">
      <h3 class="text-lg font-semibold mb-3">{project.title}</h3>
      <div class="space-y-3">
        <p class="text-sm">
          <span class="font-medium">Skills Required:</span>{project.skills}
        </p>
        <p class="text-sm">
          <!-- Change this base on AI matchi system -->
          <span class="font-medium">Match:</span>
          <span class="accent-color">90%</span>
        </p>
        <p class="text-sm">
          <span class="font-medium">Timeline:</span>
          {project.start_time} - {project.end_time}
        </p>
        <p class="text-sm">
          <span class="font-medium">Posted:</span>
          {project.created_at} by
          <!-- Change this in backend: need to pass the username into the create project handler and store it -->
          <a href="#" class="text-[#6b48ff] hover:underline"
            >{project.created_by}</a
          >
        </p>
        <p class="text-sm">
          <span class="font-medium">Description:</span>
          {project.description}
        </p>
        <!-- What the hell is this shit, is it required? -->
        <p class="text-sm">
          <span class="font-medium">Deliverables:</span> Frontend UI, Backend API,
          Documentation
        </p>
      </div>
      <!-- <button class="btn-secondary mt-4"
        >Chat with Mentor (change this button please we only support group chat
        chat)</button
      > -->
    </div>

    <!-- Proposal Form -->
    <div class="card p-4 w-1/3">
      <h3 class="text-base font-semibold mb-3">Apply to Project</h3>
      <form class="space-y-3">
        <div>
          <label class="block text-sm font-medium mb-1" for="proposal"
            >Your Proposal</label
          >
          <textarea
            id="proposal"
            class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
            rows="5"
            placeholder="Write why you're a good fit..."
          ></textarea>
        </div>
        <div class="flex space-x-2">
          <button type="button" class="btn-secondary w-1/2">
            AI Suggest
          </button>
          <button type="submit" class="btn w-1/2">Apply</button>
        </div>
      </form>
    </div>
  </div>
</main>
