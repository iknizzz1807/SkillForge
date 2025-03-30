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
    <div class="card p-6 w-1/3 h-fit">
      <h3 class="text-base font-semibold mb-4">Apply to Project</h3>
      <form class="space-y-4">
        <div>
          <label class="block text-sm font-medium mb-1" for="motivation">
            Motivation
            <span class="text-xs text-gray-500 font-normal">(Required)</span>
          </label>
          <textarea
            id="motivation"
            class="w-full p-3 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
            rows="3"
            placeholder="Why are you interested in this project? (100-300 characters)"
            required
          ></textarea>
        </div>

        <div>
          <label class="block text-sm font-medium mb-1" for="proposal">
            Detailed Proposal
            <span class="text-xs text-gray-500 font-normal">(Required)</span>
          </label>
          <textarea
            id="proposal"
            class="w-full p-3 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
            rows="6"
            placeholder="Describe your approach, timeline, and implementation strategy"
            required
          ></textarea>
          <p class="text-xs text-gray-500 mt-1">
            Supports Markdown formatting. Include sections like "Approach",
            "Timeline", and "Technical Details".
          </p>
        </div>

        <div>
          <label
            class="flex items-center space-x-2 text-sm mb-4 cursor-pointer"
          >
            <input
              type="checkbox"
              class="form-checkbox rounded text-[#6b48ff]"
            />
            <span
              >I understand that my profile information will be shared with the
              project owner</span
            >
          </label>
        </div>

        <div class="flex space-x-2">
          <button
            type="button"
            class="btn-secondary w-1/2 flex items-center justify-center space-x-1"
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
            <span>AI Suggest</span>
          </button>
          <button type="submit" class="btn w-1/2">Submit Application</button>
        </div>
      </form>

      <div class="mt-4 bg-purple-50 p-3 rounded-lg">
        <div class="flex items-center space-x-2">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-5 w-5 text-[#6b48ff]"
            viewBox="0 0 20 20"
            fill="currentColor"
          >
            <path
              fill-rule="evenodd"
              d="M11.3 1.046A1 1 0 0112 2v5h4a1 1 0 01.82 1.573l-7 10A1 1 0 018 18v-5H4a1 1 0 01-.82-1.573l7-10a1 1 0 011.12-.38z"
              clip-rule="evenodd"
            />
          </svg>
          <h4 class="text-sm font-medium text-gray-700">AI Match Score: 90%</h4>
        </div>
        <p class="text-xs text-gray-600 mt-1">
          Based on your profile skills and experience, you are highly compatible
          with this project.
        </p>
      </div>
    </div>
  </div>
</main>
