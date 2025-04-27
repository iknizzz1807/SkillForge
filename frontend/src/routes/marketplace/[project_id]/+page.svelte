<script lang="ts">
  import type { PageData } from "./$types";

  let { data }: { data: PageData } = $props();

  const role: string | null | undefined = data.role;

  // const token = data.token;

  const project = data.project;
  // Biến hasJoined và hasApplied là để dùng khi role là student, nếu đã joined hoặc apply thì sẽ không hiện form
  // apply project nữa mà sẽ điều chỉnh UI sao cho hợp lý để thông báo rằng họ đã apply hoặc joined dự án này
  let hasJoined: boolean = $state(data.hasJoined);
  let hasApplied: boolean = $state(data.hasApplied);

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

  const submitForm = async () => {
    try {
      isSubmitting = true;

      // Call the SvelteKit API endpoint instead of backend directly
      const response = await fetch(`/api/applications`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          project_id: project.id,
          motivation: motivationInput,
          detailed_proposal: proposalInput,
        }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || "Failed to submit application");
      }

      // Cập nhật hasApplied thành true sau khi submit thành công
      hasApplied = true;

      // Handle success
      successMessage = "Your application has been submitted successfully!";
      motivationInput = "";
      proposalInput = "";

      // Close the modal
      closeConfirmModal();
    } catch (error) {
      console.error("Error submitting application:", error);
      errorMessage =
        error instanceof Error ? error.message : "Failed to submit application";
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
  <div class="flex items-center mb-4">
    <a href="/marketplace">
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
      <p class="text-sm text-gray-600">Project Details</p>
    </div>
  </div>
</header>

<main class="flex-1 ml-64 pl-4 pr-4">
  <!-- Two-column layout -->
  <div class="flex flex-col md:flex-row space-y-4 md:space-y-0 md:space-x-4">
    <!-- Left column - Main content: Business Info & Project Info stacked -->
    <div class="w-full md:w-2/3 space-y-4">
      <!-- Business Info Card -->
      <div class="card p-4 h-fit border-l-4 border-l-[#6b48ff]">
        <div class="flex items-start">
          <!-- Business Avatar với hiệu ứng khung viền -->
          <div class="mr-4">
            <div
              class="w-16 h-16 rounded-full bg-gray-200 overflow-hidden border-2 border-[#6b48ff] shadow-md"
            >
              <img
                src="https://ui-avatars.com/api/?name={project.created_by_name}&background=6b48ff&color=fff"
                alt="{project.created_by_name}'s avatar"
                class="w-full h-full object-cover"
              />
            </div>
          </div>

          <!-- Business Info -->
          <div class="flex-1">
            <div class="flex justify-between items-start">
              <div>
                <h3 class="text-lg font-semibold flex items-center">
                  {project.created_by_name}
                  <span
                    class="ml-2 bg-purple-100 text-[#6b48ff] text-xs px-2 py-0.5 rounded-full"
                    >Verified</span
                  >
                </h3>
                <p class="text-sm text-gray-600">Project Creator</p>
                <p class="text-sm mt-1 flex items-center">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-4 w-4 mr-1 text-gray-500"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                  >
                    <path
                      fill-rule="evenodd"
                      d="M5.05 4.05a7 7 0 119.9 9.9L10 18.9l-4.95-4.95a7 7 0 010-9.9zM10 11a2 2 0 100-4 2 2 0 000 4z"
                      clip-rule="evenodd"
                    />
                  </svg>
                  <span class="font-medium">Location:</span> Ho Chi Minh City, Vietnam
                </p>
                <p class="text-sm flex items-center">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-4 w-4 mr-1 text-gray-500"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                  >
                    <path
                      fill-rule="evenodd"
                      d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z"
                      clip-rule="evenodd"
                    />
                  </svg>
                  <span class="font-medium">Joined:</span> January 2023
                </p>
              </div>

              <a
                href={"/profile/business/" + project.created_by_id}
                class="btn-secondary text-sm px-3 py-1.5 flex items-center shadow-sm hover:shadow transition-all"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-4 w-4 mr-1"
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
                View Profile
              </a>
            </div>

            <!-- Business Description with highlight -->
            <div
              class="mt-2 p-2 bg-gray-50 rounded-md border-l-2 border-gray-200"
            >
              <p class="text-sm text-gray-700">
                A leading tech company specializing in innovative solutions for
                education and workforce development. With over 5 years of
                experience in the industry, we've successfully completed 20+
                projects.
              </p>
            </div>

            <!-- Business Stats - với hiệu ứng highlight tốt hơn -->
            <div class="flex mt-3 space-x-4 text-sm">
              <div class="flex items-center px-2 py-1 bg-purple-50 rounded-md">
                <div class="flex items-center justify-center w-5 h-5 mr-1">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-4 w-4 text-[#6b48ff]"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                  >
                    <path
                      fill-rule="evenodd"
                      d="M6 6V5a3 3 0 013-3h2a3 3 0 013 3v1h2a2 2 0 012 2v3.57A22.952 22.952 0 0110 13a22.95 22.95 0 01-8-1.43V8a2 2 0 012-2h2zm2-1a1 1 0 011-1h2a1 1 0 011 1v1H8V5zm1 5a1 1 0 011-1h.01a1 1 0 110 2H10a1 1 0 01-1-1z"
                      clip-rule="evenodd"
                    />
                    <path
                      d="M2 13.692V16a2 2 0 002 2h12a2 2 0 002-2v-2.308A24.974 24.974 0 0110 15c-2.796 0-5.487-.46-8-1.308z"
                    />
                  </svg>
                </div>
                <span class="flex items-center"
                  ><strong>12</strong> Projects</span
                >
              </div>
              <div class="flex items-center px-2 py-1 bg-purple-50 rounded-md">
                <div class="flex items-center justify-center w-5 h-5 mr-1">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-4 w-4 text-[#6b48ff]"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                  >
                    <path
                      d="M13 6a3 3 0 11-6 0 3 3 0 016 0zM18 8a2 2 0 11-4 0 2 2 0 014 0zM14 15a4 4 0 00-8 0v3h8v-3zM6 8a2 2 0 11-4 0 2 2 0 014 0zM16 18v-3a5.972 5.972 0 00-.75-2.906A3.005 3.005 0 0119 15v3h-3zM4.75 12.094A5.973 5.973 0 004 15v3H1v-3a3 3 0 013.75-2.906z"
                    />
                  </svg>
                </div>
                <span class="flex items-center"
                  ><strong>45</strong> Students</span
                >
              </div>
              <div class="flex items-center px-2 py-1 bg-purple-50 rounded-md">
                <div class="flex items-center justify-center w-5 h-5 mr-1">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-4 w-4 text-[#6b48ff]"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                  >
                    <path
                      d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"
                    />
                  </svg>
                </div>
                <span class="flex items-center"
                  ><strong>4.8</strong>/5 Rating</span
                >
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Project Info Card (stacked below business info) -->
      <div class="card p-6 h-fit relative overflow-hidden mb-4">
        <div class="flex justify-between items-start mb-4">
          <h3 class="text-xl font-semibold text-gray-800">{project.title}</h3>
          <span
            class="px-2 py-1 bg-green-100 text-green-800 text-xs rounded-full flex items-center"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-3 w-3 mr-1"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M5 13l4 4L19 7"
              />
            </svg>
            Active
          </span>
        </div>

        <div class="grid grid-cols-2 gap-4 mb-4">
          <div class="bg-gray-50 p-3 rounded-lg flex items-center">
            <div class="bg-[#6b48ff] bg-opacity-10 p-2 rounded-full mr-3">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5 text-[#6b48ff]"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  d="M13 6a3 3 0 11-6 0 3 3 0 016 0zM18 8a2 2 0 11-4 0 2 2 0 014 0zM14 15a4 4 0 00-8 0v3h8v-3zM6 8a2 2 0 11-4 0 2 2 0 014 0zM16 18v-3a5.972 5.972 0 00-.75-2.906A3.005 3.005 0 0119 15v3h-3zM4.75 12.094A5.973 5.973 0 004 15v3H1v-3a3 3 0 013.75-2.906z"
                />
              </svg>
            </div>
            <div>
              <p class="text-xs text-gray-500">Members</p>
              <p class="text-sm font-semibold">
                {project.current_member}/{project.max_member}
              </p>
            </div>
          </div>

          <div class="bg-gray-50 p-3 rounded-lg flex items-center">
            <div class="bg-[#ff6f61] bg-opacity-10 p-2 rounded-full mr-3">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5 text-[#ff6f61]"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  fill-rule="evenodd"
                  d="M11.3 1.046A1 1 0 0112 2v5h4a1 1 0 01.82 1.573l-7 10A1 1 0 018 18v-5H4a1 1 0 01-.82-1.573l7-10a1 1 0 011.12-.38z"
                  clip-rule="evenodd"
                />
              </svg>
            </div>
            <div>
              <p class="text-xs text-gray-500">AI Match</p>
              <p class="text-sm font-semibold text-[#ff6f61]">90% Match</p>
            </div>
          </div>
        </div>

        <div class="space-y-3">
          <div class="flex items-center">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-5 w-5 text-[#6b48ff] mr-2"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fill-rule="evenodd"
                d="M2 5a2 2 0 012-2h12a2 2 0 012 2v10a2 2 0 01-2 2H4a2 2 0 01-2-2V5zm3.293 1.293a1 1 0 011.414 0l3 3a1 1 0 010 1.414l-3 3a1 1 0 01-1.414-1.414L7.586 10 5.293 7.707a1 1 0 010-1.414zM11 12a1 1 0 100 2h3a1 1 0 100-2h-3z"
                clip-rule="evenodd"
              />
            </svg>
            <p class="text-sm">
              <span class="font-bold">Skills Required:</span>
              <span class="ml-1 inline-flex flex-wrap gap-1">
                {#if typeof project.skills === "string"}
                  {#each project.skills.split(",") as skill}
                    <span
                      class="bg-purple-100 text-[#6b48ff] px-2 py-0.5 rounded text-xs"
                      >{skill.trim()}</span
                    >
                  {/each}
                {:else if Array.isArray(project.skills)}
                  {#each project.skills as skill}
                    <span
                      class="bg-purple-100 text-[#6b48ff] px-2 py-0.5 rounded text-xs"
                      >{typeof skill === "string" ? skill.trim() : skill}</span
                    >
                  {/each}
                {:else}
                  <span
                    class="bg-purple-100 text-[#6b48ff] px-2 py-0.5 rounded text-xs"
                    >{project.skills || "No skills specified"}</span
                  >
                {/if}
              </span>
            </p>
          </div>

          <div class="flex items-center">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-5 w-5 text-[#6b48ff] mr-2"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fill-rule="evenodd"
                d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z"
                clip-rule="evenodd"
              />
            </svg>
            <p class="text-sm">
              <span class="font-bold">Timeline:</span>
              <span class="ml-1"
                >{formatDate(project.start_time)} - {formatDate(
                  project.end_time
                )}</span
              >
            </p>
          </div>

          <div class="flex items-center">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-5 w-5 text-[#6b48ff] mr-2"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fill-rule="evenodd"
                d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
                clip-rule="evenodd"
              />
            </svg>
            <p class="text-sm">
              <span class="font-bold">Posted:</span>
              <span class="ml-1"
                >{formatDate(project.created_at)} by
                <a
                  href={"/profile/user/" + project.created_by_id}
                  class="text-[#6b48ff] hover:underline font-medium"
                  >{project.created_by_name}</a
                ></span
              >
            </p>
          </div>

          <!-- Project description with better styling -->
          <div class="mt-3 border-l-4 border-[#6b48ff] pl-3 py-1">
            <h4 class="font-semibold text-sm flex items-center mb-1">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5 text-[#6b48ff] mr-1"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  fill-rule="evenodd"
                  d="M4 4a2 2 0 012-2h4.586A2 2 0 0112 2.586L15.414 6A2 2 0 0116 7.414V16a2 2 0 01-2 2H6a2 2 0 01-2-2V4zm2 6a1 1 0 011-1h6a1 1 0 110 2H7a1 1 0 01-1-1zm1 3a1 1 0 100 2h6a1 1 0 100-2H7z"
                  clip-rule="evenodd"
                />
              </svg>
              Description
            </h4>
            <p class="text-sm text-gray-700">{project.description}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Right column - Proposal Form or Status Card -->
    {#if role === "student"}
      <div class="w-full md:w-1/3">
        <div class="card p-6 h-fit sticky top-4">
          {#if hasJoined}
            <!-- Already joined the project -->
            <div class="text-center">
              <div class="bg-green-100 p-4 rounded-lg mb-4">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-12 w-12 mx-auto text-green-500 mb-2"
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
                <h3 class="text-base font-semibold mb-2">
                  You're part of this project!
                </h3>
                <p class="text-sm text-gray-600">
                  You have successfully joined this project. Check your
                  dashboard for project details and tasks.
                </p>
              </div>
              <a href="/project/{project.id}" class="btn w-full"
                >Go to Project Dashboard</a
              >
            </div>
          {:else if hasApplied}
            <!-- Already applied to the project -->
            <div class="text-center">
              <div class="bg-blue-100 p-4 rounded-lg mb-4">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-12 w-12 mx-auto text-blue-500 mb-2"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
                  />
                </svg>
                <h3 class="text-base font-semibold mb-2">
                  Application Submitted
                </h3>
                <p class="text-sm text-gray-600">
                  Your application for this project is under review. We'll
                  notify you when there's an update.
                </p>
              </div>
              <a href="/project/application" class="btn w-full"
                >View My Applications</a
              >
            </div>
          {:else}
            <!-- Show application form -->
            <h3 class="text-base font-semibold mb-4">Apply to Project</h3>
            <form class="space-y-4">
              <div>
                <label class="block text-sm font-medium mb-1" for="motivation">
                  Motivation
                  <span class="text-xs text-gray-500 font-normal"
                    >(Required)</span
                  >
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
                  <span class="text-xs text-gray-500 font-normal"
                    >(Required)</span
                  >
                </label>
                <textarea
                  id="proposal"
                  class="w-full p-3 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
                  rows="6"
                  placeholder="Describe your approach, timeline, and implementation strategy"
                  required
                  bind:value={proposalInput}
                ></textarea>
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
                <button
                  type="button"
                  class="btn w-1/2"
                  onclick={openConfirmModal}>Submit Application</button
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
                <h4 class="text-sm font-medium text-gray-700">
                  AI Match Score: 90%
                </h4>
              </div>
            </div>
          {/if}
        </div>
      </div>
    {/if}
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

<!-- Error Modal -->
{#if errorMessage}
  <div
    class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 modal"
  >
    <div class="bg-white rounded-lg p-6 w-full max-w-md">
      <div class="text-center">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-12 w-12 mx-auto text-red-500"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
          />
        </svg>

        <h2 class="text-xl font-semibold mt-4">Application Error</h2>

        <p class="mt-2 text-gray-600">
          {errorMessage}
        </p>

        <div class="flex justify-center mt-6">
          <button
            class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700"
            onclick={() => (errorMessage = null)}
          >
            Close
          </button>
        </div>
      </div>
    </div>
  </div>
{/if}

<!-- Success Modal -->
{#if successMessage}
  <div
    class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 modal"
  >
    <div class="bg-white rounded-lg p-6 w-full max-w-md">
      <div class="text-center">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-12 w-12 mx-auto text-green-500"
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

        <h2 class="text-xl font-semibold mt-4">Success!</h2>

        <p class="mt-2 text-gray-600">
          {successMessage}
        </p>

        <div class="flex justify-center space-x-3 mt-6">
          <button
            class="px-4 py-2 bg-[#6b48ff] text-white rounded-md hover:bg-[#5a3dd3]"
            onclick={() => {
              successMessage = null;
              // Optional: Redirect to applications page
              window.location.href = "/project/application";
            }}
          >
            View My Applications
          </button>
          <button
            class="px-4 py-2 bg-gray-200 text-gray-800 rounded-md hover:bg-gray-300"
            onclick={() => (successMessage = null)}
          >
            Close
          </button>
        </div>
      </div>
    </div>
  </div>
{/if}
