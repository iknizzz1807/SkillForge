<script lang="ts">
  import type { PageData } from "./$types";

  let { data }: { data: PageData } = $props();

  const role: string | undefined = data.role;

  const token = data.token;

  const project = data.project;

  let motivationInput: string = $state("");
  let proposalInput: string = $state("");

  // Thêm state để kiểm soát việc hiển thị modal xác nhận
  let showConfirmModal: boolean = $state(false);
  let isSubmitting: boolean = $state(false);
  let errorMessage: string | null = $state(null);
  let successMessage: string | null = $state(null);

  // Functions to control modal
  function openConfirmModal() {
    if (!motivationInput.trim() || !proposalInput.trim()) {
      errorMessage = "Please fill out all required fields.";
      return;
    }

    // Remove error message if form is valid
    errorMessage = null;
    showConfirmModal = true;
  }

  function closeConfirmModal() {
    showConfirmModal = false;
  }

  // Complete the submitForm function
  const submitForm = async () => {
    try {
      isSubmitting = true;

      const response = await fetch(`/api/projects/${project.id}/applications`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          motivation: motivationInput,
          proposal: proposalInput,
          project_id: project.id,
        }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || "Failed to submit application");
      }

      // Handle success
      successMessage = "Your application has been submitted successfully!";
      motivationInput = "";
      proposalInput = "";

      // Close the modal
      closeConfirmModal();

      // You could redirect the user or show success message
      // window.location.href = "/project";
    } catch (error) {
      console.error("Error submitting application:", error);
      errorMessage = "Failed to submit application";
      closeConfirmModal();
    } finally {
      isSubmitting = false;
    }
  };

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

<svelte:head>
  <title>Project Detail</title>
</svelte:head>

<header class="flex justify-between items-center ml-64 pr-4 pl-4 pt-4">
  <div class="flex items-center mb-6">
    <a href="/project">
      <button class="text-gray-500 hover:text-gray-700 mr-3">
        <svg
          class="w-5 h-5"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M10 19l-7-7m0 0l7-7m-7 7h18"
          ></path>
        </svg>
      </button>
    </a>
    <div>
      <h2 class="text-xl font-semibold">{project.title}</h2>
      <p class="text-sm text-gray-600">Manage your project</p>
    </div>
  </div>
</header>

<main class="flex-1 ml-64 pl-4 pr-4">
  <div class="flex flex-col md:flex-row space-y-4 md:space-y-0 md:space-x-4">
    <!-- Project Info Card -->
    <div class="card p-4 w-full md:w-2/3 h-fit">
      <h3 class="text-lg font-semibold mb-3">{project.title}</h3>
      <div class="space-y-3">
        <p class="text-sm">
          <span class="font-bold">Members:</span>
          {project.current_member}/{project.max_member}
        </p>
        <p class="text-sm">
          <span class="font-bold">Skills Required:</span>
          {project.skills}
        </p>
        <p class="text-sm">
          <!-- Change this base on AI matchi system -->
          <span class="font-bold">Match:</span>
          <span class="accent-color">90%</span>
        </p>
        <p class="text-sm">
          <span class="font-bold">Timeline:</span>
          {formatDate(project.start_time)} - {formatDate(project.end_time)}
        </p>
        <p class="text-sm">
          <span class="font-bold">Posted:</span>
          {formatDate(project.created_at)} by
          <!-- Change this in backend: need to pass the username into the create project handler and store it -->
          <a
            href={"/profile/user/" + project.created_by_id}
            class="text-[#6b48ff] hover:underline">{project.created_by_name}</a
          >
        </p>
        <p class="text-sm">
          <span class="font-bold">Description:</span>
          {project.description}
        </p>
        <p class="text-sm">
          <span class="font-bold">Deliverables:</span> Frontend UI, Backend API,
          Documentation
        </p>
      </div>
    </div>

    <!-- Proposal Form -->
    <div class="card p-6 w-full md:w-1/3 h-fit">
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
            bind:value={motivationInput}
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
            bind:value={proposalInput}
          ></textarea>
          <!-- <p class="text-xs text-gray-500 mt-1">
            Supports Markdown formatting. Include sections like "Approach",
            "Timeline", and "Technical Details".
          </p> -->
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
          <button type="button" class="btn w-1/2" onclick={openConfirmModal}
            >Submit Application</button
          >
        </div>
      </form>

      <div class="mt-4 bg-purple-100 p-3 rounded-lg">
        <div class="flex items-center justify-center space-x-2">
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
        <!-- <p class="text-xs text-gray-600 mt-1">
          Based on your profile skills and experience, you are highly compatible
          with this project.
        </p> -->
      </div>
    </div>
  </div>
</main>

<!-- Confirmation Modal -->
{#if showConfirmModal}
  <div
    class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 modal"
  >
    <div class="bg-white rounded-lg p-6 w-full max-w-md">
      <div class="text-center">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-12 w-12 mx-auto text-blue-500"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
          />
        </svg>

        <h2 class="text-xl font-semibold mt-4">Confirm Application</h2>

        <p class="mt-2 text-gray-600">
          Are you sure you want to apply to the project "{project.title}"?
        </p>

        <div class="flex justify-center space-x-3 mt-6">
          <button
            class="px-4 py-2 bg-gray-200 text-gray-800 rounded-md hover:bg-gray-300"
            onclick={closeConfirmModal}
            disabled={isSubmitting}
          >
            Cancel
          </button>
          <button
            class="px-4 py-2 bg-[#6b48ff] text-white rounded-md hover:bg-[#5a3dd3]"
            onclick={submitForm}
            disabled={isSubmitting}
          >
            {#if isSubmitting}
              <span>Submitting...</span>
            {:else}
              <span>Submit</span>
            {/if}
          </button>
        </div>
      </div>
    </div>
  </div>
{/if}

<!-- Error/Success Messages -->
{#if errorMessage}
  <div
    class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mt-4"
    role="alert"
  >
    <span class="block sm:inline">{errorMessage}</span>
  </div>
{/if}

{#if successMessage}
  <div
    class="bg-green-100 border border-green-400 text-green-700 px-4 py-3 rounded mt-4"
    role="alert"
  >
    <span class="block sm:inline">{successMessage}</span>
  </div>
{/if}
